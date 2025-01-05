package legal_test

import (
	"testing"

	keepertest "legal/testutil/keeper"
	"legal/testutil/nullify"
	legal "legal/x/legal/module"
	"legal/x/legal/types"

	"github.com/stretchr/testify/require"
)

func TestGenesis(t *testing.T) {
	genesisState := types.GenesisState{
		Params: types.DefaultParams(),

		// this line is used by starport scaffolding # genesis/test/state
	}

	k, ctx := keepertest.LegalKeeper(t)
	legal.InitGenesis(ctx, k, genesisState)
	got := legal.ExportGenesis(ctx, k)
	require.NotNil(t, got)

	nullify.Fill(&genesisState)
	nullify.Fill(got)

	// this line is used by starport scaffolding # genesis/test/assert
}
