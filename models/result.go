package models

type Result struct {
	Geo_Lat    float64 `json: "geo_lat"`
	Geo_Lon    float64 `json: "geo_lon"`
	User_Name  string  `json: "user_name"`
	User_Score float64 `json: "user_score"`
}

type Response struct {
	User_Score float64 `json: "user_score"`
	Internal_Stats Stats
	//dados externos
}

type Stats struct{
	Average float64   `json: "average"`
	Median float64 `json: "median"`
	Max float64    `json: "max"`
	Min float64    `json: "min"`
}