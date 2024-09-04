package protowire

import (
	"github.com/AnumaNetwork/anumad/app/appmessage"
	"github.com/pkg/errors"
)

func (x *AnumadMessage_Verack) toAppMessage() (appmessage.Message, error) {
	if x == nil {
		return nil, errors.Wrapf(errorNil, "AnumadMessage_Verack is nil")
	}
	return &appmessage.MsgVerAck{}, nil
}

func (x *AnumadMessage_Verack) fromAppMessage(_ *appmessage.MsgVerAck) error {
	return nil
}
