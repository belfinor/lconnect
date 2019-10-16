package lconnect

// @author  Mikhail Kirillov <mikkirillov@yandex.ru>
// @version 1.000
// @date    2019-10-16

import (
	"bufio"
	"bytes"
	"net"
	"time"
)

type Client struct {
	addr        string
	keepAlive   time.Duration
	lastConnect int64
	con         *connect
	br          *bufio.Reader
	bw          *bufio.Writer
}

func NewClient(addr string, keepAlive time.Duration) *Client {

	client := &Client{
		addr:      addr,
		keepAlive: keepAlive,
	}

	return client
}

func (c *Client) connect() bool {

	if c.con != nil {
		return true
	}

	if c.lastConnect > time.Now().Unix()-5 {
		return false
	}

	con, err := net.DialTimeout("tcp", c.addr, time.Second*3)

	c.lastConnect = time.Now().Unix()

	if err != nil {
		return false
	}

	c.con = &connect{
		addr:      c.addr,
		id:        <-nextId,
		con:       con,
		keepAlive: c.keepAlive,
	}

	c.br = bufio.NewReader(c.con)
	c.bw = bufio.NewWriter(c.con)

	return true
}

func (c *Client) send(data []byte) bool {

	n, err := c.bw.Write(data)

	if err != nil || n != len(data) {
		return false
	}

	err = c.bw.WriteByte('\n')
	if err != nil {
		return false
	}

	return c.bw.Flush() == nil
}

func (c *Client) Write(data []byte) bool {

	data = bytes.TrimRight(data, "\r\n")

	for i := 0; i < 2; i++ {

		if !c.connect() {
			return false
		}

		if c.send(data) {
			return true
		}

		c.con.Close()
		c.con = nil
	}

	return false
}

func (c *Client) Read() ([]byte, bool) {

	if c.con == nil {
		return nil, false
	}

	data, err := c.br.ReadBytes('\n')
	if err != nil {
		c.con.Close()
		c.con = nil
		return nil, false
	}

	if len(data) == 0 {
		c.con.Close()
		c.con = nil
		return nil, false
	}

	data = bytes.TrimRight(data, "\r\n")
	return data, true
}

func (c *Client) Close() {
	if c.con != nil {
		c.con.Close()
		c.con = nil
	}
}
