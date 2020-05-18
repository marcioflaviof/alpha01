package control

import (
	"alpha01/configs"
	"alpha01/database"
	"alpha01/models"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"log"
)

//obter resposta do teste e retornar o resultado
func PostResult(w http.ResponseWriter, r *http.Request) {
	
	// instancie o objeto de resultado
	var result models.Result
	
	//pegue o resultado do teste
	body := r.Body
	bytes, err := ioutil.ReadAll(body)

	// notifique caso ocorra algum erro
	if err != nil{
		log.Printf("[ERRO] Cannot unmarshal the JSON: %v %v",err,result)
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(configs.RESPONSE_ERRO))
		return
	}

	// insira o resultado no objeto
	err = json.Unmarshal(bytes, &result)

	// notifique caso ocorra algum erro
	if err != nil {
		log.Printf("[ERRO] Cannot unmarshal the JSON: %v %v",err,result)
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

	// adicione o contador de resultados
	err = database.IncrementExamPreview(result.Exam_ID)

	// notifique se ocorrer um erro no banco de dados
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(configs.RESPONSE_ERRO))
		return
	}

	// crie a response com os dados estatísticos
	rsp := CreateResponse(result)

	// transforme em JSON
	response, err := json.Marshal(rsp)

	// notifique se ocorrer um erro com o JSON
	if err != nil {
		log.Printf("[ERRO] Cannot parse the response to JSON: %v %v",err,rsp)
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(configs.RESPONSE_ERRO))
		return
	}

	//envie para o usuário
	w.Write([]byte(response))
}

func CreateResponse(res models.Result) (r models.Response) {
	//inserir a pontuação no response http
	r.User_Score = res.User_Score

	// Inserir as estatísticas da nossa base de dados
	r.Internal_Stats = GetStats(res.Exam_ID)

	// inserir os dados de outros sites estatísticos

	return
}

func GetStats(id string)(stats models.Stats){
	//instancia um slice de results
	var results []models.Result

	// puxe todos os resultados do banco de dados
	results, err := database.SearchResultsById(id)
	if err != nil{
		return
	}

	//distribua as medidas estatisticas
	stats.Sample = len(results)
	//media
	stats.Average = AverageResults(results)
	//mediana
	stats.Median = MedianResults(results)
	a,b := BoundsResults(results)
	//valor máximo
	stats.Max = a
	//valor Mínimo
	stats.Min = b 

	return
}