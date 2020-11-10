package main

import (
	"noraclock/v2/src/database"
	"noraclock/v2/src/logger"
	"noraclock/v2/src/server"

	_ "github.com/jackc/pgx/v4/stdlib"
)

func main() {
	// The Get method returns the logger so that
	// we can call the Sync method on it.
	// https://godoc.org/go.uber.org/zap#Logger.Sync
	log, err := logger.Get()
	if err != nil {
		panic(err)
	}
	defer func() { _ = log.Sync() }()

	db, err := database.ConnectPostgreSQL()
	if err != nil {
		panic(err)
	}
	defer func() { _ = db.Close }()

	err = server.Start()
	if err != nil {
		panic(err)
	}
}
