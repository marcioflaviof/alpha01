package models

type Result struct {
	Exam_ID    string  `json: "exam_id"`
	Geo_Lat    float64 `json: "geo_lat"`
	Geo_Lon    float64 `json: "geo_lon"`
	User_Name  string  `json: "user_name"`
	User_Score float64 `json: "user_score"`
	User_Local string  `json: "user_local"`
}

type Geo struct{
	latitude   float64
	longitude  float64
}

type Response struct {
	User_Score float64 `json: "user_score"`
	Internal_Stats Stats
	//dados externos
}

type Stats struct{
	/////////////////////////////
	// numero de resultados    //
	// media de resultados     //
	// mediana de resultados   //
	// valor máximo            //
	// valor minimo            //
	/////////////////////////////
	
	Sample  int       `json: "sample"`
	Average float64   `json: "average"`
	Median  float64   `json: "median"`
	Max     float64   `json: "max"`
	Min     float64   `json: "min"`
}