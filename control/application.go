package control

import(
	"alpha01/database"
	"alpha01/configs"
	"alpha01/models"
	"encoding/json"
	"net/http"
	"log"
)

func MainMenu(w http.ResponseWriter, r *http.Request){
	var m models.Menu

	avl, err := SearchAvaliableExams()

	if err != nil{
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(configs.RESPONSE_ERRO))
		return
	}

	m.Avaliable_Exams = avl

	rsp, err := json.Marshal(m)

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(configs.RESPONSE_ERRO))
		return
	}

	w.Write([]byte(rsp))
}

func SearchAvaliableExams() (avl []string, err error){

	var exams []models.Exam

	exams, err = database.SearchAllExams()

	if err != nil {
		log.Printf("[ERROR] probleming searching avaliable exams: %v, %v",err,exams)
		return
	}

	for _,tst := range exams{
		avl = append(avl,tst.Exam_ID)
	}

	// confirm action of inserting a exam
	log.Printf("[INFO] avaliable exams: %v",avl)

	return
}