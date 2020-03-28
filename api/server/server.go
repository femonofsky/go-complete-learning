package server

import (
	"log"
	"net/http"
	"os"
	"strconv"
	"time"

	RouterFactory "github.com/femonofsky/go-complete-learning/api/router"
	"github.com/gorilla/handlers"
)

// Server - this  is the server object
type Server struct {
	Port       int
	Addr       string
	HTTPServer *http.Server
}

// Start - starts the server service
func (s *Server) Start() {
	log.Println("Server started on port ", s.Port)
	log.Fatal(s.HTTPServer.ListenAndServe())
}

// NewServer - create a new server
func NewServer(port int) *Server {
	var server Server

	server.Port = port
	server.Addr = ":" + strconv.Itoa(port)

	router := RouterFactory.NewRouter()

	handler := handlers.LoggingHandler(os.Stdout, handlers.CORS(
		handlers.AllowedOrigins([]string{'*'}),
		handlers.AllowedMethods([]string{"GET", "POST","PUT","DELETE", "OPTIONS","PATCH"}),
		handlers.AllowedHeaders([]string{"Content-Type","Origin","Cache-Control","X-App-Token"})

	)(router.Router))

	server.HTTPServer = &http.Server{
		Addr:           server.Addr,
		Handler:        router.Router,
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}

	return &server
}
