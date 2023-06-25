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
