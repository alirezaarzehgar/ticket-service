package route

import (
	"io"
	"log/slog"

	"github.com/labstack/echo/v4"

	echojwt "github.com/labstack/echo-jwt/v4"
	echoMiddleware "github.com/labstack/echo/v4/middleware"
	echoSwagger "github.com/swaggo/echo-swagger"

	"github.com/alirezaarzehgar/ticketservice/api/handler"
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

	e.POST("/register", handler.Register)
	e.POST("/login", handler.Login)

	g := e.Group("", echojwt.WithConfig(echojwt.Config{SigningKey: c.JwtSecret}))
	g.GET("/user/profile", handler.GetUserProfile, middleware.UserOnly)

	g.POST("/admin/new", handler.CreateAdmin, middleware.ForSuperAdmin)
	g.DELETE("/admin/:id", handler.DeleteAdmin, middleware.ForSuperAdmin)
	g.PUT("/admin/:id", handler.EditAdmin, middleware.ForSuperAdmin)
	g.POST("/admin/promote/:id", handler.PromoteAdmin, middleware.ForSuperAdmin)

	g.POST("/organization/new", handler.CreateOrganization, middleware.ForSuperAdmin)
	g.GET("/organization/all", handler.GetAllOrganizations)
	g.PUT("/organization/:id", handler.EditOrganization, middleware.ForAdmin)
	g.POST("/organization/hire-admin/:org_id/:user_id", handler.AssignAdminToOrganization, middleware.ForSuperAdmin)
	g.DELETE("/organization/:id", handler.DeleteOrganization, middleware.ForSuperAdmin)

	g.POST("/ticket/new", handler.SendTicket, middleware.UserOnly)
	g.GET("/ticket/:org_id", handler.GetAllTickets)
	g.POST("/ticket/:id/mail", handler.ReplyToTicket, middleware.ForSuperAdmin)

	return e
}
