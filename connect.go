package lconnect

// @author  Mikhail Kirillov <mikkirillov@yandex.ru>
// @version 1.000
// @date    2017-10-16

import (
	"fmt"
	"net"
	"time"

	"github.com/belfinor/log"
)

var nextId chan int64

func init() {

	nextId = make(chan int64, 32)

	go func() {

		id := int64(1)

		for {
			nextId <- id
			id++
		}

	}()
}

type connect struct {
	id        int64
	addr      string
	keepAlive time.Duration
	con       net.Conn
}

func (c *connect) Close() {
	if c.con != nil {
		log.Trace(fmt.Sprintf("connection #%d (%s) closed", c.id, c.addr))
		c.con.Close()
		c.con = nil
	}
}

func (c *connect) Write(p []byte) (int, error) {
	c.con.SetDeadline(time.Now().Add(c.keepAlive))
	return c.con.Write(p)
}

func (c *connect) Read(b []byte) (int, error) {
	c.con.SetDeadline(time.Now().Add(c.keepAlive))
	return c.con.Read(b)
}
