package handler_test

import (
	"log"
	"log/slog"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/alirezaarzehgar/ticketservice/api/handler"
	"github.com/alirezaarzehgar/ticketservice/config"
	"github.com/alirezaarzehgar/ticketservice/database"
	"github.com/alirezaarzehgar/ticketservice/logd"
	"github.com/alirezaarzehgar/ticketservice/model"
	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

var (
	db *gorm.DB
	e  = echo.New()
)

func TestMain(m *testing.M) {
	godotenv.Load("../../.env")
	logd.DefaultLogDir = "../../log"
	logd.InitLogger()

	dbConf, _ := config.GetDb()
	db, _ = database.Init(dbConf, log.New(logd.DefaultWriter, "", logd.DefaultLogFlags))
	database.Migrate(db, config.Admin())
	handler.SetDB(db)
	m.Run()

	db.Unscoped().Delete(&model.User{}, "username", MOCK_USER["username"])
	db.Unscoped().Delete(&model.User{}, "username", MOCK_ADMIN["username"])

	var org model.Organization
	db.First(&org, "name", MOCK_ORG["name"])
	slog.Debug("remove mockorg", "data", org)
	db.Model(&org).Association("Admins").Clear()
	db.Unscoped().Delete(&model.Organization{}, "name", MOCK_ORG["name"])
}

func nilBodyTest(t *testing.T, handler func(c echo.Context) error, method string, target string) {
	req := httptest.NewRequest(method, target, nil)
	rec := httptest.NewRecorder()
	if handler(e.NewContext(req, rec)); rec.Code != http.StatusBadRequest {
		t.Errorf("body is nil but works. code: %v, user: %v", rec.Code, rec.Body)
	}
}
