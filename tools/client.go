// +build ignore

package main

// @author  Mikhail Kirillov <mikkirillov@yandex.ru>
// @version 1.001
// @date    2019-10-16

import (
	"errors"
	"fmt"
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

	client := lconnect.NewClient(":23456")

	for {

		data := []byte(fmt.Sprintf("%d", time.Now().UnixNano()))

		if answer, ok := client.WriteRead(data); ok {

			fmt.Println("> " + string(data))
			fmt.Println("< " + string(answer))

		} else {
			fmt.Println("broken connection")
		}

		<-time.After(time.Second)

	}

}
