package handler

import (
	"fmt"
	"io"
	"log/slog"
	"net/http"
	"net/smtp"
	"os"
	"strconv"

	"github.com/labstack/echo/v4"
	"gorm.io/gorm"

	"github.com/alirezaarzehgar/ticketservice/model"
	"github.com/alirezaarzehgar/ticketservice/util"
)

var DefaultAssetDir = "./assets"

// UploadAsset godoc
//
//	@Summary		Upload asset for tickets
//	@Description	Users can upload asset and attach it's url to their tickets
//	@Tags			ticket
//	@Accept			json
//	@Produce		json
//	@Param			is_image	query		string	false	"flag for check images"
//	@Param			asset		formData	string	true	"asset"
//	@Success		200			{object}	util.Response
//	@Failure		400			{object}	util.ResponseError
//	@Failure		500			{object}	util.ResponseError
//
//	@Router			/ticket/assets [POST]
func UploadAsset(c echo.Context) error {
	var isImage bool
	if c.QueryParam("is_image") != "" {
		var err error
		isImage, err = strconv.ParseBool(c.QueryParam("is_image"))
		if err != nil {
			return c.JSON(http.StatusBadRequest, util.Response{Alert: util.ALERT_BAD_REQUEST})
		}
	}
	slog.Debug("start saving an asset")

	file, err := c.FormFile("asset")
	if err != nil {
		slog.Debug("asset form field is empty")
		return c.JSON(http.StatusBadRequest, util.Response{Alert: util.ALERT_BAD_REQUEST})
	}
	slog.Debug("recieve file", "data", file.Filename)

	src, err := file.Open()
	if err != nil {
		slog.Debug("cannot open file", "data", err)
		return c.JSON(http.StatusInternalServerError, util.Response{Alert: util.ALERT_INTERNAL})
	}
	defer src.Close()
	slog.Debug("open sent file", "data", file.Filename)

	if !util.IsValidPath(file.Filename, isImage) {
		return c.JSON(http.StatusBadRequest, util.Response{Alert: util.ALERT_BAD_REQUEST})
	}

	dirpath := util.GetUserDir(util.GetUserId(c))
	assetsDir := DefaultAssetDir + "/" + dirpath
	if _, err := os.Stat(assetsDir); err != nil {
		if err := os.Mkdir(assetsDir, os.ModePerm); err != nil {
			slog.Debug("mkdir dirpath", "err", err)
			return c.JSON(http.StatusInternalServerError, util.Response{Alert: util.ALERT_INTERNAL})
		}
		slog.Debug("create dir", "data", assetsDir)
	}

	filepath := fmt.Sprintf("%s/%s", dirpath, util.GetUniqueName(file.Filename))
	dst, err := os.Create(DefaultAssetDir + "/" + filepath)
	if err != nil {
		slog.Debug("create filepath", "path", filepath, "err", err)
		return c.JSON(http.StatusInternalServerError, util.Response{Alert: util.ALERT_INTERNAL})
	}
	defer dst.Close()
	slog.Debug("create asset on path", "path", filepath)

	if _, err = io.Copy(dst, src); err != nil {
		return err
	}
	slog.Debug("copy transfered file to assets directory")

	return c.JSON(http.StatusOK, util.Response{
		Status: true,
		Alert:  util.ALERT_SUCCESS,
		Data:   map[string]any{"path": filepath},
	})
}

// SendTicket godoc
//
//	@Summary		Create and send ticket to organization
//	@Description	Just admins can reply it though email.
//	@Tags			ticket
//	@Accept			json
//	@Produce		json
//	@Param			org_id			path		int		true	"Organization ID"
//	@Param			title			body		string	true	"Title"
//	@Param			body			body		string	true	"Body"
//	@Param			attachment_url	body		string	false	"Attachment Url"
//	@Param			website_url		body		string	false	"website Url"
//	@Success		200				{object}	util.Response
//	@Failure		400				{object}	util.ResponseError
//
//	@Router			/ticket/:org_id [POST]
func SendTicket(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("org_id"))
	if err != nil || id <= 0 {
		slog.Debug("invalid id parameter", "data", err)
		return c.JSON(http.StatusBadRequest, util.Response{Alert: util.ALERT_BAD_REQUEST})
	}

	var ticket model.Ticket
	if err := util.ParseBody(c, &ticket, []string{"title", "body"}, nil); err != nil {
		return nil
	}

	ticket.UserID = util.GetUserId(c)
	ticket.OrganizationID = uint(id)

	slog.Debug("create ticket", "data", ticket)
	err = db.Create(&ticket).Error
	if err == gorm.ErrForeignKeyViolated {
		slog.Debug("user not found", "data", err)
		return c.JSON(http.StatusNotFound, util.Response{Alert: util.ALERT_NOT_FOUND})
	} else if err != nil {
		slog.Debug("db error on create ticket", "data", err)
		return c.JSON(http.StatusInternalServerError, util.Response{Alert: util.ALERT_INTERNAL})
	}

	return c.JSON(http.StatusOK, util.Response{Status: true, Alert: util.ALERT_SUCCESS, Data: ticket})
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
//	@Success		200		{object}	util.Response
//	@Failure		400		{object}	util.ResponseError
//
//	@Router			/ticket/{org_id} [GET]
func GetAllTickets(c echo.Context) error {
	var tickets []model.Ticket

	userId := util.GetUserId(c)
	orgId, err := strconv.Atoi(c.Param("org_id"))
	if err != nil || orgId <= 0 {
		slog.Debug("invalid id parameter", "data", err)
		return c.JSON(http.StatusBadRequest, util.Response{Alert: util.ALERT_BAD_REQUEST})
	}

	role := util.GetUserRole(c)
	slog.Debug("requested role", "role", role)
	switch role {
	case model.USERS_ROLE_USER:
		db.Find(&tickets, "organization_id = ? AND user_id = ?", orgId, userId)
	case model.USERS_ROLE_ADMIN:
		var permitted int64
		oa := &model.OrgAdmin{OrganizationID: uint(orgId), UserID: userId}
		db.Where(oa).First(&model.OrgAdmin{}).Count(&permitted)

		if permitted < 1 {
			slog.Debug("user is not admin of reguested org", "user_id", userId, "org_id", orgId)
			return c.JSON(http.StatusUnauthorized, util.Response{Status: false, Alert: util.ALERT_ORG_EDIT_UNAUTHORIZED})
		}

		db.Find(&tickets, "organization_id = ?", orgId)
	case model.USERS_ROLE_SUPER_ADMIN:
		db.Find(&tickets, "organization_id = ?", orgId)
	}

	return c.JSON(http.StatusOK, util.Response{Status: true, Alert: util.ALERT_SUCCESS, Data: tickets})
}

// ReplyToTicket godoc
//
//	@Summary		Admin can reply to a ticket
//	@Description	Admins can reply to his organization tickets.
//	@Description	Super Admins can reply to every ticket.
//	@Tags			organization
//	@Accept			json
//	@Produce		json
//	@Param			id		path		int	true	"Ticket ID"
//	@Param			subject	body		int	true	"Email subject"
//	@Param			body	body		int	true	"Email body"
//	@Success		200		{object}	util.Response
//	@Failure		400		{object}	util.ResponseError
//
//	@Router			/ticket/{id}/mail [POST]
func ReplyToTicket(c echo.Context) error {
	ticketId, err := strconv.Atoi(c.Param("id"))
	if err != nil || ticketId <= 0 {
		slog.Debug("invalid id parameter", "data", err)
		return c.JSON(http.StatusBadRequest, util.Response{Alert: util.ALERT_BAD_REQUEST})
	}

	var sc util.SmtpContent
	if err := util.ParseBody(c, &sc, []string{"subject", "body"}, nil); err != nil {
		return nil
	}

	var ticket model.Ticket
	r := db.Preload("User").First(&ticket, ticketId)
	if r.Error == gorm.ErrRecordNotFound || r.RowsAffected == 0 {
		slog.Debug("ticket not found with recieved id", "data", r.Error)
		return c.JSON(http.StatusNotFound, util.Response{Alert: util.ALERT_NOT_FOUND})
	} else if err != nil {
		slog.Debug("database error", "data", err)
		return c.JSON(http.StatusInternalServerError, util.Response{Alert: util.ALERT_INTERNAL})
	}
	slog.Debug("fetched ticket", "ticket", ticket)

	ticket.Seen = true
	if err = db.Save(&ticket).Error; err != nil {
		slog.Debug("database error", "data", err)
		return c.JSON(http.StatusInternalServerError, util.Response{Alert: util.ALERT_INTERNAL})
	}
	slog.Debug("change ticket status to seen", "ticket", ticket)

	to := ticket.User.Email

	go func() {
		conf := util.SmtpConf
		msg := fmt.Sprintf("From: %s\r\n"+
			"To: %s\r\n"+
			"Subject: %s\r\n"+sc.Body,
			conf.FromAddress, to, sc.Subject,
		)
		err = smtp.SendMail(conf.Server, util.SmtpAuth, conf.FromAddress, []string{to}, []byte(msg))
		if err != nil {
			slog.Info("the system could not send your email", "email", msg, "err", err)
		}
	}()

	return c.JSON(http.StatusOK, util.Response{Status: true, Alert: util.ALERT_SUCCESS})
}
