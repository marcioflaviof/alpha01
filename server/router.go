package server

import (
	"alpha01/control"

	"github.com/gorilla/mux"
)

func createHandler() (handler *mux.Router) {

	// creats router
	handler = mux.NewRouter()

	// associate register user route
	handler.HandleFunc("/test/{id}", control.GetTestID)

	//paginaprincipal GET
	//handler.HandleFunc("/", control.MainMenu).Methods("GET")
	//questionario GET
	handler.HandleFunc("/test", control.GetTest).Methods("GET")
	//resultado POST (resultado no response)
	handler.HandleFunc("/result", control.PostResult).Methods("POST")


	// returns handle
	return
}
