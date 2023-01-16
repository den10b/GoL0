package HTTP

import (
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
func getOrder(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	err := json.NewEncoder(w).Encode(params)
	if err != nil {
		return
	}

}

func TestHttp() {
	r := mux.NewRouter()
	r.HandleFunc("/", Handlerr)
	r.HandleFunc("/adad", Handlerr2)
	r.HandleFunc("/order/{key}", getOrder)
	http.Handle("/", r)
	err := http.ListenAndServe(":80", nil)
	if err != nil {
		return
	}
}
