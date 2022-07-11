package keeper_test

import (
	"fmt"
	"github.com/cosmos/cosmos-sdk/types/query"
	"math/rand"
	"testing"

	tmproto "github.com/tendermint/tendermint/proto/tendermint/types"
	"github.com/unification-com/mainchain/app/test_helpers"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/stretchr/testify/require"
	"github.com/unification-com/mainchain/x/enterprise/types"
)

func TestSetGetTotalLockedUnd(t *testing.T) {
	app := test_helpers.Setup(false)
	ctx := app.BaseApp.NewContext(false, tmproto.Header{})
	test_helpers.SetKeeperTestParamsAndDefaultValues(app, ctx)

	denom := test_helpers.TestDenomination
	amount := int64(1000)
	locked := sdk.NewInt64Coin(denom, amount)

	err := app.EnterpriseKeeper.SetTotalLockedUnd(ctx, locked)
	require.NoError(t, err)

	lockedDb := app.EnterpriseKeeper.GetTotalLockedUnd(ctx)

	require.True(t, lockedDb.IsEqual(locked))
	require.True(t, lockedDb.Denom == denom)
	require.True(t, lockedDb.Amount.Int64() == amount)
}

func TestGetTotalUnlocked(t *testing.T) {
	app := test_helpers.Setup(false)
	ctx := app.BaseApp.NewContext(false, tmproto.Header{})
	test_helpers.SetKeeperTestParamsAndDefaultValues(app, ctx)
	test_helpers.AddTestAddrs(app, ctx, 1, sdk.NewInt(20000))

	denom := test_helpers.TestDenomination
	amount := int64(1000)
	locked := sdk.NewInt64Coin(denom, amount)

	err := app.EnterpriseKeeper.SetTotalLockedUnd(ctx, locked)
	require.NoError(t, err)

	totUnlocked := app.EnterpriseKeeper.GetTotalUnLockedUnd(ctx)
	totalSupply := app.BankKeeper.GetSupply(ctx, denom)

	diff := totalSupply.Sub(totUnlocked)

	require.Equal(t, locked, diff)
}

func TestGetTotalUndSupply(t *testing.T) {
	app := test_helpers.Setup(false)
	ctx := app.BaseApp.NewContext(false, tmproto.Header{})
	test_helpers.SetKeeperTestParamsAndDefaultValues(app, ctx)
	test_helpers.AddTestAddrs(app, ctx, 1, sdk.NewInt(20000))

	totalSupply := app.BankKeeper.GetSupply(ctx, test_helpers.TestDenomination)
	totalSupplyFromEnt := app.EnterpriseKeeper.GetTotalUndSupply(ctx)
	require.Equal(t, totalSupply, totalSupplyFromEnt)
}

func TestSetGetLockedUndForAccount(t *testing.T) {
	app := test_helpers.Setup(false)
	ctx := app.BaseApp.NewContext(false, tmproto.Header{})
	test_helpers.SetKeeperTestParamsAndDefaultValues(app, ctx)

	testAddresses := test_helpers.GenerateRandomTestAccounts(100)

	for _, addr := range testAddresses {
		amount := int64(rand.Intn(10000) + 1)
		denom := test_helpers.TestDenomination

		locked := types.LockedUnd{
			Owner:  addr.String(),
			Amount: sdk.NewInt64Coin(denom, amount),
		}

		err := app.EnterpriseKeeper.SetLockedUndForAccount(ctx, locked)
		require.NoError(t, err)

		lockedDb := app.EnterpriseKeeper.GetLockedUndForAccount(ctx, addr)

		require.True(t, locked.Owner == lockedDb.Owner)
		require.True(t, lockedDb.Amount.IsEqual(locked.Amount))

		lockedDbAmount := app.EnterpriseKeeper.GetLockedUndAmountForAccount(ctx, addr)
		require.True(t, lockedDbAmount.IsEqual(locked.Amount))
	}
}

func (suite *KeeperTestSuite) TestIsLocked() {
	app, ctx, addrs := suite.app, suite.ctx, suite.addrs

	denom := test_helpers.TestDenomination

	var (
		l    types.LockedUnd
		addr sdk.AccAddress
	)

	testCases := []struct {
		msg         string
		malleate    func()
		expIsLocked bool
	}{
		{
			"zero value",
			func() {
				addr = addrs[0]
				l = types.LockedUnd{
					Owner:  addr.String(),
					Amount: sdk.NewInt64Coin(denom, 0),
				}
			},
			false,
		},
		{
			"valid value",
			func() {
				addr = addrs[2]
				l = types.LockedUnd{
					Owner:  addr.String(),
					Amount: sdk.NewInt64Coin(denom, 100),
				}
			},
			true,
		},
	}

	for _, testCase := range testCases {
		suite.Run(fmt.Sprintf("Case %s", testCase.msg), func() {
			testCase.malleate()

			err := app.EnterpriseKeeper.SetLockedUndForAccount(ctx, l)
			suite.Require().NoError(err)

			isLocked := app.EnterpriseKeeper.IsLocked(ctx, addr)

			suite.Require().Equal(testCase.expIsLocked, isLocked)
		})
	}
}

func TestMintCoinsAndLock(t *testing.T) {
	app := test_helpers.Setup(false)
	ctx := app.BaseApp.NewContext(false, tmproto.Header{})
	test_helpers.SetKeeperTestParamsAndDefaultValues(app, ctx)

	totalAmount := int64(0)

	testAddresses := test_helpers.GenerateRandomTestAccounts(100)

	for _, addr := range testAddresses {
		amount := int64(rand.Intn(10000) + 1)
		totalAmount = totalAmount + amount

		toMint := sdk.NewInt64Coin(test_helpers.TestDenomination, amount)

		err := app.EnterpriseKeeper.MintCoinsAndLock(ctx, addr, toMint)
		require.NoError(t, err)

		isLocked := app.EnterpriseKeeper.IsLocked(ctx, addr)
		require.True(t, isLocked)

		lockedDb := app.EnterpriseKeeper.GetLockedUndForAccount(ctx, addr)
		require.True(t, lockedDb.Amount.IsEqual(toMint))
	}

	totalLocked := sdk.NewInt64Coin(test_helpers.TestDenomination, totalAmount)
	totalLockedCoins := sdk.NewCoins(totalLocked)

	totalLockedDb := app.EnterpriseKeeper.GetTotalLockedUnd(ctx)
	require.True(t, totalLockedDb.IsEqual(totalLocked))

	totalSupplyDb := app.EnterpriseKeeper.GetEnterpriseSupplyIncludingLockedUnd(ctx)
	require.True(t, totalSupplyDb.Locked == totalLocked.Amount.Uint64())

	entAccount := app.EnterpriseKeeper.GetEnterpriseAccount(ctx)
	entAccountCoins := app.BankKeeper.GetAllBalances(ctx, entAccount.GetAddress())
	require.True(t, entAccountCoins.IsEqual(totalLockedCoins))

	entAccFromAccK := app.AccountKeeper.GetModuleAccount(ctx, types.ModuleName)
	entAccFromSkCoins := app.BankKeeper.GetAllBalances(ctx, entAccFromAccK.GetAddress())
	require.True(t, entAccFromSkCoins.IsEqual(totalLockedCoins))
}

func TestUnlockCoinsForFees(t *testing.T) {
	app := test_helpers.Setup(false)
	ctx := app.BaseApp.NewContext(false, tmproto.Header{})
	test_helpers.SetKeeperTestParamsAndDefaultValues(app, ctx)

	totalAmount := int64(0)

	testAddresses := test_helpers.GenerateRandomTestAccounts(100)

	for _, addr := range testAddresses {
		amountToMint := int64(test_helpers.RandInBetween(1000, 100000))
		amountToUnlock := int64(test_helpers.RandInBetween(1, 999))
		totalAmount = totalAmount + amountToMint

		toMint := sdk.NewInt64Coin(test_helpers.TestDenomination, amountToMint)
		toUnlock := sdk.NewInt64Coin(test_helpers.TestDenomination, amountToUnlock)
		toUnlockCoins := sdk.NewCoins(toUnlock)

		_ = app.EnterpriseKeeper.MintCoinsAndLock(ctx, addr, toMint)

		err := app.EnterpriseKeeper.UnlockCoinsForFees(ctx, addr, toUnlockCoins)
		require.NoError(t, err)

		totalAmount = totalAmount - amountToUnlock

		expectedLocked := toMint.Sub(toUnlock)

		lockedDb := app.EnterpriseKeeper.GetLockedUndForAccount(ctx, addr)
		require.True(t, lockedDb.Amount.IsEqual(expectedLocked))
	}

	totalLocked := sdk.NewInt64Coin(test_helpers.TestDenomination, totalAmount)
	totalLockedCoins := sdk.NewCoins(totalLocked)

	totalLockedDb := app.EnterpriseKeeper.GetTotalLockedUnd(ctx)
	require.True(t, totalLockedDb.IsEqual(totalLocked))

	totalSupplyDb := app.EnterpriseKeeper.GetEnterpriseSupplyIncludingLockedUnd(ctx)
	require.True(t, totalSupplyDb.Locked == totalLocked.Amount.Uint64())

	entAccount := app.EnterpriseKeeper.GetEnterpriseAccount(ctx)
	entAccountCoins := app.BankKeeper.GetAllBalances(ctx, entAccount.GetAddress())
	require.True(t, entAccountCoins.IsEqual(totalLockedCoins))

	entAccFromAccK := app.AccountKeeper.GetModuleAccount(ctx, types.ModuleName)
	entAccFromSkCoins := app.BankKeeper.GetAllBalances(ctx, entAccFromAccK.GetAddress())
	require.True(t, entAccFromSkCoins.IsEqual(totalLockedCoins))
}

func TestGetTotalSupplyWithLockedNundRemoved(t *testing.T) {
	app := test_helpers.Setup(false)
	ctx := app.BaseApp.NewContext(false, tmproto.Header{})
	test_helpers.SetKeeperTestParamsAndDefaultValues(app, ctx)

	totalSupply := sdk.NewCoins(sdk.NewInt64Coin(test_helpers.TestDenomination, 0))
	totalMinted := sdk.NewInt64Coin(test_helpers.TestDenomination, 0)

	testAddresses := test_helpers.GenerateRandomTestAccounts(100)

	for _, addr := range testAddresses {
		amountToMint := int64(test_helpers.RandInBetween(1000, 100000))
		amountToUnlock := int64(test_helpers.RandInBetween(1, 999))

		toMint := sdk.NewInt64Coin(test_helpers.TestDenomination, amountToMint)
		toUnlock := sdk.NewInt64Coin(test_helpers.TestDenomination, amountToUnlock)
		toUnlockCoins := sdk.NewCoins(toUnlock)

		_ = app.EnterpriseKeeper.MintCoinsAndLock(ctx, addr, toMint)

		err := app.EnterpriseKeeper.UnlockCoinsForFees(ctx, addr, toUnlockCoins)
		require.NoError(t, err)

		totalSupply = totalSupply.Add(toUnlock)
		totalMinted = totalMinted.Add(toMint)

		pageReq := &query.PageRequest{
			Limit: 10,
		}
		totalSupplyDb, _, _ := app.EnterpriseKeeper.GetTotalSupplyWithLockedNundRemoved(ctx, pageReq)
		require.True(t, totalSupplyDb.IsEqual(totalSupply))

		totalMintedDb := app.BankKeeper.GetSupply(ctx, test_helpers.TestDenomination)
		require.True(t, totalMintedDb.IsEqual(totalMinted))
	}
}

func TestUnlockCoinsForFeesAndUsedCounter(t *testing.T) {
	app := test_helpers.Setup(false)
	ctx := app.BaseApp.NewContext(false, tmproto.Header{})
	test_helpers.SetKeeperTestParamsAndDefaultValues(app, ctx)

	totalUsed := int64(0)

	testAddresses := test_helpers.GenerateRandomTestAccounts(100)

	for _, addr := range testAddresses {
		amountToMint := int64(test_helpers.RandInBetween(1000, 100000))
		amountToUnlock := int64(test_helpers.RandInBetween(1, 999))

		toMint := sdk.NewInt64Coin(test_helpers.TestDenomination, amountToMint)
		toUnlock := sdk.NewInt64Coin(test_helpers.TestDenomination, amountToUnlock)
		toUnlockCoins := sdk.NewCoins(toUnlock)
		totalUsed = totalUsed + amountToUnlock

		_ = app.EnterpriseKeeper.MintCoinsAndLock(ctx, addr, toMint)

		err := app.EnterpriseKeeper.UnlockCoinsForFees(ctx, addr, toUnlockCoins)
		require.NoError(t, err)

		usedDb := app.EnterpriseKeeper.GetUsedUndForAccount(ctx, addr)
		require.True(t, usedDb.IsEqual(toUnlock))
	}

	expectedTotalUsedCoin := sdk.NewInt64Coin(test_helpers.TestDenomination, totalUsed)

	totalUsedDb := app.EnterpriseKeeper.GetTotalUsedUnd(ctx)
	require.True(t, totalUsedDb.IsEqual(expectedTotalUsedCoin))
}

func TestUnlockCoinsForFeesAndUsedCounterWithHalfFunds(t *testing.T) {
	app := test_helpers.Setup(false)
	ctx := app.BaseApp.NewContext(false, tmproto.Header{})
	test_helpers.SetKeeperTestParamsAndDefaultValues(app, ctx)

	totalUsed := int64(0)

	testAddresses := test_helpers.AddTestAddrs(app, ctx, 100, sdk.NewInt(10000))

	for _, addr := range testAddresses {
		amountToMint := int64(test_helpers.RandInBetween(1, 999))
		// fee is more than minted, to test using account's normal fund supply in addition to minted efund
		feeToPay := amountToMint * 2
		// only minted will count as "used"
		totalUsed = totalUsed + amountToMint

		toMint := sdk.NewInt64Coin(test_helpers.TestDenomination, amountToMint)
		fee := sdk.NewInt64Coin(test_helpers.TestDenomination, feeToPay)
		feeCoins := sdk.NewCoins(fee)

		_ = app.EnterpriseKeeper.MintCoinsAndLock(ctx, addr, toMint)

		err := app.EnterpriseKeeper.UnlockCoinsForFees(ctx, addr, feeCoins)
		require.NoError(t, err)

		usedDb := app.EnterpriseKeeper.GetUsedUndForAccount(ctx, addr)
		// fee is 2x what was minted. Only minted should count
		require.True(t, usedDb.IsEqual(toMint))
	}

	expectedTotalUsedCoin := sdk.NewInt64Coin(test_helpers.TestDenomination, totalUsed)

	totalUsedDb := app.EnterpriseKeeper.GetTotalUsedUnd(ctx)
	require.True(t, totalUsedDb.IsEqual(expectedTotalUsedCoin))
}
