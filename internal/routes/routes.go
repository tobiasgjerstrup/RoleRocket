package routes

import (
	"encoding/json"
	"net/http"
	"os"
	sqlite "rolerocket/internal/db"
	"strings"
)

func Routes(debugMode bool) *http.ServeMux {
	router := http.NewServeMux()

	router.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		if strings.HasPrefix(r.URL.Path, "/users") {
			http.NotFound(w, r)
			return
		}

		// Try to serve built files first
		publicPath := "./public" + r.URL.Path
		if _, err := os.Stat(publicPath); err == nil {
			http.ServeFile(w, r, publicPath)
			return
		}

		// Fallback to index.html for SPA routes
		http.ServeFile(w, r, "./public/index.html")
	})

	router.HandleFunc("GET /users", func(w http.ResponseWriter, r *http.Request) {
		users, err := GetUsers(w, r)
		if err != nil {
			respondWithError(w, "Failed to fetch users", http.StatusInternalServerError)
			return
		}

		jsonBytes, err := json.Marshal(users)
		if err != nil {
			respondWithError(w, "failed to encode users", http.StatusInternalServerError)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		w.Write(jsonBytes)
	})
	/* router.HandleFunc("PUT /users/{id}", func(w http.ResponseWriter, r *http.Request) {
		id := r.PathValue("id")
		w.Write([]byte("Update user with ID: " + id))
	}) */
	router.HandleFunc("POST /users", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		err := InsertUser(w, r)
		if err != nil {
			return
		}

		response := map[string]string{"message": "User created!"}
		json.NewEncoder(w).Encode(response)
	})
	/* router.HandleFunc("DELETE /users/{id}", func(w http.ResponseWriter, r *http.Request) {
		id := r.PathValue("id")
		w.Write([]byte("Delete user with ID: " + id))
	}) */
	router.HandleFunc("POST /users/token", func(w http.ResponseWriter, r *http.Request) {
		GetToken(w, r)
	})

	if debugMode {
		router.HandleFunc("GET /debug/query/{query}", func(w http.ResponseWriter, r *http.Request) {
			query := r.PathValue("query")

			rows, err := sqlite.DBInstance.Conn.Query(query)
			if err != nil {
				respondWithError(w, "Failed running debug query", http.StatusInternalServerError)
				return
			}
			defer rows.Close()

			columns, err := rows.Columns()
			if err != nil {
				respondWithError(w, "Failed to get columns", http.StatusInternalServerError)
				return
			}

			var results []map[string]interface{}

			for rows.Next() {
				values := make([]interface{}, len(columns))
				valuePtrs := make([]interface{}, len(columns))

				for i := range columns {
					valuePtrs[i] = &values[i]
				}

				if err := rows.Scan(valuePtrs...); err != nil {
					respondWithError(w, "Error scanning row", http.StatusInternalServerError)
					return
				}

				rowMap := make(map[string]interface{})
				for i, col := range columns {
					val := values[i]

					// Handle []byte as string
					if b, ok := val.([]byte); ok {
						rowMap[col] = string(b)
					} else {
						rowMap[col] = val
					}
				}
				results = append(results, rowMap)
			}

			jsonBytes, err := json.Marshal(results)
			if err != nil {
				respondWithError(w, "Failed to encode result", http.StatusInternalServerError)
				return
			}
			w.Header().Set("Content-Type", "application/json")
			w.Write(jsonBytes)
		})
	}

	return router
}
