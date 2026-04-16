package client

import (
	"context"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestHistoricalDataRequest(t *testing.T) {
	now := time.Now()
	contr := NewContractBuilder().
		SetSymbol("AAPL").
		SetSecType(STOCK).
		SetExchange(SMART).
		SetPrimaryExch(NASDAQ).Build()
	params := &QueryParams{
		StartTime:  now.Add(-720 * 7 * time.Hour),
		EndTime:    now.Add(-720 * 6 * time.Hour),
		BarSize:    ONE_HOUR,
		WhatToShow: TRADES,
	}
	req := NewHistoricalDataRequest(&ESender{}, contr, params)

	_, err := req.Send(context.TODO())
	assert.ErrorIs(t, err, ErrInvalidParam)
}
