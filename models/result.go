package models

type Result struct {
	Geo_Lat    float32 `json: "geo_lat"`
	Geo_Lon    float32 `json: "geo_lon"`
	User_Name  string  `json: "user_name"`
	User_Score float32 `json: "user_score"`
}

type Response struct {
	User_Score float32 `json: "user_score"`
	//dados internos
	//dados externos
}
