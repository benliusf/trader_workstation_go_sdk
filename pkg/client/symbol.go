package client

type SecType string

type Exchange string

type PrimaryExchange string

type Currency string

const (
	STOCK SecType = "STK"
)

const (
	SMART Exchange = "SMART"
)

const (
	NASDAQ PrimaryExchange = "NASDAQ"
)

const (
	USD Currency = "USD"
)

type Symbol struct {
	Ticker   string
	Type     SecType
	Exch     Exchange
	PrimExch PrimaryExchange
	Curr     Currency
}
