package config

import (
	"database/sql"
	"fmt"
)

var (
	db     *sql.DB
	logger *Logger
)

func Init() error {
	var err error

	db, err = InitializeSQLite()
	if err != nil {
		return fmt.Errorf("error initializing sqlite: %v", err)
	}
	return nil
}

func GetSQLite() *sql.DB {
	return db
}

func GetLogger(p string) *Logger {
	//inicializer logger
	logger = NewLogger(p)
	return logger
}
