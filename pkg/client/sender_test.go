package client

import (
	"context"
	"errors"
	"testing"

	api "github.com/benliusf/trader_workstation_go_sdk/api/v104401"
	"github.com/stretchr/testify/assert"
)

func TestSender(t *testing.T) {
	cl, _ := NewClient(TWSConfig{
		ClientId: 0,
		Host:     "localhost",
		Port:     "7497",
	}, nil)
	cl.status.setReady()

	contr := NewContractBuilder().
		SetSymbol("AAPL").
		SetSecType(STOCK).
		SetExchange(SMART).
		SetPrimaryExch(NASDAQ)

	sender, _ := NewSender(cl)
	req := NewPlaceOrderRequest(sender, 0, contr.Build(), &api.Order{})
	err := req.Send(context.TODO())
	assert.NotNil(t, err)
	assert.True(t, errors.Is(err, ErrNotAllowed))
}
