package STAN

import (
	"fmt"
	stan "github.com/nats-io/stan.go"
	"os"
	"sync"
	"time"
)

func publishh(numb int) {
	sc, _ := stan.Connect("test-cluster", fmt.Sprintf("client-t%d", numb))
	sstr := []byte(fmt.Sprintf("Hello World-%d", numb))
	err := sc.Publish("MyChannel", sstr)
	if err != nil {
		return
	}
	time.Sleep(time.Duration(numb) * time.Second)
	err = sc.Publish("MyChannel", sstr)
	if err != nil {
		return
	}
}
func TestPub() {
	var wg sync.WaitGroup

	go publishh(2)
	wg.Add(1)
	go publishh(4)
	wg.Add(1)
	go publishh(8)
	wg.Add(1)
	wg.Wait()
}
func TestJSONPub() {
	plan, _ := os.ReadFile("testJSON.json")
	sc, _ := stan.Connect("test-cluster", "sender")
	err := sc.Publish("MyChannel", plan)
	//var order DB.Orders
	//err := json.Unmarshal(plan, &order)
	if err != nil {
		return
	}
}
