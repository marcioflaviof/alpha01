package models

type Test struct{
	TestID string `json:"testid"`
	Name string `json:"name"`
	Questions []QuestionList
}

type QuestionList struct{
	Text string `json:"text"`
	Options []OptionList
}

type OptionList struct{
	Description string `json:"description"`
}