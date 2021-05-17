package keeper

import (
	"fmt"
	"github.com/MinterTeam/mhub/chain/cold_storage"
	oraclekeeper "github.com/MinterTeam/mhub/chain/x/oracle/keeper"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"math"
	"sort"
	"strconv"

	"github.com/MinterTeam/mhub/chain/x/minter/types"
	"github.com/cosmos/cosmos-sdk/codec"
	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
	paramtypes "github.com/cosmos/cosmos-sdk/x/params/types"
	"github.com/tendermint/tendermint/libs/log"
)

var defaultSender = sdk.AccAddress{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0}

// Keeper maintains the link to storage and exposes getter/setter methods for the various parts of the state machine
type Keeper struct {
	StakingKeeper types.StakingKeeper

	storeKey   sdk.StoreKey // Unexposed key to access store from sdk.Context
	paramSpace paramtypes.Subspace

	cdc          codec.BinaryMarshaler // The wire codec for binary encoding/decoding.
	bankKeeper   types.BankKeeper
	oracleKeeper oraclekeeper.Keeper

	AttestationHandler interface {
		Handle(sdk.Context, types.Attestation, types.MinterClaim) error
		SetPeggyKeeper(keeper types.PeggyKeeper)
	}
}

// NewKeeper returns a new instance of the peggy keeper
func NewKeeper(cdc codec.BinaryMarshaler, storeKey sdk.StoreKey, paramSpace paramtypes.Subspace, stakingKeeper types.StakingKeeper, bankKeeper types.BankKeeper, oracleKeeper oraclekeeper.Keeper) Keeper {
	k := Keeper{
		cdc:           cdc,
		paramSpace:    paramSpace,
		storeKey:      storeKey,
		StakingKeeper: stakingKeeper,
		bankKeeper:    bankKeeper,
		oracleKeeper:  oracleKeeper,
	}
	k.AttestationHandler = &AttestationHandler{
		keeper:     k,
		bankKeeper: bankKeeper,
	}

	// set KeyTable if it has not already been set
	if !paramSpace.HasKeyTable() {
		paramSpace = paramSpace.WithKeyTable(types.ParamKeyTable())
	}

	return k
}

/////////////////////////////
//     VALSET REQUESTS     //
/////////////////////////////

// SetValsetRequest returns a new instance of the Peggy BridgeValidatorSet
// i.e. {"nonce": 1, "memebers": [{"eth_addr": "foo", "power": 11223}]}
func (k Keeper) SetValsetRequest(ctx sdk.Context) *types.Valset {
	valset := k.GetCurrentValset(ctx)
	if len(valset.Members) == 0 {
		return nil
	}

	valset.MinterNonce = k.autoIncrementID(ctx, types.MinterNonce)
	k.storeValset(ctx, valset)

	event := sdk.NewEvent(
		types.EventTypeMultisigUpdateRequest,
		sdk.NewAttribute(sdk.AttributeKeyModule, types.ModuleName),
		sdk.NewAttribute(types.AttributeKeyContract, k.GetBridgeContractAddress(ctx)),
		sdk.NewAttribute(types.AttributeKeyBridgeChainID, strconv.Itoa(int(k.GetBridgeChainID(ctx)))),
		sdk.NewAttribute(types.AttributeKeyMultisigID, fmt.Sprint(valset.Nonce)),
		sdk.NewAttribute(types.AttributeKeyNonce, fmt.Sprint(valset.Nonce)),
	)
	ctx.EventManager().EmitEvent(event)
	return valset
}

func (k Keeper) storeValset(ctx sdk.Context, valset *types.Valset) {
	store := ctx.KVStore(k.storeKey)
	store.Set(types.GetValsetRequestKey(valset.Nonce), k.cdc.MustMarshalBinaryBare(valset))
}

func (k Keeper) deleteValset(ctx sdk.Context, valset *types.Valset) {
	store := ctx.KVStore(k.storeKey)

	// delete confirms
	confirms := k.GetAllValsetConfirmsByNonce(ctx, valset.Nonce)
	for _, valsetConf := range confirms {
		addr, _ := sdk.AccAddressFromBech32(valsetConf.Validator)
		store.Delete(types.GetValsetConfirmKey(valsetConf.Nonce, addr))
	}

	// delete valset
	store.Delete(types.GetValsetRequestKey(valset.Nonce))
}

// HasValsetRequest returns true if a valset defined by a nonce exists
func (k Keeper) HasValsetRequest(ctx sdk.Context, nonce uint64) bool {
	store := ctx.KVStore(k.storeKey)
	return store.Has(types.GetValsetRequestKey(nonce))
}

// GetValsetRequest returns a valset by nonce
func (k Keeper) GetValsetRequest(ctx sdk.Context, nonce uint64) *types.Valset {
	store := ctx.KVStore(k.storeKey)
	bz := store.Get(types.GetValsetRequestKey(nonce))

	if bz == nil {
		return nil
	}
	var valset types.Valset
	k.cdc.MustUnmarshalBinaryBare(bz, &valset)
	return &valset
}

// IterateValsetRequest retruns all valsetRequests
func (k Keeper) IterateValsetRequest(ctx sdk.Context, cb func(key []byte, val *types.Valset) bool) {
	prefixStore := prefix.NewStore(ctx.KVStore(k.storeKey), types.ValsetRequestKey)
	iter := prefixStore.ReverseIterator(nil, nil)
	defer iter.Close()
	for ; iter.Valid(); iter.Next() {
		var valset types.Valset
		k.cdc.MustUnmarshalBinaryBare(iter.Value(), &valset)
		// cb returns true to stop early
		if cb(iter.Key(), &valset) {
			break
		}
	}
}

/////////////////////////////
//     VALSET CONFIRMS     //
/////////////////////////////

// GetValsetConfirm returns a valset confirmation by a nonce and validator address
func (k Keeper) GetValsetConfirm(ctx sdk.Context, nonce uint64, validator sdk.AccAddress) *types.MsgValsetConfirm {
	store := ctx.KVStore(k.storeKey)
	entity := store.Get(types.GetValsetConfirmKey(nonce, validator))
	if entity == nil {
		return nil
	}
	confirm := types.MsgValsetConfirm{}
	k.cdc.MustUnmarshalBinaryBare(entity, &confirm)
	return &confirm
}

// SetValsetConfirm sets a valset confirmation
func (k Keeper) SetValsetConfirm(ctx sdk.Context, valsetConf types.MsgValsetConfirm) []byte {
	store := ctx.KVStore(k.storeKey)
	addr, err := sdk.AccAddressFromBech32(valsetConf.Validator)
	if err != nil {
		panic(err)
	}
	key := types.GetValsetConfirmKey(valsetConf.Nonce, addr)
	store.Set(key, k.cdc.MustMarshalBinaryBare(&valsetConf))
	return key
}

// GetAllValsetConfirmsByNonce returns all validator set confirmations by nonce
func (k Keeper) GetAllValsetConfirmsByNonce(ctx sdk.Context, nonce uint64) (confirms []*types.MsgValsetConfirm) {
	prefixStore := prefix.NewStore(ctx.KVStore(k.storeKey), types.ValsetConfirmKey)
	start, end := prefixRange(types.UInt64Bytes(nonce))
	iterator := prefixStore.Iterator(start, end)

	defer iterator.Close()

	for ; iterator.Valid(); iterator.Next() {
		confirm := types.MsgValsetConfirm{}
		k.cdc.MustUnmarshalBinaryBare(iterator.Value(), &confirm)
		confirms = append(confirms, &confirm)
	}

	return confirms
}

// IterateValsetConfirmByNonce iterates through all valset confirms by nonce in ASC order
// MARK finish-batches: this is where the key is iterated in the old (presumed working) code
// TODO: specify which nonce this is
func (k Keeper) IterateValsetConfirmByNonce(ctx sdk.Context, nonce uint64, cb func([]byte, types.MsgValsetConfirm) bool) {
	prefixStore := prefix.NewStore(ctx.KVStore(k.storeKey), types.ValsetConfirmKey)
	iter := prefixStore.Iterator(prefixRange(types.UInt64Bytes(nonce)))
	defer iter.Close()

	for ; iter.Valid(); iter.Next() {
		confirm := types.MsgValsetConfirm{}
		k.cdc.MustUnmarshalBinaryBare(iter.Value(), &confirm)
		// cb returns true to stop early
		if cb(iter.Key(), confirm) {
			break
		}
	}
}

/////////////////////////////
//      BATCH CONFIRMS     //
/////////////////////////////

// GetBatchConfirm returns a batch confirmation given its nonce, the token contract, and a validator address
func (k Keeper) GetBatchConfirm(ctx sdk.Context, nonce uint64, validator sdk.AccAddress) *types.MsgConfirmBatch {
	store := ctx.KVStore(k.storeKey)
	entity := store.Get(types.GetBatchConfirmKey(nonce, validator))
	if entity == nil {
		return nil
	}
	confirm := types.MsgConfirmBatch{}
	k.cdc.MustUnmarshalBinaryBare(entity, &confirm)
	return &confirm
}

// SetBatchConfirm sets a batch confirmation by a validator
func (k Keeper) SetBatchConfirm(ctx sdk.Context, batch *types.MsgConfirmBatch) []byte {
	store := ctx.KVStore(k.storeKey)
	acc, err := sdk.AccAddressFromBech32(batch.Validator)
	if err != nil {
		panic(err)
	}
	key := types.GetBatchConfirmKey(batch.Nonce, acc)
	store.Set(key, k.cdc.MustMarshalBinaryBare(batch))
	return key
}

// IterateBatchConfirmByNonce iterates through all batch confirmations
// MARK finish-batches: this is where the key is iterated in the old (presumed working) code
// TODO: specify which nonce this is
func (k Keeper) IterateBatchConfirmByNonce(ctx sdk.Context, nonce uint64, cb func([]byte, types.MsgConfirmBatch) bool) {
	prefixStore := prefix.NewStore(ctx.KVStore(k.storeKey), types.BatchConfirmKey)
	prefix := types.UInt64Bytes(nonce)
	iter := prefixStore.Iterator(prefixRange(prefix))
	defer iter.Close()

	for ; iter.Valid(); iter.Next() {
		confirm := types.MsgConfirmBatch{}
		k.cdc.MustUnmarshalBinaryBare(iter.Value(), &confirm)
		// cb returns true to stop early
		if cb(iter.Key(), confirm) {
			break
		}
	}
}

/////////////////////////////
//       ETH ADDRESS       //
/////////////////////////////

// SetMinterAddress sets the ethereum address for a given validator
func (k Keeper) SetMinterAddress(ctx sdk.Context, validator sdk.AccAddress, ethAddr string) {
	store := ctx.KVStore(k.storeKey)
	store.Set(types.GetMinterAddressKey(validator), []byte(ethAddr))
}

// GetMinterAddress returns the eth address for a given peggy validator
func (k Keeper) GetMinterAddress(ctx sdk.Context, validator sdk.AccAddress) string {
	store := ctx.KVStore(k.storeKey)
	return string(store.Get(types.GetMinterAddressKey(validator)))
}

// GetCurrentValset gets powers from the store and normalizes them
// into an integer percentage with a resolution of uint32 Max meaning
// a given validators 'Peggy power' is computed as
// Cosmos power / total cosmos power = x / uint32 Max
// where x is the voting power on the Peggy contract. This allows us
// to only use integer division which produces a known rounding error
// from truncation equal to the ratio of the validators
// Cosmos power / total cosmos power ratio, leaving us at uint32 Max - 1
// total voting power. This is an acceptable rounding error since floating
// point may cause consensus problems if different floating point unit
// implementations are involved.
func (k Keeper) GetCurrentValset(ctx sdk.Context) *types.Valset {
	validators := k.StakingKeeper.GetBondedValidatorsByPower(ctx)
	var bridgeValidators []*types.BridgeValidator
	var totalPower uint64
	// TODO someone with in depth info on Cosmos staking should determine
	// if this is doing what I think it's doing
	for _, validator := range validators {
		validatorAddress := validator.GetOperator()
		valAddr := sdk.AccAddress(validatorAddress)

		p := uint64(k.StakingKeeper.GetLastValidatorPower(ctx, validatorAddress))
		totalPower += p

		if minterAddr := k.GetMinterAddress(ctx, valAddr); minterAddr != "" {
			bridgeValidators = append(bridgeValidators, &types.BridgeValidator{
				Power:         p,
				MinterAddress: minterAddr,
			})
		}
	}
	// normalize power values
	for i := range bridgeValidators {
		bridgeValidators[i].Power = sdk.NewUint(bridgeValidators[i].Power).MulUint64(math.MaxUint32).QuoUint64(totalPower).Uint64()
	}
	return types.NewValset(uint64(ctx.BlockHeight()), bridgeValidators)
}

/////////////////////////////
//       PARAMETERS        //
/////////////////////////////

// GetParams returns the parameters from the store
func (k Keeper) GetParams(ctx sdk.Context) (p types.Params) {
	k.paramSpace.GetParamSet(ctx, &p)
	return
}

// setParams sets the parameters in the store
func (k Keeper) setParams(ctx sdk.Context, ps *types.Params) {
	k.paramSpace.SetParamSet(ctx, ps)
}

// GetBridgeContractAddress returns the bridge contract address on ETH
func (k Keeper) GetBridgeContractAddress(ctx sdk.Context) string {
	var a string
	k.paramSpace.Get(ctx, types.ParamsStoreKeyBridgeContractAddress, &a)
	return a
}

// GetBridgeChainID returns the chain id of the ETH chain we are running against
func (k Keeper) GetBridgeChainID(ctx sdk.Context) uint64 {
	var a uint64
	k.paramSpace.Get(ctx, types.ParamsStoreKeyBridgeContractChainID, &a)
	return a
}

// GetPeggyID returns the PeggyID (???)
func (k Keeper) GetPeggyID(ctx sdk.Context) string {
	var a string
	k.paramSpace.Get(ctx, types.ParamsStoreKeyPeggyID, &a)
	return a
}

// Set PeggyID sets the PeggyID (couldn't we reuse the ChainID here?)
func (k Keeper) setPeggyID(ctx sdk.Context, v string) {
	k.paramSpace.Set(ctx, types.ParamsStoreKeyPeggyID, v)
}

// GetStartThreshold returns the start threshold for the peggy validator set
func (k Keeper) GetStartThreshold(ctx sdk.Context) uint64 {
	var a uint64
	k.paramSpace.Get(ctx, types.ParamsStoreKeyStartThreshold, &a)
	return a
}

// setStartThreshold sets the start threshold
func (k Keeper) setStartThreshold(ctx sdk.Context, v uint64) {
	k.paramSpace.Set(ctx, types.ParamsStoreKeyStartThreshold, v)
}

// logger returns a module-specific logger.
func (k Keeper) logger(ctx sdk.Context) log.Logger {
	return ctx.Logger().With("module", fmt.Sprintf("x/%s", types.ModuleName))
}

func (k Keeper) SetPeggyKeeper(keeper types.PeggyKeeper) {
	k.AttestationHandler.SetPeggyKeeper(keeper)
}

func (k Keeper) IsStopped(ctx sdk.Context) bool {
	p := k.GetParams(ctx)
	return p.GetStopped()
}

func (k Keeper) OracleKeeper() oraclekeeper.Keeper {
	return k.oracleKeeper
}

func (k Keeper) GetValsets(ctx sdk.Context) (out []*types.Valset) {
	k.IterateValsets(ctx, func(_ []byte, val *types.Valset) bool {
		out = append(out, val)
		return false
	})
	sort.Sort(types.Valsets(out))
	return
}

// IterateValsets retruns all valsetRequests
func (k Keeper) IterateValsets(ctx sdk.Context, cb func(key []byte, val *types.Valset) bool) {
	prefixStore := prefix.NewStore(ctx.KVStore(k.storeKey), types.ValsetRequestKey)
	iter := prefixStore.ReverseIterator(nil, nil)
	defer iter.Close()
	for ; iter.Valid(); iter.Next() {
		var valset types.Valset
		k.cdc.MustUnmarshalBinaryBare(iter.Value(), &valset)
		// cb returns true to stop early
		if cb(iter.Key(), &valset) {
			break
		}
	}
}

// DeleteValset deletes the valset at a given nonce from state
func (k Keeper) DeleteValset(ctx sdk.Context, nonce uint64) {
	ctx.KVStore(k.storeKey).Delete(types.GetValsetRequestKey(nonce))
}

func (k Keeper) setLastValset(ctx sdk.Context, valset *types.Valset) {
	ctx.KVStore(k.storeKey).Set(types.LastValsetKey, k.cdc.MustMarshalBinaryBare(valset))
}

func (k Keeper) GetLastValset(ctx sdk.Context) *types.Valset {
	entity := ctx.KVStore(k.storeKey).Get(types.LastValsetKey)
	if entity == nil {
		return nil
	}

	valset := types.Valset{}
	k.cdc.MustUnmarshalBinaryBare(entity, &valset)
	return &valset
}

func (k Keeper) ColdStorageTransfer(ctx sdk.Context, c *types.ColdStorageTransferProposal) error {
	coldStorageAddr := k.GetColdStorageAddr(ctx)

	for _, coin := range c.Amount {
		minterCoin, err := types.MinterCoinFromPeggyCoin(coin, ctx, k.oracleKeeper)
		if err != nil {
			return err
		}

		vouchers := sdk.Coins{coin}
		if err := k.bankKeeper.MintCoins(ctx, types.ModuleName, vouchers); err != nil {
			return sdkerrors.Wrapf(err, "mint vouchers coins: %s", vouchers)
		}

		if err := k.bankKeeper.SendCoinsFromModuleToAccount(ctx, types.ModuleName, defaultSender, vouchers); err != nil {
			return sdkerrors.Wrap(err, "transfer vouchers")
		}

		txID, err := k.AddToOutgoingPool(ctx, defaultSender, coldStorageAddr, "", coin)
		if err != nil {
			return err
		}

		if err := k.removeFromUnbatchedTXIndex(ctx, txID); err != nil {
			return sdkerrors.Wrap(err, "fee")
		}

		nextID := k.autoIncrementID(ctx, types.KeyLastOutgoingBatchID)
		minterNonce := k.autoIncrementID(ctx, types.MinterNonce)
		batch := &types.OutgoingTxBatch{
			BatchNonce:  nextID,
			MinterNonce: minterNonce,
			Transactions: []*types.OutgoingTransferTx{
				{
					Id:          txID,
					Sender:      defaultSender.String(),
					DestAddress: coldStorageAddr,
					MinterToken: minterCoin,
					TxHash:      "",
				},
			},
		}
		k.storeBatch(ctx, batch)

		batchEvent := sdk.NewEvent(
			types.EventTypeOutgoingBatch,
			sdk.NewAttribute(sdk.AttributeKeyModule, types.ModuleName),
			sdk.NewAttribute(types.AttributeKeyContract, k.GetBridgeContractAddress(ctx)),
			sdk.NewAttribute(types.AttributeKeyBridgeChainID, strconv.Itoa(int(k.GetBridgeChainID(ctx)))),
			sdk.NewAttribute(types.AttributeKeyOutgoingBatchID, fmt.Sprint(nextID)),
			sdk.NewAttribute(types.AttributeKeyNonce, fmt.Sprint(nextID)),
		)

		ctx.EventManager().EmitEvent(batchEvent)
	}

	return nil
}

func (k Keeper) GetColdStorageAddr(ctx sdk.Context) string {
	return cold_storage.MinterStorage
}

// prefixRange turns a prefix into a (start, end) range. The start is the given prefix value and
// the end is calculated by adding 1 bit to the start value. Nil is not allowed as prefix.
// 		Example: []byte{1, 3, 4} becomes []byte{1, 3, 5}
// 				 []byte{15, 42, 255, 255} becomes []byte{15, 43, 0, 0}
//
// In case of an overflow the end is set to nil.
//		Example: []byte{255, 255, 255, 255} becomes nil
// MARK finish-batches: this is where some crazy shit happens
func prefixRange(prefix []byte) ([]byte, []byte) {
	if prefix == nil {
		panic("nil key not allowed")
	}
	// special case: no prefix is whole range
	if len(prefix) == 0 {
		return nil, nil
	}

	// copy the prefix and update last byte
	end := make([]byte, len(prefix))
	copy(end, prefix)
	l := len(end) - 1
	end[l]++

	// wait, what if that overflowed?....
	for end[l] == 0 && l > 0 {
		l--
		end[l]++
	}

	// okay, funny guy, you gave us FFF, no end to this range...
	if l == 0 && end[0] == 0 {
		end = nil
	}
	return prefix, end
}
