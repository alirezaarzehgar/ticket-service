package middleware

import "github.com/labstack/echo/v4"

func ForSuperAdmin(next echo.HandlerFunc) echo.HandlerFunc {
	return echo.HandlerFunc(func(c echo.Context) error {
		return next(c)
	})
}

func ForAdmin(next echo.HandlerFunc) echo.HandlerFunc {
	return echo.HandlerFunc(func(c echo.Context) error {
		return next(c)
	})
}
