package client

import (
	"context"
	"fmt"
	"strings"
	"time"

	api "github.com/benliusf/trader_workstation_go_sdk/api/v104401"
	"google.golang.org/protobuf/proto"
)

type apiRequest[T proto.Message] struct {
	sender *ESender
	proto  T
}

func (r *apiRequest[T]) Send(ctx context.Context) (int32, error) {
	return r.sender.Send(ctx, r.proto)
}

// NextValidId.proto

type NextValidIdRequest struct {
	*apiRequest[*api.IdsRequest]
}

func NewNextValidIdRequest(s *ESender) *NextValidIdRequest {
	numIds := int32(-1)
	return &NextValidIdRequest{
		apiRequest: &apiRequest[*api.IdsRequest]{
			sender: s,
			proto: &api.IdsRequest{
				NumIds: &numIds,
			},
		},
	}
}

func (r *NextValidIdRequest) Send(ctx context.Context) error {
	if _, err := r.apiRequest.Send(ctx); err != nil {
		return err
	}
	return nil
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
	*apiRequest[*api.AccountSummaryRequest]
}

func NewAccountSummaryRequest(s *ESender, group string, tags []AccountSummaryTag) *AccountSummaryRequest {
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
		apiRequest: &apiRequest[*api.AccountSummaryRequest]{
			sender: s,
			proto: &api.AccountSummaryRequest{
				Group: &group,
				Tags:  &tagsConcat,
			},
		},
	}
}

// AccountDataRequest.proto

type AccountDataRequest struct {
	*apiRequest[*api.AccountDataRequest]
}

func NewAccountDataRequest(s *ESender, accountId string) *AccountDataRequest {
	return &AccountDataRequest{
		apiRequest: &apiRequest[*api.AccountDataRequest]{
			sender: s,
			proto: &api.AccountDataRequest{
				AcctCode:  strPtr(accountId),
				Subscribe: boolPtr(true),
			},
		},
	}
}

func (r *AccountDataRequest) Send(ctx context.Context) error {
	if _, err := r.apiRequest.Send(ctx); err != nil {
		return err
	}
	return nil
}

type CancelAccountDataRequest struct {
	*apiRequest[*api.AccountDataRequest]
}

func NewCancelAccountDataRequest(s *ESender, accountId string) *AccountDataRequest {
	return &AccountDataRequest{
		apiRequest: &apiRequest[*api.AccountDataRequest]{
			sender: s,
			proto: &api.AccountDataRequest{
				AcctCode:  strPtr(accountId),
				Subscribe: boolPtr(false),
			},
		},
	}
}

func (r *CancelAccountDataRequest) Send(ctx context.Context) error {
	if _, err := r.apiRequest.Send(ctx); err != nil {
		return err
	}
	return nil
}

// ContractDataRequest.proto

type ContractDataRequest struct {
	*apiRequest[*api.ContractDataRequest]
}

func NewContractDataRequest(s *ESender, contr *api.Contract) *ContractDataRequest {
	return &ContractDataRequest{
		apiRequest: &apiRequest[*api.ContractDataRequest]{
			sender: s,
			proto: &api.ContractDataRequest{
				Contract: contr,
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
	*apiRequest[*api.MarketDataTypeRequest]
}

func NewMarketDataTypeRequest(s *ESender, l MarketDataLevel) *MarketDataTypeRequest {
	return &MarketDataTypeRequest{
		apiRequest: &apiRequest[*api.MarketDataTypeRequest]{
			sender: s,
			proto: &api.MarketDataTypeRequest{
				MarketDataType: int32Ptr(int32(l)),
			},
		},
	}
}

func (r *MarketDataTypeRequest) Send(ctx context.Context) error {
	if _, err := r.apiRequest.Send(ctx); err != nil {
		return err
	}
	return nil
}

// MarketDataRequest.proto

type MarketDataRequest struct {
	*apiRequest[*api.MarketDataRequest]
}

func NewMarketDataRequest(s *ESender, contr *api.Contract) *MarketDataRequest {
	return &MarketDataRequest{
		apiRequest: &apiRequest[*api.MarketDataRequest]{
			sender: s,
			proto: &api.MarketDataRequest{
				Contract:           contr,
				Snapshot:           boolPtr(false),
				RegulatorySnapshot: boolPtr(false),
			},
		},
	}
}

// HistoricalDataRequest.proto

type HistoricalDataRequest struct {
	*apiRequest[*api.HistoricalDataRequest]
	params *QueryParams
}

func NewHistoricalDataRequest(s *ESender, contr *api.Contract, params *QueryParams) *HistoricalDataRequest {
	const tsFormat = "20060102 15:04:05 US/Eastern"
	loc, _ := time.LoadLocation("US/Eastern")
	return &HistoricalDataRequest{
		apiRequest: &apiRequest[*api.HistoricalDataRequest]{
			sender: s,
			proto: &api.HistoricalDataRequest{
				Contract:       contr,
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

func (r *HistoricalDataRequest) Send(ctx context.Context) (int32, error) {
	now := time.Now()
	if r.params != nil {
		if r.params.StartTime.Before(now.Add(-720 * 6 * time.Hour)) {
			return -1, fmt.Errorf("%w: start time out of range", ErrInvalidParam)
		}
		timeRange := r.params.EndTime.Sub(r.params.StartTime)
		if timeRange.Hours() > (7 * 24) {
			return -1, fmt.Errorf("%w: time range cannot exceed one week", ErrInvalidParam)
		}
	}
	return r.apiRequest.Send(ctx)
}

// PositionsRequest.proto

type PositionsRequest struct {
	*apiRequest[*api.PositionsRequest]
}

func NewPositionsRequest(s *ESender) *PositionsRequest {
	return &PositionsRequest{
		&apiRequest[*api.PositionsRequest]{
			sender: s,
			proto:  &api.PositionsRequest{},
		},
	}
}

func (r *PositionsRequest) Send(ctx context.Context) error {
	if _, err := r.apiRequest.Send(ctx); err != nil {
		return err
	}
	return nil
}

// PlaceOrderRequest.proto

type PlaceOrderRequest struct {
	*apiRequest[*api.PlaceOrderRequest]
}

func NewPlaceOrderRequest(s *ESender, contr *api.Contract, order *api.Order) *PlaceOrderRequest {
	return &PlaceOrderRequest{
		&apiRequest[*api.PlaceOrderRequest]{
			sender: s,
			proto: &api.PlaceOrderRequest{
				Contract: contr,
				Order:    order,
			},
		},
	}
}

// CancelOrderRequest.proto

type CancelOrderRequest struct {
	*apiRequest[*api.CancelOrderRequest]
}

func NewCancelOrderRequest(s *ESender, orderId int32) *CancelOrderRequest {
	return &CancelOrderRequest{
		&apiRequest[*api.CancelOrderRequest]{
			sender: s,
			proto: &api.CancelOrderRequest{
				OrderId: &orderId,
			},
		},
	}
}

func (r *CancelOrderRequest) Send(ctx context.Context) error {
	if _, err := r.apiRequest.Send(ctx); err != nil {
		return err
	}
	return nil
}

// GlobalCancelRequest.proto

type GlobalCancelRequest struct {
	apiRequest *apiRequest[*api.GlobalCancelRequest]
}

func NewGlobalCancelRequest(s *ESender) *GlobalCancelRequest {
	return &GlobalCancelRequest{
		&apiRequest[*api.GlobalCancelRequest]{
			sender: s,
			proto:  &api.GlobalCancelRequest{},
		},
	}
}

func (r *GlobalCancelRequest) Send(ctx context.Context) error {
	if _, err := r.apiRequest.Send(ctx); err != nil {
		return err
	}
	return nil
}

// OpenOrdersRequest.proto

type OpenOrdersRequest struct {
	*apiRequest[*api.OpenOrdersRequest]
}

func NewOpenOrdersRequest(s *ESender) *OpenOrdersRequest {
	return &OpenOrdersRequest{
		&apiRequest[*api.OpenOrdersRequest]{
			sender: s,
			proto:  &api.OpenOrdersRequest{},
		},
	}
}

func (r *OpenOrdersRequest) Send(ctx context.Context) error {
	if _, err := r.apiRequest.Send(ctx); err != nil {
		return err
	}
	return nil
}

// AllOpenOrdersRequest.proto

type AllOpenOrdersRequest struct {
	*apiRequest[*api.AllOpenOrdersRequest]
}

func NewAllOpenOrdersRequest(s *ESender) *AllOpenOrdersRequest {
	return &AllOpenOrdersRequest{
		&apiRequest[*api.AllOpenOrdersRequest]{
			sender: s,
			proto:  &api.AllOpenOrdersRequest{},
		},
	}
}

func (r *AllOpenOrdersRequest) Send(ctx context.Context) error {
	if _, err := r.apiRequest.Send(ctx); err != nil {
		return err
	}
	return nil
}

// ExecutionRequest.proto

type ExecutionRequest struct {
	*apiRequest[*api.ExecutionRequest]
}

func NewExecutionRequest(s *ESender) *ExecutionRequest {
	return &ExecutionRequest{
		&apiRequest[*api.ExecutionRequest]{
			sender: s,
			proto:  &api.ExecutionRequest{},
		},
	}
}
