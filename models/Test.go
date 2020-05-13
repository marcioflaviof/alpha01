package models

type Test struct{
	Test_ID string `json:"test_id"`
	Test_Name string `json:"test_name"`
	Test_Questions []QuestionList
}

type QuestionList struct{
	Title string `json:"title"`
	Options []string `json:"options"`
}