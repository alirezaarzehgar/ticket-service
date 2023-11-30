package middleware

import (
	"net/http"

	"github.com/alirezaarzehgar/ticketservice/model"
	"github.com/alirezaarzehgar/ticketservice/util"
	"github.com/labstack/echo/v4"
)

func ForSuperAdmin(next echo.HandlerFunc) echo.HandlerFunc {
	return echo.HandlerFunc(func(c echo.Context) error {
		if util.GetUserRole(c) == model.USERS_ROLE_SUPER_ADMIN {
			return next(c)
		}
		return c.JSON(http.StatusUnauthorized, util.Response{Status: false, Alert: util.ALERT_SUPER_ADMIN_ONLY})
	})
}

func ForAdmin(next echo.HandlerFunc) echo.HandlerFunc {
	return echo.HandlerFunc(func(c echo.Context) error {
		role := util.GetUserRole(c)
		if role == model.USERS_ROLE_ADMIN || role == model.USERS_ROLE_SUPER_ADMIN {
			return next(c)
		}
		return c.JSON(http.StatusUnauthorized, util.Response{Status: false, Alert: util.ALERT_ADMIN_ONLY})
	})
}
