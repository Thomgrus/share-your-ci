package main

import (
	"log"
	"net/http"
	"os"

	"github.com/Thomgrus/share-your-ci/back/config"
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
	router := InitializeRouter()

	// Listen
	log.Fatal(http.ListenAndServe(getPort(), router))
}
