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
	AddRoute(key string, route Route)
	Serve()
}

type server struct {
	routes Routes
	router *mux.Router
	port   string
	logger *log.Logger
	prefix string
	Server
}

// Create a new server.
func NewServer(params ...mux.Router) *server {
	if len(params) == 1 {
		return &server{
			routes: make(Routes),
			port:   "3005",
			prefix: "",
			router: &params[0],
			logger: log.Default(),
		}
	}

	return &server{
		routes: make(Routes),
		port:   "3005",
		prefix: "",
		router: mux.NewRouter(),
		logger: log.Default(),
	}
}

// Set server routes.
func (s *server) SetRoutes(routes Routes) {
	s.routes = routes
}

// Set server routes.
func (s *server) GetRoutes() Routes{
	return s.routes
}

// Add a route in server.
func (s *server) AddRoute(key string, route Route) {
	s.routes[key] = route
}

// Set route prefix.
func (s *server) SetPrefix(prefix string) {
	s.prefix = fmt.Sprintf("%s", prefix)
}

// Get route prefix.
func (s *server) GetPrefix() string {
	return s.prefix
}

// Set server logger.
func (s *server) SetLogger(logger *log.Logger) {
	s.logger = logger
}

// Set server port.
func (s *server) SetPort(port string) {
	s.port = port
}

// Get server port.
func (s *server) GetPort(port string) string {
	return s.port
}

func (s *server) Serve() error {
	addr := fmt.Sprintf(":%s", s.port)
	s.logger.Infof("Servidor iniciado na porta 0.0.0.0:%s", s.port)

	for key, route := range s.routes {
		subRouter := s.router.Name(key).Subrouter()
		subRouter.Use(route.Middlewares...)

		path := fmt.Sprintf("%s%s", s.prefix, route.Path)

		subRouter.HandleFunc(path, route.Controller).Methods(route.Methods...)

		s.logger.Infof("New Route Available [%s]", path)
	}

	return http.ListenAndServe(addr, s.router)
}
