package database

import ( "go.mongodb.org/mongo-driver/bson"
		 "alpha01/models"
         "context"
         "log" )

func SearchTestID(id string) (t models.Test, err error){

	// select tests collection
	c := Db.Collection("tests")

	// try to find test with an test id
	err = c.FindOne(context.TODO(), bson.D{{"teste_id", id}}).Decode(&t)

	// checks if any error occurs searching a test
	if err != nil {
		log.Printf("[ERROR] probleming searching test: %v %v", err, t)
		return
	}

	// show the action of search a test
	log.Printf("[INFO] test searched: %v %v", err, t)

	return
}

func insertTest(t models.Test) (err error){

	// select the collection of tests
	c := Db.Collection("tests")

	// insert test in collection of tests
	test, err := c.InsertOne(context.TODO(),t)

	// check if not error occurs
	if err != nil {
		log.Printf("[ERROR] probleming in insertTest: %v, %v",err,t)
	}

	// confirm action of inserting a test
	log.Printf("[INFO] test inserted: %v",test)

	return
}

func insertResult(t models.Test) (err error){
	// select the collection of results
	c := Db.Collection("results")

	// insert the answered test in the collection
	res, err := c.InsertOne(context.TODO(),t)

	// check if none error occurs
	if err != nil {
		log.Printf("[ERROR] probleming in insertResult: %v, %v",err,t)
	}

	// show action to insert a result
	log.Printf("[INFO] Result inserted: %v",res)

	return
}