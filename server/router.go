package server

import (
	"alpha01/control"

	"github.com/gorilla/mux"
)

func createHandler() (r *mux.Router) {

	// creats router
	r = mux.NewRouter()

	// associate register user route
	r.HandleFunc("/", control.CovidData)

	// returns handle
	return
}
