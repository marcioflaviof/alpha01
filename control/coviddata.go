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
	Casos  int `json:"casos"`
	Mortes int `json:"mortes"`
}

// CovidData pega os dados da data de ontem sobre o covid
func CovidData(w http.ResponseWriter, r *http.Request) {

	var mortes []CasoseMortes

	d := time.Now().AddDate(0, 0, -1)

	url := "https://covid-api-brasil.herokuapp.com/casos/" + d.Format("2006-01-02")

	testClient := http.Client{
		Timeout: time.Second * 5, // Maximum of 2 secs
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

	//log.Println(string(body))

	err = json.Unmarshal(body, &mortes)
	if err != nil {
		log.Fatal(err)
	}

	message, _ := json.Marshal(mortes[0])

	log.Println(string(message))

	w.Write([]byte(message))

}
