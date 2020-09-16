package main

import (
	"github.com/Thomgrus/share-your-ci/back/controllers"

	"github.com/gorilla/mux"
)

// InitializeRouter : init router
func InitializeRouter() *mux.Router {
	// StrictSlash is true => redirect /Users/ to /Users
	router := mux.NewRouter().StrictSlash(true)

	// Health Check
	router.Methods("GET").Path("/up").Name("Health").HandlerFunc(controllers.HealthCheck)

	// PresPart
	router.Methods("GET").Path("/api/pres/part").Name("PresPart").HandlerFunc(controllers.GetPresPart)
	return router
}
