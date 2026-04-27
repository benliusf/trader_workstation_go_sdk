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
	"google.golang.org/protobuf/encoding/protojson"

	api "github.com/benliusf/trader_workstation_go_sdk/api/v104401"
)

// An example to demonstrate an API call to Historical Market Data -
//
//	https://www.interactivebrokers.com/campus/ibkr-api-page/twsapi-doc/#hist-md

type thisHandler struct {
	examples.ExampleHandler

	reqId int32
	data  []*api.HistoricalData

	done context.CancelFunc
}

func newThisHandler(done context.CancelFunc, logger log.Logger) *thisHandler {
	return &thisHandler{
		ExampleHandler: examples.ExampleHandler{
			Logger: logger,
		},
		done: done,
		data: []*api.HistoricalData{},
	}
}

func (h *thisHandler) setReqId(reqId int32) {
	h.reqId = reqId
}

func (h *thisHandler) HistoricalData(m *api.HistoricalData) error {
	if *m.ReqId == h.reqId {
		h.data = append(h.data, m)
	}
	return nil
}

func (h *thisHandler) HistoricalDataEnd(m *api.HistoricalDataEnd) error {
	if *m.ReqId == h.reqId {
		h.done()
	}
	return nil
}

func main() {
	conf := client.TWSConfig{
		ClientId:     0,
		Host:         "localhost",
		Port:         "7497",
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 5 * time.Second,
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

	handler := newThisHandler(done, logger)
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

	// Define params with 24 hour time range for AAPL stock.
	now := time.Now().Truncate(60 * time.Minute)
	startTime := now.Add(-24 * time.Hour)
	params := &client.QueryParams{
		StartTime:  startTime,
		EndTime:    now,
		BarSize:    client.ONE_HOUR,
		WhatToShow: client.TRADES,
	}
	contr := client.NewContractBuilder().
		SetSymbol("AAPL").
		SetSecType(client.STOCK).
		SetExchange(client.SMART).
		SetPrimaryExch(client.NASDAQ).Build()

	aaplTicker := client.NewHistoricalDataRequest(writer, contr, params)
	var send func() (int32, error) = func() (int32, error) {
		for {
			reqId, err := aaplTicker.Send(ctx)
			if err != nil {
				if errors.Is(err, client.ErrClientNotReady) {
					logger.Warn("client not ready, retrying")
					time.Sleep(1 * time.Second)
					continue
				}
				return -1, err
			}
			return reqId, nil
		}
	}

	reqId, err := send()
	if err != nil {
		panic(err)
	}
	handler.setReqId(reqId)

	select {
	case <-time.After(10 * time.Second):
		logger.Error("timed out")
		os.Exit(1)
	case <-ctx.Done():
		logger.Info("received all data")
	}

	// Print historical records -
	// ...
	// 2026/04/22 16:27:53 [INFO] {"date":"20260422 13:00:00 US/Eastern","open":272.48,"high":273.72,"low":272.47,"close":273.47,"volume":"2221787","WAP":"273.23","barCount":11840}
	// 2026/04/22 16:27:53 [INFO] {"date":"20260422 14:00:00 US/Eastern","open":273.47,"high":273.74,"low":272.1,"close":272.42,"volume":"1683503","WAP":"272.834","barCount":9310}
	// 2026/04/22 16:27:53 [INFO] {"date":"20260422 15:00:00 US/Eastern","open":272.43,"high":273.26,"low":272.26,"close":273.19,"volume":"4224526","WAP":"272.85","barCount":23745}
	// ...
	for _, d := range handler.data {
		for _, r := range d.HistoricalDataBars {
			b, _ := protojson.Marshal(r)
			logger.Info("%s", b)
		}
	}
}
