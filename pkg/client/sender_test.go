package client

import (
	"context"
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
	req := NewPlaceOrderRequest(sender, contr.Build(), &api.Order{})
	_, err := req.Send(context.TODO())
	assert.ErrorIs(t, err, ErrNotAllowed)
}
