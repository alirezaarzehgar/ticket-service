package middleware

import (
	"net/http"

	"github.com/labstack/echo/v4"

	"github.com/alirezaarzehgar/ticketservice/model"
	"github.com/alirezaarzehgar/ticketservice/util"
)

func UserOnly(next echo.HandlerFunc) echo.HandlerFunc {
	return echo.HandlerFunc(func(c echo.Context) error {
		if util.GetUserRole(c) == model.USERS_ROLE_USER {
			return next(c)
		}
		return c.JSON(http.StatusUnauthorized, util.Response{Status: false, Alert: util.ALERT_USER_ONLY})
	})
}
