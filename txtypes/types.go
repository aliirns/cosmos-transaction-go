package txtypes

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	banktypes "github.com/cosmos/cosmos-sdk/x/bank/types"
	"github.com/cosmos/ibc-go/v5/modules/apps/transfer/types"
	clienttypes "github.com/cosmos/ibc-go/v5/modules/core/02-client/types"
)

func NewIBCTransferTx(sender string, receiver string, coin sdk.Coin, channelID string, timeoutHeight clienttypes.Height) (msg sdk.Msg) {
	msg = types.NewMsgTransfer("transfer", channelID, coin, sender, receiver, timeoutHeight, 0)
	return msg
}

func NewSendTx(sender string, reciever string, coin sdk.Coins) (sdk.Msg, error) {
	fromAddr, err := sdk.AccAddressFromBech32(sender)
	if err != nil {
		return nil, err
	}
	toAddr, err := sdk.AccAddressFromBech32(reciever)
	if err != nil {
		return nil, err
	}

	return banktypes.NewMsgSend(fromAddr, toAddr, coin), nil
}
