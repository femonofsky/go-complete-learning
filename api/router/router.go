package router

import (
	"net/http"

	"github.com/gorilla/mux"
)

const (
	staticDir = "/static"
)

// RouteHandler - the handler for go routes
type RouteHandler struct {
	Router *mux.Router
}

// NewRouter - create a new router
func NewRouter() *RouteHandler {
	var router RouteHandler

	router.Router = mux.NewRouter().StrictSlash(true)

	router.Router.PathPrefix(staticDir).
		Handler(http.StripPrefix(staticDir,
			http.FileServer(http.Dir("."+staticDir))))

	return &router
}
