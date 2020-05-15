package models

type Exam struct{
	Exam_ID string `json:"exam_id"`
	Exam_Name string `json:"exam_name"`
	Exam_Questions []QuestionList
}

type QuestionList struct{
	Title string `json:"title"`
	Options []string `json:"options"`
}