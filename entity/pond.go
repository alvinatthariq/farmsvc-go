package entity

import (
	"database/sql"
	"time"
)

type Pond struct {
	ID          string       `json:"id" gorm:"primaryKey;type:varchar(36)"`
	FarmID      string       `json:"farm_id" gorm:"type:varchar(36)"`
	Name        string       `json:"name" gorm:"type:varchar(100)"`
	Description string       `json:"description" gorm:"type:varchar(150)"`
	CreatedAt   time.Time    `json:"created_at"`
	UpdatedAt   time.Time    `json:"updated_at"`
	DeletedAt   sql.NullTime `json:"-"`
	IsDeleted   sql.NullBool `json:"-"`
}
