package configs

//=========================
//	Server HTTP
//=========================
const (

	// defines ip and port address for server instance
	//SERVER_ADDR = ":"
	SERVER_ADDR = "localhost:8080"

	// host for mongo db
	MONGO_HOST = "mongodb+srv://admin:1010@cluster0-ga3ne.mongodb.net/COVIDTests?retryWrites=true&w=majority"
	//MONGO_HOST = "mongodb://localhost:27017"
)

//========================
//  DB Collections
//========================
const (
	EXAM_COLLECTION   = "exams"
	RESULT_COLLECTION = "results"
	EPREV_COLLECTION  = "eprevs"
)

//========================
//  Default Messages
//========================
const (
	RESPONSE_ERRO = `{"msg":"ocorreu um erro"}`
	MARSHAL_ERROR = `{"msg":"marshal error"}`
	UNMARSHAL_ERROR = `{"msg":"unmarshal error"}`
	DBUPD_ERROR = `{"msg":"update error in data base"}`
	DBINS_ERROR = `{"msg":"insert erro in data base"}`
	DBSRC_ERROR = `{"msg":"search error in data base"}`
)
