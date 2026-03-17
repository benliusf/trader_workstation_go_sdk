package send

import (
	"fmt"

	api "github.com/benliusf/trader_workstation_go_sdk/api/v104401"
	"github.com/benliusf/trader_workstation_go_sdk/pkg/net"
	"google.golang.org/protobuf/proto"
)

func Write(conn *net.TWSConn, m proto.Message) error {
	if m == nil {
		return fmt.Errorf("nil message")
	}
	var msgId int32
	switch t := m.(type) {
	case *api.StartApiRequest:
		msgId = START_API + PROTOBUF_MSG_ID
	case *api.IdsRequest:
		msgId = REQ_IDS + PROTOBUF_MSG_ID
	case *api.AccountSummaryRequest:
		msgId = REQ_ACCOUNT_SUMMARY + PROTOBUF_MSG_ID
	case *api.ContractDataRequest:
		msgId = REQ_CONTRACT_DATA + PROTOBUF_MSG_ID
	case *api.MarketDataTypeRequest:
		msgId = REQ_MARKET_DATA_TYPE + PROTOBUF_MSG_ID
	case *api.MarketDataRequest:
		msgId = REQ_MKT_DATA + PROTOBUF_MSG_ID
	case *api.HistoricalDataRequest:
		msgId = REQ_HISTORICAL_DATA + PROTOBUF_MSG_ID
	default:
		return fmt.Errorf("'%T' is not implemented", t)
	}
	msgBytes, err := proto.Marshal(m)
	if err != nil {
		return err
	}
	bd := newBuilder()
	bd.withMsgId(msgId).withMsgBytes(msgBytes)
	b, err := bd.build()
	if err != nil {
		return err
	}
	if _, err := conn.Write(b); err != nil {
		return err
	}
	return nil
}
