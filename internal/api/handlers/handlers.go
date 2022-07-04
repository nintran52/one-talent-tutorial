package handlers

import (
	"github.com/labstack/echo/v4"
	"github.com/nintran52/one-talent-tutorial/internal/api"
)

func AttachAllRouters(s *api.Server) {
	s.Router.Routes = []*echo.Route{}
}
