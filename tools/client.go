// +build ignore

package main

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

	client := lconnect.NewClient(":23456", time.Second*10)

	for {

		data := []byte(fmt.Sprintf("%d", time.Now().UnixNano()))

		if client.Write(data) {
			fmt.Println("> " + string(data))
			answer, ok := client.Read()
			if !ok {
				fmt.Println("broken connection")
			} else {
				fmt.Println("< " + string(answer))
			}
		} else {
			fmt.Println("broken connection")
		}

		<-time.After(time.Second)

	}

}
