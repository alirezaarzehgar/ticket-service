package handler

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

// CreateAdmin godoc
//
//	@Summary		Create new admin
//	@Description	Just super admins can create a new admin
//	@Tags			admin
//	@Accept			json
//	@Produce		json
//	@Success		200	{object}	Response
//	@Failure		400	{object}	Response
//
//	@Router			/admin/new [POST]
func CreateAdmin(c echo.Context) error {
	return c.JSON(http.StatusOK, map[any]string{})
}

// DeleteAdmin godoc
//
//	@Summary		Delete an admin
//	@Description	Super admin can delete a normal admin
//	@Tags			admin
//	@Accept			json
//	@Produce		json
//	@Param			id	path		int	true	"User ID"
//	@Success		200	{object}	Response
//	@Failure		400	{object}	ResponseError
//
//	@Router			/admin/{id} [DELETE]
func DeleteAdmin(c echo.Context) error {
	return c.JSON(http.StatusOK, map[any]string{})
}

// EditAdmin godoc
//
//	@Summary		Edit admin
//	@Description	Super admin can edit every normal admin.
//	@Description	Admin can edit hisself.
//	@Tags			admin
//	@Accept			json
//	@Produce		json
//	@Param			id	path		int	true	"User ID"
//	@Success		200	{object}	Response
//	@Failure		400	{object}	ResponseError
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
//	@Success		200	{object}	Response
//	@Failure		400	{object}	ResponseError
//
//	@Router			/admin/promote/{id} [POST]
func PromoteAdmin(c echo.Context) error {
	return c.JSON(http.StatusOK, map[any]string{})
}
