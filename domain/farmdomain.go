package domain

import (
	"errors"
	"time"

	"github.com/alvinatthariq/farmsvc-go/entity"
	"gorm.io/gorm"
)

func (d *domain) CreateFarm(v entity.CreateFarmRequest) (farm entity.Farm, err error) {
	// create farm
	farm = entity.Farm{
		ID:          v.ID,
		Name:        v.Name,
		Description: v.Description,
		CreatedAt:   time.Now().UTC(),
		UpdatedAt:   time.Now().UTC(),
	}

	// create to db
	tx := d.gorm.Create(&farm)

	return farm, tx.Error
}

func (d *domain) GetFarmByID(farmID string) (farm *entity.Farm, err error) {
	// get from db
	tx := d.gorm.First(&farm, "id = ?", farmID)
	if errors.Is(tx.Error, gorm.ErrRecordNotFound) {
		return nil, nil
	}

	return farm, tx.Error
}
