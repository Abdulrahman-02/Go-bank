package main

import (
	"github.com/Abdulrahman-02/Go-bank/internal/api"
	"github.com/Abdulrahman-02/Go-bank/internal/storage"
)

func main() {
	// Create db
	store, err := storage.NewPostgresStorage()
	if err != nil {
		panic(err)
	}

	server := api.NewAPIserver(":3000", store)
	server.Run()
}
