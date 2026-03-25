package client

import (
	"context"
	"fmt"

	"github.com/benliusf/trader_workstation_go_sdk/pkg/log"
	"github.com/benliusf/trader_workstation_go_sdk/pkg/send"
	"google.golang.org/protobuf/proto"

	api "github.com/benliusf/trader_workstation_go_sdk/api/v104401"
)

type ESender struct {
	twsClient *TWSClient
	logger    log.Logger
}

func NewSender(cl *TWSClient) (*ESender, error) {
	if cl == nil {
		return nil, fmt.Errorf("nil TWSClient")
	}
	return &ESender{
		cl, cl.logger,
	}, nil
}

func (e *ESender) StartAPI() error {
	status := e.twsClient.status
	if status.isReady() {
		return nil
	}
	req := &api.StartApiRequest{
		ClientId: &e.twsClient.clientId,
	}
	return send.Write(e.twsClient.conn, req)
}

func (e *ESender) Send(ctx context.Context, m proto.Message) error {
	if err := ctx.Err(); err != nil {
		return err
	}
	status := e.twsClient.status
	if !status.isReady() {
		return ErrClientNotReady
	}
	r := e.twsClient.privileges
	if r == nil {
		return fmt.Errorf("%w: client lacks privileges", ErrNotAllowed)
	}
	switch m.(type) {
	case *api.PlaceOrderRequest:
		if !CanCreate(r.Orders) {
			return ErrNoCreate
		}
	case *api.CancelOrderRequest:
		if !CanDelete(r.Orders) {
			return ErrNoDelete
		}
	case *api.GlobalCancelRequest:
		if !CanDelete(r.Orders) {
			return ErrNoDelete
		}
	default:
		// noop
	}
	return send.Write(e.twsClient.conn, m)
}
