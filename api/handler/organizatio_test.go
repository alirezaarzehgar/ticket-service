package handler_test

import (
	"bytes"
	"encoding/json"
	"log/slog"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/alirezaarzehgar/ticketservice/api/handler"
)

var (
	MOCK_ORG = map[string]any{
		"name":         "mockorg",
		"address":      "mockaddr",
		"phone_number": "09xxxxxxx",
	}
)

func TestCreateOrganization(t *testing.T) {
	body, _ := json.Marshal(MOCK_ORG)
	req := httptest.NewRequest(http.MethodPost, "/organization/new", bytes.NewReader(body))
	req.Header.Set("Authorization", "Bearer "+ADMIN_TOKEN)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)

	slog.Debug("test create org", "data", string(body))
	if handler.CreateOrganization(c); rec.Code != http.StatusOK {
		t.Errorf("error on creating org")
	}

	nilBodyTest(t, handler.CreateOrganization, http.MethodPost, "/organization/new")
}

func TestGetAllOrganizations(t *testing.T) {

}

func TestEditOrganization(t *testing.T) {

}

func TestAssignAdminToOrganization(t *testing.T) {

}

func TestDeleteOrganization(t *testing.T) {

}
