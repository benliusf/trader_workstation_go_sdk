package main

import (
	"context"
	"errors"
	"fmt"
	"os"
	"time"

	"github.com/benliusf/trader_workstation_go_sdk/examples"
	"github.com/benliusf/trader_workstation_go_sdk/pkg/client"
	"github.com/benliusf/trader_workstation_go_sdk/pkg/log"

	api "github.com/benliusf/trader_workstation_go_sdk/api/v104401"
)

// An example to demonstrate an API call to Place Pending Order -
//	https://www.interactivebrokers.com/campus/ibkr-api-page/twsapi-doc/#place-order
//
// An order in pending status won't be fulfilled unless "transmitted".
//
// This example creates an order to buy "AAPL" stock and cancels it.

// Implement a custom handler to capture response for account summary.
type thisHandler struct {
	examples.ExampleHandler

	accountSummary *api.AccountSummary
}

func newThisHandler(logger log.Logger) *thisHandler {
	return &thisHandler{
		ExampleHandler: examples.ExampleHandler{
			Logger: logger,
		},
	}
}

func (h *thisHandler) AccountSummary(m *api.AccountSummary) error {
	h.accountSummary = m
	return nil
}

func main() {
	// Write permissions are required for placing orders.
	rw := client.ReadAndWrite()
	conf := client.TWSConfig{
		ClientId:     0,
		Host:         "localhost",
		Port:         "7497",
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 5 * time.Second,
		Privileges:   rw,
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
	defer done()

	handler := newThisHandler(logger)
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

	// This is for testing purposes only.
	// We request account data to verify it's paper trading.
	accountSummary := client.NewAccountSummaryRequest(writer, "", []client.AccountSummaryTag{
		client.AccountType,
	})
	for {
		if _, err := accountSummary.Send(ctx); err != nil {
			if errors.Is(err, client.ErrClientNotReady) {
				logger.Warn("client not ready, retrying")
				time.Sleep(1 * time.Second)
				continue
			}
			panic(err)
		}
		break
	}

	// Wait for account data response and assert it's a paper trading account.
	accountSummaryCh := make(chan *api.AccountSummary)
	go func() {
		for {
			if handler.accountSummary != nil {
				accountSummaryCh <- handler.accountSummary
				return
			}
		}
	}()
	select {
	case <-time.After(10 * time.Second):
		logger.Error("timed out waiting for account summary")
		os.Exit(1)
	case a := <-accountSummaryCh:
		if !client.IsPaperTrading(a.GetAccount()) {
			logger.Error("account is not paper trading")
			os.Exit(1)
		}
		logger.Info("account_id=%v", a.GetAccount())
	}

	// Place open order to buy "AAPL" stock.
	// The order will be pending because SetTransmit() is not called.
	contract := client.NewContractBuilder().
		SetSymbol("AAPL").
		SetSecType(client.STOCK).
		SetExchange(client.SMART).
		SetPrimaryExch(client.NASDAQ).Build()
	order := client.NewMarketOrderBuilder().
		SetAction(client.BUY).
		SetQuantity(10).
		SetTimeInForce(client.GTC).
		//SetTransmit().
		Build()
	placeOrder := client.NewPlaceOrderRequest(writer, contract, order)
	orderId, err := placeOrder.Send(ctx)
	if err != nil {
		panic(err)
	}

	time.Sleep(5 * time.Second)

	cancelOrder := client.NewCancelOrderRequest(writer, orderId)
	if err := cancelOrder.Send(ctx); err != nil {
		panic(err)
	}
	logger.Info("canceled order_id=%v", orderId)

	done()
}
