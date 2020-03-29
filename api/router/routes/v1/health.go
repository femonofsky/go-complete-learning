package v1

import "net/http"

// Health - health routes
func Health() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("V1 Routes are active! \n"))
	}
}
