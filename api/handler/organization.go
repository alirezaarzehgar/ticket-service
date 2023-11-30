package handler

import (
	"net/http"

	"github.com/labstack/echo/v4"

	"github.com/alirezaarzehgar/ticketservice/util"
)

// CreateOrganization godoc
//
//	@Summary		Create Organization
//	@Description	Super admins can create organization
//	@Tags			organization
//	@Accept			json
//	@Produce		json
//
//	@Param			name			body		string	false	"Name"
//	@Param			address			body		string	false	"Address"
//	@Param			phone_number	body		string	false	"Phone number"
//	@Param			website_url		body		string	false	"Website URL"
//
//	@Success		200				{object}	util.Response
//	@Failure		400				{object}	util.ResponseError
//
//	@Router			/organization/new [POST]
func CreateOrganization(c echo.Context) error {
	return c.JSON(http.StatusOK, util.Response{Status: false, Alert: util.ALERT_SUCCESS, Data: map[string]any{}})
}

// GetAllOrganizations godoc
//
//	@Summary		List all organizations
//	@Description	Response every registred organization.
//	@Tags			organization
//	@Accept			json
//	@Produce		json
//	@Success		200	{object}	util.Response
//	@Failure		400	{object}	util.ResponseError
//
//	@Router			/organization/all [GET]
func GetAllOrganizations(c echo.Context) error {
	return c.JSON(http.StatusOK, map[string]any{})
}

// EditOrganization godoc
//
//	@Summary		Edit organization
//	@Description	Admins can edit organizations information.
//	@Tags			organization
//	@Accept			json
//	@Produce		json
//	@Param			name			body		string	false	"Name"
//	@Param			address			body		string	false	"Address"
//	@Param			phone_number	body		string	false	"Phone number"
//	@Param			website_url		body		string	false	"Website URL"
//	@Success		200				{object}	util.Response
//	@Failure		400				{object}	util.ResponseError
//
//	@Router			/organization/{id} [PUT]
func EditOrganization(c echo.Context) error {
	return c.JSON(http.StatusOK, map[string]any{})
}

// AssignAdminToOrganization godoc
//
//	@Summary		Add admin to organization
//	@Description	Assign organization to an admin. Just super admins can do it.
//	@Tags			organization
//	@Accept			json
//	@Produce		json
//	@Param			org_id	path		string	true	"Organization Path"
//	@Param			user_id	path		string	true	"User ID"
//	@Success		200		{object}	util.Response
//	@Failure		400		{object}	util.ResponseError
//
//	@Router			/organization/hire-admin/{org_id}/{user_id} [POST]
func AssignAdminToOrganization(c echo.Context) error {
	return c.JSON(http.StatusOK, map[string]any{})
}

// DeleteOrganization godoc
//
//	@Summary		Delete Organization
//	@Description	Super admin can delete an organization.
//	@Tags			organization
//	@Accept			json
//	@Produce		json
//	@Param			id	path		string	true	"Organization ID"
//	@Success		200	{object}	util.Response
//	@Failure		400	{object}	util.ResponseError
//
//	@Router			/organization/{id} [DELETE]
func DeleteOrganization(c echo.Context) error {
	return c.JSON(http.StatusOK, map[string]any{})
}
