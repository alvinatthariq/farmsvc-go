package entity

import (
	"database/sql"
	"strings"
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

func (p Pond) Validate() error {
	p.ID = strings.TrimSpace(p.ID)
	if len(p.ID) < 1 {
		return ErrorPondIDRequired
	} else if len(p.ID) > 36 {
		return ErrorPondIDMaxLength
	}

	p.FarmID = strings.TrimSpace(p.FarmID)
	if len(p.FarmID) < 1 {
		return ErrorFarmIDRequired
	} else if len(p.FarmID) > 36 {
		return ErrorFarmIDMaxLength
	}

	p.Name = strings.TrimSpace(p.Name)
	if len(p.Name) < 1 {
		return ErrorPondNameRequired
	} else if len(p.Name) > 100 {
		return ErrorPondNameMaxLength
	}

	p.Description = strings.TrimSpace(p.Description)
	if len(p.Description) < 1 {
		return ErrorPondDescriptionRequired
	} else if len(p.Description) > 150 {
		return ErrorPondDescriptionMaxLength
	}

	return nil
}

type CreatePondRequest struct {
	ID          string `json:"id"`
	FarmID      string `json:"farm_id"`
	Name        string `json:"name"`
	Description string `json:"description"`
}

type UpdatePondRequest struct {
	FarmID      string `json:"farm_id"`
	Name        string `json:"name"`
	Description string `json:"description"`
}
