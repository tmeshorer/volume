package server

import (
	"context"
	"github.com/gin-contrib/cors"
	"github.com/gin-contrib/logger"
	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog"
	apiv1 "github.com/tmeshorer/volume/pkg/api/v1"
	"net"
	"net/http"
	"os"
	"os/signal"
	"time"
)

const (
	Address = ":8080"
	Mode    = "release"
)

type Server struct {
	srv    *http.Server
	router *gin.Engine
}

// create a new server
func New() (s *Server, err error) {

	// Set the global level
	zerolog.SetGlobalLevel(zerolog.InfoLevel)

	// Create the server and prepare to serve
	s = &Server{}

	// Create the router
	gin.SetMode(Mode)
	s.router = gin.New()
	if err = s.setupRoutes(); err != nil {
		return nil, err
	}

	// Create the http server
	s.srv = &http.Server{
		Addr:         Address,
		Handler:      s.router,
		ErrorLog:     nil,
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 15 * time.Second,
		IdleTimeout:  30 * time.Second,
	}
	return s, nil
}

func (s *Server) Serve() error {
	// catch os for gracefull shutdown
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt)
	go func() {
		<-quit
	}()

	// bind the listener
	var sock net.Listener
	var err error
	if sock, err = net.Listen("tcp", Address); err != nil {
		return err
	}

	// Listen for HTTP requests on the specified address and port
	if err = s.srv.Serve(sock); err != nil && err != http.ErrServerClosed {
		return err
	}

	return nil

}

func (s *Server) Shutdown() error {

	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	if err := s.srv.Shutdown(ctx); err != nil {
		return err
	}

	return nil
}

func (s *Server) setupRoutes() (err error) {
	// Microservice Middleware

	middlewares := []gin.HandlerFunc{
		// logger
		logger.SetLogger(),
		// help the service recover from panic
		gin.Recovery(),

		// set cors. currently allow all domains.
		cors.New(cors.Config{
			AllowOrigins: []string{"*"},
			AllowMethods: []string{"GET"},
			AllowHeaders: []string{"Origin", "Content-Length", "Content-Type", "Authorization", "X-CSRF-TOKEN"},
			MaxAge:       12 * time.Hour,
		}),
	}

	// Add the middlewares
	for _, middleware := range middlewares {
		if middleware != nil {
			s.router.Use(middleware)
		}
	}

	// Set the v1 api.
	v1 := s.router.Group("/v1")
	{
		v1.GET("/calculate", s.Calculate)
	}

	// NotFound and NotAllowed routes
	s.router.NoRoute(apiv1.NotFound)
	s.router.NoMethod(apiv1.NotAllowed)
	return nil
}
