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

// An example to demonstrate an API call to Place Open Order -
//	https://www.interactivebrokers.com/campus/ibkr-api-page/twsapi-doc/#place-order
//
// An "open order" is a order in pending status that won't be fulfilled unless "transmitted".
//
// This example creates an open order to buy "AAPL" stock and cancels it.

// Implement a custom handler to capture the response for the order.
type thisHandler struct {
	examples.ExampleHandler

	accountSummary *api.AccountSummary
	orderStatus    *api.OrderStatus
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

func (h *thisHandler) OrderStatus(m *api.OrderStatus) error {
	h.orderStatus = m
	return nil
}

func main() {
	// Write permissions are required for placing orders.
	rw := client.ReadAndWrite()
	conf := client.TWSConfig{
		ClientId:     0,
		Host:         "localhost",
		Port:         "7497",
		ReadTimeout:  2 * time.Second,
		WriteTimeout: 2 * time.Second,
		Privileges:   &rw,
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
	if err := writer.StartAPI(); err != nil {
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
	if _, err := placeOrder.Send(ctx); err != nil {
		panic(err)
	}

	// Get open orders to find the order_id from the above placeOrder.
	// Wait for the response and cancel the open order.
	openOrders := client.NewAllOpenOrdersRequest(writer)
	if err := openOrders.Send(ctx); err != nil {
		panic(err)
	}
	orderStatusCh := make(chan *api.OrderStatus)
	go func() {
		for {
			if handler.orderStatus != nil {
				orderStatusCh <- handler.orderStatus
				return
			}
		}
	}()
	select {
	case <-time.After(10 * time.Second):
		logger.Error("timed out waiting for open order")
		os.Exit(1)
	case o := <-orderStatusCh:
		cancelOrder := client.NewCancelOrderRequest(writer, o.GetOrderId())
		if err := cancelOrder.Send(ctx); err != nil {
			panic(err)
		}
		logger.Info("canceled order_id=%v", o.GetOrderId())
	}
}
