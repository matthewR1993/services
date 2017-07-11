// Service description

package service

import (
	log "github.com/sirupsen/logrus"
	"github.com/julienschmidt/httprouter"
	"net/http"
	"fmt"
	. "github.com/matthewR1993/services/endpoints"
)

type Config struct {
	Host string
	Port string
	AssetsDir string	
}

type Service struct {
	Name string
	Config Config
	ServerMiddlewares []Middleware
	Endpoints []Endpoint
}

// Initializer
func Init() *Service {	
	return &Service{}
}

// Set config
func (s *Service) SetConf(conf Config) {
	s.Config = conf
}

// Add endpoints
func (s *Service) AddEndpoints(eps ...Endpoint) {
	s.Endpoints = append(s.Endpoints, eps...)
}

// Get all endpoints
func (s Service) GetEndpoints() []Endpoint {
	return s.Endpoints
}

// Add middlewares
func (s *Service) AddMiddlewares(mdlwres ...Middleware) {
	s.ServerMiddlewares = append(s.ServerMiddlewares, mdlwres...)
}

// Get all middlewares
func (s Service) GetServerMiddlewares() []Middleware {
	return s.ServerMiddlewares
}

func (s Service) GetName() string {
	return s.Name
}

// Handle all included endpoints with middleware
func (s Service) HandleAllEndpoints(router *httprouter.Router) {
	for _, endpt := range s.Endpoints {
		var handler Handler = endpt.Handler

		// Add endpoint specific middlewares
		for _, mdlwre := range endpt.EndpointMiddlewares {
			handler = mdlwre(handler)
		}
	
		// Add shared server middlewares
		for _, mdlwre := range s.ServerMiddlewares {
			handler = mdlwre(handler)
		}

		router.Handle(endpt.Method, endpt.Path, httprouter.Handle(handler))  // (endpt.Handler)
	}
}

// Start server
func (s Service) Run() {
	router := httprouter.New()
	s.HandleAllEndpoints(router)

	// Set static files directory	
	router.ServeFiles("/static/*filepath", http.Dir(s.Config.AssetsDir))

	fmt.Println(fmt.Sprintf(" Listening on %v:%s", s.Config.Host, s.Config.Port))

	// Simple http 1.1
	log.Fatal(http.ListenAndServe(fmt.Sprintf("%v:%s", s.Config.Host, s.Config.Port), router))

	// With TLS
	//log.Fatal(http.ListenAndServeTLS(fmt.Sprintf("%v:%s", s.Config.Host, s.Config.Port), "crypto/server.crt", "crypto/server.key", router))
}

// Service interface description
type Server interface {
	GetName() string
	AddEndpoints(eps ...Endpoint)
	GetEndpoints() []Endpoint
	Endpoint(name string) (Endpoint, bool)

	// Shared middlewares are triggered first
	GetServerMiddlewares() []Middleware
	AddMiddlewares(mdlwre ...Middleware)
	
	Run()
}


