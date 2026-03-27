package client

import (
	"context"
	"fmt"
	"reflect"

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

func (e *ESender) Send(ctx context.Context, m proto.Message) (int32, error) {
	const (
		requestId = "ReqId"
		orderId   = "OrderId"
	)
	if err := ctx.Err(); err != nil {
		return -1, err
	}
	status := e.twsClient.status
	if !status.isReady() {
		return -1, ErrClientNotReady
	}
	r := e.twsClient.privileges
	if r == nil {
		return -1, fmt.Errorf("%w: client lacks privileges", ErrNotAllowed)
	}
	if err := ValidateRequestACL(r, m); err != nil {
		return -1, err
	}
	idFieldName := ""
	switch m.(type) {
	case *api.AccountSummaryRequest, *api.ContractDataRequest, *api.MarketDataRequest, *api.HistoricalDataRequest:
		idFieldName = requestId
	case *api.PlaceOrderRequest:
		idFieldName = orderId
	}
	if idFieldName == "" {
		return -1, send.Write(e.twsClient.conn, m)
	}
	s := reflect.ValueOf(m)
	if s.Kind() == reflect.Ptr {
		s = s.Elem()
	}
	f := s.FieldByName(idFieldName)
	if !f.IsValid() || !f.CanSet() {
		return -1, fmt.Errorf("unable to set '%s'", idFieldName)
	}
	fPtr := reflect.New(f.Type().Elem())
	var id int32 = -1
	switch idFieldName {
	case requestId:
		id = e.twsClient.GetNextReqId()
	case orderId:
		id = e.twsClient.GetNextOrderId()
	}
	fPtr.Elem().SetInt(int64(id))
	f.Set(fPtr)
	return id, send.Write(e.twsClient.conn, m)
}
