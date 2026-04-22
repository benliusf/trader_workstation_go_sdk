package send

import (
	"fmt"
	"reflect"

	api "github.com/benliusf/trader_workstation_go_sdk/api/v104401"
	"github.com/benliusf/trader_workstation_go_sdk/pkg/net"
	"google.golang.org/protobuf/proto"
)

var msgIdMap map[reflect.Type]int32 = map[reflect.Type]int32{}

func init() {
	ids := map[proto.Message]int32{
		&api.StartApiRequest{}:       START_API,
		&api.IdsRequest{}:            REQ_IDS,
		&api.AccountSummaryRequest{}: REQ_ACCOUNT_SUMMARY,
		&api.AccountDataRequest{}:    REQ_ACCT_DATA,
		&api.ContractDataRequest{}:   REQ_CONTRACT_DATA,
		&api.MarketDataTypeRequest{}: REQ_MARKET_DATA_TYPE,
		&api.MarketDataRequest{}:     REQ_MKT_DATA,
		&api.HeadTimestampRequest{}:  REQ_HEAD_TIMESTAMP,
		&api.HistoricalDataRequest{}: REQ_HISTORICAL_DATA,
		&api.PositionsRequest{}:      REQ_POSITIONS,
		&api.PlaceOrderRequest{}:     PLACE_ORDER,
		&api.CancelOrderRequest{}:    CANCEL_ORDER,
		&api.GlobalCancelRequest{}:   REQ_GLOBAL_CANCEL,
		&api.OpenOrdersRequest{}:     REQ_OPEN_ORDERS,
		&api.AllOpenOrdersRequest{}:  REQ_ALL_OPEN_ORDERS,
		&api.ExecutionRequest{}:      REQ_EXECUTIONS,
	}
	for k, v := range ids {
		msgIdMap[reflect.TypeOf(k)] = v
	}
}

func Write(conn *net.TWSConn, m proto.Message) error {
	if m == nil {
		return fmt.Errorf("nil message")
	}
	t := reflect.TypeOf(m)
	msgId, ok := msgIdMap[t]
	if !ok {
		return fmt.Errorf("'%T' is not implemented", m)
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
