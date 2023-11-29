package handler

import (
	"net/http"

	"github.com/labstack/echo/v4"
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
//	@Success		200				{object}	Response
//	@Failure		400				{object}	ResponseError
//
//	@Router			/organization/new [POST]
func CreateOrganization(c echo.Context) error {
	return c.JSON(http.StatusOK, map[any]string{})
}

// GetAllOrganizations godoc
//
//	@Summary		List all organizations
//	@Description	Response every registred organization.
//	@Tags			organization
//	@Accept			json
//	@Produce		json
//	@Success		200	{object}	Response
//	@Failure		400	{object}	ResponseError
//
//	@Router			/organization/all [GET]
func GetAllOrganizations(c echo.Context) error {
	return c.JSON(http.StatusOK, map[any]string{})
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
//	@Success		200				{object}	Response
//	@Failure		400				{object}	ResponseError
//
//	@Router			/organization/{id} [PUT]
func EditOrganization(c echo.Context) error {
	return c.JSON(http.StatusOK, map[any]string{})
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
//	@Success		200		{object}	Response
//	@Failure		400		{object}	ResponseError
//
//	@Router			/organization/hire-admin/{org_id}/{user_id} [POST]
func AssignAdminToOrganization(c echo.Context) error {
	return c.JSON(http.StatusOK, map[any]string{})
}

// DeleteOrganization godoc
//
//	@Summary		Delete Organization
//	@Description	Super admin can delete an organization.
//	@Tags			organization
//	@Accept			json
//	@Produce		json
//	@Param			id	path		string	true	"Organization ID"
//	@Success		200	{object}	Response
//	@Failure		400	{object}	ResponseError
//
//	@Router			/organization/{id} [DELETE]
func DeleteOrganization(c echo.Context) error {
	return c.JSON(http.StatusOK, map[any]string{})
}
