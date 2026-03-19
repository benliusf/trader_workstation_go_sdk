package main

import (
	"context"
	"errors"
	"fmt"
	"time"

	"github.com/benliusf/trader_workstation_go_sdk/examples"
	"github.com/benliusf/trader_workstation_go_sdk/pkg/client"
)

// This example demonstrates the read and write functionality of the SDK to interact with TWS API.
// All responses are printed to stdout from the test handler defined in example_handler.go
// The code calls the writer to make a single request for the NextValidId as documented -
//	https://www.interactivebrokers.com/campus/ibkr-api-page/twsapi-doc/#next-valid-id

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

	// Use test handler from example_handler.go to print all the messages into stdout.
	handler := examples.NewExampleHandler(logger)

	// Create the reader to pull messages from client connection.
	reader, err := client.NewReader(cl)
	if err != nil {
		panic(err)
	}

	// Create the client writer to send API requests.
	writer, err := client.NewSender(cl)
	if err != nil {
		panic(err)
	}
	// You must explicitly send this request to inform the server that this connection serves API calls.
	if err := writer.StartAPI(); err != nil {
		panic(err)
	}

	// Run the reader asynchronously to continuously listen for new messages.
	// Each received message is passed to the handler for processing.
	go func() {
		if err := reader.Read(ctx, handler); err != nil {
			logger.Error(fmt.Sprintf("read error: %v", err))
		}
		time.Sleep(1 * time.Second)
	}()

	// Make a request to the server for the NextValidId.
	// The send() is inside a loop to retry until the server is ready to accept API calls.
	nextValidId := client.NewNextValidIdRequest(writer)
	for {
		if err := nextValidId.Send(ctx); err != nil {
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

	// Cancel the context to stop the reader from receiving messages.
	done()
}
