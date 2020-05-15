package control

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"strings"
	"time"

	"github.com/gorilla/mux"
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

func CovidbyState(w http.ResponseWriter, r *http.Request) {
	var cases []Caso

	params := mux.Vars(r)

	uf := params["uf"]

	log.Println(uf)

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

	err = json.Unmarshal(body, &cases)
	if err != nil {
		log.Fatal(err)
	}

	for _, estado := range cases[0].Estados {

		if estado.UF == strings.ToUpper(uf) {

			message, _ := json.Marshal(estado)

			log.Println(string(message))

			w.Write([]byte(message))

		}

	}

}
