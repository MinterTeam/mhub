package types

import (
	"encoding/binary"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"math"
	"sort"
	"strconv"
)

// UInt64FromBytes create uint from binary big endian representation
func UInt64FromBytes(s []byte) uint64 {
	return binary.BigEndian.Uint64(s)
}

// UInt64Bytes uses the SDK byte marshaling to encode a uint64
func UInt64Bytes(n uint64) []byte {
	return sdk.Uint64ToBigEndian(n)
}

// UInt64FromString to parse out a uint64 for a nonce
func UInt64FromString(s string) (uint64, error) {
	return strconv.ParseUint(s, 10, 64)
}

//////////////////////////////////////
//      BRIDGE VALIDATOR(S)         //
//////////////////////////////////////

// ValidateBasic performs stateless checks on validity
func (b *BridgeValidator) ValidateBasic() error {
	if b.Power == 0 {
		return sdkerrors.Wrap(ErrEmpty, "power")
	}
	if err := ValidateMinterAddress(b.MinterAddress); err != nil {
		return sdkerrors.Wrap(err, "minter address")
	}
	if b.MinterAddress == "" {
		return sdkerrors.Wrap(ErrEmpty, "address")
	}
	return nil
}

// BridgeValidators is the sorted set of validator data for Ethereum bridge MultiSig set
type BridgeValidators []*BridgeValidator

// Sort sorts the validators by power
func (b BridgeValidators) Sort() {
	sort.Slice(b, func(i, j int) bool {
		if b[i].Power == b[j].Power {
			// Secondary sort on eth address in case powers are equal
			return MinterAddrLessThan(b[i].MinterAddress, b[j].MinterAddress)
		}
		return b[i].Power > b[j].Power
	})
}

// HasDuplicates returns true if there are duplicates in the set
func (b BridgeValidators) HasDuplicates() bool {
	m := make(map[string]struct{}, len(b))
	for i := range b {
		m[b[i].MinterAddress] = struct{}{}
	}
	return len(m) != len(b)
}

// GetPowers returns only the power values for all members
func (b BridgeValidators) GetPowers() []uint64 {
	r := make([]uint64, len(b))
	for i := range b {
		r[i] = b[i].Power
	}
	return r
}

// ValidateBasic performs stateless checks
func (b BridgeValidators) ValidateBasic() error {
	// TODO: check if the set is sorted here?
	if len(b) == 0 {
		return ErrEmpty
	}
	for i := range b {
		if err := b[i].ValidateBasic(); err != nil {
			return sdkerrors.Wrapf(err, "member %d", i)
		}
	}
	if b.HasDuplicates() {
		return sdkerrors.Wrap(ErrDuplicate, "addresses")
	}
	return nil
}

// PowerDiff returns the difference in power between two bridge validator sets
// TODO: this needs to be potentially refactored
func (b BridgeValidators) PowerDiff(c BridgeValidators) float64 {
	powers := map[string]int64{}
	var totalB int64
	// loop over b and initialize the map with their powers
	for _, bv := range b {
		powers[bv.MinterAddress] = int64(bv.Power)
		totalB += int64(bv.Power)
	}

	// subtract c powers from powers in the map, initializing
	// uninitialized keys with negative numbers
	for _, bv := range c {
		if val, ok := powers[bv.MinterAddress]; ok {
			powers[bv.MinterAddress] = val - int64(bv.Power)
		} else {
			powers[bv.MinterAddress] = -int64(bv.Power)
		}
	}

	var delta float64
	for _, v := range powers {
		// NOTE: we care about the absolute value of the changes
		delta += math.Abs(float64(v))
	}

	return math.Abs(delta / float64(totalB))
}

// TotalPower returns the total power in the bridge validator set
func (b BridgeValidators) TotalPower() (out uint64) {
	for _, v := range b {
		out += v.Power
	}
	return
}

// NewValset returns a new valset
func NewValset(nonce uint64, members BridgeValidators) *Valset {
	members.Sort()
	var mem []*BridgeValidator
	for _, val := range members {
		mem = append(mem, val)
	}
	return &Valset{Nonce: uint64(nonce), Members: mem}
}

// WithoutEmptyMembers returns a new Valset without member that have 0 power or an empty Ethereum address.
func (v *Valset) WithoutEmptyMembers() *Valset {
	if v == nil {
		return nil
	}
	r := Valset{Nonce: v.Nonce, Members: make([]*BridgeValidator, 0, len(v.Members))}
	for i := range v.Members {
		if err := v.Members[i].ValidateBasic(); err == nil {
			r.Members = append(r.Members, v.Members[i])
		}
	}
	return &r
}

// Valsets is a collection of valset
type Valsets []*Valset

func (v Valsets) Len() int {
	return len(v)
}

func (v Valsets) Less(i, j int) bool {
	return v[i].Nonce > v[j].Nonce
}

func (v Valsets) Swap(i, j int) {
	v[i], v[j] = v[j], v[i]
}
