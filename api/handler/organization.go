package handler

import (
	"fmt"
	"log/slog"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
	"gorm.io/gorm"

	"github.com/alirezaarzehgar/ticketservice/model"
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
//	@Failure		409				{object}	util.ResponseError"
//	@Failure		500				{object}	util.ResponseError"
//
//	@Router			/organization/new [POST]
func CreateOrganization(c echo.Context) error {
	var org model.Organization
	if err := util.ParseBody(c, &org, []string{"name", "address", "phone_number"}, nil); err != nil {
		return nil
	}

	org.Admins = append(org.Admins, model.User{Model: gorm.Model{ID: util.GetUserId(c)}})
	r := db.Create(&org)
	if r.Error == gorm.ErrDuplicatedKey {
		slog.Debug("conflict on database", "data", r.Error)
		return c.JSON(http.StatusConflict, util.Response{Alert: util.ALERT_ORG_CONFLICT})
	} else if r.Error != nil {
		slog.Debug("db error on create org", "data", r.Error)
		return c.JSON(http.StatusInternalServerError, util.Response{Alert: util.ALERT_INTERNAL})
	}
	slog.Debug("org created", "data", org)

	db.Preload("Admins", func(db *gorm.DB) *gorm.DB {
		return db.Omit("password")
	}).First(&org)
	return c.JSON(http.StatusOK, util.Response{Status: false, Alert: util.ALERT_SUCCESS, Data: org})
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
	var orgs []model.Organization
	db.Preload("Admins", func(db *gorm.DB) *gorm.DB {
		return db.Omit("password")
	}).Find(&orgs)
	return c.JSON(http.StatusOK, util.Response{Status: true, Alert: util.ALERT_SUCCESS, Data: orgs})
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
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil || id <= 0 {
		slog.Debug("invalid id parameter", "data", err)
		return c.JSON(http.StatusBadRequest, util.Response{Alert: util.ALERT_BAD_REQUEST})
	}

	var permitted int64
	oa := model.OrgAdmin{OrganizationID: uint(id), UserID: util.GetUserId(c)}
	slog.Debug("search options for OrgAdmin", "data", oa)

	db.Where(oa).First(&oa).Count(&permitted)
	if permitted <= 0 {
		slog.Debug("admin is not member of organization")
		return c.JSON(http.StatusUnauthorized, util.Response{Status: false, Alert: util.ALERT_ORG_EDIT_UNAUTHORIZED})
	}
	slog.Debug("admin is member of organization", "permitted", permitted, "data", oa)

	var org model.Organization
	if err := util.ParseBody(c, &org, nil, nil); err != nil {
		return nil
	}

	org.ID = uint(id)
	slog.Debug("organization to update", "data", org)
	r := db.Updates(&org)
	if r.Error == gorm.ErrRecordNotFound || r.RowsAffected == 0 {
		slog.Debug("organization not found with recieved id", "data", r.Error)
		return c.JSON(http.StatusNotFound, util.Response{Alert: util.ALERT_NOT_FOUND})
	} else if err != nil {
		slog.Debug("database error", "data", err)
		return c.JSON(http.StatusInternalServerError, util.Response{Alert: util.ALERT_INTERNAL})
	}

	db.Preload("Admins", func(db *gorm.DB) *gorm.DB {
		return db.Omit("password")
	}).First(&org, id)
	return c.JSON(http.StatusOK, util.Response{Status: true, Alert: util.ALERT_SUCCESS, Data: org})
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
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil || id <= 0 {
		slog.Debug("invalid id parameter", "data", err)
		return c.JSON(http.StatusBadRequest, util.Response{Alert: util.ALERT_BAD_REQUEST})
	}

	var org model.Organization
	r := db.Delete(&org, id)
	if r.Error == gorm.ErrRecordNotFound || r.RowsAffected == 0 {
		slog.Debug("organization not found with recieved id", "data", r.Error)
		return c.JSON(http.StatusNotFound, util.Response{Alert: util.ALERT_NOT_FOUND})
	} else if err != nil {
		slog.Debug("database error", "data", err)
		return c.JSON(http.StatusInternalServerError, util.Response{Alert: util.ALERT_INTERNAL})
	}

	db.Unscoped().First(&org, id)
	slog.Debug("organization to edit", "data", org)
	org.Name = fmt.Sprintf("%s-%s", util.CreateSHA256(fmt.Sprint(id)), org.Name)
	db.Unscoped().Save(&org)

	return c.JSON(http.StatusOK, util.Response{Status: true, Alert: util.ALERT_SUCCESS, Data: org})
}
