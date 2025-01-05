package keeper_test

import (
	"testing"

	"github.com/stretchr/testify/require"

	keepertest "legal/testutil/keeper"
	"legal/x/legal/types"
)

func TestGetParams(t *testing.T) {
	k, ctx := keepertest.LegalKeeper(t)
	params := types.DefaultParams()

	require.NoError(t, k.SetParams(ctx, params))
	require.EqualValues(t, params, k.GetParams(ctx))
}
