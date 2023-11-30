package handler_test

import (
	"bytes"
	"encoding/json"
	"log/slog"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/alirezaarzehgar/ticketservice/api/handler"
	"github.com/alirezaarzehgar/ticketservice/model"
	"github.com/alirezaarzehgar/ticketservice/util"
)

var (
	MOCK_ADMIN = map[string]any{
		"username": "mockadmin",
		"password": "pass",
		"email":    "mockadmin@example.com",
	}

	ADMIN_TOKEN = util.CreateUserToken(1, "admin@example.com", "admin", model.USERS_ROLE_SUPER_ADMIN)
)

func TestCreateAdmin(t *testing.T) {
	var res util.Response
	body, _ := json.Marshal(MOCK_ADMIN)
	req := httptest.NewRequest(http.MethodPost, "/admin/new", bytes.NewReader(body))
	rec := httptest.NewRecorder()
	if handler.CreateAdmin(e.NewContext(req, rec)); rec.Code != http.StatusOK {
		t.Errorf("error on creating admin")
	}
	json.NewDecoder(rec.Body).Decode(&res)
	slog.Debug("test case body", "data", res)

	if res.Data == nil {
		t.Errorf("Empty res.Data")
		return
	}
	u := res.Data.(map[string]any)

	if u["role"] != model.USERS_ROLE_ADMIN {
		t.Errorf("Created admin haven't admin role: %v", u)
		return
	}

	MOCK_ADMIN = u
}

func TestEditAdmin(t *testing.T) {

}

func TestPromoteAdmin(t *testing.T) {

}
