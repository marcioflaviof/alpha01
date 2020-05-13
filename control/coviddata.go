package control

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"time"
)

type Estados struct {
	UF     string `json:"uf"`
	Casos  int    `json:"casos"`
	Mortes int    `json:"mortes"`
}

type Caso struct {
	Data    string `json:"data"`
	Casos   int    `json:"casos"`
	Mortes  int    `json:"mortes"`
	Estados []Estados
}

type CasoseMortes struct {
	Casos  int `json:"casos"`
	Mortes int `json:"mortes"`
}

// CovidData pega os dados da data de ontem sobre o covid
func CovidData(w http.ResponseWriter, r *http.Request) {

	var mortes []CasoseMortes

	d := time.Now().AddDate(0, 0, -1)

	log.Println(d.Format("2006-01-02"))

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

	jsonErr := json.Unmarshal(body, &mortes)
	if jsonErr != nil {
		log.Fatal(jsonErr)
	}

	message, _ := json.Marshal(mortes[0])

	log.Println(string(message))

	w.Write([]byte(message))

}
