package database

import (
	"database/sql"

	_ "github.com/mattn/go-sqlite3"
	"github.com/viniciusbls9/go-movie/pkg/utils"
)

func OpenDatabaseConnection() (*sql.DB, error) {
	dbURL := utils.GetEnv()

	db, err := sql.Open("sqlite3", dbURL)
	if err != nil {
		return nil, err
	}

	return db, nil
}
