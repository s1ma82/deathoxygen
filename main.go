package main

import (
	"embed"
	"fmt"
	"net/http"

	"github.com/labstack/echo/v4"
	api "github.com/s1ma82/deathoxygen/pkg/api"
	"github.com/s1ma82/deathoxygen/pkg/chat"
)

//go:embed pkg/openapi/index.html
//go:embed pkg/openapi/openapi.yaml
var swaggerUI embed.FS

func main() {
	e := echo.New()

	// chat handlers
	api.RegisterHandlers(e, chat.NewServer())

	e.GET("/docs/*", echo.WrapHandler(http.StripPrefix("/docs/", http.FileServer(http.FS(swaggerUI)))))
	e.Start("localhost:8080")
	fmt.Println("Жизнь это игра в режиме хардкор, другого шанса уже не будет)")

}
