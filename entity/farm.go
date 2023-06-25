package entity

import (
	"database/sql"
	"time"
)

type Farm struct {
	ID          string       `json:"id" gorm:"primaryKey;type:varchar(36)"`
	Name        string       `json:"name" gorm:"type:varchar(100)"`
	Description string       `json:"description" gorm:"type:varchar(150)"`
	CreatedAt   time.Time    `json:"created_at"`
	UpdatedAt   time.Time    `json:"updated_at"`
	DeletedAt   sql.NullTime `json:"-"`
	IsDeleted   sql.NullBool `json:"-"`
}

type UpdateFarmRequest struct {
	Name        string `json:"name"`
	Description string `json:"description"`
}

type CreateFarmRequest struct {
	ID          string `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
}

type HTTPFarmResp struct {
	Meta Meta         `json:"meta"`
	Data HTTPFarmData `json:"data"`
}

type HTTPFarmData struct {
	Farm Farm `json:"farm"`
}

type HTTPFarmsResp struct {
	Meta       Meta          `json:"meta"`
	Data       HTTPFarmsData `json:"data"`
	Pagination *Pagination   `json:"pagination,omitempty"`
}

type HTTPFarmsData struct {
	Farms []Farm `json:"farms"`
}

type Meta struct {
	Path       string `json:"path"`
	StatusCode int    `json:"status_code"`
	Status     string `json:"status"`
	Message    string `json:"message"`
	Timestamp  string `json:"timestamp"`
	Error      string `json:"error,omitempty"`
}

type HTTPEmptyResp struct {
	Meta Meta `json:"metadata"`
}
