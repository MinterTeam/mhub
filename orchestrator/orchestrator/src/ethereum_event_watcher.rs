//! Ethereum Event watcher watches for events such as a deposit to the Peggy Ethereum contract or a validator set update
//! or a transaction batch update. It then responds to these events by performing actions on the Cosmos chain if required

use clarity::{Address as EthAddress, Uint256};
use contact::client::Contact;
use cosmos_peggy::{query::get_last_event_nonce, send::send_ethereum_claims};
use deep_space::{coin::Coin, private_key::PrivateKey as CosmosPrivateKey};
use peggy_proto::peggy::query_client::QueryClient as PeggyQueryClient;
use peggy_utils::{
    error::PeggyError,
    types::{
        SendToCosmosEvent, SendToMinterEvent, TransactionBatchExecutedEvent, ValsetUpdatedEvent,
    },
};
use std::ops::Sub;
use tonic::transport::Channel;
use web30::client::Web3;
use web30::jsonrpc::error::Web3Error;

pub async fn check_for_events(
    web3: &Web3,
    contact: &Contact,
    grpc_client: &mut PeggyQueryClient<Channel>,
    peggy_contract_address: EthAddress,
    our_private_key: CosmosPrivateKey,
    fee: Coin,
    starting_block: Uint256,
) -> Result<Uint256, PeggyError> {
    let our_cosmos_address = our_private_key.to_public_key().unwrap().to_address();
    let latest_block = web3.eth_block_number().await?.sub(5u64.into());

    let deposits = web3
        .check_for_events(
            starting_block.clone(),
            Some(latest_block.clone()),
            vec![peggy_contract_address],
            vec!["SendToHubEvent(address,address,bytes32,uint256,uint256)"],
        )
        .await;
    trace!("Deposits {:?}", deposits);

    let transfers = web3
        .check_for_events(
            starting_block.clone(),
            Some(latest_block.clone()),
            vec![peggy_contract_address],
            vec!["SendToMinterEvent(address,address,bytes32,uint256,uint256)"],
        )
        .await;
    trace!("Transfers {:?}", transfers);

    let batches = web3
        .check_for_events(
            starting_block.clone(),
            Some(latest_block.clone()),
            vec![peggy_contract_address],
            vec!["TransactionBatchExecutedEvent(uint256,address,address,uint256)"],
        )
        .await;
    trace!("Batches {:?}", batches);

    let valsets = web3
        .check_for_events(
            starting_block.clone(),
            Some(latest_block.clone()),
            vec![peggy_contract_address],
            vec!["ValsetUpdatedEvent(uint256,address[],uint256[])"],
        )
        .await;
    trace!("Valsets {:?}", valsets);

    if let (Ok(valsets), Ok(batches), Ok(deposits), Ok(transfers)) =
        (valsets, batches, deposits, transfers)
    {
        let valsets = ValsetUpdatedEvent::from_logs(&valsets)?;
        trace!("parsed valsets {:?}", valsets);
        let withdraws = TransactionBatchExecutedEvent::from_logs(&batches)?;
        trace!("parsed batches {:?}", batches);
        let deposits = SendToCosmosEvent::from_logs(&deposits)?;
        trace!("parsed deposits {:?}", deposits);

        let transfers = SendToMinterEvent::from_logs(&transfers)?;
        trace!("parsed deposits {:?}", deposits);

        // note that starting block overlaps with our last checked block, because we have to deal with
        // the possibility that the relayer was killed after relaying only one of multiple events in a single
        // block, so we also need this routine so make sure we don't send in the first event in this hypothetical
        // multi event block again. In theory we only send all events for every block and that will pass of fail
        // atomicly but lets not take that risk.
        let last_event_nonce = get_last_event_nonce(grpc_client, our_cosmos_address).await?;
        let deposits = SendToCosmosEvent::filter_by_event_nonce(last_event_nonce, &deposits);
        let transfers = SendToMinterEvent::filter_by_event_nonce(last_event_nonce, &transfers);
        let withdraws =
            TransactionBatchExecutedEvent::filter_by_event_nonce(last_event_nonce, &withdraws);

        if !deposits.is_empty() {
            info!(
                "Oracle observed deposit with sender {}, destination {}, amount {}, and event nonce {}",
                deposits[0].sender, deposits[0].destination, deposits[0].amount, deposits[0].event_nonce
            )
        }

        if !transfers.is_empty() {
            info!(
                "Oracle observed transfer with sender {}, destination {}, amount {}, and event nonce {}",
                transfers[0].sender, transfers[0].destination, transfers[0].amount, transfers[0].event_nonce
            )
        }

        if !withdraws.is_empty() {
            info!(
                "Oracle observed batch with nonce {}, contract {}, and event nonce {}",
                withdraws[0].batch_nonce, withdraws[0].erc20, withdraws[0].event_nonce
            )
        }

        if !deposits.is_empty() || !withdraws.is_empty() || !transfers.is_empty() {
            let _res = send_ethereum_claims(
                contact,
                our_private_key,
                deposits,
                withdraws,
                transfers,
                fee,
            )
            .await?;
            let new_event_nonce = get_last_event_nonce(grpc_client, our_cosmos_address).await?;
            // since we can't actually trust that the above txresponse is correct we have to check here
            // we may be able to trust the tx response post grpc
            if new_event_nonce == last_event_nonce {
                return Err(PeggyError::InvalidBridgeStateError(
                    "Claims did not process, trying again in a moment".to_string(),
                ));
            }
        }
        Ok(latest_block)
    } else {
        error!("Failed to get events");
        Err(PeggyError::EthereumRestError(Web3Error::BadResponse(
            "Failed to get logs!".to_string(),
        )))
    }
}
