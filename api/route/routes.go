package route

import (
	"io"

	echojwt "github.com/labstack/echo-jwt/v4"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func todo(c echo.Context) error { return nil }

type RouteConfig struct {
	LogWriter io.Writer
	DebugMode bool
	JwtSecret []byte
}

func Init(c RouteConfig) *echo.Echo {
	e := echo.New()

	if c.DebugMode {
		middleware.DefaultLoggerConfig.Output = c.LogWriter
		e.Use(middleware.Logger())
	}

	e.POST("/register", todo)
	e.POST("/login", todo)

	g := e.Group("", echojwt.WithConfig(echojwt.Config{SigningKey: c.JwtSecret}))
	g.GET("/secret", todo)

	return e
}
