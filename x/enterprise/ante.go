package enterprise

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/cosmos/cosmos-sdk/x/auth"
	"github.com/unification-com/mainchain-cosmos/x/wrkchain"
)

var (
	_ FeeTx = (*auth.StdTx)(nil) // assert StdTx implements FeeTx
)

// FeeTx defines the interface to be implemented by Tx to use the FeeDecorators
type FeeTx interface {
	sdk.Tx
	GetGas() uint64
	GetFee() sdk.Coins
	FeePayer() sdk.AccAddress
}

type CheckLockedUndDecorator struct {
	ak   auth.AccountKeeper
	entk Keeper
}

func NewCheckLockedUndDecorator(ak auth.AccountKeeper, entk Keeper) CheckLockedUndDecorator {
	return CheckLockedUndDecorator{
		ak:  ak,
		entk: entk,
	}
}

func (ld CheckLockedUndDecorator) AnteHandle(ctx sdk.Context, tx sdk.Tx, simulate bool, next sdk.AnteHandler) (sdk.Context, error) {
	feeTx, ok := tx.(FeeTx)

	if !ok {
		return ctx, sdkerrors.Wrap(sdkerrors.ErrTxDecode, "Tx must be a FeeTx")
	}

	if checkIsWrkChainTx(feeTx) {
		// no need to check locked UND. Continue
		return next(ctx, tx, simulate)
	}

	if !ld.entk.IsLocked(ctx, feeTx.FeePayer()) {
		// no locked UND - continue
		return next(ctx, tx, simulate)
	}

	// todo - check tx (value + fees) < (spendable coins - locked und)
	// for message types - Send, create validator, delegate
	// "getTxValue(msg)" function to get Amount/Value from msg and return as coins
	return ctx, sdkerrors.Wrap(sdkerrors.ErrInsufficientCoins, "Locked UND can only be used for paying WRKChain or BEACON fees")

	//return next(ctx, tx, simulate)
}

func checkIsWrkChainTx(tx FeeTx) bool {
	msgs := tx.GetMsgs()
	for _, msg := range msgs {
		switch msg.(type) {
		case wrkchain.MsgRegisterWrkChain:
			return true
		case wrkchain.MsgRecordWrkChainBlock:
			return true
		}
	}
	return false
}
