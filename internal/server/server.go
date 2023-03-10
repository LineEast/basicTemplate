package server

import (
	"kmfRedirect/internal/database"

	"github.com/fasthttp/router"
	"github.com/valyala/fasthttp"
)

type (
	Server struct {
		router   *router.Router
		database *database.Database

		configuration *Configuration
	}

	Configuration struct {
		Host string `toml:"host"`
		Port string `toml:"port"`
	}
)

func New(database *database.Database, configuration *Configuration) *Server {
	server := &Server{
		database:      database,
		configuration: configuration,
	}

	r := router.New()

	api := r.Group("/api")
	api.POST("/createUser", server.CreateUser())
	api.GET("/getAllUserList", server.GetAllUserList())
	api.GET("/getUser", server.GetUser())

	server.router = r

	return server
}

func (s *Server) Run() error {
	d := fasthttp.Server{
		Handler:           s.router.Handler,
		StreamRequestBody: true,
	}

	return d.ListenAndServe(s.configuration.Host + ":" + s.configuration.Port)
}
