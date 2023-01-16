package main

import (
	"fmt"
	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
	"log"
	"time"
)

var schema = `
CREATE TABLE person (
    first_name text,
    last_name text,
    email text
);

CREATE TABLE place (
    country text,
    city text NULL,
    telcode integer
)`

const (
	host     = "localhost"
	port     = 5432
	user     = "postgres"
	password = "1234"
	dbname   = "postgres"
	sslmode  = "disable"
)

type Orders struct {
	Id      uuid.UUID `db:"id" json:"id"`
	Numb    string    `db:"track_number" json:"track_number"`
	Entry   string    `db:"entry" json:"entry"`
	Locale  string    `db:"locale" json:"locale"`
	InSig   string    `db:"internal_signature" json:"internal_signature"`
	CustId  string    `db:"customer_id" json:"customer_id"`
	DelServ string    `db:"delivery_service" json:"delivery_service"`
	ShK     string    `db:"shardkey" json:"shardkey"`
	SmId    int       `db:"sm_id" json:"sm_id"`
	DateCr  time.Time `db:"date_created" json:"date_created"`
	OofSha  string    `db:"oof_shard" json:"oof_shard"`
}
type Payments struct {
	Transaction  uuid.UUID `db:"transaction" json:"id"`
	OrderUid     string    `db:"order_uid" json:"track_number"`
	RequestId    string    `db:"request_id" json:"entry"`
	Currency     string    `db:"currency" json:"locale"`
	Provider     string    `db:"provider" json:"internal_signature"`
	Amount       string    `db:"amount" json:"customer_id"`
	PaymentDT    string    `db:"payment_dt" json:"delivery_service"`
	Bank         string    `db:"bank" json:"shardkey"`
	DeliveryCost int       `db:"delivery_cost" json:"sm_id"`
	GoodsTotal   time.Time `db:"goods_total" json:"date_created"`
	CustomFee    string    `db:"custom_fee" json:"oof_shard"`
}

func main() {

	t := "host=%s port=%d user=%s password=%s dbname=%s sslmode=%s"
	connectionString := fmt.Sprintf(t, host, port, user, password, dbname, sslmode)
	db, err := sqlx.Open("postgres", connectionString)
	if err != nil {
		log.Fatalln(err)
	}
	var p Orders
	var pp []Orders
	err = db.Get(&p, "select * from public.orders LIMIT 1")
	// Select записывает в pp массив полученных строк.
	querr := fmt.Sprintf("select * from public.orders where id = '%s'", p.Id.String())
	err = db.Select(&pp, querr)
	i := 1
	i = i + 1

}
