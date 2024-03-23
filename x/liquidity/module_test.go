package liquidity_test

import (
	"testing"

	abcitypes "github.com/cometbft/cometbft/abci/types"
	tmproto "github.com/cometbft/cometbft/proto/tendermint/types"
	"github.com/stretchr/testify/require"

	lapp "github.com/Victor118/liquidity/app"
)

func TestItCreatesModuleAccountOnInitBlock(t *testing.T) {
	app := lapp.Setup(t, false)
	ctx := app.BaseApp.NewContext(false, tmproto.Header{})

	app.InitChain(
		abcitypes.RequestInitChain{
			AppStateBytes: []byte("{}"),
			ChainId:       "test-chain-id",
		},
	)
	params := app.LiquidityKeeper.GetParams(ctx)
	require.NotNil(t, params)
}
