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
	Privileges   Role
}

type TWSClient struct {
	conf TWSConfig

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
		conf:   conf,
		logger: logger,
		status: &clientState{},
	}
	if cl.conf.ReadTimeout <= 0 {
		cl.conf.ReadTimeout = 10 * time.Second
	}
	if cl.conf.WriteTimeout <= 0 {
		cl.conf.WriteTimeout = 10 * time.Second
	}
	if cl.conf.Privileges.None() {
		cl.conf.Privileges = ReadOnly()
	}
	if cl.logger == nil {
		cl.logger = &log.EmptyLogger{}
	}
	return cl, nil
}

func (c *TWSClient) Connect() (err error) {
	conn := &net.TWSConn{
		ReadTimeout:  c.conf.ReadTimeout,
		WriteTimeout: c.conf.WriteTimeout,
	}
	defer func() {
		if err != nil {
			c.Disconnect()
		}
	}()
	addr := fmt.Sprintf("%v:%v", c.conf.Host, c.conf.Port)
	if err := conn.Open(addr); err != nil {
		return err
	}
	c.conn = conn
	res, err := read.ServerVersion(c.conn)
	if err != nil {
		return fmt.Errorf("failed to get server version: %v", err)
	}
	c.serverVersion = res.ServerVersion
	c.logger.Info("client=%d connected to %s ver=%d ts=%v",
		c.conf.ClientId, addr, c.serverVersion, time.Unix(res.Timestamp, 0))
	return nil
}

func (c *TWSClient) Disconnect() (err error) {
	defer func() {
		if err != nil {
			c.logger.Error("disconnect error: %v", err)
			return
		}
		c.logger.Info("client=%d disconnected", c.conf.ClientId)
	}()
	if c != nil &&
		c.conn != nil {
		return c.conn.Close()
	}
	return nil
}

func (c *TWSClient) ClientId() int32 {
	return c.conf.ClientId
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
