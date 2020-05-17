package control

import (
	"alpha01/configs"
	"alpha01/database"
	"alpha01/models"
	"encoding/json"
	"io/ioutil"
	"net/http"
)

//obter resposta do teste e retornar o resultado
func PostResult(w http.ResponseWriter, r *http.Request) {
	//pegue o resultado do teste
	body := r.Body
	bytes, _ := ioutil.ReadAll(body)

	// instancie o objeto de resultado
	var result models.Result

	// insira o resultado no objeto
	err := json.Unmarshal(bytes, &result)

	// notifique caso ocorra algum erro
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(configs.RESPONSE_ERRO))
		return
	}

	//insira a resposta no banco de dados
	err = database.InsertResult(result)

	// notifique se ocorrer um erro no banco de dados
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(configs.RESPONSE_ERRO))
		return
	}

	//calcule o resultado
	rsp := CreateResponse(result.User_Score)

	// transforme em JSON
	response, err := json.Marshal(rsp)

	// notifique se ocorrer um erro com o JSON
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(configs.RESPONSE_ERRO))
		return
	}

	//envie para o usuário
	w.Write([]byte(response))
}

func CreateResponse(score float64) (r models.Response) {
	//inserir a pontuação no response http
	r.User_Score = score

	// Inserir as estatísticas da nossa base de dados
	r.Internal_Stats = GetStats()

	// inserir os dados de outros sites estatísticos

	return
}

func GetStats()(stats models.Stats){
	var results []models.Result

	results, err := database.SearchAllResults()

	if err != nil{
		return
	}

	stats.Average = AverageResults(results)
	stats.Median = MedianResults(results)
	a,b := BoundsResults(results)
	stats.Max = a
	stats.Min = b 

	return
}
