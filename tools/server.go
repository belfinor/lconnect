// +build ignore

package main

import (
	"errors"
	"time"

	"github.com/belfinor/lconnect"
)

type Echo struct {
}

func (e *Echo) OnMessage(data []byte) ([]byte, error) {

	if string(data) == "quit" {
		return nil, errors.New("bye")
	}

	return data, nil
}

func main() {

	if err := lconnect.Server(":23456", time.Second*30, func() lconnect.Handler {
		return &Echo{}
	}); err != nil {

		panic(err)
	}

}
