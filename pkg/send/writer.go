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
		msgId = START_API
	case *api.IdsRequest:
		msgId = REQ_IDS
	case *api.AccountSummaryRequest:
		msgId = REQ_ACCOUNT_SUMMARY
	case *api.AccountDataRequest:
		msgId = REQ_ACCT_DATA
	case *api.ContractDataRequest:
		msgId = REQ_CONTRACT_DATA
	case *api.MarketDataTypeRequest:
		msgId = REQ_MARKET_DATA_TYPE
	case *api.MarketDataRequest:
		msgId = REQ_MKT_DATA
	case *api.HistoricalDataRequest:
		msgId = REQ_HISTORICAL_DATA
	case *api.PositionsRequest:
		msgId = REQ_POSITIONS
	case *api.PlaceOrderRequest:
		msgId = PLACE_ORDER
	case *api.CancelOrderRequest:
		msgId = CANCEL_ORDER
	case *api.GlobalCancelRequest:
		msgId = REQ_GLOBAL_CANCEL
	case *api.OpenOrdersRequest:
		msgId = REQ_OPEN_ORDERS
	case *api.AllOpenOrdersRequest:
		msgId = REQ_ALL_OPEN_ORDERS
	case *api.ExecutionRequest:
		msgId = REQ_EXECUTIONS
	default:
		return fmt.Errorf("'%T' is not implemented", t)
	}
	msgId += PROTOBUF_MSG_ID
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
