package control

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"strings"
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

func RegisterUser(w http.ResponseWriter, r *http.Request) {

	message := r.URL.Path
	message = strings.TrimPrefix(message, "/")
	message = "Hello " + message
	w.Write([]byte(message))
}

func CovidData(w http.ResponseWriter, r *http.Request) {

	//var cases interface{}
	var casos []Caso
	var mortes []CasoseMortes

	url := "https://covid-api-brasil.herokuapp.com/casos/2020-05-06"

	testClient := http.Client{
		Timeout: time.Second * 5, // Maximum of 2 secs
	}

	req, err := http.NewRequest(http.MethodGet, url, nil)
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

	cases1 := casos

	//log.Println(string(body))

	jsonErr := json.Unmarshal(body, &cases1)
	if jsonErr != nil {
		log.Fatal(jsonErr)
	}

	jsonErr = json.Unmarshal(body, &mortes)
	if jsonErr != nil {
		log.Fatal(jsonErr)
	}

	message, _ := json.Marshal(mortes[0])

	log.Println(string(message))

	w.Write([]byte(message))

}
