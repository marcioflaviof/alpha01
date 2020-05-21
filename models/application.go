package models

type Menu struct {
	Avaliable_Exams []ExamPreview `json: "avaliable_exams"`
}

type ExamPreview struct{
	Exam_ID string      `json: "exam_id"`
	Number_Results int  `json: "number_results"`
}

type ExamPreviewUpdate struct{
	Updated ExamPreview `json: "updated"`
}