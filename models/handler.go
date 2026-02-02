package models

import (
	"database/sql"
	"encoding/json"
	"net/http"
)

func GetModels(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodGet {
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
			return
		}

		rows, err := db.Query(`
			SELECT
				id,
				description,
				fipe,
				brand_id
			FROM models
		`)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		defer rows.Close()

		models := []Model{}

		for rows.Next() {
			var m Model
			if err := rows.Scan(
				&m.Id,
				&m.Description,
				&m.Fipe,
				&m.Brand_id,
			); err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}
			models = append(models, m)
		}

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(models)
	}
}