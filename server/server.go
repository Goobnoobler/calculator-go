package server

import (
	"fmt"
	"net/http"

	"github.com/goobnoobler/server/calculate"
	"github.com/gorilla/mux"
)

func Request() http.Handler {
	r := mux.NewRouter()

	r.HandleFunc("/calculate", func(w http.ResponseWriter, r *http.Request) {
		expression := r.URL.Query().Get("expression")
		if expression == "" {
			http.Error(w, "missing expression parameter", http.StatusBadRequest)
			return
		}

		output, errShunt := calculate.Shunt(expression)
		if errShunt != nil {
			http.Error(w, errShunt.Error(), http.StatusBadRequest)
			return
		}

		ans, errEval := calculate.Eval(output)
		if errEval != nil {
			http.Error(w, errEval.Error(), http.StatusBadRequest)
			return
		}

		w.Header().Set("Content-Type", "text/plain; charset=utf-8")
		fmt.Fprintf(w, "The answer is: %g", ans)
	})
	return r

}
