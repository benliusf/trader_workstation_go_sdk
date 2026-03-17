package main

import (
	"time"

	"github.com/benliusf/trader_workstation_go_sdk/examples"
	"github.com/benliusf/trader_workstation_go_sdk/pkg/client"
)

// An example to test the client connection to TWS localhost:7496
// You should see the following output in stdout -
//		2026/03/17 13:36:49 [INFO] client=0 connected to localhost:7496 ver=222 ts=2026-03-17 06:36:49 -0700 PDT
//		2026/03/17 13:36:49 [INFO] client=0 disconnected

func main() {
	conf := client.TWSConfig{
		ClientID:     0,
		Host:         "localhost",
		Port:         "7497",
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 5 * time.Second,
	}
	cl, err := client.NewClient(conf, examples.NewExampleLogger())
	if err != nil {
		panic(err)
	}
	defer cl.Disconnect()
	if err := cl.Connect(); err != nil {
		panic(err)
	}
}
