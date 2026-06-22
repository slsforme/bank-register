package postgres

import (
	"bankapp/config"

	"github.com/jmoiron/sqlx"
)

func getConnection(driverName string, dsn string) *sqlx.DB {
	db, err := sqlx.Connect(driverName, dsn)

	if err != nil {
		config.Logger.Error("error occurred while connecting to database:", "error", err)
		panic("")
	}

	return db
}
