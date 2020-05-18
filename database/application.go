package database

import (
	"go.mongodb.org/mongo-driver/bson"
	"alpha01/configs"
	"alpha01/models"
	"context"
	"log"
)

func CreateExamsPreview(id string) (err error){
	// instancia o objeto ExamPreview
	var ep models.ExamPreview

	// cria os valores iniciais do Exam Preview
	ep.Exam_ID = id
	ep.Number_Results = 0

	//utilize a coleção de Previsão de resultados
	c := Db.Collection(configs.EPREV_COLLECTION)

	// insira o resultado
	eprev, err := c.InsertOne(context.TODO(), ep)

	// notifique caso ocorra algum erro
	if err != nil {
		log.Printf("[ERRO] cannot insert Exam Preview: %v %v", err, eprev)
		return
	}

	// informe que finalizou a operação
	log.Printf("[INFO] a result was inserted: %v", ep)

	return
}

func SearchAllExamsPreview() (eps []models.ExamPreview, err error){

	// utilize a coleção de Exam Preview
	c := Db.Collection(configs.EPREV_COLLECTION)

	// Busque no banco de dados
	cur, err := c.Find(context.TODO(),bson.D{{}})
	if err != nil {
		log.Printf("[ERRO] cannot search Exam Preview %v %v", err, cur)
		return
	}

	// Coloque no modelo
	err = cur.All(context.TODO(),&eps)
	if err != nil{
		log.Printf("[ERRO] cannot search Exam Preview %v %v", err, eps)
		return
	}

	// informe que finalizou a operação
	log.Printf("[INFO] searched all Exam Preview")

	return
}

func IncrementExamPreview(id string)(err error){

	// instancie o modelo
	var ep models.ExamPreview

	// utilize a coleção de exam preview
	c := Db.Collection(configs.EPREV_COLLECTION)

	// busque no banco de dados o preview solicitado e insira no modelo
	err = c.FindOne(context.TODO(),bson.D{{"exam_id",id}}).Decode(&ep)

	// notifique caso ocorra um erro
	if err != nil{
		log.Printf("[ERRO] cannot find the Exam Preview: %v %v", err, ep)
	}

	// incremente o numero de resultados do modelo
	ep.Number_Results += 1

	// atualize o banco de dados com o modelo utilizado
	err = c.FindOneAndUpdate(context.TODO(), bson.D{{"exam_id",ep.Exam_ID}}, bson.D{{"$set",bson.D{{"number_results",ep.Number_Results}}}}).Decode(&ep)
	
	// notifique caso ocorra um erro
	if err != nil{
		log.Printf("[ERRO] cannot update the Exam Preview: %v %v",err, ep)
	}

	// informe que a operação foi realizada
	log.Print("[INFO] updated Exam Preview: %v",ep)

	return
}