package server

import (
	"alpha01/configs"
	"log"
	"net/http"
	"time"
)

// Start comeca o codigo
func Start() {

	h := createHandler()

	s := createServer()

	s.Handler = h

	log.Fatal(s.ListenAndServe())

}

func createServer() (server *http.Server) {

	server = &http.Server{

		Addr:         configs.SERVER_ADDR,
		IdleTimeout:  1000 * time.Millisecond,
		ReadTimeout:  1000 * time.Millisecond,
		WriteTimeout: 1000 * time.Millisecond,
	}

	return
}
