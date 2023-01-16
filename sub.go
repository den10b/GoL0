package main

import (
	"bufio"
	"fmt"
	stan "github.com/nats-io/stan.go"
	"os"
)

func main() {

	sc, err := stan.Connect("test-cluster", "client-123")
	if err != nil {
		panic(err)
	}
	sub, _ := sc.Subscribe("MyChannel", func(m *stan.Msg) {
		fmt.Printf("Received a message: %s\n", string(m.Data))
	})

	input := bufio.NewScanner(os.Stdin)
	input.Scan()
	sub.Unsubscribe()
	sc.Close()
}
