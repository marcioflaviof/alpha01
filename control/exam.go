package control

import (
	"alpha01/database"
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
)

func GetExamID(w http.ResponseWriter, r *http.Request) {
	//get the ID in URL
	vars := mux.Vars(r)
	num_id := vars["id"]
	//n_id, _ := strconv.Atoi(num_id)

	//search the exam with the ID
	exam, err := database.SearchExamID(num_id)

	//check if none searching error occurs
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		return
	}

	//parse exam to JSON
	resp, err := json.Marshal(exam)

	//check if none parsing error occurs
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
	}

	//return the response with the JSON
	w.Write([]byte(resp))
}

//pegar o exame do cornavírus
func GetExam(w http.ResponseWriter, r *http.Request) {
	//Pegue o examo no banco de dados
	exam, err := database.SearchExamID("1")

	// verifique se não houve um erro
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		return
	}

	// transforme o exame em JSON
	rsp, err := json.Marshal(exam)

	// verifique se tudo ocorreu tudo bem
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
	}

	// envie o exame ao solicitante
	w.Write([]byte(rsp))
}
