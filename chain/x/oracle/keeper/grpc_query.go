package keeper

import (
	"context"

	"github.com/MinterTeam/mhub/chain/x/oracle/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

const gweiInEth = 1e9

var _ types.QueryServer = Keeper{}

func (k Keeper) Coins(context context.Context, request *types.QueryCoinsRequest) (*types.QueryCoinsResponse, error) {
	ctx := sdk.UnwrapSDKContext(context)

	return &types.QueryCoinsResponse{Coins: k.GetCoins(ctx).List()}, nil
}

func (k Keeper) CurrentEpoch(context context.Context, request *types.QueryCurrentEpochRequest) (*types.QueryCurrentEpochResponse, error) {
	ctx := sdk.UnwrapSDKContext(context)

	currentEpoch := types.Epoch{
		Nonce: k.GetCurrentEpoch(ctx),
	}

	att := k.GetAttestation(ctx, currentEpoch.Nonce, &types.MsgPriceClaim{})
	votes := att.GetVotes()
	for _, valaddr := range votes {
		validator, _ := sdk.ValAddressFromBech32(valaddr)
		oracle := sdk.AccAddress(validator).String()
		priceClaim := k.GetClaim(ctx, oracle, currentEpoch.Nonce).(*types.GenericClaim).GetPriceClaim()
		currentEpoch.Votes = append(currentEpoch.Votes, &types.Vote{
			Oracle: oracle,
			Claim:  priceClaim,
		})
	}

	return &types.QueryCurrentEpochResponse{Epoch: &currentEpoch}, nil
}

func (k Keeper) EthFee(context context.Context, request *types.QueryEthFeeRequest) (*types.QueryEthFeeResponse, error) {
	ctx := sdk.UnwrapSDKContext(context)

	gasPrice, err := k.GetEthGasPrice(ctx)
	if err != nil {
		return nil, sdkerrors.Wrap(err, "gas price")
	}

	ethPrice, err := k.GetEthPrice(ctx)
	if err != nil {
		return nil, sdkerrors.Wrap(err, "eth price")
	}

	return &types.QueryEthFeeResponse{
		Min:  gasPrice.Mul(ethPrice).MulRaw(int64(k.GetMinSingleWithdrawGas(ctx))).QuoRaw(gweiInEth),
		Fast: gasPrice.Mul(ethPrice).MulRaw(int64(k.GetMinBatchGas(ctx))).QuoRaw(gweiInEth),
	}, nil
}
