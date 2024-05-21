package keeper_test

import (
	"github.com/cometbft/cometbft/crypto/ed25519"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"math/rand"

	simapp "github.com/unification-com/mainchain/app"
	"github.com/unification-com/mainchain/x/enterprise/types"
)

func createRandomAccounts(accNum int) []sdk.AccAddress {
	testAddrs := make([]sdk.AccAddress, accNum)
	for i := 0; i < accNum; i++ {
		pk := ed25519.GenPrivKey().PubKey()
		testAddrs[i] = sdk.AccAddress(pk.Address())
	}

	return testAddrs
}

func ParamsEqual(paramsA, paramsB types.Params) bool {
	return paramsA == paramsB
}

func LockedUndEqual(lA, lB types.LockedUnd) bool {
	return lA == lB
}

func RandomDecision() types.PurchaseOrderStatus {
	rnd := rand.Intn(100)
	if rnd >= 50 {
		return types.StatusAccepted
	}
	return types.StatusRejected
}

func RandomStatus() types.PurchaseOrderStatus {
	rnd := simapp.RandInBetween(1, 5)
	switch rnd {
	case 1:
		return types.StatusRaised
	case 2:
		return types.StatusAccepted
	case 3:
		return types.StatusRejected
	case 4:
		return types.StatusCompleted
	default:
		return types.StatusRaised
	}
}

func AddressInDecisions(addr sdk.AccAddress, decisions types.PurchaseOrderDecisions) bool {
	for _, d := range decisions {
		if d.Signer == addr.String() {
			return true
		}
	}
	return false
}
