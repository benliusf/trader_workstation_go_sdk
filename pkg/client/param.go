package client

import (
	"fmt"
	"math"
	"time"
)

type Temporal string

const (
	SECOND = "S"
	DAY    = "D"
	WEEK   = "W"
	MONTH  = "M"
)

type Duration struct {
	N int
	T Temporal
}

func (d Duration) String() string {
	return fmt.Sprintf("%d %s", d.N, d.T)
}

type BarSize string

const (
	ONE_SECOND     BarSize = "1 sec"
	FIVE_SECOND    BarSize = "5 secs"
	FIFTEEN_SECOND BarSize = "15 secs"
	THIRTY_SECOND  BarSize = "30 secs"
	ONE_MINUTE     BarSize = "1 min"
	TWO_MINUTE     BarSize = "2 mins"
	THREE_MINUTE   BarSize = "3 mins"
	FOUR_MINUTE    BarSize = "4 mins"
	FIVE_MINUTE    BarSize = "5 mins"
	FIFTEEN_MINUTE BarSize = "15 mins"
	THIRTY_MINUTE  BarSize = "30 mins"
	ONE_HOUR       BarSize = "1 hour"
	ONE_DAY        BarSize = "1 day"
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
	StartTime  time.Time
	EndTime    time.Time
	BarSize    BarSize
	WhatToShow DisplayType
}

func (p *QueryParams) Duration() *Duration {
	res := &Duration{}
	diff := p.EndTime.Sub(p.StartTime)
	switch {
	case diff.Hours() < 24:
		res.N = int(diff.Seconds())
		res.T = SECOND
	case diff.Hours() < 168:
		res.N = int(math.Ceil(diff.Hours() / 24))
		res.T = DAY
	case diff.Hours() < 720:
		res.N = int(math.Ceil(diff.Hours() / 168))
		res.T = WEEK
	default:
		res.N = int(math.Ceil(diff.Hours() / 720))
		res.T = MONTH
	}
	return res
}
