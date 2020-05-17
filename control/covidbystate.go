package control

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"time"

	"github.com/gorilla/mux"
)

//CovidEstados is how many covid cases exists in states
type CovidEstados struct {
	Casos  int `json:"casos"`
	Mortes int `json:"mortes"`
}

//CovidbyState picks covid by state
func CovidbyState(w http.ResponseWriter, r *http.Request) {
	var cases CovidEstados

	params := mux.Vars(r)

	uf := params["uf"]

	log.Println(uf)

	url := "https://covid-api-brasil.herokuapp.com/" + uf

	testClient := http.Client{
		Timeout: time.Second * 10, // Maximum of 2 secs
	}

	req, err := http.NewRequest(http.MethodGet, string(url), nil)
	if err != nil {
		log.Fatal(err)
	}

	res, err := testClient.Do(req)
	if err != nil {
		log.Fatal(err)
	}

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		log.Fatal(err)
	}

	err = json.Unmarshal(body, &cases)
	if err != nil {
		log.Fatal(err)
	}

	message, _ := json.Marshal(cases)

	log.Println(string(message))

	w.Write([]byte(message))

}
