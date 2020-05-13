package database

import ( "go.mongodb.org/mongo-driver/bson"
		 "alpha01/models"
		 "alpha01/configs"
         "context"
         "log" )

func SearchTestID(id string) (t models.Test, err error){

	// select tests collection
	c := Db.Collection(configs.TEST_COLLECTION)

	// try to find test with an test id
	fnd := c.FindOne(context.TODO(), bson.D{{"test_id", id}})
	err = fnd.Decode(&t)

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
	c := Db.Collection(configs.TEST_COLLECTION)

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