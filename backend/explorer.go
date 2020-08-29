package main

import (
	"log"

	"github.com/shinecloudnet/explorer/backend/rest"
	"github.com/shinecloudnet/explorer/backend/task"
)

func main() {
	task.Start()
	server := rest.NewApiServer()

	if err := server.Start(); err != nil {
		log.Fatal("ListenAndServe Failed: ", err)
	}
}
