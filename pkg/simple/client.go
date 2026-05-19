package simple

import (
	"context"
	"fmt"
	"time"

	api "github.com/benliusf/trader_workstation_go_sdk/api/v104401"
	"github.com/benliusf/trader_workstation_go_sdk/pkg/client"
	"github.com/benliusf/trader_workstation_go_sdk/pkg/log"
)

type SimpleClient struct {
	*client.TWSClient

	sender *client.ESender
	reader *client.EReader
	logger log.Logger
}

func NewClient(conf client.TWSConfig, logger log.Logger) (*SimpleClient, error) {
	twsClient, err := client.NewClient(conf, logger)
	if err != nil {
		return nil, err
	}
	return &SimpleClient{
		TWSClient: twsClient,
		logger:    logger,
	}, nil
}

func (c *SimpleClient) Connect(timeout time.Duration) error {
	var err error
	if err = c.TWSClient.Connect(); err != nil {
		return err
	}
	if c.reader, err = client.NewReader(c.TWSClient); err != nil {
		return err
	}
	if c.sender, err = client.NewSender(c.TWSClient); err != nil {
		return err
	}
	if err = c.sender.StartAPI(timeout); err != nil {
		return err
	}
	return nil
}

func (c *SimpleClient) GetNextValidId(ctx context.Context) (*api.NextValidId, error) {
	req := client.NewNextValidIdRequest(c.sender)
	if err := req.Send(ctx); err != nil {
		return nil, err
	}
	ctx, cancel := context.WithCancel(ctx)
	done := make(chan struct{})
	handler := newSimpleHandler(done, -1)
	go func() {
		if err := c.reader.Read(ctx, handler); err != nil {
			c.logger.Error(fmt.Sprintf("read error: %v", err))
		}
	}()
	<-handler.done
	cancel()
	return handler.nextValidId, nil
}

func (c *SimpleClient) GetAccountSummary(ctx context.Context, group string, tags []client.AccountSummaryTag) ([]*api.AccountSummary, error) {
	req := client.NewAccountSummaryRequest(c.sender, group, tags)
	id, err := req.Send(ctx)
	if err != nil {
		return nil, err
	}
	ctx, cancel := context.WithCancel(ctx)
	done := make(chan struct{})
	handler := newSimpleHandler(done, id)
	go func() {
		if err := c.reader.Read(ctx, handler); err != nil {
			c.logger.Error(fmt.Sprintf("read error: %v", err))
		}
	}()
	<-handler.done
	cancel()
	return handler.accountSummary, nil
}
