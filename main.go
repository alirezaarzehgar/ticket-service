package main

import (
	"log"
	"log/slog"
	"os"

	"github.com/joho/godotenv"

	"github.com/alirezaarzehgar/ticketservice/api/handler"
	"github.com/alirezaarzehgar/ticketservice/api/middleware"
	"github.com/alirezaarzehgar/ticketservice/api/route"
	"github.com/alirezaarzehgar/ticketservice/config"
	"github.com/alirezaarzehgar/ticketservice/database"
	"github.com/alirezaarzehgar/ticketservice/logd"
)

func main() {
	go logd.HandleInterrupt()
	if err := godotenv.Load(".env"); err != nil {
		log.Fatal("faild to load .env: ", err)
	}
	logd.InitLogger()

	dbConf, err := config.GetDb()
	if err != nil {
		slog.Error(".env: ", err)
		os.Exit(1)
	}

	db, err := database.Init(dbConf)
	if err != nil {
		slog.Error("database: ", err)
		os.Exit(1)
	}

	if err := database.Migrate(db); err != nil {
		slog.Error("migrate: ", err)
		os.Exit(1)
	}

	middleware.SetDB(db)
	handler.SetDB(db)

	slog.Info("Start application")
	if err := route.Init().Start(config.ListenerAddr()); err != nil {
		slog.Error("echo start:", err)
		os.Exit(1)
	}
}
