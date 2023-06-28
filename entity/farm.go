package entity

import (
	"database/sql"
	"strings"
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

type FarmParam struct {
	ID    string
	Name  string
	Limit int `gorm:"-"`
	Page  int `gorm:"-"`
}

func (f Farm) Validate() error {
	f.ID = strings.TrimSpace(f.ID)
	if len(f.ID) < 1 {
		return ErrorFarmIDRequired
	} else if len(f.ID) > 36 {
		return ErrorFarmIDMaxLength
	}

	f.Name = strings.TrimSpace(f.Name)
	if len(f.Name) < 1 {
		return ErrorFarmNameRequired
	} else if len(f.Name) > 100 {
		return ErrorFarmNameMaxLength
	}

	f.Description = strings.TrimSpace(f.Description)
	if len(f.Description) < 1 {
		return ErrorFarmDescriptionRequired
	} else if len(f.Description) > 150 {
		return ErrorFarmDescriptionMaxLength
	}

	return nil
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
