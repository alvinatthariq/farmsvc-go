package domain

import (
	"database/sql"
	"errors"
	"time"

	"github.com/alvinatthariq/farmsvc-go/entity"
	"github.com/go-sql-driver/mysql"
	"gorm.io/gorm"
)

func (d *domain) CreatePond(v entity.CreatePondRequest) (pond entity.Pond, err error) {
	// get farm by id
	farm, err := d.GetFarmByID(v.FarmID)
	if err != nil {
		return pond, err
	}

	if farm == nil {
		// return error if farm not found
		return pond, entity.ErrorFarmNotFound
	}

	// create Pond
	pond = entity.Pond{
		ID:          v.ID,
		FarmID:      v.FarmID,
		Name:        v.Name,
		Description: v.Description,
		CreatedAt:   time.Now().UTC(),
		UpdatedAt:   time.Now().UTC(),
	}

	err = pond.Validate()
	if err != nil {
		return pond, err
	}

	// create to db
	err = d.gorm.Create(&pond).Error
	if err != nil {
		var mysqlError *mysql.MySQLError
		if errors.As(err, &mysqlError) {
			// check duplicate constraint
			if mysqlError.Number == entity.CodeMySQLDuplicateEntry {
				return pond, entity.ErrorPondAlreadyExist
			}
		}
	}

	return pond, err
}

func (d *domain) GetPondByID(pondID string) (pond *entity.Pond, err error) {
	// get from db
	err = d.gorm.First(&pond, "id = ?", pondID).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, nil
	}

	return pond, err
}

func (d *domain) GetPond(param entity.PondParam) (ponds []entity.Pond, err error) {
	// get from db
	err = d.gorm.Where("is_deleted is null").Where(&param).Find(&ponds).Error
	if err != nil {
		return ponds, err
	}

	return ponds, nil
}

func (d *domain) UpdatePond(pondID string, v entity.UpdatePondRequest) (pond entity.Pond, err error) {
	// get farm by id
	farm, err := d.GetFarmByID(v.FarmID)
	if err != nil {
		return pond, err
	}

	if farm == nil {
		// return error if farm not found
		return pond, entity.ErrorFarmNotFound
	}

	pondRes, err := d.GetPondByID(pondID)
	if err != nil {
		return pond, err
	} else if pondRes == nil {
		// create if not exist
		pond, err = d.CreatePond(entity.CreatePondRequest{
			ID:          pondID,
			FarmID:      v.FarmID,
			Name:        v.Name,
			Description: v.Description,
		})
		if err != nil {
			return pond, err
		}
	} else {
		// update if exist
		pond = *pondRes
		pond.FarmID = v.FarmID
		pond.Name = v.Name
		pond.Description = v.Description
		pond.UpdatedAt = time.Now().UTC()

		err = pond.Validate()
		if err != nil {
			return pond, err
		}

		err := d.gorm.Save(&pond).Error
		if err != nil {
			return pond, err
		}
	}

	return pond, nil
}

func (d *domain) DeletePondByID(pondID string) (err error) {
	var pond entity.Pond
	pondRes, err := d.GetPondByID(pondID)
	if err != nil {
		return err
	} else if pondRes == nil {
		return entity.ErrorPondNotFound
	} else {
		pond = *pondRes
		if !pond.IsDeleted.Bool {
			// soft delete
			pond.IsDeleted = sql.NullBool{Bool: true, Valid: true}
			pond.DeletedAt = sql.NullTime{Time: time.Now().UTC(), Valid: true}
			err = d.gorm.Save(&pond).Error
			if err != nil {
				return err
			}
		}
	}

	return nil
}
