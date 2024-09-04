package rpchandlers

import (
	"github.com/AnumaNetwork/anumad/app/appmessage"
	"github.com/AnumaNetwork/anumad/app/rpc/rpccontext"
	"github.com/AnumaNetwork/anumad/infrastructure/network/netadapter/router"
)

// HandleGetCurrentNetwork handles the respectively named RPC command
func HandleGetCurrentNetwork(context *rpccontext.Context, _ *router.Router, _ appmessage.Message) (appmessage.Message, error) {
	response := appmessage.NewGetCurrentNetworkResponseMessage(context.Config.ActiveNetParams.Net.String())
	return response, nil
}
