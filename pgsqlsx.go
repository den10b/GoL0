package main

import (
	"encoding/json"
	"fmt"
	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
	"log"
	"time"
)

const (
	host     = "localhost"
	port     = 5432
	user     = "postgres"
	password = "1234"
	dbname   = "postgres"
	sslmode  = "disable"
)

type Orders struct {
	Id                uuid.UUID `db:"id" json:"id"`
	TrackNumber       string    `db:"track_number" json:"track_number"`
	Entry             string    `db:"entry" json:"entry"`
	Locale            string    `db:"locale" json:"locale"`
	InternalSignature string    `db:"internal_signature" json:"internal_signature"`
	CustomerId        string    `db:"customer_id" json:"customer_id"`
	DeliveryService   string    `db:"delivery_service" json:"delivery_service"`
	Shardkey          string    `db:"shardkey" json:"shardkey"`
	SmId              int       `db:"sm_id" json:"sm_id"`
	DateCreated       time.Time `db:"date_created" json:"date_created"`
	OofShard          string    `db:"oof_shard" json:"oof_shard"`
}
type Payments struct {
	Transaction  uuid.UUID `db:"transaction" json:"transaction"`
	OrderUid     uuid.UUID `db:"order_uid" json:"order_uid"`
	RequestId    string    `db:"request_id" json:"request_id"`
	Currency     string    `db:"currency" json:"currency"`
	Provider     string    `db:"provider" json:"provider"`
	Amount       int       `db:"amount" json:"amount"`
	PaymentDT    int       `db:"payment_dt" json:"payment_dt"`
	Bank         string    `db:"bank" json:"bank"`
	DeliveryCost int       `db:"delivery_cost" json:"delivery_cost"`
	GoodsTotal   int       `db:"goods_total" json:"goods_total"`
	CustomFee    int       `db:"custom_fee" json:"custom_fee"`
}

type Items struct {
	Uid         uuid.UUID `db:"uid" json:"uid"`
	OrderUid    uuid.UUID `db:"order_uid" json:"order_uid"`
	ChrtId      int       `db:"chrt_id" json:"chrt_id"`
	TrackNumber string    `db:"track_number" json:"track_number"`
	Price       int       `db:"price" json:"price"`
	Rid         string    `db:"rid" json:"rid"`
	Name        string    `db:"name" json:"name"`
	Sale        int       `db:"sale" json:"sale"`
	Size        string    `db:"size" json:"size"`
	TotalPrice  int       `db:"total_price" json:"total_price"`
	NmId        int       `db:"nm_id" json:"nm_id"`
	Brand       string    `db:"brand" json:"brand"`
	Status      int       `db:"status" json:"status"`
}

type Delivery struct {
	Uid      uuid.UUID `db:"uid" json:"uid"`
	OrderUid uuid.UUID `db:"order_uid" json:"order_uid"`
	Name     string    `db:"name" json:"chrt_id"`
	Phone    string    `db:"phone" json:"track_number"`
	Zip      string    `db:"zip" json:"price"`
	City     string    `db:"city" json:"rid"`
	Address  string    `db:"address" json:"name"`
	Region   string    `db:"region" json:"sale"`
	Email    string    `db:"email" json:"size"`
}

func main() {

	t := "host=%s port=%d user=%s password=%s dbname=%s sslmode=%s"
	connectionString := fmt.Sprintf(t, host, port, user, password, dbname, sslmode)
	db, err := sqlx.Open("postgres", connectionString)
	if err != nil {
		log.Fatalln(err)
	}
	var order Orders
	var order1 Orders

	err = db.Get(&order, "select * from public.orders")
	query := fmt.Sprintf("select * from public.orders where id = '%s'", order.Id.String())
	err = db.Get(&order1, query)

	var orders []Orders
	err = db.Select(&orders, "select * from public.orders")

	for i, myOrder := range orders {
		var delivery Delivery
		query = fmt.Sprintf("select * from public.delivery where delivery.order_uid = '%s'", myOrder.Id.String())
		err = db.Get(&delivery, query)

		var items []Items
		query = fmt.Sprintf("select * from public.items where items.order_uid = '%s'", myOrder.Id.String())
		err = db.Select(&items, query)

		var payment Payments
		query = fmt.Sprintf("select * from public.payments where payments.order_uid = '%s'", myOrder.Id.String())
		err = db.Get(&payment, query)

		fmt.Println(i)
		marOrder, _ := json.Marshal(myOrder)
		fmt.Println(string(marOrder))
		marDelivery, _ := json.Marshal(delivery)
		fmt.Println(string(marDelivery))
		marItems, _ := json.Marshal(items)
		fmt.Println(string(marItems))
		marPayment, _ := json.Marshal(payment)
		fmt.Println(string(marPayment))
	}

}
