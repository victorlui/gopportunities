package openings

import (
	"fmt"
	"time"
)

func paramIsRequired(name, typ string) error {
	return fmt.Errorf("param: %s (type: %s) is required", name, typ)
}

// createOpning
type CreateOpeningRequest struct {
	Role      string    `json:"role"`
	Company   string    `json:"company"`
	Location  string    `json:"location"`
	Remote    *bool     `json:"remote"`
	Link      string    `json:"link"`
	Salary    int64     `json:"salary"`
	CreatedAt time.Time `json:"createdAt"`
}

func (r *CreateOpeningRequest) Validate() error {
	if r.Role == "" && r.Company == "" && r.Location == "" && r.Remote == nil && r.Salary <= 0 {
		return fmt.Errorf("request body is empty")
	}

	if r.Role == "" {
		return paramIsRequired("role", "string")
	}
	if r.Company == "" {
		return paramIsRequired("company", "string")
	}
	if r.Location == "" {
		return paramIsRequired("location", "string")
	}
	if r.Link == "" {
		return paramIsRequired("link", "string")
	}
	if r.Remote == nil {
		return paramIsRequired("remote", "bool")
	}
	if r.Salary <= 0 {
		return paramIsRequired("salary", "int64")
	}

	return nil
}
