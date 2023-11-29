package util_test

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/alirezaarzehgar/ticketservice/model"
	"github.com/alirezaarzehgar/ticketservice/util"
	"github.com/labstack/echo/v4"
)

var (
	mockSha256Value = "123"
	mockSha256Hash  = "a665a45920422f9d417e4867efdc4fb8a04a1f3fff1fa07e998e86f7f7a27ae3"
	// token payload: id: 1, email: "admin@example.com", user: "admin", role: super admin
	mockTokenID uint = 1
	mockToken        = "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJpc3MiOiJhZG1pbkBleGFtcGxlLmNvbSIsInN1YiI6ImFkbWluIiwiYXVkIjpbImFkbWluIl0sImV4cCI6MTcwMzg4NTY3OSwianRpIjoiMSJ9.LoYd4E66YP8j5mQbWYYi9FBNRwGq0Zj3NNXi1Nohvb0"
)

func TestCreateSHA256(t *testing.T) {
	if real := util.CreateSHA256(mockSha256Value); real != mockSha256Hash {
		t.Errorf("%s != %s", real, mockSha256Value)
	}
}

func TestCreateUserToken(t *testing.T) {
	expectedHeaderPayload := "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJpc3MiOiJhZG1pbkBleGFtcGxlLmNvbSIsInN1YiI6ImFkbWluIiwiYXVkIjpbImFkbWluIl0sImV4cCI6MTcwMzg4NTY3OSwianRpIjoiMSJ9."
	if real := util.CreateUserToken(1, "user@example.com", "user", model.USERS_ROLE_USER); strings.HasPrefix(real, expectedHeaderPayload) {
		t.Errorf("generated token haven't right header and payload: %s", real)
	}
}
func TestGetToken(t *testing.T) {
	e := echo.New()

	req := httptest.NewRequest(http.MethodPost, "/path", nil)
	req.Header.Set("Authorization", "Bearer "+mockToken)
	rec := httptest.NewRecorder()

	if real := util.GetToken(e.NewContext(req, rec)); real != mockToken {
		t.Errorf("{%s} != {%s}", real, mockToken)
	}
}
func TestGetUserId(t *testing.T) {
	e := echo.New()
	req := httptest.NewRequest(http.MethodPost, "/path", nil)
	req.Header.Set("Authorization", "Bearer "+mockToken)
	rec := httptest.NewRecorder()

	if id := util.GetUserId(e.NewContext(req, rec)); id != mockTokenID {
		t.Errorf("Wrong user id!")
	}
}

func TestParseBody(t *testing.T) {
	type CModel struct {
		Req, Ign, Opt int
	}
	body := CModel{1, 2, 3}
	jsonBody, _ := json.Marshal(body)

	e := echo.New()
	req := httptest.NewRequest(http.MethodPost, "/path", bytes.NewReader(jsonBody))
	req.Header.Set("Authorization", "Bearer "+mockToken)
	rec := httptest.NewRecorder()

	var d2p CModel
	if err := util.ParseBody(e.NewContext(req, rec), &d2p, []string{"Req"}, []string{"Ign"}); err != nil {
		t.Errorf("error: %v", err)
	}
	if !(d2p.Req == 1 && d2p.Ign == 0 && d2p.Opt == 3) {
		t.Errorf("%v != %v", d2p, CModel{1, 0, 3})
	}

	if err := util.ParseBody(e.NewContext(req, rec), &d2p, []string{"bad field"}, nil); err == nil {
		t.Errorf("function should return error when a required field does not exists")
	}
}

func TestGetUserRole(t *testing.T) {
	e := echo.New()
	req := httptest.NewRequest(http.MethodPost, "/path", nil)
	req.Header.Set("Authorization", "Bearer "+mockToken)
	rec := httptest.NewRecorder()
	real := util.GetUserRole(e.NewContext(req, rec))

	if real != model.USERS_ROLE_ADMIN {
		t.Errorf("%v is not super admin", real)
	}
}
