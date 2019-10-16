package lconnect

// @author  Mikhail Kirillov <mikkirillov@yandex.ru>
// @version 1.001
// @date    2017-10-16

import (
	"bufio"
	"bytes"
	"fmt"
	"net"
	"time"

	"github.com/belfinor/log"
)

type Handler interface {
	OnMessage([]byte) ([]byte, error)
}

type HandlerMaker func() Handler

func Server(addr string, keepAlive time.Duration, maker HandlerMaker) error {
	ln, err := net.Listen("tcp", addr)

	if err != nil {
		return err
	}

	for {

		conn, err := ln.Accept()

		if err != nil {
			continue
		}

		con := &connect{
			addr:      conn.RemoteAddr().String(),
			keepAlive: keepAlive,
			con:       conn,
		}

		proto := maker()

		go serverHandler(con, proto)
	}

	return nil
}

func serverHandler(conn *connect, proto Handler) {

	log.Trace(fmt.Sprintf("connection #%d (%s) opened", conn.id, conn.addr))

	defer conn.Close()

	br := bufio.NewReaderSize(conn, 10240)
	bw := bufio.NewWriterSize(conn, 10240)

	for {

		data, err := br.ReadBytes('\n')

		size := len(data)

		if size == 0 {
			break
		}

		data = bytes.TrimRight(data, "\r\n")

		if len(data) == 0 {
			continue
		}

		data, err = proto.OnMessage(data)
		if err != nil {
			break
		}

		if len(data) > 0 {

			_, err = bw.Write(data)
			if err != nil {
				break
			}
			if bw.WriteByte('\n') != nil {
				break
			}
			if bw.Flush() != nil {
				break
			}
		}

	}
}
