package main

import (
	"context"
	"time"

	"github.com/benliusf/trader_workstation_go_sdk/examples"
	"github.com/benliusf/trader_workstation_go_sdk/pkg/client"
	"github.com/benliusf/trader_workstation_go_sdk/pkg/simple"
	"google.golang.org/protobuf/encoding/protojson"
)

func main() {
	conf := client.TWSConfig{
		Host:         "localhost",
		Port:         "7497",
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 5 * time.Second,
		ClientId:     0,
		Privileges:   client.ReadAndWrite(),
	}
	logger := examples.NewExampleLogger()
	cl, err := simple.NewClient(conf, logger)
	if err != nil {
		panic(err)
	}
	defer cl.Disconnect()
	if err := cl.Connect(10 * time.Second); err != nil {
		panic(err)
	}
	ctx := context.Background()

	accountSummary, err := cl.GetAccountSummary(ctx, "", client.AllTags)
	if err != nil {
		panic(err)
	}
	for _, m := range accountSummary {
		tmp, _ := protojson.Marshal(m)
		logger.Info("received account summary data: %s", tmp)
	}

	contractData, err := cl.GetContractData(ctx, client.NewContractBuilder().
		SetSymbol("AAPL").
		SetSecType(client.STOCK).
		SetExchange(client.SMART).
		SetPrimaryExch(client.NASDAQ).Build())
	if err != nil {
		panic(err)
	}
	tmp, _ := protojson.Marshal(contractData)
	logger.Info("received contract data: %s", tmp)
}
