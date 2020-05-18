package control

import (
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
		w.Write([]byte(configs.RESPONSE_ERRO))
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
		w.Write([]byte(configs.RESPONSE_ERRO))
		return
	}

	// retorne com a resposta
	w.Write([]byte(rsp))
}