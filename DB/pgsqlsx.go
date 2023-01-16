package DB

import (
	"encoding/json"
	"fmt"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
	"log"
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

func CloseConn(db *sqlx.DB) {
	err := db.Close()
	if err != nil {
		log.Fatalln(err)
	}
}

func TestSQL() {

	t := "host=%s port=%d user=%s password=%s dbname=%s sslmode=%s"
	connectionString := fmt.Sprintf(t, host, port, user, password, dbname, sslmode)
	var err error
	Db, err = sqlx.Open("postgres", connectionString)

	if err != nil {
		log.Fatalln(err)
	}

	defer CloseConn(Db)

	var order Orders
	var order1 Orders

	err = Db.Get(&order, "select * from public.orders")
	query := fmt.Sprintf("select * from public.orders where id = '%s'", order.Id.String())
	err = Db.Get(&order1, query)

	var orders []Orders
	err = Db.Select(&orders, "select * from public.orders")

	for i, myOrder := range orders {
		var delivery Delivery
		query = fmt.Sprintf("select * from public.delivery where delivery.order_uid = '%s'", myOrder.Id.String())
		err = Db.Get(&delivery, query)

		var items []Items
		query = fmt.Sprintf("select * from public.items where items.order_uid = '%s'", myOrder.Id.String())
		err = Db.Select(&items, query)

		var payment Payments
		query = fmt.Sprintf("select * from public.payments where payments.order_uid = '%s'", myOrder.Id.String())
		err = Db.Get(&payment, query)

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
