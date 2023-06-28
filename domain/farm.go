package domain

import (
	"database/sql"
	"errors"
	"time"

	"github.com/alvinatthariq/farmsvc-go/entity"
	"github.com/go-sql-driver/mysql"
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

	err = farm.Validate()
	if err != nil {
		return farm, err
	}

	// create to db
	err = d.gorm.Create(&farm).Error
	if err != nil {
		var mysqlError *mysql.MySQLError
		if errors.As(err, &mysqlError) {
			// check duplicate constraint
			if mysqlError.Number == entity.CodeMySQLDuplicateEntry {
				return farm, entity.ErrorFarmAlreadyExist
			}
		}
	}

	return farm, err
}

func (d *domain) GetFarmByID(farmID string) (farm *entity.Farm, err error) {
	// get from db
	err = d.gorm.First(&farm, "id = ?", farmID).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, nil
	}

	return farm, err
}

func (d *domain) GetFarm() (farms []entity.Farm, err error) {
	// get from db
	err = d.gorm.Where("is_deleted is null").Find(&farms).Error
	if err != nil {
		return farms, err
	}

	return farms, nil
}

func (d *domain) UpdateFarm(farmID string, v entity.UpdateFarmRequest) (farm entity.Farm, err error) {
	farmRes, err := d.GetFarmByID(farmID)
	if err != nil {
		return farm, err
	} else if farmRes == nil {
		// create if not exist
		farm, err = d.CreateFarm(entity.CreateFarmRequest{
			ID:          farmID,
			Name:        v.Name,
			Description: v.Description,
		})
		if err != nil {
			return farm, err
		}
	} else {
		// update if exist
		farm = *farmRes
		farm.Name = v.Name
		farm.Description = v.Description
		farm.UpdatedAt = time.Now().UTC()

		err = farm.Validate()
		if err != nil {
			return farm, err
		}

		err = d.gorm.Save(&farm).Error
		if err != nil {
			return farm, err
		}
	}

	return farm, nil
}

func (d *domain) DeleteFarmByID(farmID string) (err error) {
	var farm entity.Farm
	farmRes, err := d.GetFarmByID(farmID)
	if err != nil {
		return err
	} else if farmRes == nil {
		return entity.ErrorFarmNotFound
	} else {
		farm = *farmRes
		if !farm.IsDeleted.Bool {
			// soft delete
			farm.IsDeleted = sql.NullBool{Bool: true, Valid: true}
			farm.DeletedAt = sql.NullTime{Time: time.Now().UTC(), Valid: true}
			if err := d.gorm.Save(&farm).Error; err != nil {
				return err
			}
		}
	}

	return nil
}
