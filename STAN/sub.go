package STAN

import (
	"GoL0/DB"
	"bufio"
	"encoding/json"
	"fmt"
	stan "github.com/nats-io/stan.go"
	"os"
)

func MsgReciever(m *stan.Msg) {
	var order DB.Orders
	err := json.Unmarshal(m.Data, &order)
	if err != nil {
		fmt.Printf("Received a message but not a json: %s\n", string(m.Data))
		return
	}
	fmt.Printf(order.Id.String())
	fmt.Printf("Received an: %s\n", string(m.Data))
}

func TestSub() {

	sc, err := stan.Connect("test-cluster", "client-123")
	if err != nil {
		panic(err)
	}
	sub, _ := sc.Subscribe("MyChannel", MsgReciever)

	input := bufio.NewScanner(os.Stdin)
	input.Scan()
	sub.Unsubscribe()
	sc.Close()
}
