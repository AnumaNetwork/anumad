package protowire

import (
	"github.com/AnumaNetwork/anumad/app/appmessage"
	"github.com/pkg/errors"
)

func (x *AnumadMessage_RequestPruningPointAndItsAnticone) toAppMessage() (appmessage.Message, error) {
	if x == nil {
		return nil, errors.Wrapf(errorNil, "AnumadMessage_RequestPruningPointAndItsAnticone is nil")
	}
	return &appmessage.MsgRequestPruningPointAndItsAnticone{}, nil
}

func (x *AnumadMessage_RequestPruningPointAndItsAnticone) fromAppMessage(_ *appmessage.MsgRequestPruningPointAndItsAnticone) error {
	return nil
}
