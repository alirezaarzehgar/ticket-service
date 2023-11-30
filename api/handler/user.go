package handler

import (
	"fmt"
	"log/slog"
	"net/http"
	"strconv"

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
		return nil
	}
	slog.Debug("recieved body", "data", user)

	user.Password = util.CreateSHA256(user.Password)
	r := db.Create(&user)
	if r.Error == gorm.ErrDuplicatedKey {
		slog.Debug("conflict on database", "data", r.Error)
		return c.JSON(http.StatusConflict, util.Response{Alert: util.ALERT_USER_CONFLICT})
	} else if r.Error != nil {
		slog.Debug("db error on create user", "data", r.Error)
		return c.JSON(http.StatusInternalServerError, util.Response{Alert: util.ALERT_INTERNAL})
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
		return nil
	}

	user.Password = util.CreateSHA256(user.Password)
	db.Where(user).First(&user).Count(&loggedin)
	if loggedin == 0 {
		slog.Debug("user not found", "data", user)
		return c.JSON(http.StatusUnauthorized, util.Response{Alert: util.ALERT_LOGIN_UNAUTHORIZED})
	}

	token := util.CreateUserToken(user.ID, user.Email, user.Username, user.Role)
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
//	@Success		200	{object}	util.Response
//	@Failure		400	{object}	util.ResponseError
//
//	@Router			/user/profile [GET]
func GetUserProfile(c echo.Context) error {
	var user model.User
	if err := db.First(&user, util.GetUserId(c)).Error; err != nil {
		slog.Debug("invalid id", "data", c.Param("id"))
		return c.JSON(http.StatusInternalServerError, util.Response{Alert: util.ALERT_INTERNAL})
	}

	user.Password = ""
	return c.JSON(http.StatusOK, util.Response{Status: true, Alert: util.ALERT_SUCCESS, Data: user})
}

// DeleteUser godoc
//
//	@Summary		Delete a user or admin
//	@Description	Super admin can delete users and admins.
//	@Description	Super admin can'n remove super admins/
//	@Description	Admins can't remove another admins or users.
//	@Tags			user
//	@Accept			json
//	@Produce		json
//	@Param			id	path		int	true	"User ID"
//	@Success		200	{object}	util.Response
//	@Failure		400	{object}	util.ResponseError
//
//	@Router			/user/{id} [DELETE]
func DeleteUser(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil || id <= 0 {
		slog.Debug("invalid id parameter", "data", err)
		return c.JSON(http.StatusBadRequest, util.Response{Alert: util.ALERT_BAD_REQUEST})
	}

	var u model.User
	r := db.Delete(&u, id)
	if r.Error == gorm.ErrRecordNotFound || r.RowsAffected == 0 {
		slog.Debug("user not found with recieved id", "data", r.Error)
		return c.JSON(http.StatusNotFound, util.Response{Alert: util.ALERT_NOT_FOUND})
	} else if err != nil {
		slog.Debug("database error", "data", err)
		return c.JSON(http.StatusInternalServerError, util.Response{Alert: util.ALERT_INTERNAL})
	}

	db.Unscoped().First(&u, id)
	slog.Debug("user to edit", "data", u)
	uniquePrefix := util.CreateSHA256(fmt.Sprint(id))
	u.Username = fmt.Sprintf("%s-%s", uniquePrefix, u.Username)
	u.Email = fmt.Sprintf("%s-%s", uniquePrefix, u.Username)
	db.Unscoped().Save(&u)

	return c.JSON(http.StatusOK, util.Response{Status: true, Alert: util.ALERT_SUCCESS, Data: u})
}
