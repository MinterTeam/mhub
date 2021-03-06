package types

const (
	EventTypeObservation              = "eth_observation"
	EventTypeOutgoingBatch            = "eth_outgoing_batch"
	EventTypeOutgoingBatchExecuted    = "eth_outgoing_batch_executed"
	EventTypeMultisigBootstrap        = "eth_multisig_bootstrap"
	EventTypeMultisigUpdateRequest    = "eth_multisig_update_request"
	EventTypeOutgoingBatchCanceled    = "eth_outgoing_batch_canceled"
	EventTypeBridgeWithdrawalReceived = "eth_withdrawal_received"
	EventTypeBridgeDepositReceived    = "eth_deposit_received"
	EventTypeRefund                   = "eth_refund"

	AttributeKeyAttestationID    = "attestation_id"
	AttributeKeyAttestationIDs   = "attestation_ids"
	AttributeKeyBatchConfirmKey  = "batch_confirm_key"
	AttributeKeyValsetConfirmKey = "valset_confirm_key"
	AttributeKeyMultisigID       = "multisig_id"
	AttributeKeyOutgoingBatchID  = "batch_id"
	AttributeKeyOutgoingTXID     = "outgoing_tx_id"
	AttributeKeyAttestationType  = "attestation_type"
	AttributeKeyContract         = "bridge_contract"
	AttributeKeyNonce            = "nonce"
	AttributeKeyValsetNonce      = "valset_nonce"
	AttributeKeyBatchNonce       = "batch_nonce"
	AttributeKeyBridgeChainID    = "bridge_chain_id"
	AttributeKeySetOperatorAddr  = "set_operator_address"
	AttributeKeyTxHash           = "tx_hash"
	AttributeKeyBatchTxHash      = "batch_tx_hash"
)
