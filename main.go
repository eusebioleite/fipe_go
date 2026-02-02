package main

import (
	"fmt"
	"net/http"

	"github.com/eusebioleite/fipe_go/brands"
	"github.com/eusebioleite/fipe_go/db"
	"github.com/eusebioleite/fipe_go/models"
	"github.com/eusebioleite/fipe_go/references"
	"github.com/eusebioleite/fipe_go/types"
	"github.com/eusebioleite/fipe_go/years"
	"github.com/go-chi/chi/v5"
)

func main() {
	database := db.Open()
	r := chi.NewRouter()

	// GET Endpoints
	r.Get("/types", types.GetTypes(database))
	r.Get("/references", references.GetReferences(database))
	r.Get("/brands", brands.GetBrands(database))
	r.Get("/models", models.GetModels(database))
	r.Get("/years", years.GetYears(database))


	fmt.Println("Server running on http://localhost:8080/types")
	fmt.Println("Server running on http://localhost:8080/references")
	fmt.Println("Server running on http://localhost:8080/brands")
	fmt.Println("Server running on http://localhost:8080/models")
	fmt.Println("Server running on http://localhost:8080/years")
	http.ListenAndServe(":8080", r)
}
