// +build ignore

package main

// @author  Mikhail Kirillov <mikkirillov@yandex.ru>
// @version 1.001
// @date    2019-10-16

import (
	"errors"

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

	if err := lconnect.Server(":23456", func() lconnect.Handler {
		return &Echo{}
	}); err != nil {

		panic(err)
	}

}
