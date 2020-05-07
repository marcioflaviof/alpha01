package server

import (
	"alpha01/control"

	"github.com/gorilla/mux"
)

func createHandler() (handler *mux.Router) {

	// creats router
	handler = mux.NewRouter()

	// associate register user route
	handler.HandleFunc("/", control.RegisterUser)
	handler.HandleFunc("/test/{id}", control.GetTestID)

	// returns handle
	return
}
