package chat

import (
	"net/http"

	"github.com/labstack/echo/v4"
	api "github.com/s1ma82/deathoxygen/pkg/api"
)

type Server struct{}

func NewServer() *Server {
	return &Server{}
}

func (s *Server) SendPing(ctx echo.Context) error {
	r := api.PingPong{
		Message: "Pong!",
		Status:  true,
	}

	return ctx.JSON(http.StatusOK, r)
}
