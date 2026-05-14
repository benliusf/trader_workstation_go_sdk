package main

import (
	"context"
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
		Host:         "localhost",
		Port:         "7497",
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 5 * time.Second,
		ClientId:     0,
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
	reqId, err := aaplTicker.Send(ctx)
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
	// 2026/05/13 20:05:52 [INFO] {"date":"1778684400", "open":296.51, "high":298.01, "low":296.25, "close":296.94, "volume":"3103085", "WAP":"297.123", "barCount":26442}
	// 2026/05/13 20:05:52 [INFO] {"date":"1778688000", "open":296.93, "high":299.42, "low":296.78, "close":299.15, "volume":"4144633", "WAP":"298.52", "barCount":34107}
	// 2026/05/13 20:05:52 [INFO] {"date":"1778691600", "open":299.16, "high":300.92, "low":299.12, "close":300.17, "volume":"5893813", "WAP":"300.051", "barCount":40970}
	// ...
	for _, d := range handler.data {
		for _, r := range d.HistoricalDataBars {
			b, _ := protojson.Marshal(r)
			logger.Info("%s", b)
		}
	}
}
