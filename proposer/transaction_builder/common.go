package builder

import (
	"context"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/taikoxyz/taiko-client/pkg/rpc"
)

// getParentMetaHash returns the meta hash of the parent block of the latest proposed block in protocol.
func getParentMetaHash(ctx context.Context, rpc *rpc.Client) (common.Hash, error) {
	state, err := rpc.TaikoL1.State(&bind.CallOpts{Context: ctx})
	if err != nil {
		return common.Hash{}, err
	}

	parent, err := rpc.TaikoL1.GetBlock(&bind.CallOpts{Context: ctx}, state.SlotB.NumBlocks-1)
	if err != nil {
		return common.Hash{}, err
	}

	return parent.Blk.MetaHash, nil
}