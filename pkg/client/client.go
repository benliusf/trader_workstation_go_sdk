package client

import (
	"fmt"
	"time"

	"github.com/benliusf/trader_workstation_go_sdk/pkg/log"
	"github.com/benliusf/trader_workstation_go_sdk/pkg/net"
	"github.com/benliusf/trader_workstation_go_sdk/pkg/read"
)

const (
	DEFAULT_READ_TIMEOUT  = 10 * time.Second
	DEFAULT_WRITE_TIMEOUT = 10 * time.Second
)

type TWSConfig struct {
	ClientID     int32
	Host         string
	Port         string
	ReadTimeout  time.Duration
	WriteTimeout time.Duration
}

type TWSClient struct {
	Conf *TWSConfig

	logger log.Logger

	conn *net.TWSConn

	ServerVersion int32

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
		Conf: &TWSConfig{
			ClientID:     conf.ClientID,
			Host:         conf.Host,
			Port:         conf.Port,
			ReadTimeout:  conf.ReadTimeout,
			WriteTimeout: conf.WriteTimeout,
		},
		logger: logger,
		status: &clientState{},
	}
	if cl.logger == nil {
		cl.logger = &log.EmptyLogger{}
	}
	if cl.Conf.ReadTimeout <= 0 {
		cl.Conf.ReadTimeout = DEFAULT_READ_TIMEOUT
	}
	if cl.Conf.WriteTimeout <= 0 {
		cl.Conf.WriteTimeout = DEFAULT_WRITE_TIMEOUT
	}
	return cl, nil
}

func (c *TWSClient) Connect() (err error) {
	conn := &net.TWSConn{
		ReadTimeout:  c.Conf.ReadTimeout,
		WriteTimeout: c.Conf.WriteTimeout,
	}
	defer func() {
		if err != nil {
			c.Disconnect()
		}
	}()
	addr := fmt.Sprintf("%v:%v", c.Conf.Host, c.Conf.Port)
	if err := conn.Open(addr); err != nil {
		return err
	}
	c.conn = conn
	res, err := read.ServerVersion(c.conn)
	if err != nil {
		return fmt.Errorf("failed to get server version: %v", err)
	}
	c.ServerVersion = res.ServerVersion
	c.logger.Info("client=%d connected to %v ver=%d ts=%v", c.Conf.ClientID, addr, c.ServerVersion, time.Unix(res.Timestamp, 0))
	return nil
}

func (c *TWSClient) Disconnect() error {
	c.logger.Info("client=%d disconnected", c.Conf.ClientID)
	if c != nil &&
		c.conn != nil {
		return c.conn.Close()
	}
	return nil
}

func (c *TWSClient) GetNextReqId() int32 {
	return c.status.getNextValidId()
}
