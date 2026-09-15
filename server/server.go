package server

import (
	"encoding/json"
	"net/http"

	"github.com/goobnoobler/server/calculate"
	"github.com/gorilla/mux"
)

type Expression struct {
	Exp string `json:"expression"`
}

type Answer struct {
	Result float64 `json:"result"`
}

func Request() http.Handler {
	r := mux.NewRouter()

	r.HandleFunc("/api/calculate", func(w http.ResponseWriter, r *http.Request) {
		var expression Expression

		if jsonErr := json.NewDecoder(r.Body).Decode(&expression); jsonErr != nil {
			http.Error(w, "malformed json", http.StatusBadRequest)
			return
		}

		output, errShunt := calculate.Shunt(expression.Exp)
		if errShunt != nil {
			http.Error(w, errShunt.Error(), http.StatusBadRequest)
			return
		}

		ans, errEval := calculate.Eval(output)
		if errEval != nil {
			http.Error(w, errEval.Error(), http.StatusBadRequest)
			return
		}

		response := Answer{
			Result: ans,
		}

		w.Header().Set("Content-Type", "application/json")
		if encodeErr := json.NewEncoder(w).Encode(response); encodeErr != nil {
			http.Error(w, "encode error", http.StatusInternalServerError)
			return
		}
	})
	return r

}
