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

type Difference struct {
	DiffWeekC  int `json:"diffweekc"`
	DiffMonthC int `json:"diffmonthc"`
	DiffWeekM  int `json:"diffweekm"`
	DiffMonthM int `json:"diffmonthm"`
}

func CovidDifference(w http.ResponseWriter, r *http.Request) {

	var diff Difference

	yesterdayCases, yesterdayDeaths := covidTime(0, 1, r)
	lastWeekCases, weekDeaths := covidTime(0, 8, r)
	monthCases, monthDeaths := covidTime(1, 1, r)

	diff.DiffWeekC = yesterdayCases - lastWeekCases

	diff.DiffMonthC = yesterdayCases - monthCases

	diff.DiffWeekM = yesterdayDeaths - weekDeaths

	diff.DiffMonthM = yesterdayDeaths - monthDeaths

	byteDiff, err := json.Marshal(diff)

	if err != nil {
		log.Println(err)
	}

	w.Write(byteDiff)

}

func covidTime(tempoMes int, tempoDia int, r *http.Request) (int, int) {

	var cases []Caso

	params := mux.Vars(r)

	uf := params["uf"]

	//log.Println(uf)

	d := time.Now().AddDate(0, -tempoMes, -tempoDia)

	url := "https://covid-api-brasil.herokuapp.com/casos/" + d.Format("2006-01-02")

	testClient := http.Client{
		Timeout: time.Second * 5, // Maximum of 2 secs
	}

	req, err := http.NewRequest(http.MethodGet, string(url), nil)
	if err != nil {
		log.Println(err)
	}

	res, err := testClient.Do(req)
	if err != nil {
		log.Println(err)
	}

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		log.Println(err)
	}

	err = json.Unmarshal(body, &cases)
	if err != nil {
		log.Println(err)
	}

	for _, estado := range cases[0].Estados {

		if estado.UF == strings.ToUpper(uf) {

			return estado.Casos, estado.Mortes
		}

	}
	return 0, 0
}
