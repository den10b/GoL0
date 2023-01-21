package HTTP

import (
	"GoL0/Cache"
	"github.com/gorilla/mux"
	"html/template"
	"net/http"
	"os"
)

func emptyHandler(w http.ResponseWriter, r *http.Request) {
	tpl, err := template.ParseFiles("./Static/index.html")
	if err != nil {
		panic(err)
	}
	err = tpl.Execute(w, "nil")
	if err != nil {
		panic(err)
	}
}
func cssHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/css")
	stuff, _ := os.ReadFile("./Static/style.css")
	_, err := w.Write(stuff)
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
func allItemsCache(w http.ResponseWriter, r *http.Request) {
	orders, err := Cache.GetOrders()
	tpl, err := template.ParseFiles("./Static/items.html")
	if err != nil {
		panic(err)
	}
	err = tpl.Execute(w, orders)
	if err != nil {
		panic(err)
	}
}
func allDeliveriesCache(w http.ResponseWriter, r *http.Request) {
	orders, err := Cache.GetOrders()
	tpl, err := template.ParseFiles("./Static/deliveries.html")
	if err != nil {
		panic(err)
	}
	err = tpl.Execute(w, orders)
	if err != nil {
		panic(err)
	}
}
func allPaymentsCache(w http.ResponseWriter, r *http.Request) {
	orders, err := Cache.GetOrders()
	tpl, err := template.ParseFiles("./Static/payments.html")
	if err != nil {
		panic(err)
	}
	err = tpl.Execute(w, orders)
	if err != nil {
		panic(err)
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

//	func getOrderDB(w http.ResponseWriter, r *http.Request) {
//		w.Header().Set("Content-Type", "application/json")
//		params := mux.Vars(r)
//		order, err := DB.GetOrder(params["order_id"])
//		err = json.NewEncoder(w).Encode(order)
//
//		if err != nil {
//			return
//		}
//
// }

func InitHttp() {
	r := mux.NewRouter()
	r.HandleFunc("/style.css", cssHandler)
	r.HandleFunc("/order/style.css", cssHandler)
	r.HandleFunc("/", emptyHandler)
	r.HandleFunc("/deliveries", allDeliveriesCache)
	r.HandleFunc("/items", allItemsCache)
	r.HandleFunc("/payments", allPaymentsCache)
	r.HandleFunc("/order", allOrdersCache)
	r.HandleFunc("/order/{order_id}", getOrderCache)
	http.Handle("/", r)
	err := http.ListenAndServe(":80", nil)
	if err != nil {
		panic(err)
	}
}
