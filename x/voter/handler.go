package voter

import (
	"fmt"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/victoryeo/voter/x/voter/keeper"
	"github.com/victoryeo/voter/x/voter/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

// NewHandler ...
func NewHandler(k keeper.Keeper) sdk.Handler {
	return func(ctx sdk.Context, msg sdk.Msg) (*sdk.Result, error) {
		ctx = ctx.WithEventManager(sdk.NewEventManager())
		switch msg := msg.(type) {
    // this line is used by starport scaffolding # 1
		case types.MsgCreatePoll:
			return handleMsgCreatePoll(ctx, k, msg)
		case types.MsgSetPoll:
			return handleMsgSetPoll(ctx, k, msg)
		case types.MsgDeletePoll:
			return handleMsgDeletePoll(ctx, k, msg)
		default:
			errMsg := fmt.Sprintf("unrecognized %s message type: %T", types.ModuleName, msg)
			return nil, sdkerrors.Wrap(sdkerrors.ErrUnknownRequest, errMsg)
		}
	}
}
