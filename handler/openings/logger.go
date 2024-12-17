package openings

import (
	"database/sql"

	"github.com/victorlui/gopportunities/config"
)

var (
	logger *config.Logger
	db     *sql.DB
)
