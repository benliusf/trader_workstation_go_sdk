package client

import (
	api "github.com/benliusf/trader_workstation_go_sdk/api/v104401"
)

type SecurityType string

type Exchange string

type PrimaryExchange string

type Currency string

const (
	STOCK         SecurityType = "STK"
	OPTION        SecurityType = "OPT"
	FUTURE_OPTION SecurityType = "FOP"
	FUTURE        SecurityType = "FUT"
	FOREX_PAIR    SecurityType = "CASH"
	COMMODITY     SecurityType = "CMDTY"
	INDEX         SecurityType = "IND"
)

const (
	SMART Exchange = "SMART"
)

const (
	NASDAQ PrimaryExchange = "NASDAQ"
	NYSE   PrimaryExchange = "NYSE"
)

const (
	USD Currency = "USD"
)

type ContractBuilder struct {
	c *api.Contract
}

func NewContractBuilder() *ContractBuilder {
	return &ContractBuilder{
		c: &api.Contract{
			Currency: strPtr(string(USD)),
		},
	}
}

func (b *ContractBuilder) Build() *api.Contract {
	return b.c
}

func (b *ContractBuilder) SetId(v int32) *ContractBuilder {
	b.c.ConId = int32Ptr(v)
	return b
}

func (b *ContractBuilder) SetSymbol(v string) *ContractBuilder {
	b.c.Symbol = &v
	return b
}

func (b *ContractBuilder) SetSecType(v SecurityType) *ContractBuilder {
	b.c.SecType = strPtr(string(v))
	return b
}

func (b *ContractBuilder) SetExchange(v Exchange) *ContractBuilder {
	b.c.Exchange = strPtr(string(v))
	return b
}
func (b *ContractBuilder) SetPrimaryExch(v PrimaryExchange) *ContractBuilder {
	b.c.PrimaryExch = strPtr(string(v))
	return b
}

func (b *ContractBuilder) SetCurrency(v Currency) *ContractBuilder {
	b.c.Currency = strPtr(string(v))
	return b
}
