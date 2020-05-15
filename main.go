package main

import ("alpha01/server"
		"alpha01/database")

func main() {
	database.CreateClient()
	server.Start()
}
