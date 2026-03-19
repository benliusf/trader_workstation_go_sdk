package client

import (
	"context"
	"fmt"
	"strings"
	"time"

	api "github.com/benliusf/trader_workstation_go_sdk/api/v104401"
	"google.golang.org/protobuf/proto"
)

type apiRequest struct {
	sender *ESender
	proto  proto.Message
}

func (r *apiRequest) Send(ctx context.Context) error {
	if err := ctx.Err(); err != nil {
		return err
	}
	return r.sender.Send(ctx, r.proto)
}

// NextValidId.proto

type NextValidIdRequest struct {
	*apiRequest
}

func NewNextValidIdRequest(s *ESender) *NextValidIdRequest {
	numIds := int32(-1)
	return &NextValidIdRequest{
		apiRequest: &apiRequest{
			sender: s,
			proto: &api.IdsRequest{
				NumIds: &numIds,
			},
		},
	}
}

// AccountSummaryRequest.proto

const All string = "All"

type AccountSummaryTag string

const (
	AccountType                    AccountSummaryTag = "AccountType"
	NetLiquidation                 AccountSummaryTag = "NetLiquidation"
	TotalCashValue                 AccountSummaryTag = "TotalCashValue"
	SettledCash                    AccountSummaryTag = "SettledCash"
	AccruedCash                    AccountSummaryTag = "AccruedCash"
	BuyingPower                    AccountSummaryTag = "BuyingPower"
	EquityWithLoanValue            AccountSummaryTag = "EquityWithLoanValue"
	PreviousDayEquityWithLoanValue AccountSummaryTag = "PreviousDayEquityWithLoanValue"
	GrossPositionValue             AccountSummaryTag = "GrossPositionValue"
	ReqTEquity                     AccountSummaryTag = "ReqTEquity"
	ReqTMargin                     AccountSummaryTag = "ReqTMargin"
	SMA                            AccountSummaryTag = "SMA"
	InitMarginReq                  AccountSummaryTag = "InitMarginReq"
	MaintMarginReq                 AccountSummaryTag = "MaintMarginReq"
	AvailableFunds                 AccountSummaryTag = "AvailableFunds"
	ExcessLiquidity                AccountSummaryTag = "ExcessLiquidity"
	Cushion                        AccountSummaryTag = "Cushion"
	FullInitMarginReq              AccountSummaryTag = "FullInitMarginReq"
	FullMaintMarginReq             AccountSummaryTag = "FullMaintMarginReq"
	FullAvailableFunds             AccountSummaryTag = "FullAvailableFunds"
	FullExcessLiquidity            AccountSummaryTag = "FullExcessLiquidity"
	LookAheadNextChange            AccountSummaryTag = "LookAheadNextChange"
	LookAheadInitMarginReq         AccountSummaryTag = "LookAheadInitMarginReq"
	LookAheadMaintMarginReq        AccountSummaryTag = "LookAheadMaintMarginReq"
	LookAheadAvailableFunds        AccountSummaryTag = "LookAheadAvailableFunds"
	LookAheadExcessLiquidity       AccountSummaryTag = "LookAheadExcessLiquidity"
	HighestSeverity                AccountSummaryTag = "HighestSeverity"
	DayTradesRemaining             AccountSummaryTag = "DayTradesRemaining"
	Leverage                       AccountSummaryTag = "Leverage"
)

var AllTags = []AccountSummaryTag{
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
	*apiRequest
}

func NewAccountSummaryRequest(s *ESender, reqId int32, group string, tags []AccountSummaryTag) *AccountSummaryRequest {
	if group == "" {
		group = All
	}
	if tags == nil || len(tags) == 0 {
		tags = AllTags
	}
	tmp := make([]string, len(tags))
	for i, t := range tags {
		tmp[i] = string(t)
	}
	tagsConcat := strings.Join(tmp, ",")
	return &AccountSummaryRequest{
		apiRequest: &apiRequest{
			sender: s,
			proto: &api.AccountSummaryRequest{
				ReqId: &reqId,
				Group: &group,
				Tags:  &tagsConcat,
			},
		},
	}
}

// ContractDataRequest.proto

type ContractDataRequest struct {
	*apiRequest
}

func NewContractDataRequest(s *ESender, reqId int32, symb *Symbol) *ContractDataRequest {
	return &ContractDataRequest{
		apiRequest: &apiRequest{
			sender: s,
			proto: &api.ContractDataRequest{
				ReqId: &reqId,
				Contract: &api.Contract{
					Symbol:      &symb.Ticker,
					SecType:     strPtr(string(symb.Type)),
					Exchange:    strPtr(string(symb.Exch)),
					PrimaryExch: strPtr(string(symb.PrimExch)),
					Currency:    strPtr(string(symb.Curr)),
				},
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
	*apiRequest
}

func NewMarketDataTypeRequest(s *ESender, l MarketDataLevel) *MarketDataRequest {
	return &MarketDataRequest{
		apiRequest: &apiRequest{
			sender: s,
			proto: &api.MarketDataTypeRequest{
				MarketDataType: int32Ptr(int32(l)),
			},
		},
	}
}

// MarketDataRequest.proto

type MarketDataRequest struct {
	*apiRequest
}

func NewMarketDataRequest(s *ESender, reqId int32, symb *Symbol) *MarketDataRequest {
	return &MarketDataRequest{
		apiRequest: &apiRequest{
			sender: s,
			proto: &api.MarketDataRequest{
				ReqId: &reqId,
				Contract: &api.Contract{
					Symbol:      &symb.Ticker,
					SecType:     strPtr(string(symb.Type)),
					Exchange:    strPtr(string(symb.Exch)),
					PrimaryExch: strPtr(string(symb.PrimExch)),
					Currency:    strPtr(string(symb.Curr)),
				},
				Snapshot:           boolPtr(false),
				RegulatorySnapshot: boolPtr(false),
			},
		},
	}
}

// HistoricalDataRequest.proto

type HistoricalDataRequest struct {
	*apiRequest
	params *QueryParams
}

func NewHistoricalDataRequest(s *ESender, reqId int32, symb *Symbol, params *QueryParams) *HistoricalDataRequest {
	const tsFormat = "20060102 15:04:05 US/Eastern"
	loc, _ := time.LoadLocation("US/Eastern")
	return &HistoricalDataRequest{
		apiRequest: &apiRequest{
			sender: s,
			proto: &api.HistoricalDataRequest{
				ReqId: &reqId,
				Contract: &api.Contract{
					Symbol:      &symb.Ticker,
					SecType:     strPtr(string(symb.Type)),
					Exchange:    strPtr(string(symb.Exch)),
					PrimaryExch: strPtr(string(symb.PrimExch)),
					Currency:    strPtr(string(symb.Curr)),
				},
				EndDateTime:    strPtr(params.EndTime.In(loc).Format(tsFormat)),
				Duration:       strPtr(params.Duration().String()),
				BarSizeSetting: strPtr(string(params.BarSize)),
				WhatToShow:     strPtr(string(params.WhatToShow)),
				UseRTH:         boolPtr(true),
				FormatDate:     int32Ptr(1),
				KeepUpToDate:   boolPtr(false),
			},
		},
		params: params,
	}
}

func (r *HistoricalDataRequest) Send(ctx context.Context) error {
	now := time.Now()
	if r.params != nil {
		if r.params.StartTime.Before(now.Add(-720 * 6 * time.Hour)) {
			return fmt.Errorf("%w: start time out of range", ErrInvalidParam)
		}
		timeRange := r.params.EndTime.Sub(r.params.StartTime)
		if timeRange.Hours() > (7 * 24) {
			return fmt.Errorf("%w: time range cannot exceed one week", ErrInvalidParam)
		}
	}
	return r.apiRequest.Send(ctx)
}
