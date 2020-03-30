package v1

import (
	"net/http"

	Routes "github.com/femonofsky/go-complete-learning/api/router/routes"
	"xorm.io/xorm"
)

// Middleware - this is the middleware for v1 application
func Middleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		next.ServeHTTP(w, r)
	})
}

// GetRoutes - get v1 routes
func GetRoutes(db *xorm.Engine) (SubRoute map[string]Routes.SubRoutePackage) {
	SubRoute = map[string]Routes.SubRoutePackage{
		"/v1": {
			Routes: Routes.Routes{
				Routes.Route{
					Name:        "V1HealthRoute",
					Method:      "GET",
					Pattern:     "/health",
					HandlerFunc: Health(db),
				},
			},
			Middleware: Middleware,
		},
	}
	return
}
