package client

import (
	"context"
	"strings"
	"time"

	api "github.com/benliusf/trader_workstation_go_sdk/api/v104401"
	"google.golang.org/protobuf/proto"
)

type request struct {
	sender *ESender
	proto  proto.Message
}

func (r *request) Send(ctx context.Context) error {
	if err := ctx.Err(); err != nil {
		return err
	}
	return r.sender.Send(ctx, r.proto)
}

// NextValidId.proto

type NextValidIdRequest struct {
	*request
}

func NewNextValidIdRequest(s *ESender) *NextValidIdRequest {
	numIds := int32(-1)
	return &NextValidIdRequest{
		request: &request{
			sender: s,
			proto: &api.IdsRequest{
				NumIds: &numIds,
			},
		},
	}
}

// AccountSummaryRequest.proto

const All string = "All"
const (
	AccountType                    = "AccountType"
	NetLiquidation                 = "NetLiquidation"
	TotalCashValue                 = "TotalCashValue"
	SettledCash                    = "SettledCash"
	AccruedCash                    = "AccruedCash"
	BuyingPower                    = "BuyingPower"
	EquityWithLoanValue            = "EquityWithLoanValue"
	PreviousDayEquityWithLoanValue = "PreviousDayEquityWithLoanValue"
	GrossPositionValue             = "GrossPositionValue"
	ReqTEquity                     = "ReqTEquity"
	ReqTMargin                     = "ReqTMargin"
	SMA                            = "SMA"
	InitMarginReq                  = "InitMarginReq"
	MaintMarginReq                 = "MaintMarginReq"
	AvailableFunds                 = "AvailableFunds"
	ExcessLiquidity                = "ExcessLiquidity"
	Cushion                        = "Cushion"
	FullInitMarginReq              = "FullInitMarginReq"
	FullMaintMarginReq             = "FullMaintMarginReq"
	FullAvailableFunds             = "FullAvailableFunds"
	FullExcessLiquidity            = "FullExcessLiquidity"
	LookAheadNextChange            = "LookAheadNextChange"
	LookAheadInitMarginReq         = "LookAheadInitMarginReq"
	LookAheadMaintMarginReq        = "LookAheadMaintMarginReq"
	LookAheadAvailableFunds        = "LookAheadAvailableFunds"
	LookAheadExcessLiquidity       = "LookAheadExcessLiquidity"
	HighestSeverity                = "HighestSeverity"
	DayTradesRemaining             = "DayTradesRemaining"
	Leverage                       = "Leverage"
)

var AllTags = []string{
	AccountType,
	NetLiquidation,
	TotalCashValue,
	SettledCash,
	AccruedCash,
	BuyingPower,
	EquityWithLoanValue,
	PreviousDayEquityWithLoanValue,
	GrossPositionValue,
	ReqTEquity,
	ReqTMargin,
	SMA,
	InitMarginReq,
	MaintMarginReq,
	AvailableFunds,
	ExcessLiquidity,
	Cushion,
	FullInitMarginReq,
	FullMaintMarginReq,
	FullAvailableFunds,
	FullExcessLiquidity,
	LookAheadNextChange,
	LookAheadInitMarginReq,
	LookAheadMaintMarginReq,
	LookAheadAvailableFunds,
	LookAheadExcessLiquidity,
	HighestSeverity,
	DayTradesRemaining,
	Leverage,
}

type AccountSummaryRequest struct {
	*request
}

func NewAccountSummaryRequest(s *ESender, reqId int32, group, tags string) *AccountSummaryRequest {
	if group == "" {
		group = All
	}
	if tags == "" {
		tags = strings.Join(AllTags, ",")
	}
	m := &api.AccountSummaryRequest{
		ReqId: &reqId,
		Group: &group,
		Tags:  &tags,
	}
	return &AccountSummaryRequest{
		request: &request{
			sender: s,
			proto:  m,
		},
	}
}

// ContractDataRequest.proto

type ContractDataRequest struct {
	*request
}

func NewContractDataRequest(s *ESender, reqId int32, symb *Symbol) *ContractDataRequest {
	secType := string(symb.Type)
	exch := string(symb.Exch)
	primExch := string(symb.PrimExch)
	curr := string(symb.Curr)
	c := &api.Contract{
		Symbol:      &symb.Ticker,
		SecType:     &secType,
		Exchange:    &exch,
		PrimaryExch: &primExch,
		Currency:    &curr,
	}
	return &ContractDataRequest{
		request: &request{
			sender: s,
			proto: &api.ContractDataRequest{
				ReqId:    &reqId,
				Contract: c,
			},
		},
	}
}

// MarketDataTypeRequest.proto

type MarketDataLevel int32

const (
	MARKET_DATA_LIVE           MarketDataLevel = 1
	MARKET_DATA_FROZEN         MarketDataLevel = 2
	MARKET_DATA_DELAYED        MarketDataLevel = 3
	MARKET_DATA_DELAYED_FROZEN MarketDataLevel = 4
)

type MarketDataTypeRequest struct {
	*request
}

func NewMarketDataTypeRequest(s *ESender, l MarketDataLevel) *MarketDataRequest {
	return &MarketDataRequest{
		request: &request{
			sender: s,
			proto: &api.MarketDataTypeRequest{
				MarketDataType: int32Ptr(int32(l)),
			},
		},
	}
}

// MarketDataRequest.proto

type MarketDataRequest struct {
	*request
}

func NewMarketDataRequest(s *ESender, reqId int32, symb *Symbol) *MarketDataRequest {
	secType := string(symb.Type)
	exch := string(symb.Exch)
	primExch := string(symb.PrimExch)
	curr := string(symb.Curr)
	c := &api.Contract{
		Symbol:      &symb.Ticker,
		SecType:     &secType,
		Exchange:    &exch,
		PrimaryExch: &primExch,
		Currency:    &curr,
	}

	return &MarketDataRequest{
		request: &request{
			sender: s,
			proto: &api.MarketDataRequest{
				ReqId:              &reqId,
				Contract:           c,
				Snapshot:           boolPtr(false),
				RegulatorySnapshot: boolPtr(false),
			},
		},
	}
}

// HistoricalDataRequest.proto

type HistoricalDataRequest struct {
	*request
}

func NewHistoricalDataRequest(s *ESender, reqId int32, symb *Symbol, param *QueryParams) *MarketDataRequest {
	const tsFormat = "20060102 15:04:05 US/Eastern"
	loc, _ := time.LoadLocation("US/Eastern")

	secType := string(symb.Type)
	exch := string(symb.Exch)
	primExch := string(symb.PrimExch)
	curr := string(symb.Curr)
	c := &api.Contract{
		Symbol:      &symb.Ticker,
		SecType:     &secType,
		Exchange:    &exch,
		PrimaryExch: &primExch,
		Currency:    &curr,
	}
	return &MarketDataRequest{
		request: &request{
			sender: s,
			proto: &api.HistoricalDataRequest{
				ReqId:          &reqId,
				Contract:       c,
				EndDateTime:    stringPtr(param.EndTime.In(loc).Format(tsFormat)),
				Duration:       stringPtr(string(param.Duration)),
				BarSizeSetting: stringPtr(string(param.BarSize)),
				WhatToShow:     stringPtr(string(param.WhatToShow)),
				UseRTH:         boolPtr(true),
				FormatDate:     int32Ptr(1),
				KeepUpToDate:   boolPtr(false),
			},
		},
	}
}
