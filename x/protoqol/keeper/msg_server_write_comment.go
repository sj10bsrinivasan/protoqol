package keeper

import (
	"context"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"protoqol/x/protoqol/types"
)

func (k msgServer) WriteComment(goCtx context.Context, msg *types.MsgWriteComment) (*types.MsgWriteCommentResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	// TODO: Handling the message
	_ = ctx

	return &types.MsgWriteCommentResponse{}, nil
}
