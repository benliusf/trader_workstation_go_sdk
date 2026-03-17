package examples

import (
	api "github.com/benliusf/trader_workstation_go_sdk/api/v104401"
	"github.com/benliusf/trader_workstation_go_sdk/pkg/log"
	"github.com/benliusf/trader_workstation_go_sdk/pkg/read"

	"google.golang.org/protobuf/encoding/protojson"
)

type ExampleHandler struct {
	logger log.Logger
}

func NewExampleHandler(logger log.Logger) *ExampleHandler {
	return &ExampleHandler{
		logger,
	}
}

func (h *ExampleHandler) NextValidId(m *api.NextValidId) error {
	h.logger.Info("received next valid id: %v", m.GetOrderId())
	return nil
}

func (h *ExampleHandler) AccountSummary(m *api.AccountSummary) error {
	d, _ := protojson.Marshal(m)
	h.logger.Info("received account summary data: %s", d)
	return nil
}

func (h *ExampleHandler) ContractData(m *api.ContractData) error {
	d, _ := protojson.Marshal(m)
	h.logger.Info("received contract data: %s", d)
	return nil
}

func (h *ExampleHandler) TickPrice(m *api.TickPrice) error {
	d, _ := protojson.Marshal(m)
	h.logger.Info("received tick price: %s", d)
	return nil
}

func (h *ExampleHandler) TickSize(m *api.TickSize) error {
	d, _ := protojson.Marshal(m)
	h.logger.Info("received tick size: %s", d)
	return nil
}

func (h *ExampleHandler) TickString(m *api.TickString) error {
	d, _ := protojson.Marshal(m)
	h.logger.Info("received tick string: %s", d)
	return nil
}

func (h *ExampleHandler) HistoricalData(m *api.HistoricalData) error {
	d, _ := protojson.Marshal(m)
	h.logger.Info("received historical data: %s", d)
	return nil
}

func (h *ExampleHandler) HistoricalDataEnd(m *api.HistoricalDataEnd) error {
	d, _ := protojson.Marshal(m)
	h.logger.Info("received historical data end: %s", d)
	return nil
}

func (h *ExampleHandler) ErrorMessage(m *api.ErrorMessage) error {
	h.logger.Info("received error message: %v", m.GetErrorMsg())
	return nil
}

func (h *ExampleHandler) Unsupported(m *read.Message) error {
	msgId := int32(-1)
	if tmp, err := m.ReadMsgId(); err == nil {
		msgId = tmp
	}
	h.logger.Info("received unsupported message id=%d", msgId)
	return nil
}
