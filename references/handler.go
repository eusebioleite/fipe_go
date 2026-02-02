package references

import (
	"database/sql"
	"encoding/json"
	"net/http"
)

func GetReferences(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodGet {
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
			return
		}

		rows, err := db.Query(`
			SELECT
				id,
				month,
				year,
				fipe
			FROM "references"
		`)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		defer rows.Close()

		references := []Reference{}

		for rows.Next() {
			var r Reference
			if err := rows.Scan(
				&r.Id,
				&r.Month,
				&r.Year,
				&r.Fipe,
			); err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}
			references = append(references, r)
		}

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(references)
	}
}