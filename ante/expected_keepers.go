package ante

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
)

// BankKeeper is a modified version of the default found in Cosmos SDK x/auth/types/expected_keepers.go
// and adds the GetAllBalances and SpendableCoins methods required by the beacon and wrkchain modules' ante handlers
type BankKeeper interface {
	SendCoinsFromAccountToModule(ctx sdk.Context, senderAddr sdk.AccAddress, recipientModule string, amt sdk.Coins) error
	GetAllBalances(ctx sdk.Context, address sdk.AccAddress) sdk.Coins
	SpendableCoins(ctx sdk.Context, address sdk.AccAddress) sdk.Coins
}
