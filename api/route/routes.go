package route

import (
	"io"
	"log/slog"

	echojwt "github.com/labstack/echo-jwt/v4"
	echoSwagger "github.com/swaggo/echo-swagger"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"

	_ "github.com/alirezaarzehgar/ticketservice/docs"
	"github.com/alirezaarzehgar/ticketservice/logd"
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
		slog.Debug("enable logger and swagger cause to debug mode")
		middleware.DefaultLoggerConfig.Output = c.LogWriter
		e.Use(middleware.Logger())

		logd.RegisterHandlers(e.Group(""))

		e.GET("/swagger/*", echoSwagger.WrapHandler)
	}

	e.POST("/register", todo)
	e.POST("/login", todo)

	g := e.Group("", echojwt.WithConfig(echojwt.Config{SigningKey: c.JwtSecret}))
	g.GET("/secret", todo)

	return e
}
