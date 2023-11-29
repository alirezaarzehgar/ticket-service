package handler

import (
	"net/http"

	"github.com/labstack/echo/v4"

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
//
//	@Router			/admin/new [POST]
func CreateAdmin(c echo.Context) error {

	return c.JSON(http.StatusOK, util.Response{Status: false, Alert: util.ALERT_SUCCESS, Data: map[any]string{}})
}

// DeleteAdmin godoc
//
//	@Summary		Delete an admin
//	@Description	Super admin can delete a normal admin
//	@Tags			admin
//	@Accept			json
//	@Produce		json
//	@Param			id	path		int	true	"User ID"
//	@Success		200	{object}	util.Response
//	@Failure		400	{object}	util.ResponseError
//
//	@Router			/admin/{id} [DELETE]
func DeleteAdmin(c echo.Context) error {
	return c.JSON(http.StatusOK, map[any]string{})
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
