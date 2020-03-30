package routes

import (
	"net/http"

	"xorm.io/xorm"
)

// Health - health routes
func Health(db *xorm.Engine) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Server is Healthy!\n"))
	}
}
