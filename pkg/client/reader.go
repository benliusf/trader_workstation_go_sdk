package client

import (
	"context"
	"fmt"

	"github.com/benliusf/trader_workstation_go_sdk/pkg/log"
	"github.com/benliusf/trader_workstation_go_sdk/pkg/read"

	api "github.com/benliusf/trader_workstation_go_sdk/api/v104401"
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/reflect/protoreflect"
)

var deserialize func(m protoreflect.ProtoMessage, b []byte) error = func(m protoreflect.ProtoMessage, b []byte) error {
	if err := proto.Unmarshal(b, m); err != nil {
		return err
	}
	return nil
}

type EReader struct {
	twsClient *TWSClient
	logger    log.Logger
}

func NewReader(cl *TWSClient) (*EReader, error) {
	if cl == nil {
		return nil, fmt.Errorf("nil TWSClient")
	}
	return &EReader{
		cl, cl.logger,
	}, nil
}

func (e *EReader) Read(ctx context.Context, handler EHandler) error {
	for {
		select {
		case <-ctx.Done():
			e.logger.Debug("reader has stopped")
			return nil
		default:
			msg, err := read.Next(e.twsClient.conn)
			if err != nil {
				e.logger.Error("failed to read next message: %v", err)
				continue
			}
			if err := e.Process(msg, handler); err != nil {
				e.logger.Error("failed to process message: %v", err)
			}
		}
	}
}

func (e *EReader) Process(msg *read.Message, handler EHandler) error {
	id, err := msg.ReadMsgId()
	if err != nil {
		return err
	}
	b, err := msg.ReadBytes()
	if err != nil {
		return err
	}
	switch id {
	case read.NEXT_VALID_ID:
		return e.handleNextValidId(b, handler)
	case read.ACCOUNT_SUMMARY:
		return e.handleAccountSummary(b, handler)
	case read.ACCOUNT_SUMMARY_END:
		return e.handleAccountSummaryEnd(b, handler)
	case read.ACCT_VALUE:
		return e.handleAccountValue(b, handler)
	case read.ACCT_UPDATE_TIME:
		return e.handleAccountUpdateTime(b, handler)
	case read.ACCT_DOWNLOAD_END:
		return e.handleAccountDataEnd(b, handler)
	case read.CONTRACT_DATA:
		return e.handleContractData(b, handler)
	case read.CONTRACT_DATA_END:
		return e.handleContractDataEnd(b, handler)
	case read.TICK_PRICE:
		return e.handleTickPrice(b, handler)
	case read.TICK_SIZE:
		return e.handleTickSize(b, handler)
	case read.TICK_STRING:
		return e.handleTickString(b, handler)
	case read.HISTORICAL_DATA:
		return e.handleHistoricalData(b, handler)
	case read.HISTORICAL_DATA_END:
		return e.handleHistoricalDataEnd(b, handler)
	case read.POSITION_DATA:
		return e.handlePosition(b, handler)
	case read.POSITION_END:
		return e.handlePositionEnd(b, handler)
	case read.OPEN_ORDER:
		return e.handleOpenOrder(b, handler)
	case read.OPEN_ORDER_END:
		return e.handleOpenOrdersEnd(b, handler)
	case read.ORDER_STATUS:
		return e.handleOrderStatus(b, handler)
	case read.ERR_MSG:
		return e.handleErrorMessage(b, handler)
	default:
		return handler.Unsupported(msg)
	}
}

func (e *EReader) handleNextValidId(b []byte, handler EHandler) error {
	m := &api.NextValidId{}
	if err := deserialize(m, b); err != nil {
		return err
	}
	status := e.twsClient.status
	status.setNextOrderId(*m.OrderId)
	if !status.isReady() {
		status.setReady()
	}
	return handler.NextValidId(m)
}

func (e *EReader) handleAccountSummary(b []byte, handler EHandler) error {
	m := &api.AccountSummary{}
	if err := deserialize(m, b); err != nil {
		return err
	}
	return handler.AccountSummary(m)
}

func (e *EReader) handleAccountSummaryEnd(b []byte, handler EHandler) error {
	m := &api.AccountSummaryEnd{}
	if err := deserialize(m, b); err != nil {
		return err
	}
	return handler.AccountSummaryEnd(m)
}

func (e *EReader) handleAccountValue(b []byte, handler EHandler) error {
	m := &api.AccountValue{}
	if err := deserialize(m, b); err != nil {
		return err
	}
	return handler.AccountValue(m)
}

func (e *EReader) handleAccountUpdateTime(b []byte, handler EHandler) error {
	m := &api.AccountUpdateTime{}
	if err := deserialize(m, b); err != nil {
		return err
	}
	return handler.AccountUpdateTime(m)
}

func (e *EReader) handleAccountDataEnd(b []byte, handler EHandler) error {
	m := &api.AccountDataEnd{}
	if err := deserialize(m, b); err != nil {
		return err
	}
	return handler.AccountDataEnd(m)
}

func (e *EReader) handleContractData(b []byte, handler EHandler) error {
	m := &api.ContractData{}
	if err := deserialize(m, b); err != nil {
		return err
	}
	return handler.ContractData(m)
}

func (e *EReader) handleContractDataEnd(b []byte, handler EHandler) error {
	m := &api.ContractDataEnd{}
	if err := deserialize(m, b); err != nil {
		return err
	}
	return handler.ContractDataEnd(m)
}

func (e *EReader) handleTickPrice(b []byte, handler EHandler) error {
	m := &api.TickPrice{}
	if err := deserialize(m, b); err != nil {
		return err
	}
	return handler.TickPrice(m)
}

func (e *EReader) handleTickSize(b []byte, handler EHandler) error {
	m := &api.TickSize{}
	if err := deserialize(m, b); err != nil {
		return err
	}
	return handler.TickSize(m)
}

func (e *EReader) handleTickString(b []byte, handler EHandler) error {
	m := &api.TickString{}
	if err := deserialize(m, b); err != nil {
		return err
	}
	return handler.TickString(m)
}

func (e *EReader) handleHistoricalData(b []byte, handler EHandler) error {
	m := &api.HistoricalData{}
	if err := deserialize(m, b); err != nil {
		return err
	}
	return handler.HistoricalData(m)
}

func (e *EReader) handleHistoricalDataEnd(b []byte, handler EHandler) error {
	m := &api.HistoricalDataEnd{}
	if err := deserialize(m, b); err != nil {
		return err
	}
	return handler.HistoricalDataEnd(m)
}

func (e *EReader) handlePosition(b []byte, handler EHandler) error {
	m := &api.Position{}
	if err := deserialize(m, b); err != nil {
		return err
	}
	return handler.Position(m)
}

func (e *EReader) handlePositionEnd(b []byte, handler EHandler) error {
	m := &api.PositionEnd{}
	if err := deserialize(m, b); err != nil {
		return err
	}
	return handler.PositionEnd(m)
}

func (e *EReader) handleOpenOrder(b []byte, handler EHandler) error {
	m := &api.OpenOrder{}
	if err := deserialize(m, b); err != nil {
		return err
	}
	return handler.OpenOrder(m)
}

func (e *EReader) handleOpenOrdersEnd(b []byte, handler EHandler) error {
	m := &api.OpenOrdersEnd{}
	if err := deserialize(m, b); err != nil {
		return err
	}
	return handler.OpenOrdersEnd(m)
}

func (e *EReader) handleOrderStatus(b []byte, handler EHandler) error {
	m := &api.OrderStatus{}
	if err := deserialize(m, b); err != nil {
		return err
	}
	return handler.OrderStatus(m)
}

func (e *EReader) handleErrorMessage(b []byte, handler EHandler) error {
	m := &api.ErrorMessage{}
	if err := deserialize(m, b); err != nil {
		return err
	}
	return handler.ErrorMessage(m)
}
