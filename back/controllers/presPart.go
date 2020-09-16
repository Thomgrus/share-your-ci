package controllers

import (
	"encoding/json"
	"net/http"

	"github.com/Thomgrus/share-your-ci/back/models"
)

// GetPresPart : endpoint for get presPart by Code
func GetPresPart(w http.ResponseWriter, r *http.Request) {
	query := r.URL.Query()
	code := query.Get("code")

	presPart := models.FindPresPartByCode(code)

	// We test patient uid
	if presPart == nil {
		w.WriteHeader(http.StatusNotFound)
		return
	}

	w.Header().Set("Content-type", "application/json;charset=UTF-8")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(presPart)
}
