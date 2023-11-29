package main

import (
	"log"
	"log/slog"
	"os"

	"github.com/joho/godotenv"

	"github.com/alirezaarzehgar/ticketservice/api/handler"
	"github.com/alirezaarzehgar/ticketservice/api/route"
	"github.com/alirezaarzehgar/ticketservice/config"
	"github.com/alirezaarzehgar/ticketservice/database"
	"github.com/alirezaarzehgar/ticketservice/logd"
)

//	@version	2.0
//	@title		ticket service

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
	slog.Debug("get database config", "data", dbConf)

	db, err := database.Init(dbConf, log.New(logd.DefaultWriter, "", logd.DefaultLogFlags))
	if err != nil {
		slog.Error("database: ", err)
		os.Exit(1)
	}

	if err := database.Migrate(db, config.Admin()); err != nil {
		slog.Error("migrate: ", err)
		os.Exit(1)
	}

	handler.SetDB(db)

	rc := route.RouteConfig{
		LogWriter: logd.DefaultWriter,
		DebugMode: config.Debug(),
		JwtSecret: config.JwtSecret(),
	}
	slog.Debug("route config", "data", rc)

	slog.Info("Start application")
	e := route.Init(rc)
	if err := e.Start(config.ListenerAddr()); err != nil {
		slog.Error("echo start:", err)
		os.Exit(1)
	}
}
