package protowire

import (
	"github.com/AnumaNetwork/anumad/app/appmessage"
	"github.com/pkg/errors"
)

func (x *AnumadMessage_DonePruningPointUtxoSetChunks) toAppMessage() (appmessage.Message, error) {
	if x == nil {
		return nil, errors.Wrapf(errorNil, "AnumadMessage_DonePruningPointUtxoSetChunks is nil")
	}
	return &appmessage.MsgDonePruningPointUTXOSetChunks{}, nil
}

func (x *AnumadMessage_DonePruningPointUtxoSetChunks) fromAppMessage(_ *appmessage.MsgDonePruningPointUTXOSetChunks) error {
	x.DonePruningPointUtxoSetChunks = &DonePruningPointUtxoSetChunksMessage{}
	return nil
}
