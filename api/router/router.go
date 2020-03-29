package router

import (
	"net/http"

	Routes "github.com/femonofsky/go-complete-learning/api/router/routes"
	V1Routes "github.com/femonofsky/go-complete-learning/api/router/routes/v1"
	"github.com/gorilla/mux"
)

const (
	staticDir = "/static"
)

// RouteHandler - the handler for go routes
type RouteHandler struct {
	Router *mux.Router
}

// AttachSubRouterWithMiddleware - allows you to attach
// 		subroutes to router
func (r *RouteHandler) AttachSubRouterWithMiddleware(
	path string,
	subroutes Routes.Routes,
	middleware mux.MiddlewareFunc,
) (SubRouter *mux.Router) {

	SubRouter = r.Router.PathPrefix(path).Subrouter()
	SubRouter.Use(middleware)

	for _, route := range subroutes {
		SubRouter.
			Methods(route.Method).
			Path(route.Pattern).
			Name(route.Name).
			Handler(route.HandlerFunc)
	}

	return SubRouter

}

// NewRouter - create a new router
func NewRouter() *RouteHandler {
	var router RouteHandler

	router.Router = mux.NewRouter().StrictSlash(true)

	router.Router.PathPrefix(staticDir).
		Handler(http.StripPrefix(staticDir,
			http.FileServer(http.Dir("."+staticDir))))

	router.Router.Use(Routes.Middleware)

	routes := Routes.GetRoutes()

	for _, route := range routes {
		router.Router.
			Methods(route.Method).
			Path(route.Pattern).
			Name(route.Name).
			Handler(route.HandlerFunc)
	}
	v1subroutes := V1Routes.GetRoutes()

	for name, pack := range v1subroutes {
		router.AttachSubRouterWithMiddleware(name, pack.Routes, pack.Middleware)
	}

	return &router
}
