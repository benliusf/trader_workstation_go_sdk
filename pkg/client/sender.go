package client

import (
	"context"
	"fmt"
	"reflect"
	"time"

	"github.com/benliusf/trader_workstation_go_sdk/pkg/log"
	"github.com/benliusf/trader_workstation_go_sdk/pkg/read"
	"github.com/benliusf/trader_workstation_go_sdk/pkg/send"
	"google.golang.org/protobuf/proto"

	api "github.com/benliusf/trader_workstation_go_sdk/api/v104401"
)

type startHandler struct {
	nextValidIdCh chan int32
}

func (h *startHandler) NextValidId(m *api.NextValidId) error {
	h.nextValidIdCh <- *m.OrderId
	return nil
}

func (h *startHandler) AccountSummary(m *api.AccountSummary) error {
	return nil
}

func (h *startHandler) AccountSummaryEnd(m *api.AccountSummaryEnd) error {
	return nil
}

func (h *startHandler) AccountValue(m *api.AccountValue) error {
	return nil
}

func (h *startHandler) AccountUpdateTime(m *api.AccountUpdateTime) error {
	return nil
}

func (h *startHandler) AccountDataEnd(m *api.AccountDataEnd) error {
	return nil
}

func (h *startHandler) ContractData(m *api.ContractData) error {
	return nil
}

func (h *startHandler) ContractDataEnd(m *api.ContractDataEnd) error {
	return nil
}

func (h *startHandler) TickPrice(m *api.TickPrice) error {
	return nil
}

func (h *startHandler) TickSize(m *api.TickSize) error {
	return nil
}

func (h *startHandler) TickString(m *api.TickString) error {
	return nil
}

func (h *startHandler) HeadTimestamp(m *api.HeadTimestamp) error {
	return nil
}

func (h *startHandler) HistoricalData(m *api.HistoricalData) error {
	return nil
}

func (h *startHandler) HistoricalDataEnd(m *api.HistoricalDataEnd) error {
	return nil
}

func (h *startHandler) Position(m *api.Position) error {
	return nil
}

func (h *startHandler) PositionEnd(m *api.PositionEnd) error {
	return nil
}

func (h *startHandler) OpenOrder(m *api.OpenOrder) error {
	return nil
}

func (h *startHandler) OpenOrdersEnd(m *api.OpenOrdersEnd) error {
	return nil
}

func (h *startHandler) OrderStatus(m *api.OrderStatus) error {
	return nil
}

func (h *startHandler) ExecutionDetails(m *api.ExecutionDetails) error {
	return nil
}

func (h *startHandler) ExecutionDetailsEnd(m *api.ExecutionDetailsEnd) error {
	return nil
}

func (h *startHandler) ErrorMessage(m *api.ErrorMessage) error {
	return nil
}

func (h *startHandler) Unsupported(m *read.Message) error {
	return nil
}

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

func (e *ESender) StartAPI(timeout time.Duration) error {
	status := e.twsClient.status
	if status.isReady() {
		return nil
	}

	ctx, done := context.WithCancel(context.Background())
	defer done()
	reader, err := NewReader(e.twsClient)
	if err != nil {
		return err
	}
	nextValidIdCh := make(chan int32)
	handler := &startHandler{nextValidIdCh}

	go func() {
		if err := reader.Read(ctx, handler); err != nil {
			e.logger.Debug(fmt.Sprintf("start api read error: %v", err))
		}
		time.Sleep(10 * time.Millisecond)
	}()

	req := &api.StartApiRequest{
		ClientId: &e.twsClient.conf.ClientId,
	}
	if err := send.Write(e.twsClient.conn, req); err != nil {
		return err
	}

	select {
	case <-time.After(timeout):
		return fmt.Errorf("timed out")
	case id := <-nextValidIdCh:
		status.setNextOrderId(id)
		status.setReady()
	}
	return nil
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
	r := e.twsClient.conf.Privileges
	if r.None() {
		return -1, fmt.Errorf("%w: client lacks privileges", ErrNotAllowed)
	}
	if err := ValidateRequestACL(r, m); err != nil {
		return -1, err
	}
	idFieldName := ""
	switch m.(type) {
	case *api.AccountSummaryRequest, *api.ContractDataRequest, *api.MarketDataRequest, *api.HistoricalDataRequest, *api.ExecutionRequest:
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
		id = e.twsClient.NextReqId()
	case orderId:
		id = e.twsClient.NextOrderId()
	}
	fPtr.Elem().SetInt(int64(id))
	f.Set(fPtr)
	return id, send.Write(e.twsClient.conn, m)
}
