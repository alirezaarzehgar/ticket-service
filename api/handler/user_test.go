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
	MOCK_USER = map[string]any{
		"username": "mockuser",
		"password": "pass",
		"email":    "mockuser@example.com",
	}
	// token payload: id: 1, email: "user@example.com", user: "user"
	mockToken        = "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJpc3MiOiJ1c2VyQGV4YW1wbGUuY29tIiwic3ViIjoidXNlciIsImV4cCI6MTcwMzg2Mjk1NywianRpIjoiMSJ9.Jx_mEygZjnkTNif2VEgWsFxAn7soV8oKYih51ZZ7I-w"
	mockTokenID uint = 1
)

func TestRegister(t *testing.T) {
	body, _ := json.Marshal(MOCK_USER)
	req := httptest.NewRequest(http.MethodPost, "/register", bytes.NewReader(body))
	rec := httptest.NewRecorder()

	if err := handler.Register(e.NewContext(req, rec)); err != nil {
		t.Errorf("error: %v", err)
	}
	if rec.Code != http.StatusOK {
		t.Errorf("status code: %d != %d", rec.Code, http.StatusOK)
	}

	nilBodyTest(t, handler.Register, http.MethodPost, "/register")
}

func TestLogin(t *testing.T) {
	body, _ := json.Marshal(MOCK_USER)
	req := httptest.NewRequest(http.MethodPost, "/login", bytes.NewReader(body))
	rec := httptest.NewRecorder()

	if err := handler.Login(e.NewContext(req, rec)); err != nil {
		t.Errorf("error: %v", err)
	}
	if rec.Code != http.StatusOK {
		t.Errorf("status code: %d != %d", rec.Code, http.StatusOK)
	}

	res := struct {
		Data struct {
			Token string `json:"token"`
		}
	}{}
	json.NewDecoder(rec.Body).Decode(&res)
	if len(res.Data.Token) < 10 {
		t.Errorf("invalid token: %v", res.Data.Token)
	}

	nilBodyTest(t, handler.Login, http.MethodPost, "/login")
}

func TestGetUserProfile(t *testing.T) {
	req := httptest.NewRequest(http.MethodGet, "/user/profile", nil)
	req.Header.Set("Authorization", "Bearer "+mockToken)
	rec := httptest.NewRecorder()

	if err := handler.GetUserProfile(e.NewContext(req, rec)); err != nil {
		t.Errorf("error: %v", err)
	}
	if rec.Code != http.StatusOK {
		t.Errorf("status code: %d != %d", rec.Code, http.StatusOK)
	}

	res := struct {
		ID uint `json:"id"`
	}{}
	json.NewDecoder(rec.Body).Decode(&res)
	if res.ID == mockTokenID {
		t.Errorf("invalid id: %v", res.ID)
	}
}

func TestGetUser(t *testing.T) {
	req := httptest.NewRequest(http.MethodGet, "/user/1", nil)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	c.SetParamNames("id")
	c.SetParamValues("1")

	if err := handler.GetUser(c); err != nil {
		t.Errorf("error: %v", err)
	}
	if rec.Code != http.StatusOK {
		t.Errorf("status code: %d != %d", rec.Code, http.StatusOK)
	}
}

func TestDeleteUser(t *testing.T) {
	var u model.User
	db.Select("id").Last(&u)

	req := httptest.NewRequest(http.MethodDelete, "/user/1", nil)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	c.SetParamNames("id")
	c.SetParamValues(fmt.Sprint(u.ID))

	if err := handler.DeleteUser(c); err != nil {
		t.Errorf("error: %v", err)
	}

	body, _ := json.Marshal(MOCK_USER)
	req = httptest.NewRequest(http.MethodPost, "/register", bytes.NewReader(body))
	rec = httptest.NewRecorder()

	if err := handler.Register(e.NewContext(req, rec)); err != nil {
		t.Errorf("error: %v", err)
	}
	if rec.Code != http.StatusOK {
		t.Errorf("status code: %d != %d", rec.Code, http.StatusOK)
	}
}

func TestEditUser(t *testing.T) {
	var user model.User
	db.Select("id").Last(&user)

	newName := "updated user"
	MOCK_USER["username"] = newName
	var res util.Response
	body, _ := json.Marshal(MOCK_USER)
	req := httptest.NewRequest(http.MethodPut, "/user/1", bytes.NewReader(body))
	rec := httptest.NewRecorder()
	req.Header.Set("Authorization", "Bearer "+ADMIN_TOKEN)

	c := e.NewContext(req, rec)
	c.SetParamNames("id")
	c.SetParamValues(fmt.Sprint(user.ID))

	if handler.EditUser(c); rec.Code != http.StatusOK {
		t.Errorf("error on udpate user: %v", rec.Code)
	}
	json.NewDecoder(rec.Body).Decode(&res)
	slog.Debug("test case body", "data", res)

	if res.Data == nil {
		t.Errorf("Empty res.Data")
		return
	}
	u := res.Data.(map[string]any)

	if u["username"] != newName {
		t.Errorf("Created admin haven't admin role: %v", u)
		return
	}

	MOCK_USER = u
}
