package enterprise_test

import (
	simapp "github.com/unification-com/mainchain/app"
	"testing"

	tmproto "github.com/cometbft/cometbft/proto/tendermint/types"
	"github.com/stretchr/testify/require"

	authtypes "github.com/cosmos/cosmos-sdk/x/auth/types"
	"github.com/unification-com/mainchain/x/enterprise/types"
)

func TestItCreatesModuleAccountOnInitBlock(t *testing.T) {
	app := simapp.Setup(t, false)
	ctx := app.BaseApp.NewContext(false, tmproto.Header{})

	//app.InitChain(
	//	abcitypes.RequestInitChain{
	//		AppStateBytes: []byte("{}"),
	//		ChainId:       "test-chain-id",
	//	},
	//)

	acc := app.AccountKeeper.GetAccount(ctx, authtypes.NewModuleAddress(types.ModuleName))
	require.NotNil(t, acc)
}
