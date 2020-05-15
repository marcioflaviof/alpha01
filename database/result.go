package database

import (
	"alpha01/configs"
	"alpha01/models"
	"context"
	"log"

	"go.mongodb.org/mongo-driver/bson"
)

func InsertResult(r models.Result) (err error) {

	//utilize a coleção de resultados
	c := Db.Collection(configs.RESULT_COLLECTION)

	// insira o resultado
	rslt, err := c.InsertOne(context.TODO(), r)

	// notifique caso ocorra algum erro
	if err != nil {
		log.Printf("[ERRO] cannot insert Result: %v %v", r, err)
		return
	}

	// informe que finalizou a operação
	log.Printf("[INFO] a result was inserted: %v", rslt)

	return
}

func SearchResult(name string) (r models.Result, err error) {

	// utilize a coleção de resultados
	c := Db.Collection(configs.RESULT_COLLECTION)

	// busque o resultado com o nome passado
	err = c.FindOne(context.TODO(), bson.D{{"user_name", name}}).Decode(&r)

	// notifique caso ocorra algum erro
	if err != nil {
		log.Printf("[ERROR] cannot find a Result: %v %v", r, err)
		return
	}

	// apresente que ação foi realizada
	log.Printf("[INFO] test searched: %v", r)

	return
}
