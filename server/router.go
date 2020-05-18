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
	handler.HandleFunc("/", control.MainMenu).Methods(http.MethodGet)

	//exams routes
	handler.HandleFunc("/exam", control.GetExam).Methods(http.MethodGet)
	handler.HandleFunc("/exam/{id}", control.GetExamID).Methods(http.MethodGet)
	handler.HandleFunc("/exam", control.PostExam).Methods(http.MethodPost)

	//results routes
	handler.HandleFunc("/result", control.PostResult).Methods(http.MethodPost)

	//covid numbers routes
	handler.HandleFunc("/covid", control.CovidData).Methods(http.MethodGet)
	handler.HandleFunc("/covid/{uf}", control.CovidbyState).Methods(http.MethodGet)

	// returns handle
	return
}
