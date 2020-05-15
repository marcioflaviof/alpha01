package server

import (
	"alpha01/control"
	"net/http"

	"github.com/gorilla/mux"
)

func createHandler() (handler *mux.Router) {

	// creats router
	handler = mux.NewRouter()

	//application routes
	handler.HandleFunc("/", control.MainMenu).Methods("GET")

	//exams routes
	handler.HandleFunc("/exam", control.GetExam).Methods("GET")
	handler.HandleFunc("/exam/{id}", control.GetExamID).Methods("GET")

	//results routes
	handler.HandleFunc("/result", control.PostResult).Methods("POST")

	//covid numbers routes
	handler.HandleFunc("/covid", control.CovidData).Methods(http.MethodGet)
	handler.HandleFunc("/covid/{uf}", control.CovidbyState).Methods(http.MethodGet)

	// returns handle
	return
}
