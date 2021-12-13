package main

import (
	"flag"
	"github.com/inspectorvitya/note-storage/internal/application"
	httpserver "github.com/inspectorvitya/note-storage/internal/server"
	"github.com/inspectorvitya/note-storage/internal/storage/memory"
	"log"
)

func main() {
	var port string
	flag.StringVar(&port, "port", "8080", "port for http servers")
	flag.Parse()

	storage := memory.New()
	app := application.New(storage)
	server := httpserver.New(port, app)

	if err := server.Start(); err != nil {
		log.Fatalln(err)
	}
}
