package keeper_test

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/unification-com/mainchain/x/stream/types"
)

func (s *KeeperTestSuite) TestParamsQuery() {
	defaultFee := sdk.NewDecWithPrec(1, 2)
	newFee := sdk.NewDecWithPrec(24, 2)

	req1 := &types.QueryParamsRequest{}
	expRes1 := &types.QueryParamsResponse{Params: types.DefaultParams()}

	res1, err1 := s.app.StreamKeeper.Params(s.ctx, req1)

	s.Require().NoError(err1)
	s.Require().Equal(expRes1, res1)

	req2 := &types.QueryParamsRequest{}
	expRes2 := &types.QueryParamsResponse{Params: types.Params{ValidatorFee: defaultFee}}

	res2, err2 := s.app.StreamKeeper.Params(s.ctx, req2)

	s.Require().NoError(err2)
	s.Require().Equal(expRes2, res2)

	_ = s.app.StreamKeeper.SetParams(s.ctx, types.NewParams(newFee))
	req3 := &types.QueryParamsRequest{}
	expRes3 := &types.QueryParamsResponse{Params: types.Params{ValidatorFee: newFee}}

	res3, err3 := s.app.StreamKeeper.Params(s.ctx, req3)

	s.Require().NoError(err3)
	s.Require().Equal(expRes3, res3)
}
