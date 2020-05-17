package control

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"time"
)

//CasoseMortes is a model from deaths and cases in covid
type CasoseMortes struct {
	Casos  int `json:"totalCasos"`
	Mortes int `json:"totalMortes"`
}

// CovidData pega os dados da data de ontem sobre o covid
func CovidData(w http.ResponseWriter, r *http.Request) {

	var mortes CasoseMortes

	url := "https://covid-api-brasil.herokuapp.com/casos"

	testClient := http.Client{
		Timeout: time.Second * 10, // Maximum of 2 secs
	}

	req, err := http.NewRequest(http.MethodGet, string(url), nil)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("Get url error"))
	}

	res, err := testClient.Do(req)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("Requisition error"))
	}

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("Body reading error"))
	}

	//log.Println(string(body))

	err = json.Unmarshal(body, &mortes)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("Unmarshall error"))
	}

	message, _ := json.Marshal(mortes)

	log.Println(string(message))

	w.Write([]byte(message))

}
