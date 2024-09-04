package protowire

import (
	"github.com/AnumaNetwork/anumad/app/appmessage"
	"github.com/pkg/errors"
)

func (x *AnumadMessage_Pong) toAppMessage() (appmessage.Message, error) {
	if x == nil {
		return nil, errors.Wrapf(errorNil, "AnumadMessage_Pong is nil")
	}
	return x.Pong.toAppMessage()
}

func (x *PongMessage) toAppMessage() (appmessage.Message, error) {
	if x == nil {
		return nil, errors.Wrapf(errorNil, "PongMessage is nil")
	}
	return &appmessage.MsgPong{
		Nonce: x.Nonce,
	}, nil
}

func (x *AnumadMessage_Pong) fromAppMessage(msgPong *appmessage.MsgPong) error {
	x.Pong = &PongMessage{
		Nonce: msgPong.Nonce,
	}
	return nil
}
