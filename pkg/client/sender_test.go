package client

import (
	"context"
	"testing"

	api "github.com/benliusf/trader_workstation_go_sdk/api/v104401"
	"github.com/benliusf/trader_workstation_go_sdk/pkg/net"
	"github.com/stretchr/testify/assert"
)

func TestSender(t *testing.T) {
	cl, _ := NewClient(TWSConfig{
		ClientId: 0,
		Host:     "localhost",
		Port:     "7497",
	}, nil)
	cl.conn = net.NewMockConn()
	cl.status.setReady()

	sender, _ := NewSender(cl)
	ctx := context.TODO()

	tests := []struct {
		send func() (int32, error)
		id   int32
		err  error
	}{
		{
			send: func() (int32, error) {
				contr := NewContractBuilder().
					SetSymbol("AAPL").
					SetSecType(STOCK).
					SetExchange(SMART).
					SetPrimaryExch(NASDAQ).Build()
				req := NewPlaceOrderRequest(sender, contr, &api.Order{})
				return req.Send(ctx)
			},
			id:  -1,
			err: ErrNotAllowed,
		},
		{
			send: func() (int32, error) {
				req := NewAccountDataRequest(sender, "")
				return -1, req.Send(ctx)
			},
			id:  -1,
			err: nil,
		},
	}

	for _, tt := range tests {
		id, err := tt.send()
		assert.Equal(t, tt.id, id)
		if err != nil {
			assert.ErrorIs(t, err, tt.err)
		}
	}
}
