package main

import (
	"context"
	"errors"
	"fmt"
	"time"

	"github.com/benliusf/trader_workstation_go_sdk/examples"
	"github.com/benliusf/trader_workstation_go_sdk/pkg/client"
)

// An example to demonstrate an API call to Delayed Market Data -
//
//	https://www.interactivebrokers.com/campus/ibkr-api-page/twsapi-doc/#delayed-market-data
func main() {
	conf := client.TWSConfig{
		ClientId:     0,
		Host:         "localhost",
		Port:         "7497",
		ReadTimeout:  2 * time.Second,
		WriteTimeout: 2 * time.Second,
	}
	logger := examples.NewExampleLogger()
	cl, err := client.NewClient(conf, logger)
	if err != nil {
		panic(err)
	}
	defer cl.Disconnect()
	if err := cl.Connect(); err != nil {
		panic(err)
	}
	ctx, done := context.WithCancel(context.Background())

	handler := examples.NewExampleHandler(logger)
	reader, err := client.NewReader(cl)
	if err != nil {
		panic(err)
	}
	writer, err := client.NewSender(cl)
	if err != nil {
		panic(err)
	}
	if err := writer.StartAPI(10 * time.Second); err != nil {
		panic(err)
	}

	go func() {
		if err := reader.Read(ctx, handler); err != nil {
			logger.Error(fmt.Sprintf("read error: %v", err))
		}
		time.Sleep(1 * time.Second)
	}()

	// Tell the server that we're requesting DELAYED data -
	//	https://www.interactivebrokers.com/campus/ibkr-api-page/twsapi-doc/#delayed-market-data
	marketLevel := client.NewMarketDataTypeRequest(writer, client.MARKET_DATA_DELAYED)

	// Request for `GOOG` stock symbol
	googTicker := client.NewMarketDataRequest(writer,
		client.NewContractBuilder().
			SetSymbol("GOOG").
			SetSecType(client.STOCK).
			SetExchange(client.SMART).
			SetPrimaryExch(client.NASDAQ).Build())
	// Request for `NVDA` stock symbol
	nvdaTicker := client.NewMarketDataRequest(writer,
		client.NewContractBuilder().
			SetSymbol("NVDA").
			SetSecType(client.STOCK).
			SetExchange(client.SMART).
			SetPrimaryExch(client.NASDAQ).Build())
	for {
		if err := marketLevel.Send(ctx); err != nil {
			if errors.Is(err, client.ErrClientNotReady) {
				logger.Warn("client not ready, retrying")
				time.Sleep(1 * time.Second)
				continue
			}
			panic(err)
		}
		if _, err := googTicker.Send(ctx); err != nil {
			panic(err)
		}
		if _, err := nvdaTicker.Send(ctx); err != nil {
			panic(err)
		}
		break
	}

	time.Sleep(5 * time.Second)
	done()
}
