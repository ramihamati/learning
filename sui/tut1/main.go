package main

import (
	"fmt"
	"github.com/block-vision/sui-go-sdk/constant"
	"github.com/block-vision/sui-go-sdk/sui"
)

//TIP <p>To run your code, right-click the code and select <b>Run</b>.</p> <p>Alternatively, click
// the <icon src="AllIcons.Actions.Execute"/> icon in the gutter and select the <b>Run</b> menu item from here.</p>

func main() {
	// configure your endpoint here or use BlockVision's free Sui RPC endpoint
	cli := sui.NewSuiClient("https://sui-testnet-endpoint.blockvision.org")
	faucetHost, err := sui.GetFaucetHost(constant.SuiDevnet)
	if err != nil {
		fmt.Println("GetFaucetHost err:", err)
		return
	}
	fmt.Println("faucetHost:", faucetHost)
	recipient := "0xaf9f4d20c205f26051a7e1758601c4c47b9f99df3f9823f70926c17c80882d36"

	header := map[string]string{}
	err = sui.RequestSuiFromFaucet(faucetHost, recipient, header)
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	// the successful transaction block url: https://suiexplorer.com/txblock/91moaxbXsQnJYScLP2LpbMXV43ZfngS2xnRgj1CT7jLQ?network=devnet
	fmt.Println("Request DevNet Sui From Faucet success")

	fmt.Println("Hello and welcome, %s!", cli)

}
