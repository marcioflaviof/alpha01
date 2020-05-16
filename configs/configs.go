package configs

//=========================
//	Server HTTP
//=========================
const (

	// defines ip and port address for server instance
	SERVER_ADDR = ":"

	// host for mongo db
	MONGO_HOST = "mongodb+srv://admin:1010@cluster0-ga3ne.mongodb.net/COVIDTests?retryWrites=true&w=majority"
)

//========================
//  DB Collections
//========================
const (
	EXAM_COLLECTION   = "exams"
	RESULT_COLLECTION = "results"
)

//========================
//  Default Messages
//========================
const (
	RESPONSE_ERRO = `{"msg":"ocorreu um erro"}`
)
