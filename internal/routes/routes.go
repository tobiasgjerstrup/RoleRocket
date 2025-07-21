package routes

import (
	"encoding/json"
	"net/http"
)

func Routes() *http.ServeMux {
	router := http.NewServeMux()

	router.HandleFunc("GET /users", func(w http.ResponseWriter, r *http.Request) {
		users, err := GetUsers(r.Context())
		if err != nil {
			http.Error(w, "Failed to fetch users", http.StatusInternalServerError)
			return
		}

		jsonBytes, err := json.Marshal(users)
		if err != nil {
			http.Error(w, "failed to encode users", http.StatusInternalServerError)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		w.Write(jsonBytes)
	})
	router.HandleFunc("PUT /users/{id}", func(w http.ResponseWriter, r *http.Request) {
		id := r.PathValue("id")
		w.Write([]byte("Update user with ID: " + id))
	})
	router.HandleFunc("POST /users", func(w http.ResponseWriter, r *http.Request) {
		InsertUser(r.Context())
		w.Write([]byte("Create user"))
	})
	router.HandleFunc("DELETE /users/{id}", func(w http.ResponseWriter, r *http.Request) {
		id := r.PathValue("id")
		w.Write([]byte("Delete user with ID: " + id))
	})

	return router
}
