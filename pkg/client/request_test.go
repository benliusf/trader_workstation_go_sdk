package client

import (
	"context"
	"testing"
	"time"

	"github.com/benliusf/trader_workstation_go_sdk/pkg/net"
	"github.com/stretchr/testify/assert"
)

func TestHistoricalDataRequest(t *testing.T) {
	cl, err := NewClient(TWSConfig{
		Host:     "localhost",
		Port:     "7497",
		ClientId: 0,
	}, nil)
	assert.NoError(t, err)

	cl.conn = net.NewMockConn()
	cl.status.setReady()
	cl.status.reqId = 101

	sender, err := NewSender(cl)
	assert.NoError(t, err)

	now := time.Now().Truncate(60 * time.Minute)
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
	id, err := req.Send(context.TODO())
	assert.NoError(t, err)
	assert.Equal(t, int32(101), id)
}
