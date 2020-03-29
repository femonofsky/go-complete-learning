package routes

import "net/http"

// Health - health routes
func Health() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Server is Healthy!\n"))
	}
}
