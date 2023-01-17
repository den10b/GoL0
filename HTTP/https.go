package HTTP

import (
	"GoL0/DB"
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
)

func Handlerr(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello World!")
}
func Handlerr2(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello World!22222")
}
func allOrders(w http.ResponseWriter, r *http.Request) {
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
func getOrder(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	order, err := DB.GetOrder(params["order_id"])
	err = json.NewEncoder(w).Encode(order)

	if err != nil {
		return
	}

}

func TestHttp() {
	r := mux.NewRouter()
	r.HandleFunc("/", Handlerr)
	r.HandleFunc("/adad", Handlerr2)
	r.HandleFunc("/order", allOrders)
	r.HandleFunc("/order/{order_id}", getOrder)
	http.Handle("/", r)
	err := http.ListenAndServe(":80", nil)
	if err != nil {
		return
	}
}
