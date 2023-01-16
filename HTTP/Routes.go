package HTTP

//
//import (
//	"fmt"
//	"net/http"
//)
//
//type HelloResponse struct {
//Message string `json:"message"`
//}
//
//// HelloName returns a personalized JSON message
//func HelloName(w http.ResponseWriter, r *http.Request) {
//name := chi.URLParam(r, "name")
//response := HelloResponse{
//Message: fmt.Sprintf("Hello %s!", name),
//}
//jsonResponse(w, response, http.StatusOK)
//}
//
//// NewRouter returns an HTTP handler that implements the routes for the API
//func NewRouter() http.Handler {
//r := chi.NewRouter()
//r.Get("/{name}", HelloName)
//return r
