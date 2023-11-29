package util_test

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/alirezaarzehgar/ticketservice/util"
	"github.com/labstack/echo/v4"
)

var (
	mockSha256Value = "123"
	mockSha256Hash  = "a665a45920422f9d417e4867efdc4fb8a04a1f3fff1fa07e998e86f7f7a27ae3"
	// token payload: id: 1, email: "user@example.com", user: "user"
	mockTokenID uint = 1
	mockToken        = "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJpc3MiOiJ1c2VyQGV4YW1wbGUuY29tIiwic3ViIjoidXNlciIsImV4cCI6MTcwMzg2Mjk1NywianRpIjoiMSJ9.Jx_mEygZjnkTNif2VEgWsFxAn7soV8oKYih51ZZ7I-w"
)

func TestCreateSHA256(t *testing.T) {
	if real := util.CreateSHA256(mockSha256Value); real != mockSha256Hash {
		t.Errorf("%s != %s", real, mockSha256Value)
	}
}

func TestCreateUserToken(t *testing.T) {
	expectedHeaderPayload := "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJpc3MiOiJ1c2VyQGV4YW1wbGUuY29tIiwic3ViIjoidXNlciIsImV4cCI6MTcwMzg2Mjk1OSwianRpIjoiMSJ9."
	if real := util.CreateUserToken(1, "user@example.com", "user"); strings.HasPrefix(real, expectedHeaderPayload) {
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
