package server

// import (
// 	"net/http"

// 	"github.com/gorilla/mux"
// )

// func Request() http.Handler {
// 	r := mux.NewRouter()

// 	r.HandleFunc("/calculate/{expression}", func(w http.ResponseWriter, r *http.Request) {
// 		vars := mux.Vars(r)
// 		expression := vars["expression"]
// 		EvalExpression(expression)

// 	})
// 	return r
// }
