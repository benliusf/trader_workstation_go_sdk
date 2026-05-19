package simple

import (
	"context"
	"fmt"
	"sync"
	"time"

	api "github.com/benliusf/trader_workstation_go_sdk/api/v104401"
	"github.com/benliusf/trader_workstation_go_sdk/pkg/client"
	"github.com/benliusf/trader_workstation_go_sdk/pkg/log"
)

type SimpleClient struct {
	*client.TWSClient

	mu sync.Mutex

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

func (c *SimpleClient) get(ctx context.Context, send func(ctx context.Context) (int32, error)) (res *simpleHandler, err error) {
	c.mu.Lock()
	defer c.mu.Unlock()

	id, err := send(ctx)
	if err != nil {
		return nil, err
	}
	ctx, cancel := context.WithCancel(ctx)
	done := make(chan struct{})
	handler := newSimpleHandler(id)
	go func() {
		if err := c.reader.Read(ctx, handler); err != nil {
			c.logger.Error(fmt.Sprintf("read error: %v", err))
		}
		done <- struct{}{}
	}()
	select {
	case <-handler.done:
		res = handler
	case err = <-handler.errs:
	}
	cancel()
	<-done
	return res, err
}

func (c *SimpleClient) GetNextValidId(ctx context.Context) (*api.NextValidId, error) {
	send := func(ctx context.Context) (int32, error) {
		req := client.NewNextValidIdRequest(c.sender)
		return -1, req.Send(ctx)
	}
	res, err := c.get(ctx, send)
	if err != nil {
		return nil, err
	}
	return res.nextValidId, nil
}

func (c *SimpleClient) GetAccountSummary(ctx context.Context, group string, tags []client.AccountSummaryTag) ([]*api.AccountSummary, error) {
	req := client.NewAccountSummaryRequest(c.sender, group, tags)
	res, err := c.get(ctx, req.Send)
	if err != nil {
		return nil, err
	}
	return res.accountSummary, nil
}

func (c *SimpleClient) GetContractData(ctx context.Context, contr *api.Contract) (*api.ContractData, error) {
	req := client.NewContractDataRequest(c.sender, contr)
	res, err := c.get(ctx, req.Send)
	if err != nil {
		return nil, err
	}
	return res.contractData, nil
}

func (c *SimpleClient) GetHeadTimestamp(ctx context.Context, contr *api.Contract, whatToShow client.DisplayType) (*api.HeadTimestamp, error) {
	req := client.NewHeadTimestampRequest(c.sender, contr, whatToShow)
	res, err := c.get(ctx, req.Send)
	if err != nil {
		return nil, err
	}
	return res.headTimestamp, nil
}

func (c *SimpleClient) GetHistoricalData(ctx context.Context, contr *api.Contract, params *client.QueryParams) (*api.HistoricalData, error) {
	req := client.NewHistoricalDataRequest(c.sender, contr, params)
	res, err := c.get(ctx, req.Send)
	if err != nil {
		return nil, err
	}
	return res.historicalData, nil
}

func (c *SimpleClient) GetPositions(ctx context.Context) ([]*api.Position, error) {
	send := func(ctx context.Context) (int32, error) {
		req := client.NewPositionsRequest(c.sender)
		return -1, req.Send(ctx)
	}
	res, err := c.get(ctx, send)
	if err != nil {
		return nil, err
	}
	return res.position, nil
}

func (c *SimpleClient) GetOpenOrders(ctx context.Context) ([]*api.OpenOrder, error) {
	send := func(ctx context.Context) (int32, error) {
		req := client.NewOpenOrdersRequest(c.sender)
		return -1, req.Send(ctx)
	}
	res, err := c.get(ctx, send)
	if err != nil {
		return nil, err
	}
	return res.openOrder, nil
}
