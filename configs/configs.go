package configs

//=========================
//	Server HTTP
//=========================
const (

	// defines ip and port address for server instance
	SERVER_ADDR = "localhost:8080"

	// host for mongo db
	MONGO_HOST = "mongodb://localhost:27017"
)

//=========================
//	HTTP Routes
//=========================
const (
	USER_PATH  = "/user/"
	LOGIN_PATH = "/login/"
)

//========================
//  DB Collections
//========================
const (
	TEST_COLLECTION = "tests"
	RESULT_COLLECTION = "results"
)

//========================
//  Default Messages
//========================
const (
	RESPONSE_ERRO = `{"msg":"ocorreu um erro"}`
)
