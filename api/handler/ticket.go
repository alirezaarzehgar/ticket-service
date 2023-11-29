package handler

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

// SendTicket godoc
//
//	@Summary		Create and send ticket to organization
//	@Description	Just admins can reply it though email.
//	@Tags			ticket
//	@Accept			json
//	@Produce		json
//	@Param			user_id			body		uint	true	"User ID"
//	@Param			org_id			body		uint	true	"Organize ID"
//	@Param			title			body		string	true	"Title"
//	@Param			body			body		string	true	"Body"
//	@Param			attachment_url	body		string	true	"Attachment Url"
//	@Success		200				{object}	Response
//	@Failure		400				{object}	ResponseError
//
//	@Router			/ticket/new [POST]
func SendTicket(c echo.Context) error {
	return c.JSON(http.StatusOK, map[any]string{})
}

// GetAllTickets godoc
//
//	@Summary		Get tickets by priviledge
//	@Description	If you are a user you can see your tickets about an organization.
//	@Description	If you are an admin you can see tickets of all users for allorganizations.
//	@Tags			ticket
//	@Accept			json
//	@Produce		json
//	@Param			org_id	path		int	true	"Organization ID"
//	@Success		200		{object}	Response
//	@Failure		400		{object}	ResponseError
//
//	@Router			/ticket/{org_id} [GET]
func GetAllTickets(c echo.Context) error {
	return c.JSON(http.StatusOK, map[any]string{})
}

// ReplyToTicket godoc
//
//	@Summary		Admin can reply to a ticket
//	@Description	Admins can reply to his organization tickets.
//	@Description	Super Admins can reply to every ticket.
//	@Tags			organization
//	@Accept			json
//	@Produce		json
//	@Param			id	path		int	true	"User ID"
//	@Success		200	{object}	Response
//	@Failure		400	{object}	ResponseError
//
//	@Router			/ticket/{id}/mail [POST]
func ReplyToTicket(c echo.Context) error {
	return c.JSON(http.StatusOK, map[any]string{})
}
