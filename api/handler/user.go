package handler

import (
	"log/slog"
	"net/http"

	"github.com/alirezaarzehgar/ticketservice/model"
	"github.com/labstack/echo/v4"
	"gorm.io/gorm"

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
//	@Failure		409			{object}	util.ResponseError"
//	@Failure		500			{object}	util.ResponseError"
//
//	@Router			/register [POST]
func Register(c echo.Context) error {
	var user model.User
	if err := util.ParseBody(c, &user, []string{"username", "password", "email"}, []string{"role"}); err != nil {
		slog.Debug("parse body failed", "data", err, "body", c.Request().Body)
		return err
	}
	slog.Debug("recieved body", "data", user)

	if user.Email == "" || user.Password == "" {
		return c.JSON(http.StatusBadRequest, util.Response{Status: false, Alert: util.ALERT_BAD_REQUEST})
	}

	user.Password = util.CreateSHA256(user.Password)
	r := db.Create(&user)
	if r.Error == gorm.ErrDuplicatedKey {
		slog.Debug("conflict on database", "data", r.Error)
		return c.JSON(http.StatusConflict, util.Response{Status: false, Alert: util.ALERT_USER_CONFLICT})
	} else if r.Error != nil {
		slog.Debug("db error on create user", "data", r.Error)
		return c.JSON(http.StatusInternalServerError, util.Response{Status: false, Alert: util.ALERT_INTERNAL})
	}
	slog.Debug("user created", "data", user)

	return c.JSON(http.StatusOK, util.Response{Status: true, Alert: util.ALERT_SUCCESS, Data: user})
}

// Login godoc
//
//	@Summary		Login user
//	@Description	Check user&pass and util.Response JWT Token
//	@Tags			user
//	@Accept			json
//	@Produce		json
//	@Param			password	body		string	true	"Password"
//	@Param			email		body		string	true	"Email"
//	@Success		200			{object}	util.Response
//	@Failure		400			{object}	util.ResponseError
//	@Failure		401			{object}	util.ResponseError
//
//	@Router			/login [POST]
func Login(c echo.Context) error {
	var loggedin int64
	var user model.User
	if err := util.ParseBody(c, &user, []string{"email", "password"}, []string{"role"}); err != nil {
		slog.Debug("parse body failed", "data", err, "body", c.Request().Body)
		return err
	}

	user.Password = util.CreateSHA256(user.Password)
	db.Where(user).First(&user).Count(&loggedin)
	if loggedin == 0 {
		slog.Debug("user not found", "data", user)
		return c.JSON(http.StatusUnauthorized, util.Response{Status: false, Alert: util.ALERT_LOGIN_UNAUTHORIZED})
	}

	token := util.CreateUserToken(user.ID, user.Email, user.Username)
	slog.Debug("create token for", "data", user)
	return c.JSON(http.StatusOK, util.Response{
		Status: true,
		Alert:  util.ALERT_SUCCESS,
		Data:   map[string]string{"token": token},
	})
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
