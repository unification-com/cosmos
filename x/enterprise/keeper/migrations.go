package keeper

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/unification-com/mainchain/x/enterprise/exported"
	v3 "github.com/unification-com/mainchain/x/enterprise/migrations/v3"
)

// Migrator is a struct for handling in-place store migrations.
type Migrator struct {
	keeper         Keeper
	legacySubspace exported.Subspace
}

// NewMigrator returns a new Migrator.
func NewMigrator(keeper Keeper, ss exported.Subspace) Migrator {
	return Migrator{
		keeper:         keeper,
		legacySubspace: ss,
	}
}

// Migrate1to2 migrates from version 1 to 2.
//func (m Migrator) Migrate1to2(ctx sdk.Context) error {
//	return v045.MigrateStore(ctx, m.keeper.storeKey, m.keeper.paramSpace, m.keeper.cdc)
//}

// Migrate2to3 migrates the x/beacon module state from the consensus version 2 to
// version 3. Specifically, it takes the parameters that are currently stored
// and managed by the x/params modules and stores them directly into the x/mint
// module state.
func (m Migrator) Migrate2to3(ctx sdk.Context) error {
	return v3.Migrate(ctx, ctx.KVStore(m.keeper.storeKey), m.legacySubspace, m.keeper.cdc)
}
