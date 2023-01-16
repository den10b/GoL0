package main

import (
	"fmt"
	"net/http"
)

func Handlerr(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello World!")
}
func Handlerr2(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello World!22222")
}

func main() {
	http.HandleFunc("/", Handlerr)
	http.HandleFunc("/adad", Handlerr2)
	err := http.ListenAndServe(":80", nil)
	if err != nil {
		return
	}
}
