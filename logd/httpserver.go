package logd

import (
	"fmt"
	"io/fs"
	"net/http"
	"path/filepath"
	"strings"
	"time"

	"github.com/labstack/echo/v4"
)

func ShowLogs(c echo.Context) error {
	list := ""
	filepath.Walk(DefaultLogDir, func(path string, info fs.FileInfo, err error) error {
		if info.IsDir() || !strings.HasSuffix(path, ".log") {
			return nil
		}
		path = filepath.Base(path)
		list += fmt.Sprintf("<a href=\"%s\">%s</a><br>\n", path, path)
		return nil
	})
	return c.HTML(http.StatusOK, list)
}

func ShowCurrentLogs(c echo.Context) error {
	url := fmt.Sprintf("/logs/%s.log", time.Now().Format("2006-01-02"))
	return c.Redirect(http.StatusTemporaryRedirect, url)
}

func RegisterHandlers(e *echo.Group) {
	e.Static("/logs/", DefaultLogDir)
	e.GET("/logs/list", ShowLogs)
	e.GET("/logs/current", ShowCurrentLogs)
}
