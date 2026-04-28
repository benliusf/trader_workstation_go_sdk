package main

import (
	"context"
	"errors"
	"fmt"
	"os"
	"strconv"
	"time"

	"github.com/benliusf/trader_workstation_go_sdk/examples"
	"github.com/benliusf/trader_workstation_go_sdk/pkg/client"
	"github.com/benliusf/trader_workstation_go_sdk/pkg/log"

	api "github.com/benliusf/trader_workstation_go_sdk/api/v104401"
)

// An example to demonstrate an API call to Positions -
//	https://www.interactivebrokers.com/campus/ibkr-api-page/twsapi-doc/#positions
//
// In this example, we get all the current positions in the account and create "SELL" orders for all the positions.

type thisHandler struct {
	examples.ExampleHandler

	accountSummary *api.AccountSummary
	positions      []*api.Position
	positionEnd    bool
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

func (h *thisHandler) Position(m *api.Position) error {
	h.positions = append(h.positions, m)
	return nil
}

func (h *thisHandler) PositionEnd(m *api.PositionEnd) error {
	h.positionEnd = true
	return nil
}

func main() {
	conf := client.TWSConfig{
		ClientId:     0,
		Host:         "localhost",
		Port:         "7497",
		ReadTimeout:  2 * time.Second,
		WriteTimeout: 2 * time.Second,
		Privileges:   client.ReadAndWrite(),
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
	}()

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
		logger.Info("account_id=%v", *a.Account)
	}

	// Request to get all positions in the account.
	positions := client.NewPositionsRequest(writer)
	if err := positions.Send(ctx); err != nil {
		panic(err)
	}
	positionsEndCh := make(chan bool)
	go func() {
		for {
			if handler.positionEnd {
				positionsEndCh <- true
				return
			}
		}
	}()
	select {
	case <-time.After(10 * time.Second):
		logger.Error("timed out waiting for positions")
		os.Exit(1)
	case <-positionsEndCh:
		logger.Info("received all positions")
	}

	// Function to close a position. It places a "SELL" order on the position.
	var sell func(p *api.Position) error = func(p *api.Position) error {
		contr := client.NewContractBuilder().
			SetExchange(client.SMART).
			SetId(p.GetContract().GetConId()).Build()
		qty, err := strconv.ParseFloat(p.GetPosition(), 64)
		if err != nil {
			return err
		}
		order := client.NewMarketOrderBuilder().
			SetAction(client.SELL).
			SetQuantity(qty).
			SetTimeInForce(client.DAY_ONLY).
			SetTransmit().Build()
		placeOrder := client.NewPlaceOrderRequest(writer, contr, order)
		if _, err := placeOrder.Send(ctx); err != nil {
			panic(err)
		}
		return nil
	}
	for _, p := range handler.positions {
		if err := sell(p); err != nil {
			panic(err)
		}
		logger.Info("sold position in %v", *p.GetContract().Symbol)
	}

	time.Sleep(5 * time.Second)
}
