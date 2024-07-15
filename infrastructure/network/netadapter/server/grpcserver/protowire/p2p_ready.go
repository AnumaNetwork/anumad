package protowire

import (
	"github.com/AnumaNetwork/anumad/app/appmessage"
	"github.com/pkg/errors"
)

func (x *AnumadMessage_Ready) toAppMessage() (appmessage.Message, error) {
	if x == nil {
		return nil, errors.Wrapf(errorNil, "AnumadMessage_Ready is nil")
	}
	return &appmessage.MsgReady{}, nil
}

func (x *AnumadMessage_Ready) fromAppMessage(_ *appmessage.MsgReady) error {
	return nil
}
