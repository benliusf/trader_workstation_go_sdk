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
	var deserialize func(m protoreflect.ProtoMessage, b []byte) error = func(m protoreflect.ProtoMessage, b []byte) error {
		if err := proto.Unmarshal(b, m); err != nil {
			return err
		}
		return nil
	}
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
		m := &api.NextValidId{}
		if err := deserialize(m, b); err != nil {
			return err
		}
		e.twsClient.status.setNextValidId(*m.OrderId)
		if !e.twsClient.status.isReady() {
			e.twsClient.status.setReady()
		}
		return handler.NextValidId(m)
	case read.ACCOUNT_SUMMARY:
		m := &api.AccountSummary{}
		if err := deserialize(m, b); err != nil {
			return err
		}
		return handler.AccountSummary(m)
	case read.CONTRACT_DATA:
		m := &api.ContractData{}
		if err := deserialize(m, b); err != nil {
			return err
		}
		return handler.ContractData(m)
	case read.TICK_PRICE:
		m := &api.TickPrice{}
		if err := deserialize(m, b); err != nil {
			return err
		}
		return handler.TickPrice(m)
	case read.TICK_SIZE:
		m := &api.TickSize{}
		if err := deserialize(m, b); err != nil {
			return err
		}
		return handler.TickSize(m)
	case read.TICK_STRING:
		m := &api.TickString{}
		if err := deserialize(m, b); err != nil {
			return err
		}
		return handler.TickString(m)
	case read.HISTORICAL_DATA:
		m := &api.HistoricalData{}
		if err := deserialize(m, b); err != nil {
			return err
		}
		return handler.HistoricalData(m)
	case read.HISTORICAL_DATA_END:
		m := &api.HistoricalDataEnd{}
		if err := deserialize(m, b); err != nil {
			return err
		}
		return handler.HistoricalDataEnd(m)
	case read.ERR_MSG:
		m := &api.ErrorMessage{}
		if err := deserialize(m, b); err != nil {
			return err
		}
		return handler.ErrorMessage(m)
	default:
		return handler.Unsupported(msg)
	}
}
