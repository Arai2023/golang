package api

import (
	"encoding/json"
	"net/http"
	"strings"

	"github.com/gorilla/mux"
	"github.com/yourusername/go-wizard-app/internal/wizard"
)

var wizards = []wizard.Wizard{
	{...},
	{...},
	{...},
}

// HealthCheckHandler returns a simple health check message.
func HealthCheckHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Wizard Web App is running! - Created by OpenAI's ChatGPT"))
}

// ListWizardsHandler returns a JSON list of wizards.
func ListWizardsHandler(w http.ResponseWriter, r *http.Request) {
	// ...
}

// GetWizardHandler returns JSON information about a specific wizard.
func GetWizardHandler(w http.ResponseWriter, r *http.Request) {
	// ...
}

// Router returns the configured router for the API.
func Router() *mux.Router {
	router := mux.NewRouter()
	router.HandleFunc("/health-check", HealthCheckHandler).Methods("GET")
	router.HandleFunc("/wizards", ListWizardsHandler).Methods("GET")
	router.HandleFunc("/wizards/{name}", GetWizardHandler).Methods("GET")
	return router
}
