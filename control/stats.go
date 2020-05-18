package control

import(
	"alpha01/models"
	"sort")

// ordene a lista de resultados
func SortResults(results []models.Result){
	sort.SliceStable(results,func(i, j int) bool { return results[i].User_Score < results[j].User_Score })
	return
}

// retorne o somatório dos valores de uma lista de resultados
func SumResults(results []models.Result)(sum float64){
	for _,rslt := range results{
		sum+=rslt.User_Score
	}
	return
}


// retorne o maior e o menor valor de uma lista de resultados
func BoundsResults(results []models.Result)(max float64,min float64){
	SortResults(results)
	max = results[(len(results)-1)].User_Score
	min = results[0].User_Score
	return
}

// retorne a média de uma lista de resultados
func AverageResults(results []models.Result)(avg float64){
	sum := SumResults(results)
	total := float64(len(results))
	avg = sum/total
	return
}

// retorne a mediana de uma lista de resultados
func MedianResults(results []models.Result)(mdn float64){
	SortResults(results)
	length := len(results)
	middle := length/2
	if length & 2 == 0 {
		mdn = results[(length/2)].User_Score
	}else{
		a := results[middle].User_Score
		b := results[(middle-1)].User_Score
		mdn = (a+b)/2.0
	}
	return
}