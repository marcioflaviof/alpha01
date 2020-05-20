package control

import (
	"github.com/gorilla/mux"
	"alpha01/configs"
	"alpha01/database"
	"alpha01/models"
	"encoding/json"
	"log"
	"net/http"
)

func MainMenu(w http.ResponseWriter, r *http.Request) {
	
	// instancie o modelo do menu principal
	var m models.Menu

	// busque todos os ExamPreviews
	avl, err := database.SearchAllExamsPreview()

	// notifique caso o ocorra algum erro
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(configs.DBSRC_ERROR))
		return
	}

	// coloque os exames disponiveis no modelo
	m.Avaliable_Exams = avl

	// converta o modelo em JSON
	rsp, err := json.Marshal(m)

	// notifique se ocorrer algum erro
	if err != nil {
		log.Printf("[ERRO] Cannot parse the menu to JSON: %v %v", err, m)
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(configs.MARSHAL_ERROR))
		return
	}

	// retorne com a resposta
	w.Write([]byte(rsp))
}

func UpdateExamPreview(w http.ResponseWriter, r *http.Request){
	var results []models.Result
	var updated models.ExamPreviewUpdate

	//get the ID in URL
	vars := mux.Vars(r)
	num_id := vars["id"]

	results, err := database.SearchResultsById(num_id)

	//notique caso ocorra algum erro
	if err != nil{
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(configs.DBSRC_ERROR))
		return
	}

	new_num := len(results)

	updt, err := database.UpdateExamPreviewNumber(num_id,new_num)

	//notique caso ocorra algum erro
	if err != nil{
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(configs.DBUPD_ERROR))
		return
	}

	updated.Updated = updt

	res, err := json.Marshal(updated)

	//notique caso ocorra algum erro
	if err != nil{
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(configs.MARSHAL_ERROR))
		return
	}

	w.Write(res)
}