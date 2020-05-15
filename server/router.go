package server

import (
	"alpha01/control"
	"net/http"

	"github.com/gorilla/mux"
)

func createHandler() (r *mux.Router) {

	// creats router
	r = mux.NewRouter()

	// associate register user route
	r.HandleFunc("/", control.CovidData).Methods(http.MethodGet)
	r.HandleFunc("/{uf}", control.CovidbyState).Methods(http.MethodGet)

	// returns handle
	return
}
