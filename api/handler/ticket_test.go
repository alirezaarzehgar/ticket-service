package handler_test

import (
	"net/http"
	"testing"

	"github.com/alirezaarzehgar/ticketservice/api/handler"
)

func TestUploadAsset(t *testing.T) {
	nilBodyTest(t, handler.UploadAsset, http.MethodPost, "/ticket/asset")
}

func TestSendTicket(t *testing.T) {
	nilBodyTest(t, handler.CreateOrganization, http.MethodPost, "/organization/new")
}

func TestGetAllTickets(t *testing.T) {
	nilBodyTest(t, handler.CreateOrganization, http.MethodPost, "/organization/new")
}

func TestReplyToTicket(t *testing.T) {
	nilBodyTest(t, handler.ReplyToTicket, http.MethodPost, "/ticket/1/mail")
}
