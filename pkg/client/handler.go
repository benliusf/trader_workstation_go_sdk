package client

import (
	api "github.com/benliusf/trader_workstation_go_sdk/api/v104401"
	"github.com/benliusf/trader_workstation_go_sdk/pkg/read"
)

type EHandler interface {
	NextValidId(m *api.NextValidId) error
	AccountSummary(m *api.AccountSummary) error
	ContractData(m *api.ContractData) error
	TickPrice(m *api.TickPrice) error
	TickSize(m *api.TickSize) error
	TickString(m *api.TickString) error
	HistoricalData(m *api.HistoricalData) error
	HistoricalDataEnd(m *api.HistoricalDataEnd) error
	ErrorMessage(m *api.ErrorMessage) error
	Unsupported(m *read.Message) error
}
