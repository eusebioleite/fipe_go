package types

import (
	"database/sql"
	"encoding/json"
	"net/http"
)

func GetTypes(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodGet {
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
			return
		}

		rows, err := db.Query(`
			SELECT
				id,
				description
			FROM types
		`)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		defer rows.Close()

		types := []Type{}

		for rows.Next() {
			var t Type
			if err := rows.Scan(
				&t.Id,
				&t.Description,
			); err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}
			types = append(types, t)
		}

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(types)
	}
}