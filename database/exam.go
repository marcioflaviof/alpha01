package database

import (
	"alpha01/configs"
	"alpha01/models"
	"context"
	"log"

	"go.mongodb.org/mongo-driver/bson"
)

func SearchExamID(id string) (t models.Exam, err error) {

	// select exams collection
	c := Db.Collection(configs.EXAM_COLLECTION)

	// try to find exam with an exam id
	fnd := c.FindOne(context.TODO(), bson.D{{"exam_id", id}})
	err = fnd.Decode(&t)

	// checks if any error occurs searching a exam
	if err != nil {
		log.Printf("[ERROR] probleming searching exam: %v %v", err, t)
		return
	}

	// show the action of search a exam
	log.Printf("[INFO] exam searched: %v %v", err, t)

	return
}

func insertExam(t models.Exam) (err error) {

	// select the collection of exams
	c := Db.Collection(configs.EXAM_COLLECTION)

	// insert exam in collection of exams
	exam, err := c.InsertOne(context.TODO(), t)

	// check if not error occurs
	if err != nil {
		log.Printf("[ERROR] probleming in insertexam: %v, %v", err, t)
	}

	// confirm action of inserting a exam
	log.Printf("[INFO] exam inserted: %v", exam)

	return
}

func SearchAllExams() (exams []models.Exam, err error) {
	c := Db.Collection(configs.EXAM_COLLECTION)

	cursor, err := c.Find(context.TODO(), bson.D{})

	// check if not error occurs
	if err != nil {
		log.Printf("[ERROR] probleming searching All exams: %v, %v", err, cursor)
		return
	}

	err = cursor.All(context.TODO(), &exams)

	// check if not error occurs
	if err != nil {
		log.Printf("[ERROR] probleming searching All exams: %v, %v", err, cursor)
		return
	}

	log.Print("[INFO] All exam searched")

	return
}
