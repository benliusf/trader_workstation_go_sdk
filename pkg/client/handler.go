package client

import (
	api "github.com/benliusf/trader_workstation_go_sdk/api/v104401"
	"github.com/benliusf/trader_workstation_go_sdk/pkg/read"
)

type EHandler interface {
	NextValidId(m *api.NextValidId) error

	AccountSummary(m *api.AccountSummary) error

	AccountSummaryEnd(m *api.AccountSummaryEnd) error

	AccountValue(m *api.AccountValue) error

	AccountUpdateTime(m *api.AccountUpdateTime) error

	AccountDataEnd(m *api.AccountDataEnd) error

	ContractData(m *api.ContractData) error

	ContractDataEnd(m *api.ContractDataEnd) error

	TickPrice(m *api.TickPrice) error

	TickSize(m *api.TickSize) error

	TickString(m *api.TickString) error

	HeadTimestamp(m *api.HeadTimestamp) error

	HistoricalData(m *api.HistoricalData) error

	HistoricalDataEnd(m *api.HistoricalDataEnd) error

	Position(m *api.Position) error

	PositionEnd(m *api.PositionEnd) error

	OpenOrder(m *api.OpenOrder) error

	OpenOrdersEnd(m *api.OpenOrdersEnd) error

	OrderStatus(m *api.OrderStatus) error

	ExecutionDetails(m *api.ExecutionDetails) error

	ExecutionDetailsEnd(m *api.ExecutionDetailsEnd) error

	ErrorMessage(m *api.ErrorMessage) error

	Unsupported(m *read.Message) error
}
