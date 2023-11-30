package handler_test

import (
	"bytes"
	"encoding/json"
	"fmt"
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
		t.Fatalf("Empty res.Data")
	}
	u := res.Data.(map[string]any)

	if u["role"] != model.USERS_ROLE_ADMIN {
		t.Fatalf("Created admin haven't admin role: %v", u)
	}

	MOCK_ADMIN["id"] = u["ID"]
	MOCK_ADMIN["role"] = u["role"]
}

func TestPromoteAdmin(t *testing.T) {
	body, _ := json.Marshal(MOCK_USER)
	req := httptest.NewRequest(http.MethodPut, "/user/1", bytes.NewReader(body))
	rec := httptest.NewRecorder()
	req.Header.Set("Authorization", "Bearer "+ADMIN_TOKEN)

	adminId := fmt.Sprint(MOCK_ADMIN["id"])
	c := e.NewContext(req, rec)
	c.SetParamNames("id")
	c.SetParamValues(adminId)

	if handler.PromoteAdmin(c); rec.Code != http.StatusOK {
		t.Errorf("error on udpate user: %v", rec.Code)
	}

	var u model.User
	db.First(&u, adminId)

	if u.Role != model.USERS_ROLE_SUPER_ADMIN {
		t.Fatalf("Admin doesn't promote: %v", u)
	}

	MOCK_ADMIN["role"] = model.USERS_ROLE_SUPER_ADMIN
}
