package handler_test

import (
	"log"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/alirezaarzehgar/ticketservice/api/handler"
	"github.com/alirezaarzehgar/ticketservice/api/middleware"
	"github.com/alirezaarzehgar/ticketservice/config"
	"github.com/alirezaarzehgar/ticketservice/database"
	"github.com/alirezaarzehgar/ticketservice/logd"
	"github.com/alirezaarzehgar/ticketservice/model"
	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"
)

func TestMain(m *testing.M) {
	godotenv.Load("../../.env")
	logd.DefaultLogDir = "../../log"
	logd.InitLogger()

	dbConf, _ := config.GetDb()
	db, _ := database.Init(dbConf, log.New(logd.DefaultWriter, "", logd.DefaultLogFlags))
	database.Migrate(db, config.Admin())
	middleware.SetDB(db)
	handler.SetDB(db)
	m.Run()

	db.Unscoped().Delete(&model.User{}, "username", MOCK_USER["username"])
}

func nilBodyTest(t *testing.T, handler func(c echo.Context) error) {
	req := httptest.NewRequest(http.MethodPost, "/register", nil)
	rec := httptest.NewRecorder()
	if err := handler(e.NewContext(req, rec)); err == nil || rec.Code != http.StatusBadRequest {
		t.Errorf("body is nil but works. code: %v, err: %v, user: %v", rec.Code, err, rec.Body)
	}
}
