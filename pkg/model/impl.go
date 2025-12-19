package api

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

type Server struct{}

func NewServer() *Server {
	return &Server{}
}

func (s *Server) SendPing(ctx echo.Context) error {
	r := PingPong{
		Message: "Pong!",
		Status:  true,
	}

	return ctx.JSON(http.StatusOK, r)
}

