package control

import (
	"github.com/gorilla/mux"
	"alpha01/database"
	"alpha01/models"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"log"
)

func GetExamID(w http.ResponseWriter, r *http.Request) {
	//get the ID in URL
	vars := mux.Vars(r)
	num_id := vars["id"]

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

func PostExam(w http.ResponseWriter, r *http.Request){

	// instancie o modelo de exame
	var ex models.Exam

	//pegue o resultado do teste
	body := r.Body
	bytes, err := ioutil.ReadAll(body)

	//notifique caso ocorra algum erro
	if err != nil{
		log.Printf("[ERRO] Cannot read the body request: %v %v",err,ex)
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	// insira o body no modelo
	err = json.Unmarshal(bytes,&ex)

	// notifique caso ocorra algum erro
	if err != nil{
		log.Printf("[ERRO] Cannot unmarshal the JSON: %v %v",err,ex)
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	// insira o exam no banco de dados
	err = database.InsertExam(ex)

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	// crie um exame preview para ter o controle de resposta
	err = database.CreateExamsPreview(ex.Exam_ID)

	if err != nil{
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	// notifique a operação na resposta
	w.Write([]byte(`{"msg":"Exam was inserted"`))

	return
}