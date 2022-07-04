package api

import "github.com/labstack/echo/v4"

type Router struct {
	Routes    []*echo.Route
	APIV1Auth *echo.Group
}

type Server struct {
	Echo   *echo.Echo
	Router *Router
}

func NewServer() *Server {
	s := &Server{
		Echo:   nil,
		Router: nil,
	}
	return s
}

func (s *Server) Start() error {
	ListenAddress := "http://localhost:8080"
	return s.Echo.Start(ListenAddress)
}
