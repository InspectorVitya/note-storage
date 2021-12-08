package main

import (
	"github.com/inspectorvitya/note-storage/internal/application"
	httpserver "github.com/inspectorvitya/note-storage/internal/server"
	"github.com/inspectorvitya/note-storage/internal/storage/memory"
	"log"
)

func main() {
	storage := memory.New()
	app := application.New(storage)
	server := httpserver.New("8080", app)
	if err := server.Start(); err != nil {
		log.Fatalln(err)
	}
}
