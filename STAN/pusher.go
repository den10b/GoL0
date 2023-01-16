package STAN

import (
	"fmt"
	stan "github.com/nats-io/stan.go"
	"sync"
	"time"
)

func publishh(numb int) {
	sc, _ := stan.Connect("test-cluster", fmt.Sprintf("client-t%d", numb))
	sstr := []byte(fmt.Sprintf("Hello World-%d", numb))
	sc.Publish("MyChannel", sstr)
	time.Sleep(time.Duration(numb) * time.Second)
	sc.Publish("MyChannel", sstr)
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
