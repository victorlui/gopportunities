package schemas

import (
	"database/sql"
	"time"
)

type OpeningResponse struct {
	ID        uint         `json:"id"`
	CreatedAt time.Time    `json:"createdAt"`
	UpdatedAt time.Time    `json:"updatedAt"`
	DeletedAt sql.NullTime `json:"deletedAt"`
	Role      string       `json:"role"`
	Company   string       `json:"company"`
	Location  string       `json:"location"`
	Remote    bool         `json:"remote"`
	Link      string       `json:"link"`
	Salary    int64        `json:"salary"`
}

type OpeningResponseSwagger struct {
	ID        uint      `json:"id"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
	Role      string    `json:"role"`
	Company   string    `json:"company"`
	Location  string    `json:"location"`
	Remote    bool      `json:"remote"`
	Link      string    `json:"link"`
	Salary    int64     `json:"salary"`
}
