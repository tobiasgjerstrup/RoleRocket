package routes

import "net/http"

func Routes() *http.ServeMux {
	router := http.NewServeMux()

	router.HandleFunc("GET /users", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("List users"))
	})
	router.HandleFunc("PUT /users/{id}", func(w http.ResponseWriter, r *http.Request) {
		id := r.PathValue("id")
		w.Write([]byte("Update user with ID: " + id))
	})
	router.HandleFunc("POST /users", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Create user"))
	})
	router.HandleFunc("DELETE /users/{id}", func(w http.ResponseWriter, r *http.Request) {
		id := r.PathValue("id")
		w.Write([]byte("Delete user with ID: " + id))
	})

	return router
}
