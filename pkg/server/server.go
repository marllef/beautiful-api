package server

import (
	"fmt"
	log "marllef/beautiful-api/pkg/mylogger"
	"net/http"

	"github.com/gorilla/mux"
)

type Server interface {
	SetRoutes(routes Routes)
	SetPrefix(prefix string)
	SetLogger(logger *log.Logger)
	SetPort(port string)
	Serve()
}

type server struct {
	routes Routes
	port   string
	logger *log.Logger
	prefix string
}

// Create a new server.
func NewServer() *server {
	return &server{
		routes: make(Routes),
		port:   "3005",
		prefix: "",
		logger: log.Default(),
	}
}

// Set server routes.
func (s *server) SetRoutes(routes Routes) {
	s.routes = routes
}

// Set route prefix.
func (s *server) SetPrefix(prefix string) {
	s.prefix = fmt.Sprintf("%s", prefix)
}

// Set server logger.
func (s *server) SetLogger(logger *log.Logger) {
	s.logger = logger
}

// Set server port.
func (s *server) SetPort(port string) {
	s.port = port
}

func (s *server) Serve() {
	addr := fmt.Sprintf(":%s", s.port)
	s.logger.Infof("Servidor iniciado na porta 0.0.0.0:%s", s.port)

	router := mux.NewRouter()

	for key, route := range s.routes {
		subRouter := router.Name(key).Subrouter()
		subRouter.Use(route.Middlewares...)

		path := fmt.Sprintf("%s%s", s.prefix, route.Path)

		subRouter.HandleFunc(path, route.Controller).Methods(route.Methods...)

		s.logger.Infof("New Route Available [%s]", path)
	}

	http.ListenAndServe(addr, router)
}
