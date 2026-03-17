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
				e.logger.Error("%v", err)
				continue
			}
			if err := e.Process(msg, handler); err != nil {
				e.logger.Error("%v", err)
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
	case read.CONTRACT_DATA:
		return e.handleContractData(b, handler)
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
	e.twsClient.status.setNextValidId(*m.OrderId)
	if !e.twsClient.status.isReady() {
		e.twsClient.status.setReady()
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

func (e *EReader) handleContractData(b []byte, handler EHandler) error {
	m := &api.ContractData{}
	if err := deserialize(m, b); err != nil {
		return err
	}
	return handler.ContractData(m)
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

func (e *EReader) handleErrorMessage(b []byte, handler EHandler) error {
	m := &api.ErrorMessage{}
	if err := deserialize(m, b); err != nil {
		return err
	}
	return handler.ErrorMessage(m)
}
