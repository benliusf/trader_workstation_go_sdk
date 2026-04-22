package examples

import (
	api "github.com/benliusf/trader_workstation_go_sdk/api/v104401"
	"github.com/benliusf/trader_workstation_go_sdk/pkg/log"
	"github.com/benliusf/trader_workstation_go_sdk/pkg/read"

	"google.golang.org/protobuf/encoding/protojson"
)

// This example handler implements the EHandler interface from handler.go
// and merely prints every api message onto stdout.

type ExampleHandler struct {
	Logger log.Logger
}

func NewExampleHandler(logger log.Logger) *ExampleHandler {
	return &ExampleHandler{
		logger,
	}
}

func (h *ExampleHandler) NextValidId(m *api.NextValidId) error {
	h.Logger.Info("received next valid id: %v", m.GetOrderId())
	return nil
}

func (h *ExampleHandler) AccountSummary(m *api.AccountSummary) error {
	d, _ := protojson.Marshal(m)
	h.Logger.Info("received account summary data: %s", d)
	return nil
}

func (h *ExampleHandler) AccountSummaryEnd(m *api.AccountSummaryEnd) error {
	d, _ := protojson.Marshal(m)
	h.Logger.Info("received account summary end: %s", d)
	return nil
}

func (h *ExampleHandler) AccountValue(m *api.AccountValue) error {
	d, _ := protojson.Marshal(m)
	h.Logger.Info("received account value: %s", d)
	return nil
}

func (h *ExampleHandler) AccountUpdateTime(m *api.AccountUpdateTime) error {
	d, _ := protojson.Marshal(m)
	h.Logger.Info("received account update time: %s", d)
	return nil
}

func (h *ExampleHandler) AccountDataEnd(m *api.AccountDataEnd) error {
	d, _ := protojson.Marshal(m)
	h.Logger.Info("received account data end: %s", d)
	return nil
}

func (h *ExampleHandler) ContractData(m *api.ContractData) error {
	d, _ := protojson.Marshal(m)
	h.Logger.Info("received contract data: %s", d)
	return nil
}

func (h *ExampleHandler) ContractDataEnd(m *api.ContractDataEnd) error {
	d, _ := protojson.Marshal(m)
	h.Logger.Info("received contract data end: %s", d)
	return nil
}

func (h *ExampleHandler) TickPrice(m *api.TickPrice) error {
	d, _ := protojson.Marshal(m)
	h.Logger.Info("received tick price: %s", d)
	return nil
}

func (h *ExampleHandler) TickSize(m *api.TickSize) error {
	d, _ := protojson.Marshal(m)
	h.Logger.Info("received tick size: %s", d)
	return nil
}

func (h *ExampleHandler) TickString(m *api.TickString) error {
	d, _ := protojson.Marshal(m)
	h.Logger.Info("received tick string: %s", d)
	return nil
}

func (h *ExampleHandler) HeadTimestamp(m *api.HeadTimestamp) error {
	d, _ := protojson.Marshal(m)
	h.Logger.Info("received head timestamp: %s", d)
	return nil
}

func (h *ExampleHandler) HistoricalData(m *api.HistoricalData) error {
	d, _ := protojson.Marshal(m)
	h.Logger.Info("received historical data: %s", d)
	return nil
}

func (h *ExampleHandler) HistoricalDataEnd(m *api.HistoricalDataEnd) error {
	d, _ := protojson.Marshal(m)
	h.Logger.Info("received historical data end: %s", d)
	return nil
}

func (h *ExampleHandler) Position(m *api.Position) error {
	d, _ := protojson.Marshal(m)
	h.Logger.Info("received position data: %s", d)
	return nil
}

func (h *ExampleHandler) PositionEnd(m *api.PositionEnd) error {
	d, _ := protojson.Marshal(m)
	h.Logger.Info("received position end: %s", d)
	return nil
}

func (h *ExampleHandler) OpenOrder(m *api.OpenOrder) error {
	d, _ := protojson.Marshal(m)
	h.Logger.Info("received open order: %s", d)
	return nil
}

func (h *ExampleHandler) OpenOrdersEnd(m *api.OpenOrdersEnd) error {
	d, _ := protojson.Marshal(m)
	h.Logger.Info("received open orders end: %s", d)
	return nil
}

func (h *ExampleHandler) OrderStatus(m *api.OrderStatus) error {
	d, _ := protojson.Marshal(m)
	h.Logger.Info("received order status: %s", d)
	return nil
}

func (h *ExampleHandler) ExecutionDetails(m *api.ExecutionDetails) error {
	d, _ := protojson.Marshal(m)
	h.Logger.Info("received execution details: %s", d)
	return nil
}

func (h *ExampleHandler) ExecutionDetailsEnd(m *api.ExecutionDetailsEnd) error {
	d, _ := protojson.Marshal(m)
	h.Logger.Info("received execution details end: %s", d)
	return nil
}

func (h *ExampleHandler) ErrorMessage(m *api.ErrorMessage) error {
	h.Logger.Info("received error message: %v", m.GetErrorMsg())
	return nil
}

func (h *ExampleHandler) Unsupported(m *read.Message) error {
	msgId := int32(-1)
	if tmp, err := m.ReadMsgId(); err == nil {
		msgId = tmp
	}
	h.Logger.Info("received unsupported message id=%d", msgId)
	return nil
}
