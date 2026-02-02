package years

import (
	"database/sql"
	"encoding/json"
	"net/http"
)

func GetYears(db *sql.DB) http.HandlerFunc {
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
				model_id
			FROM years
		`)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		defer rows.Close()

		years := []Year{}

		for rows.Next() {
			var y Year
			if err := rows.Scan(
				&y.Id,
				&y.Description,
				&y.Fipe,
				&y.Model_id,
			); err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}
			years = append(years, y)
		}

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(years)
	}
}