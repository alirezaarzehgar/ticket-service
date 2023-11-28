package route

import (
	echojwt "github.com/labstack/echo-jwt/v4"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"

	"github.com/alirezaarzehgar/ticketservice/config"
	"github.com/alirezaarzehgar/ticketservice/logd"
)

func todo(c echo.Context) error { return nil }

func Init() *echo.Echo {
	e := echo.New()

	if config.Debug() {
		middleware.DefaultLoggerConfig.Output = logd.DefaultWriter
		e.Use(middleware.Logger())
	}

	e.POST("/register", todo)
	e.POST("/login", todo)

	g := e.Group("", echojwt.WithConfig(echojwt.Config{SigningKey: config.JwtSecret()}))
	g.GET("/secret", todo)

	return e
}
