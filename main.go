package main

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

func main() {

	//var cases interface{}
	var Casos []Caso

	url := "https://covid-api-brasil.herokuapp.com/casos/2020-05-06"

	testClient := http.Client{
		Timeout: time.Second * 2, // Maximum of 2 secs
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

	cases1 := Casos

	log.Println(string(body))

	jsonErr := json.Unmarshal(body, &cases1)
	if jsonErr != nil {
		log.Fatal(jsonErr)
	}
	//fmt.Printf("HTTP: %s\n", res.Status)
	//fmt.Println(cases1)

}
