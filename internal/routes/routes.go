package routes

import (
	"encoding/json"
	"net/http"
)

func Routes() *http.ServeMux {
	router := http.NewServeMux()

	router.HandleFunc("GET /users", func(w http.ResponseWriter, r *http.Request) {
		users, err := GetUsers(w, r)
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
		InsertUser(w, r)
	})
	router.HandleFunc("DELETE /users/{id}", func(w http.ResponseWriter, r *http.Request) {
		id := r.PathValue("id")
		w.Write([]byte("Delete user with ID: " + id))
	})
	router.HandleFunc("POST /users/token", func(w http.ResponseWriter, r *http.Request) {
		GetToken(w, r)
	})

	return router
}
