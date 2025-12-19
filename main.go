package main

import (
	"fmt"

	"github.com/labstack/echo/v4"
	api "github.com/s1ma82/deathoxygen/pkg/model"
)

func main() {
	e := echo.New()
	srv := api.NewServer()
	api.RegisterHandlers(e, srv)
	e.Start("localhost:8080")
	fmt.Println("Жизнь это игра в режиме хардкор, другого шанса уже не будет)")
}
