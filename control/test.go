package control

import (
	"github.com/gorilla/mux"
	"alpha01/database"
	"encoding/json"
	"net/http"
	"strings"
)

func RegisterUser(w http.ResponseWriter, r *http.Request) {

	message := r.URL.Path
	message = strings.TrimPrefix(message, "/")
	message = "Hello " + message
	w.Write([]byte(message))
}

func GetTestID(w http.ResponseWriter, r *http.Request){
	//get the ID in URL
	vars := mux.Vars(r)
	num_id := vars["id"]

	//search the test with the ID
	test, err := database.SearchTestID(num_id)

	//check if none searching error occurs
	if err != nil{
		w.WriteHeader(http.StatusNotFound)
		return
	}

	//parse test to JSON
	resp, err := json.Marshal(test)

	//check if none parsing error occurs
	if err != nil{
		w.WriteHeader(http.StatusInternalServerError)
	}

	//return the response with the JSON
	w.Write([]byte(resp))
}
