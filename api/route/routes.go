package route

import (
	"io"
	"log/slog"

	"github.com/labstack/echo/v4"

	echojwt "github.com/labstack/echo-jwt/v4"
	echoMiddleware "github.com/labstack/echo/v4/middleware"
	echoSwagger "github.com/swaggo/echo-swagger"

	"github.com/alirezaarzehgar/ticketservice/api/middleware"
	"github.com/alirezaarzehgar/ticketservice/logd"

	_ "github.com/alirezaarzehgar/ticketservice/docs"
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
		echoMiddleware.DefaultLoggerConfig.Output = c.LogWriter
		e.Use(echoMiddleware.Logger())

		logd.RegisterHandlers(e.Group(""))

		e.GET("/swagger/*", echoSwagger.WrapHandler)
	}

	e.POST("/register", todo)
	e.POST("/login", todo)

	g := e.Group("", echojwt.WithConfig(echojwt.Config{SigningKey: c.JwtSecret}))

	g.POST("/admin/new", todo, middleware.ForSuperAdmin)
	g.DELETE("/admin/:id", todo, middleware.ForSuperAdmin)
	g.PUT("/admin/:id", todo, middleware.ForSuperAdmin)
	g.POST("/admin/promote/:id", todo, middleware.ForSuperAdmin)

	g.GET("/organize/all", todo)
	g.PUT("/organize/:id", todo, middleware.ForAdmin)
	g.POST("/organize/new", todo, middleware.ForSuperAdmin)
	g.POST("/organize/hire-admin/:org_id/:user_id", todo, middleware.ForSuperAdmin)
	g.DELETE("/organize/:id", todo, middleware.ForSuperAdmin)

	g.POST("/ticket/new", todo, middleware.UserOnly)
	g.GET("/tickets/all", todo)
	g.POST("/ticket/:id/mail", todo, middleware.ForSuperAdmin)

	return e
}
