package client

import (
	"fmt"

	api "github.com/benliusf/trader_workstation_go_sdk/api/v104401"
)

type OrderType string

const (
	MARKET OrderType = "MKT"
	LIMIT  OrderType = "LMT"
)

type Action string

const (
	BUY  Action = "BUY"
	SELL Action = "SELL"
)

type TimeInForce string

const (
	DAY_ONLY TimeInForce = "DAY"
	GTC      TimeInForce = "GTC"
)

type MarketOrderBuilder struct {
	o *api.Order
}

func NewMarketOrderBuilder() *MarketOrderBuilder {
	return &MarketOrderBuilder{
		o: &api.Order{
			OrderType: strPtr(string(MARKET)),
		},
	}
}

func (b *MarketOrderBuilder) SetAction(v Action) *MarketOrderBuilder {
	b.o.Action = strPtr(string(v))
	return b
}

func (b *MarketOrderBuilder) SetQuantity(v float64) *MarketOrderBuilder {
	b.o.TotalQuantity = strPtr(fmt.Sprintf("%.2f", v))
	return b
}

func (b *MarketOrderBuilder) SetTimeInForce(v TimeInForce) *MarketOrderBuilder {
	b.o.Tif = strPtr(string(v))
	return b
}

func (b *MarketOrderBuilder) SetTransmit() *MarketOrderBuilder {
	b.o.Transmit = boolPtr(true)
	return b
}

func (b *MarketOrderBuilder) Build() *api.Order {
	return b.o
}

type LimitOrderBuilder struct {
	o *api.Order
}

func NewLimitOrderBuilder() *LimitOrderBuilder {
	return &LimitOrderBuilder{
		o: &api.Order{
			OrderType: strPtr(string(LIMIT)),
		},
	}
}

func (b *LimitOrderBuilder) SetAction(v Action) *LimitOrderBuilder {
	b.o.Action = strPtr(string(v))
	return b
}

func (b *LimitOrderBuilder) SetQuantity(v float64) *LimitOrderBuilder {
	b.o.TotalQuantity = strPtr(fmt.Sprintf("%.2f", v))
	return b
}

func (b *LimitOrderBuilder) SetPrice(v float64) *LimitOrderBuilder {
	b.o.LmtPrice = &v
	return b
}

func (b *LimitOrderBuilder) SetTimeInForce(v TimeInForce) *LimitOrderBuilder {
	b.o.Tif = strPtr(string(v))
	return b
}

func (b *LimitOrderBuilder) SetTransmit() *LimitOrderBuilder {
	b.o.Transmit = boolPtr(true)
	return b
}

func (b *LimitOrderBuilder) Build() *api.Order {
	return b.o
}
