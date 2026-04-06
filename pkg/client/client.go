package client

import (
	"fmt"
	"time"

	"github.com/benliusf/trader_workstation_go_sdk/pkg/log"
	"github.com/benliusf/trader_workstation_go_sdk/pkg/net"
	"github.com/benliusf/trader_workstation_go_sdk/pkg/read"
)

type TWSConfig struct {
	ClientId     int32
	Host         string
	Port         string
	ReadTimeout  time.Duration
	WriteTimeout time.Duration
	Privileges   *Role
}

type TWSClient struct {
	clientId     int32
	host         string
	port         string
	readTimeout  time.Duration
	writeTimeout time.Duration
	privileges   *Role

	logger log.Logger

	conn *net.TWSConn

	serverVersion int32

	status *clientState
}

func NewClient(conf TWSConfig, logger log.Logger) (*TWSClient, error) {
	if conf.Host == "" {
		return nil, fmt.Errorf("host not defined!")
	}
	if conf.Port == "" {
		return nil, fmt.Errorf("port not defined!")
	}
	cl := &TWSClient{
		clientId:     conf.ClientId,
		host:         conf.Host,
		port:         conf.Port,
		readTimeout:  conf.ReadTimeout,
		writeTimeout: conf.WriteTimeout,
		privileges:   conf.Privileges,
		logger:       logger,
		status:       &clientState{},
	}
	if cl.readTimeout <= 0 {
		cl.readTimeout = 10 * time.Second
	}
	if cl.writeTimeout <= 0 {
		cl.writeTimeout = 10 * time.Second
	}
	if cl.privileges == nil {
		tmp := ReadOnly()
		cl.privileges = &tmp
	}
	if cl.logger == nil {
		cl.logger = &log.EmptyLogger{}
	}
	return cl, nil
}

func (c *TWSClient) Connect() (err error) {
	conn := &net.TWSConn{
		ReadTimeout:  c.readTimeout,
		WriteTimeout: c.writeTimeout,
	}
	defer func() {
		if err != nil {
			c.Disconnect()
		}
	}()
	addr := fmt.Sprintf("%v:%v", c.host, c.port)
	if err := conn.Open(addr); err != nil {
		return err
	}
	c.conn = conn
	res, err := read.ServerVersion(c.conn)
	if err != nil {
		return fmt.Errorf("failed to get server version: %v", err)
	}
	c.serverVersion = res.ServerVersion
	c.logger.Info("client=%d connected to %v ver=%d ts=%v", c.clientId, addr, c.serverVersion, time.Unix(res.Timestamp, 0))
	return nil
}

func (c *TWSClient) Disconnect() (err error) {
	defer func() {
		if err != nil {
			c.logger.Error("disconnect error: %v", err)
			return
		}
		c.logger.Info("client=%d disconnected", c.clientId)
	}()
	if c != nil &&
		c.conn != nil {
		return c.conn.Close()
	}
	return nil
}

func (c *TWSClient) ClientId() int32 {
	return c.clientId
}

func (c *TWSClient) ServerVersion() int32 {
	return c.serverVersion
}

func (c *TWSClient) NextReqId() int32 {
	return c.status.getNextReqId()
}

func (c *TWSClient) NextOrderId() int32 {
	return c.status.getNextOrderId()
}
