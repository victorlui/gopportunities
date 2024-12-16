package config

import (
	"os"

	"database/sql"

	_ "modernc.org/sqlite"
)

func InitializeSQLite() (*sql.DB, error) {
	logger := GetLogger("sqlite")
	dbPath := "./db/opportunities.db"

	//check if the database file exists
	_, err := os.Stat(dbPath)
	if os.IsNotExist(err) {
		logger.Info("database file not found, creating...")
		err = os.MkdirAll("./db", os.ModePerm)
		if err != nil {
			return nil, err
		}

		file, err := os.Create(dbPath)

		if err != nil {
			return nil, err
		}

		// fechar o arquivo
		file.Close()
	}

	// Create db sqlite and connect
	db, err := sql.Open("sqlite", "file:./db/opportunities.db")
	logger.InfoF("db aberto")

	if err != nil {
		logger.ErrF("sqlite opening error: %v", err)
		return nil, err
	}

	createTableQuery := `
    CREATE TABLE IF NOT EXISTS openings (
        id INTEGER PRIMARY KEY AUTOINCREMENT,
        role TEXT NOT NULL,
        company TEXT NOT NULL,
        location TEXT NOT NULL,
        remote BOOLEAN NOT NULL,
        link TEXT,
        salary INTEGER,
		created_at DATETIME,
        updated_at DATETIME,
        deleted_at DATETIME
    );
    `

	_, err = db.Exec(createTableQuery)

	if err != nil {
		logger.ErrF("sqlite create table error: %v", err)
		return nil, err
	}

	return db, nil
}
