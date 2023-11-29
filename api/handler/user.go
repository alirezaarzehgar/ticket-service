package handler

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

// Register godoc
//
//	@Summary		Register user
//	@Description	Create user on database based on body.
//	@Tags			user
//	@Accept			json
//	@Produce		json
//	@Param			username	body		string	true	"Username"
//	@Param			password	body		string	true	"Password"
//	@Param			email		body		string	true	"Email"
//	@Success		200			{object}	Response
//	@Failure		400			{object}	ResponseError"
//
//	@Router			/register [POST]
func Register(c echo.Context) error {
	return c.JSON(http.StatusOK, Response{false, ALERT_SUCCESS, nil})
}

// Login godoc
//
//	@Summary		Login user
//	@Description	Check user&pass and response JWT Token
//	@Tags			user
//	@Accept			json
//	@Produce		json
//	@Param			username	body		string	true	"Username"
//	@Param			password	body		string	true	"Password"
//	@Param			email		body		string	true	"Email"
//	@Success		200			{object}	Response
//	@Failure		400			{object}	ResponseError
//
//	@Router			/login [POST]
func Login(c echo.Context) error {
	return c.JSON(http.StatusOK, Response{false, ALERT_SUCCESS, "token"})
}

// GetUserProfile godoc
//
//	@Summary		Get user profile
//	@Description	Fetch a user by ID
//	@Tags			user
//	@Accept			json
//	@Produce		json
//	@Param			id	path		int	true	"User ID"
//	@Success		200	{object}	Response
//	@Failure		400	{object}	ResponseError
//
//	@Router			/user/profile/{id} [GET]
func GetUserProfile(c echo.Context) error {
	return c.JSON(http.StatusOK, Response{false, ALERT_SUCCESS, nil})
}
