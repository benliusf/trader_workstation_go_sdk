package simple

import (
	"fmt"

	api "github.com/benliusf/trader_workstation_go_sdk/api/v104401"

	"github.com/benliusf/trader_workstation_go_sdk/pkg/read"
)

type simpleHandler struct {
	done chan struct{}
	errs chan error

	reqId int32

	nextValidId      *api.NextValidId
	accountSummary   []*api.AccountSummary
	accountValue     []*api.AccountValue
	contractData     *api.ContractData
	headTimestamp    *api.HeadTimestamp
	historicalData   *api.HistoricalData
	position         []*api.Position
	openOrder        []*api.OpenOrder
	executionDetails *api.ExecutionDetails
}

func newSimpleHandler(reqId int32) *simpleHandler {
	return &simpleHandler{done: make(chan struct{}), errs: make(chan error), reqId: reqId}
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
	if m.GetReqId() == h.reqId {
		h.accountSummary = append(h.accountSummary, m)
	}
	return nil
}

func (h *simpleHandler) AccountSummaryEnd(m *api.AccountSummaryEnd) error {
	if m.GetReqId() == h.reqId {
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
	if m.GetReqId() == h.reqId {
		h.contractData = m
		h.done <- struct{}{}
	}
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
	if m.GetReqId() == h.reqId {
		h.headTimestamp = m
		h.done <- struct{}{}
	}
	return nil
}

func (h *simpleHandler) HistoricalData(m *api.HistoricalData) error {
	if m.GetReqId() == h.reqId {
		h.historicalData = m
		h.done <- struct{}{}
	}
	return nil
}

func (h *simpleHandler) HistoricalDataEnd(m *api.HistoricalDataEnd) error {
	return nil
}

func (h *simpleHandler) Position(m *api.Position) error {
	h.position = append(h.position, m)
	return nil
}

func (h *simpleHandler) PositionEnd(m *api.PositionEnd) error {
	h.done <- struct{}{}
	return nil
}

func (h *simpleHandler) OpenOrder(m *api.OpenOrder) error {
	h.openOrder = append(h.openOrder, m)
	return nil
}

func (h *simpleHandler) OpenOrdersEnd(m *api.OpenOrdersEnd) error {
	h.done <- struct{}{}
	return nil
}

func (h *simpleHandler) OrderStatus(m *api.OrderStatus) error {
	return nil
}

func (h *simpleHandler) ExecutionDetails(m *api.ExecutionDetails) error {
	if m.GetReqId() == h.reqId {
		h.executionDetails = m
		h.done <- struct{}{}
	}
	return nil
}

func (h *simpleHandler) ExecutionDetailsEnd(m *api.ExecutionDetailsEnd) error {
	return nil
}

func (h *simpleHandler) ErrorMessage(m *api.ErrorMessage) error {
	if (m.GetErrorCode() >= 100 && m.GetErrorCode() < 1000) ||
		(m.GetErrorCode() >= 10000) {
		h.errs <- fmt.Errorf(m.GetErrorMsg())
	}
	return nil
}

func (h *simpleHandler) Unsupported(m *read.Message) error {
	return nil
}
