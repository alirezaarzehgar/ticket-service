package handler

import (
	"log/slog"
	"net/http"

	"github.com/alirezaarzehgar/ticketservice/model"
	"github.com/labstack/echo/v4"

	"github.com/alirezaarzehgar/ticketservice/util"
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
//	@Success		200			{object}	util.Response
//	@Failure		400			{object}	util.ResponseError"
//
//	@Router			/register [POST]
func Register(c echo.Context) error {
	var user model.User
	if err := util.ParseBody(c, &user, []string{"username", "password", "email"}, []string{"role"}); err != nil {
		slog.Debug("parse body failed", "data", err, "body", c.Request().Body)
	}
	slog.Debug("recieved body", "data", user)

	return c.JSON(http.StatusOK, util.Response{Status: false, Alert: util.ALERT_SUCCESS, Data: map[string]any{}})
}

// Login godoc
//
//	@Summary		Login user
//	@Description	Check user&pass and util.Response JWT Token
//	@Tags			user
//	@Accept			json
//	@Produce		json
//	@Param			username	body		string	true	"Username"
//	@Param			password	body		string	true	"Password"
//	@Param			email		body		string	true	"Email"
//	@Success		200			{object}	util.Response
//	@Failure		400			{object}	util.ResponseError
//
//	@Router			/login [POST]
func Login(c echo.Context) error {
	return c.JSON(http.StatusOK, map[any]string{})
}

// GetUserProfile godoc
//
//	@Summary		Get user profile
//	@Description	Fetch a user by ID
//	@Tags			user
//	@Accept			json
//	@Produce		json
//	@Param			id	path		int	true	"User ID"
//	@Success		200	{object}	util.Response
//	@Failure		400	{object}	util.ResponseError
//
//	@Router			/user/profile/{id} [GET]
func GetUserProfile(c echo.Context) error {
	return c.JSON(http.StatusOK, map[any]string{})
}
