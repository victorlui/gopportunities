package schemas

import (
	"database/sql"
	"time"
)

type Opening struct {
	//gorm.Model
	Role      string
	Company   string
	Location  string
	Remote    bool
	Link      string
	Salary    int64
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt sql.NullTime
}
