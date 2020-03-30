package routes

import (
	"log"
	"net/http"

	"xorm.io/xorm"
)

// Login - Login route
func Login(db *xorm.Engine) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		r.ParseForm()

		email := r.FormValue("email")
		password := r.FormValue("password")

		if len(email) == 0 || len(password) {
			http.Error(w, "An email and password are required", http.StatusBadRequest)
			return
		}

		Users{}

		log.Println(email)
		log.Println(password)
	}
}
