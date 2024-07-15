package main

import (
	"fmt"
	"reflect"
	"strings"

	"github.com/AnumaNetwork/anumad/infrastructure/network/netadapter/server/grpcserver/protowire"
)

var commandTypes = []reflect.Type{
	reflect.TypeOf(protowire.AnumadMessage_AddPeerRequest{}),
	reflect.TypeOf(protowire.AnumadMessage_GetConnectedPeerInfoRequest{}),
	reflect.TypeOf(protowire.AnumadMessage_GetPeerAddressesRequest{}),
	reflect.TypeOf(protowire.AnumadMessage_GetCurrentNetworkRequest{}),
	reflect.TypeOf(protowire.AnumadMessage_GetInfoRequest{}),

	reflect.TypeOf(protowire.AnumadMessage_GetBlockRequest{}),
	reflect.TypeOf(protowire.AnumadMessage_GetBlocksRequest{}),
	reflect.TypeOf(protowire.AnumadMessage_GetHeadersRequest{}),
	reflect.TypeOf(protowire.AnumadMessage_GetBlockCountRequest{}),
	reflect.TypeOf(protowire.AnumadMessage_GetBlockDagInfoRequest{}),
	reflect.TypeOf(protowire.AnumadMessage_GetSelectedTipHashRequest{}),
	reflect.TypeOf(protowire.AnumadMessage_GetVirtualSelectedParentBlueScoreRequest{}),
	reflect.TypeOf(protowire.AnumadMessage_GetVirtualSelectedParentChainFromBlockRequest{}),
	reflect.TypeOf(protowire.AnumadMessage_ResolveFinalityConflictRequest{}),
	reflect.TypeOf(protowire.AnumadMessage_EstimateNetworkHashesPerSecondRequest{}),

	reflect.TypeOf(protowire.AnumadMessage_GetBlockTemplateRequest{}),
	reflect.TypeOf(protowire.AnumadMessage_SubmitBlockRequest{}),

	reflect.TypeOf(protowire.AnumadMessage_GetMempoolEntryRequest{}),
	reflect.TypeOf(protowire.AnumadMessage_GetMempoolEntriesRequest{}),
	reflect.TypeOf(protowire.AnumadMessage_GetMempoolEntriesByAddressesRequest{}),

	reflect.TypeOf(protowire.AnumadMessage_SubmitTransactionRequest{}),

	reflect.TypeOf(protowire.AnumadMessage_GetUtxosByAddressesRequest{}),
	reflect.TypeOf(protowire.AnumadMessage_GetBalanceByAddressRequest{}),
	reflect.TypeOf(protowire.AnumadMessage_GetCoinSupplyRequest{}),

	reflect.TypeOf(protowire.AnumadMessage_BanRequest{}),
	reflect.TypeOf(protowire.AnumadMessage_UnbanRequest{}),
}

type commandDescription struct {
	name       string
	parameters []*parameterDescription
	typeof     reflect.Type
}

type parameterDescription struct {
	name   string
	typeof reflect.Type
}

func commandDescriptions() []*commandDescription {
	commandDescriptions := make([]*commandDescription, len(commandTypes))

	for i, commandTypeWrapped := range commandTypes {
		commandType := unwrapCommandType(commandTypeWrapped)

		name := strings.TrimSuffix(commandType.Name(), "RequestMessage")
		numFields := commandType.NumField()

		var parameters []*parameterDescription
		for i := 0; i < numFields; i++ {
			field := commandType.Field(i)

			if !isFieldExported(field) {
				continue
			}

			parameters = append(parameters, &parameterDescription{
				name:   field.Name,
				typeof: field.Type,
			})
		}
		commandDescriptions[i] = &commandDescription{
			name:       name,
			parameters: parameters,
			typeof:     commandTypeWrapped,
		}
	}

	return commandDescriptions
}

func (cd *commandDescription) help() string {
	sb := &strings.Builder{}
	sb.WriteString(cd.name)
	for _, parameter := range cd.parameters {
		_, _ = fmt.Fprintf(sb, " [%s]", parameter.name)
	}
	return sb.String()
}
