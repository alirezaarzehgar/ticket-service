package logd

import (
	"log/slog"
	"os"
	"os/signal"
	"syscall"
)

func HandleInterrupt() {
	ch := make(chan os.Signal, 1)
	signal.Notify(ch, os.Interrupt, syscall.SIGINT, syscall.SIGTERM)
	<-ch

	slog.Info("Stop application")
	stopLogger()
	os.Exit(1)
}
