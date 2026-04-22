package client

import (
	"context"
	"testing"
	"time"

	"github.com/benliusf/trader_workstation_go_sdk/pkg/net"
	"github.com/stretchr/testify/assert"
)

func TestHistoricalDataRequest(t *testing.T) {
	cl, _ := NewClient(TWSConfig{
		ClientId: 0,
		Host:     "localhost",
		Port:     "7497",
	}, nil)
	cl.conn = net.NewMockConn()
	cl.status.setReady()
	sender, _ := NewSender(cl)

	now := time.Now()
	contr := NewContractBuilder().
		SetSymbol("AAPL").
		SetSecType(STOCK).
		SetExchange(SMART).
		SetPrimaryExch(NASDAQ).Build()
	params := &QueryParams{
		StartTime:  now.Add(-720 * 7 * time.Hour),
		EndTime:    now,
		BarSize:    ONE_HOUR,
		WhatToShow: TRADES,
	}

	req := NewHistoricalDataRequest(sender, contr, params)
	_, err := req.Send(context.TODO())
	assert.NoError(t, err)
}
