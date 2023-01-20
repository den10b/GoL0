package DB

import (
	"encoding/json"
	"fmt"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
	"log"
	"time"
)

var Db *sqlx.DB

const (
	host     = "localhost"
	port     = 5432
	user     = "postgres"
	password = "1234"
	dbname   = "postgres"
	sslmode  = "disable"
)

//func getItems(uuid uuid.UUID) []Items {
//	var items []Items
//	query := fmt.Sprintf("select * from public.items where items.order_uid = '%s'", uuid.String())
//	err = db.Select(&items, query)
//
//}

func CloseConn() {
	err := Db.Close()
	if err != nil {
		log.Fatalln(err)
	}
}
func OpenConn() {
	t := "host=%s port=%d user=%s password=%s dbname=%s sslmode=%s"
	connectionString := fmt.Sprintf(t, host, port, user, password, dbname, sslmode)
	var err error
	Db, err = sqlx.Open("postgres", connectionString)

	if err != nil {
		log.Fatalln(err)
	}
}
func GetAllOrders() ([]Orders, error) {
	var orders []Orders
	err := Db.Select(&orders, "select * from public.orders")
	if err != nil {
		return nil, err
	}
	return orders, nil
}
func GetOrder(order_id string) (Orders, error) {
	var order Orders
	err := Db.Get(&order, "select * from public.orders where orders.id = $1", order_id)
	if err != nil {
		return Orders{}, err
	}
	return order, nil
}
func SetOrder(orderId string, trackNumber string, entry string, locale string, internalSignature string, customerId string, deliveryService string, shardKey string, smId int, dateCreated time.Time, oofShard string) (Orders, error) {
	var order Orders
	err := Db.QueryRow("INSERT INTO public.orders VALUES($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11) RETURNING id,track_number,entry,locale,internal_signature,customer_id,delivery_service,shardkey,sm_id,date_created,oof_shard", orderId, trackNumber, entry, locale, internalSignature, customerId, deliveryService, shardKey, smId, dateCreated, oofShard).Scan(&order)
	if err != nil {
		return Orders{}, err
	}
	return order, nil
}
func GetAllItems() ([]Items, error) {
	var items []Items
	err := Db.Select(&items, "select * from public.orders")
	if err != nil {
		return nil, err
	}
	return items, nil
}

func TestSQL() {

	defer CloseConn()
	var err error
	var order Orders
	var order1 Orders

	err = Db.Get(&order, "select * from public.orders")
	if err != nil {
		return
	}

	//query := fmt.Sprintf("select * from public.orders where id = '%s'", order.Id.String())
	err = Db.Get(&order1, "select * from public.orders where id = $1", order.Id.String())
	getOrder, err := GetOrder(order.Id.String())
	if err != nil {
		return
	}
	fmt.Println(getOrder)
	var orders []Orders
	err = Db.Select(&orders, "select * from public.orders")

	for i, myOrder := range orders {
		var delivery Delivery
		err = Db.Get(&delivery, "select * from public.delivery where delivery.order_uid = $1", myOrder.Id.String())

		var items []Items
		err = Db.Select(&items, "select * from public.items where items.order_uid = $1", myOrder.Id.String())

		var payment Payments
		err = Db.Get(&payment, "select * from public.payments where payments.order_uid = $1", myOrder.Id.String())

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
