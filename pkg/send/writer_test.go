package send

import (
	"reflect"
	"testing"

	api "github.com/benliusf/trader_workstation_go_sdk/api/v104401"
	"github.com/benliusf/trader_workstation_go_sdk/pkg/net"
	"github.com/stretchr/testify/assert"
)

func TestReqIds(t *testing.T) {
	id, ok := msgIdMap[reflect.TypeOf(&api.StartApiRequest{})]
	assert.True(t, ok)
	assert.Equal(t, START_API, id)
}

func TestWriter(t *testing.T) {
	conn := net.NewMockConn()
	assert.NoError(t, Write(conn, &api.StartApiRequest{}))
	assert.Error(t, Write(conn, &api.ApiConfig{}))
}
