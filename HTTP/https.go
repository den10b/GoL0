package HTTP

import (
	"GoL0/Cache"
	"GoL0/DB"
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"html/template"
	"net/http"
)

func Handlerr(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello World!")
}
func Handlerr2(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello World!22222")
}
func allOrdersDB(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	orders, err := DB.GetAllOrders()
	if err != nil {
		return
	}
	err = json.NewEncoder(w).Encode(orders)
	if err != nil {
		return
	}

}
func allOrdersCache(w http.ResponseWriter, r *http.Request) {
	orders, err := Cache.GetOrders()
	tpl, err := template.ParseFiles("./Static/orders.html")
	if err != nil {
		panic(err)
	}
	err = tpl.Execute(w, orders)
	if err != nil {
		panic(err)
	}

}

func getOrderDB(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	order, err := DB.GetOrder(params["order_id"])
	err = json.NewEncoder(w).Encode(order)

	if err != nil {
		return
	}

}
func getOrderCache(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	order, err := Cache.GetOrder(params["order_id"])
	tpl, err := template.ParseFiles("./Static/order.html")
	if err != nil {
		panic(err)
	}
	err = tpl.Execute(w, order)
	if err != nil {
		panic(err)
	}

}

func TestHttp() {
	r := mux.NewRouter()
	r.HandleFunc("/", Handlerr)
	r.HandleFunc("/adad", Handlerr2)
	r.HandleFunc("/order", allOrdersCache)
	r.HandleFunc("/order/{order_id}", getOrderCache)
	http.Handle("/", r)
	err := http.ListenAndServe(":80", nil)
	if err != nil {
		return
	}
}
