package handler

import (
	"log/slog"
	"net/http"

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

// EditAdmin godoc
//
//	@Summary		Edit admin
//	@Description	Super admin can edit every normal admin.
//	@Description	Admin can edit hisself. But cannot change his role.
//	@Tags			admin
//	@Accept			json
//	@Produce		json
//	@Param			id			path		int		true	"User ID"
//	@Param			username	body		string	false	"Username"
//	@Param			password	body		string	false	"Password"
//	@Param			email		body		string	false	"Email"
//	@Param			role		body		string	false	"User role (super_admin|admin|user)"
//	@Success		200			{object}	util.Response
//	@Failure		400			{object}	util.ResponseError
//
//	@Router			/admin/{id} [PUT]
func EditAdmin(c echo.Context) error {
	return c.JSON(http.StatusOK, map[any]string{})
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
	return c.JSON(http.StatusOK, map[any]string{})
}
