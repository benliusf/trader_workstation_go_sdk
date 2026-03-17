package client

import "time"

type Duration string

const (
	ONE_DAY  Duration = "1 D"
	ONE_WEEK Duration = "1 W"
)

type BarSize string

const (
	FIFTEEN_SECOND   BarSize = "15 secs"
	ONE_MINUTE       BarSize = "1 mins"
	FIVE_MINUTE      BarSize = "5 mins"
	THIRTY_MINUTE    BarSize = "30 mins"
	ONE_HOUR         BarSize = "1 hour"
	ONE_DAY_BAR_SIZE BarSize = "1 day"
)

type DisplayType string

const (
	TRADES                    DisplayType = "TRADES"
	MIDPOINT                  DisplayType = "MIDPOINT"
	BID                       DisplayType = "BID"
	ASK                       DisplayType = "ASK"
	BID_ASK                   DisplayType = "BID_ASK"
	HISTORICAL_VOLATILITY     DisplayType = "HISTORICAL_VOLATILITY"
	OPTION_IMPLIED_VOLATILITY DisplayType = "OPTION_IMPLIED_VOLATILITY"
	SCHEDULE                  DisplayType = "SCHEDULE"
)

type QueryParams struct {
	EndTime    time.Time
	Duration   Duration
	BarSize    BarSize
	WhatToShow DisplayType
}
