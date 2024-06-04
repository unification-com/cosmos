package keeper_test

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/unification-com/mainchain/x/stream/types"
)

func (s *KeeperTestSuite) TestParams() {
	testCases := []struct {
		name      string
		input     types.Params
		expectErr bool
	}{
		{
			name: "set full valid params",
			input: types.Params{
				ValidatorFee: sdk.MustNewDecFromStr("0.24"),
			},
			expectErr: false,
		},
		{
			name: "set invalid params > 100%",
			input: types.Params{
				ValidatorFee: sdk.MustNewDecFromStr("1.01"),
			},
			expectErr: true,
		},
		{
			name: "set invalid params negative value",
			input: types.Params{
				ValidatorFee: sdk.MustNewDecFromStr("-0.01"),
			},
			expectErr: true,
		},
		{
			name: "set invalid params nil value",
			input: types.Params{
				ValidatorFee: sdk.Dec{},
			},
			expectErr: true,
		},
	}

	for _, tc := range testCases {
		tc := tc

		s.Run(tc.name, func() {
			expected := s.app.StreamKeeper.GetParams(s.ctx)
			err := s.app.StreamKeeper.SetParams(s.ctx, tc.input)
			if tc.expectErr {
				s.Require().Error(err)
			} else {
				expected = tc.input
				s.Require().NoError(err)
			}

			p := s.app.StreamKeeper.GetParams(s.ctx)
			s.Require().Equal(expected, p)
		})
	}
}
