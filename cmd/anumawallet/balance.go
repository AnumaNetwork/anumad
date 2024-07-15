package main

import (
	"context"
	"fmt"

	"github.com/AnumaNetwork/anumad/cmd/anumawallet/daemon/client"
	"github.com/AnumaNetwork/anumad/cmd/anumawallet/daemon/pb"
	"github.com/AnumaNetwork/anumad/cmd/anumawallet/utils"
)

func balance(conf *balanceConfig) error {
	daemonClient, tearDown, err := client.Connect(conf.DaemonAddress)
	if err != nil {
		return err
	}
	defer tearDown()

	ctx, cancel := context.WithTimeout(context.Background(), daemonTimeout)
	defer cancel()
	response, err := daemonClient.GetBalance(ctx, &pb.GetBalanceRequest{})
	if err != nil {
		return err
	}

	pendingSuffix := ""
	if response.Pending > 0 {
		pendingSuffix = " (pending)"
	}
	if conf.Verbose {
		pendingSuffix = ""
		println("Address                                                                       Available             Pending")
		println("-----------------------------------------------------------------------------------------------------------")
		for _, addressBalance := range response.AddressBalances {
			fmt.Printf("%s %s %s\n", addressBalance.Address, utils.FormatAnum(addressBalance.Available), utils.FormatAnum(addressBalance.Pending))
		}
		println("-----------------------------------------------------------------------------------------------------------")
		print("                                                 ")
	}
	fmt.Printf("Total balance, ANUM %s %s%s\n", utils.FormatAnum(response.Available), utils.FormatAnum(response.Pending), pendingSuffix)

	return nil
}
