package main

import (
	"context"
	"errors"
	"fmt"
	"time"

	"github.com/benliusf/trader_workstation_go_sdk/examples"
	"github.com/benliusf/trader_workstation_go_sdk/pkg/client"
)

// An example to demonstrate read and write functionality by making a request for the NextValidId -
//	https://www.interactivebrokers.com/campus/ibkr-api-page/twsapi-doc/#next-valid-id
//
// The following response will be printed to stdout via the test handler -
//		2026/04/01 15:51:45 [INFO] received next valid id: 10

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

	// Use test handler from example_handler.go to print all the messages onto stdout.
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
	// You must explicitly call this to inform the server that this connection serves API calls.
	if err := writer.StartAPI(10 * time.Second); err != nil {
		panic(err)
	}

	// Run the reader asynchronously to continuously listen for new messages.
	// Every message is passed to the handler for processing.
	go func() {
		if err := reader.Read(ctx, handler); err != nil {
			logger.Error(fmt.Sprintf("read error: %v", err))
		}
	}()

	// Make and send request for NextValidId.
	// The send() is inside a retry loop until the server is ready to accept API calls.
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
