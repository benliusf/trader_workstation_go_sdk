package client

import (
	"context"
	"errors"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestHistoricalDataRequest(t *testing.T) {
	ctx := context.TODO()
	now := time.Now()
	contr := NewContractBuilder().
		SetSymbol("AAPL").
		SetSecType(STOCK).
		SetExchange(SMART).
		SetPrimaryExch(NASDAQ)
	params := &QueryParams{
		StartTime:  now.Add(-720 * 7 * time.Hour),
		EndTime:    now.Add(-720 * 6 * time.Hour),
		BarSize:    ONE_HOUR,
		WhatToShow: TRADES,
	}
	req := NewHistoricalDataRequest(&ESender{}, contr.Build(), params)
	_, err := req.Send(ctx)
	assert.NotNil(t, err)
	assert.True(t, errors.Is(err, ErrInvalidParam))
}
