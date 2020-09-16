package main

import (
	"log"
	"net/http"
	"os"

	"github.com/Thomgrus/share-your-ci/back/config"
	"github.com/Thomgrus/share-your-ci/back/models"
)

func getPort() string {
	p := os.Getenv("PORT")
	if p != "" {
		return ":" + p
	}
	return ":8080"
}

func main() {
	config.DatabaseInit()

	// Initialize Data
	models.InitPresPartData()

	router := InitializeRouter()

	// Listen
	log.Fatal(http.ListenAndServe(getPort(), router))
}
