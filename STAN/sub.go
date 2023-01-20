package STAN

import (
	"GoL0/Cache"
	"GoL0/DB"
	"encoding/json"
	"github.com/google/uuid"
	stan "github.com/nats-io/stan.go"
	"log"
)

var sc stan.Conn
var sub stan.Subscription

func MsgReciever(m *stan.Msg) {
	var order DB.Orders
	err := json.Unmarshal(m.Data, &order)
	if err != nil {
		log.Printf("Received a message but not a json: %s\n", string(m.Data))
		return
	}
	_, err = Cache.GetOrder(order.Id.String())
	if err == nil {
		log.Printf("Order %s already added", order.Id.String())
		return
	}
	order.Delivery.OrderUid = order.Id
	order.Delivery.Uid = uuid.New()

	order.Payment.OrderUid = order.Id

	for i, _ := range order.Items {
		order.Items[i].OrderUid = order.Id
		order.Items[i].Uid = uuid.New()
	}
	err = DB.SetOrder(order)
	if err != nil {
		log.Printf("Error adding order to DB: %s", err)
	}
	err = DB.SetItems(order.Items)
	if err != nil {
		log.Printf("Error adding items to DB: %s", err)
	}
	err = DB.SetDelivery(order.Delivery)
	if err != nil {
		log.Printf("Error adding delivery to DB: %s", err)
	}
	err = DB.SetPayment(order.Payment)
	if err != nil {
		log.Printf("Error adding payment to DB: %s", err)
	}
	err = Cache.AddOrder(order)
	if err != nil {
		log.Printf("Error adding order to Cache: %s", err)
	}

	log.Printf("Received an order: %s\n", order.Id.String())
}

func InitSub() {
	var err error
	sc, err = stan.Connect("test-cluster", "client-123")
	if err != nil {
		panic(err)
	}
	channelName := "MyChannel"
	sub, err = sc.Subscribe(channelName, MsgReciever)
	if err != nil {
		panic(err)
	}
	log.Printf("Subscribed to channel: %s\n", channelName)
}

func QuitSub() {
	err := sub.Unsubscribe()
	if err != nil {
		log.Printf("Error unsubscribing: %s", err)
	}
	err = sc.Close()
	if err != nil {
		log.Printf("Error closing: %s", err)
	}
}
