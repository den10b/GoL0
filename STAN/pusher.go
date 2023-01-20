package STAN

import (
	stan "github.com/nats-io/stan.go"
	"log"
	"os"
)

//	func publishh(numb int) {
//		sc, _ := stan.Connect("test-cluster", fmt.Sprintf("client-t%d", numb))
//		sstr := []byte(fmt.Sprintf("Hello World-%d", numb))
//		err := sc.Publish("MyChannel", sstr)
//		if err != nil {
//			return
//		}
//		time.Sleep(time.Duration(numb) * time.Second)
//		err = sc.Publish("MyChannel", sstr)
//		if err != nil {
//			return
//		}
//	}
//
//	func TestPub() {
//		var wg sync.WaitGroup
//
//		go publishh(2)
//		wg.Add(1)
//		go publishh(4)
//		wg.Add(1)
//		go publishh(8)
//		wg.Add(1)
//		wg.Wait()
//	}
func TestJSONPub(filename string) {
	stuff, _ := os.ReadFile(filename)
	sc, _ := stan.Connect("test-cluster", "sender")
	err := sc.Publish("MyChannel", stuff)
	if err != nil {
		log.Printf("Error publishing to channel: %s", err)
	}
	log.Printf("%s was published!", filename)
}
