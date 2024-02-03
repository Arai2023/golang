package main

import (
	"fmt"
	"net/http"

	"github.com/yourusername/go-wizard-app/api"
)

func main() {
	fmt.Println("Wizard Web App is running on :8080")
	http.ListenAndServe(":8080", api.Router())
}
