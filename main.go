package main

import (
	"fmt"
	"time"

	"github.com/maneeshsagar/message-broker/core"
)

func main() {

	broker := core.NewMessageBroker(2, 1)
	timeChan := time.NewTicker(time.Second * time.Duration(broker.Interval))

	broker.AddMsg([]string{"jkhg", "jkhgf", "jkhgfcdx", "jhgfcdx"})
	for {
		select {
		case <-timeChan.C:

			msg := broker.Emit()
			fmt.Println(msg)
		}
	}
}
