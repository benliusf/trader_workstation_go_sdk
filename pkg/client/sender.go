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
	req := &api.StartApiRequest{
		ClientId: &e.twsClient.Conf.ClientID,
	}
	return send.Write(e.twsClient.conn, req)
}

func (e *ESender) Send(ctx context.Context, m proto.Message) error {
	if err := ctx.Err(); err != nil {
		return err
	}
	if !e.twsClient.status.isReady() {
		return ErrClientNotReady
	}
	return send.Write(e.twsClient.conn, m)
}
