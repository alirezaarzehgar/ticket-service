package logd

import (
	"fmt"
	"log"
	"log/slog"
	"os"
	"time"

	"github.com/jasonlvhit/gocron"
)

var (
	DefaultWriter     *os.File
	DefaultLogDir     = "./log"
	DefaultDateFormat = "2006-01-02"
	DefaultLogFlags   = log.Ltime | log.Lshortfile
	DefaultLogLevel   = slog.LevelDebug
)

func changeLogWriter() {
	var err error
	DefaultWriter.Close()
	logpath := fmt.Sprintf("%s/%s.log", DefaultLogDir, time.Now().Format(DefaultDateFormat))
	DefaultWriter, err = os.OpenFile(logpath, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		log.Println("Unable to open logfile:", err)
	}
	log.SetOutput(DefaultWriter)
}

func InitLogger() {
	changeLogWriter()
	log.SetFlags(DefaultLogFlags)
	gocron.Every(1).Day().Do(changeLogWriter)
	gocron.Start()

	slog.SetDefault(slog.New(slog.NewJSONHandler(DefaultWriter, &slog.HandlerOptions{Level: DefaultLogLevel})))
}

func stopLogger() {
	DefaultWriter.Close()
}
