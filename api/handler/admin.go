package handler

import (
	"log/slog"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
	"gorm.io/gorm"

	"github.com/alirezaarzehgar/ticketservice/model"
	"github.com/alirezaarzehgar/ticketservice/util"
)

// CreateAdmin godoc
//
//	@Summary		Create new admin
//	@Description	Just super admins can create a new admin
//	@Tags			admin
//	@Accept			json
//	@Produce		json
//	@Param			username	body		string	true	"Username"
//	@Param			password	body		string	true	"Password"
//	@Param			email		body		string	true	"Email"
//	@Success		200			{object}	util.Response
//	@Failure		400			{object}	util.Response
//	@Failure		409			{object}	util.ResponseError"
//	@Failure		500			{object}	util.ResponseError"
//
//	@Router			/admin/new [POST]
func CreateAdmin(c echo.Context) error {
	var admin model.User
	if err := util.ParseBody(c, &admin, []string{"username", "password", "email"}, nil); err != nil {
		return nil
	}
	slog.Debug("recieved body", "data", admin)

	admin.Password = util.CreateSHA256(admin.Password)
	if admin.Role == "" || admin.Role == model.USERS_ROLE_USER {
		admin.Role = model.USERS_ROLE_ADMIN
	}

	r := db.Create(&admin)
	if r.Error == gorm.ErrDuplicatedKey {
		slog.Debug("conflict on database", "data", r.Error)
		return c.JSON(http.StatusConflict, util.Response{Alert: util.ALERT_USER_CONFLICT})
	} else if r.Error != nil {
		slog.Debug("db error on create admin", "data", r.Error)
		return c.JSON(http.StatusInternalServerError, util.Response{Alert: util.ALERT_INTERNAL})
	}
	slog.Debug("user created", "data", admin)

	return c.JSON(http.StatusOK, util.Response{Status: true, Alert: util.ALERT_SUCCESS, Data: admin})
}

// PromoteAdmin godoc
//
//	@Summary		Promote admin to super user
//	@Description	Super admin can convert an admin to super admin.
//	@Tags			admin
//	@Accept			json
//	@Produce		json
//	@Param			id	path		int	true	"User ID"
//	@Success		200	{object}	util.Response
//	@Failure		400	{object}	util.ResponseError
//
//	@Router			/admin/promote/{id} [POST]
func PromoteAdmin(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil || id <= 0 {
		slog.Debug("invalid id parameter", "data", err)
		return c.JSON(http.StatusBadRequest, util.Response{Alert: util.ALERT_BAD_REQUEST})
	}

	r := db.Model(&model.User{}).
		Where("role = ?", model.USERS_ROLE_ADMIN).
		Where(id).Update("role", model.USERS_ROLE_SUPER_ADMIN)
	if r.Error == gorm.ErrRecordNotFound || r.RowsAffected == 0 {
		slog.Debug("user not found with recieved id", "data", r.Error)
		return c.JSON(http.StatusNotFound, util.Response{Alert: util.ALERT_NOT_FOUND})
	} else if err != nil {
		slog.Debug("database error", "data", err)
		return c.JSON(http.StatusInternalServerError, util.Response{Alert: util.ALERT_INTERNAL})
	}

	return c.JSON(http.StatusOK, util.Response{Status: true, Alert: util.ALERT_SUCCESS})
}
