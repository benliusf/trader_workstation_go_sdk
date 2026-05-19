package simple

import (
	api "github.com/benliusf/trader_workstation_go_sdk/api/v104401"

	"github.com/benliusf/trader_workstation_go_sdk/pkg/read"
)

type simpleHandler struct {
	done chan struct{}

	reqId int32

	nextValidId    *api.NextValidId
	accountSummary []*api.AccountSummary
	accountValue   []*api.AccountValue
}

func newSimpleHandler(done chan struct{}, reqId int32) *simpleHandler {
	return &simpleHandler{done: done, reqId: reqId}
}

func (h *simpleHandler) NextValidId(m *api.NextValidId) error {
	h.nextValidId = m
	h.done <- struct{}{}
	return nil
}

func (h *simpleHandler) ManagedAccounts(m *api.ManagedAccounts) error {
	return nil
}

func (h *simpleHandler) AccountSummary(m *api.AccountSummary) error {
	if *m.ReqId == h.reqId {
		h.accountSummary = append(h.accountSummary, m)
	}
	return nil
}

func (h *simpleHandler) AccountSummaryEnd(m *api.AccountSummaryEnd) error {
	if *m.ReqId == h.reqId {
		h.done <- struct{}{}
	}
	return nil
}

func (h *simpleHandler) AccountValue(m *api.AccountValue) error {
	h.accountValue = append(h.accountValue, m)
	return nil
}

func (h *simpleHandler) AccountUpdateTime(m *api.AccountUpdateTime) error {
	return nil
}

func (h *simpleHandler) AccountDataEnd(m *api.AccountDataEnd) error {
	h.done <- struct{}{}
	return nil
}

func (h *simpleHandler) ContractData(m *api.ContractData) error {
	return nil
}

func (h *simpleHandler) ContractDataEnd(m *api.ContractDataEnd) error {
	return nil
}

func (h *simpleHandler) TickPrice(m *api.TickPrice) error {
	return nil
}

func (h *simpleHandler) TickSize(m *api.TickSize) error {
	return nil
}

func (h *simpleHandler) TickString(m *api.TickString) error {
	return nil
}

func (h *simpleHandler) HeadTimestamp(m *api.HeadTimestamp) error {
	return nil
}

func (h *simpleHandler) HistoricalData(m *api.HistoricalData) error {
	return nil
}

func (h *simpleHandler) HistoricalDataEnd(m *api.HistoricalDataEnd) error {
	return nil
}

func (h *simpleHandler) Position(m *api.Position) error {
	return nil
}

func (h *simpleHandler) PositionEnd(m *api.PositionEnd) error {
	return nil
}

func (h *simpleHandler) OpenOrder(m *api.OpenOrder) error {
	return nil
}

func (h *simpleHandler) OpenOrdersEnd(m *api.OpenOrdersEnd) error {
	return nil
}

func (h *simpleHandler) OrderStatus(m *api.OrderStatus) error {
	return nil
}

func (h *simpleHandler) ExecutionDetails(m *api.ExecutionDetails) error {
	return nil
}

func (h *simpleHandler) ExecutionDetailsEnd(m *api.ExecutionDetailsEnd) error {
	return nil
}

func (h *simpleHandler) ErrorMessage(m *api.ErrorMessage) error {
	return nil
}

func (h *simpleHandler) Unsupported(m *read.Message) error {
	return nil
}
