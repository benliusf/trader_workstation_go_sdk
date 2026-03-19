package main

import (
	"context"
	"errors"
	"fmt"
	"time"

	"github.com/benliusf/trader_workstation_go_sdk/examples"
	"github.com/benliusf/trader_workstation_go_sdk/pkg/client"
)

// An example to demonstrate a API call to Account Summary -
//	https://www.interactivebrokers.com/campus/ibkr-api-page/twsapi-doc/#sync-account-summary
//
// Every write request needs a unique reqId. The code below calls cl.GetNextReqId()
// to get the next reqId before making the API call.

func main() {
	conf := client.TWSConfig{
		ClientID:     0,
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
	if err := writer.StartAPI(); err != nil {
		panic(err)
	}

	go func() {
		if err := reader.Read(ctx, handler); err != nil {
			logger.Error(fmt.Sprintf("read error: %v", err))
		}
		time.Sleep(1 * time.Second)
	}()

	// Use the TWSClient to get the next valid request id which is used to send write requests.
	reqId := cl.GetNextReqId()

	// Create new request using the reqId.
	accountSummary := client.NewAccountSummaryRequest(writer, reqId, "", []client.AccountSummaryTag{})
	for {
		if err := accountSummary.Send(ctx); err != nil {
			if errors.Is(err, client.ErrClientNotReady) {
				logger.Warn("client not ready, retrying")
				time.Sleep(1 * time.Second)
				continue
			}
			panic(err)
		}
		break
	}

	time.Sleep(5 * time.Second)
	done()
}
