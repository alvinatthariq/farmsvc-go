package domain

import (
	"database/sql"
	"errors"
	"math"
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

	// create to db
	tx := d.gorm.Create(&farm)
	if tx.Error != nil {
		var mysqlError *mysql.MySQLError
		if errors.As(tx.Error, &mysqlError) {
			// check duplicate constraint
			if mysqlError.Number == entity.CodeMySQLDuplicateEntry {
				return farm, entity.ErrorFarmAlreadyExist
			}
		}
	}

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

func (d *domain) GetFarm() (farms []entity.Farm, err error) {
	var pagination entity.Pagination

	// get from db
	tx := d.gorm.Scopes(paginate(farms, &pagination, d.gorm)).Where("is_deleted is null").Find(&farms)
	if tx.Error != nil {
		return farms, tx.Error
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

		tx := d.gorm.Save(&farm)
		if tx.Error != nil {
			return farm, tx.Error
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
			d.gorm.Save(&farm)
		}
	}

	return nil
}

func paginate(value interface{}, pagination *entity.Pagination, db *gorm.DB) func(db *gorm.DB) *gorm.DB {
	var totalRows int64
	db.Model(value).Count(&totalRows)

	pagination.TotalRows = totalRows
	totalPages := int(math.Ceil(float64(totalRows) / float64(pagination.Limit)))
	pagination.TotalPages = totalPages

	return func(db *gorm.DB) *gorm.DB {
		return db.Offset(pagination.GetOffset()).Limit(pagination.GetLimit()).Order(pagination.GetSort())
	}
}
