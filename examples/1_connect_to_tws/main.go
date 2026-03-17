package main

import (
	"time"

	"github.com/benliusf/trader_workstation_go_sdk/examples"
	"github.com/benliusf/trader_workstation_go_sdk/pkg/client"
)

func main() {
	conf := client.TWSConfig{
		ClientID:     0,
		Host:         "localhost",
		Port:         "7496",
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
