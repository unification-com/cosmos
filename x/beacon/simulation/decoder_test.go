package simulation_test

import (
	"fmt"
	simapp "github.com/unification-com/mainchain/app"
	"testing"
	"time"

	"github.com/stretchr/testify/require"

	"github.com/cometbft/cometbft/crypto/ed25519"
	"github.com/cosmos/cosmos-sdk/types/kv"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/unification-com/mainchain/x/beacon/simulation"
	"github.com/unification-com/mainchain/x/beacon/types"
)

var (
	bPk1   = ed25519.GenPrivKey().PubKey()
	bAddr1 = sdk.AccAddress(bPk1.Address())
)

func TestDecodeStore(t *testing.T) {
	testApp := simapp.Setup(t, false)
	cdc := testApp.AppCodec()
	dec := simulation.NewDecodeStore(cdc)

	beacon, err := types.NewBeacon(1, "beacon1", "Test BEACON 1", 0, bAddr1.String())
	require.NoError(t, err)

	beaconTs, err := types.NewBeaconTimestamp(1, uint64(time.Now().Unix()), "arbitraryblockhashvalue")
	require.NoError(t, err)

	beaconBz, err := cdc.Marshal(&beacon)
	require.NoError(t, err)

	beaconTsBz, err := cdc.Marshal(&beaconTs)
	require.NoError(t, err)

	kvPairs := kv.Pairs{
		Pairs: []kv.Pair{
			{Key: types.BeaconKey(1), Value: beaconBz},
			{Key: types.BeaconTimestampKey(1, 1), Value: beaconTsBz},
			{Key: []byte{0x99}, Value: []byte{0x99}},
		},
	}

	tests := []struct {
		name        string
		expectedLog string
	}{
		{"beacon", fmt.Sprintf("%v\n%v", beacon, beacon)},
		{"beacon timestamp", fmt.Sprintf("%v\n%v", beaconTs, beaconTs)},
		{"other", ""},
	}

	for i, tt := range tests {
		i, tt := i, tt
		t.Run(tt.name, func(t *testing.T) {
			switch i {
			case len(tests) - 1:
				require.Panics(t, func() { dec(kvPairs.Pairs[i], kvPairs.Pairs[i]) }, tt.name)
			default:
				require.Equal(t, tt.expectedLog, dec(kvPairs.Pairs[i], kvPairs.Pairs[i]), tt.name)
			}
		})
	}
}
