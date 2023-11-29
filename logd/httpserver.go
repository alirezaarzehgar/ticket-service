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

// ShowLogs godoc
//
//	@Summary		List all log paths to see
//	@Description	Find logs in log directory and link to them. List them and show to user for redirecting to suitable log by date.
//	@Tags			log
//	@Produce		html
//	@Success		200	{object}	string "<a href="2023-11-28.log">2023-11-28.log</a><br><a href="2023-11-29.log">2023-11-29.log</a><br>"
//
//	@Router			/logs/list [GET]
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

// ShowCurrentLogs godoc
//
//	@Summary		Show logs of current date
//	@Description	Generate url with log path and date. Then redirect to it.
//	@Tags			log
//	@Produce		html
//	@Success		200	{object}	string "{ log }<br/>{ log }<br/>{ log }"
//
//	@Router			/logs/current [GET]
func ShowCurrentLogs(c echo.Context) error {
	url := fmt.Sprintf("/logs/%s.log", time.Now().Format("2006-01-02"))
	return c.Redirect(http.StatusTemporaryRedirect, url)
}

func RegisterHandlers(e *echo.Group) {
	e.Static("/logs/", DefaultLogDir)
	e.GET("/logs/list", ShowLogs)
	e.GET("/logs/current", ShowCurrentLogs)
}
