package brands

import (
	"database/sql"
	"encoding/json"
	"net/http"
)

func GetBrands(db *sql.DB) http.HandlerFunc {
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
				type_id,
				ref_id
			FROM brands
		`)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		defer rows.Close()

		brands := []Brand{}

		for rows.Next() {
			var b Brand
			if err := rows.Scan(
				&b.Id,
				&b.Description,
				&b.Fipe,
				&b.Type_id,
				&b.Ref_id,
			); err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}
			brands = append(brands, b)
		}

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(brands)
	}
}