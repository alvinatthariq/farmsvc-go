package controllers

import (
	"database/sql"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"time"

	"github.com/alvinatthariq/farmsvc-go/entity"

	"github.com/go-sql-driver/mysql"
	"github.com/gorilla/mux"
)

func (c *controller) CreateFarm(w http.ResponseWriter, r *http.Request) {
	// parse request body
	var createFarmRequest entity.CreateFarmRequest
	if err := json.NewDecoder(r.Body).Decode(&createFarmRequest); err != nil {
		httpRespError(w, r, fmt.Errorf("Error Decode Request Body : %w", err), http.StatusInternalServerError)
		return
	}

	farm, err := c.domain.CreateFarm(createFarmRequest)
	if err != nil {
		var mysqlError *mysql.MySQLError
		if errors.As(err, &mysqlError) {
			// check duplicate constraint
			if mysqlError.Number == entity.CodeMySQLDuplicateEntry {
				httpRespError(w, r, err, http.StatusConflict)
				return
			}
		}
		httpRespError(w, r, err, http.StatusInternalServerError)
		return
	}

	httpRespSuccess(w, r, http.StatusCreated, farm, nil)
}

func (c *controller) GetFarmById(w http.ResponseWriter, r *http.Request) {
	farmID := mux.Vars(r)["id"]

	farmRes, err := c.domain.GetFarmByID(farmID)
	if err != nil {
		httpRespError(w, r, fmt.Errorf("Error when get farm by id : %w", err), http.StatusInternalServerError)
		return
	} else if farmRes == nil {
		httpRespError(w, r, fmt.Errorf("Farm Not Found !"), http.StatusNotFound)
		return
	}

	httpRespSuccess(w, r, http.StatusOK, *farmRes, nil)
}

func (c *controller) GetFarm(w http.ResponseWriter, r *http.Request) {
	var (
		farms      []entity.Farm
		pagination entity.Pagination
	)

	// get from db
	tx := c.gorm.Scopes(paginate(farms, &pagination, c.gorm)).Where("is_deleted is null").Find(&farms)
	if tx.Error != nil {
		httpRespError(w, r, tx.Error, http.StatusInternalServerError)
		return
	}

	if len(farms) < 1 {
		httpRespError(w, r, fmt.Errorf("Farm Not Found !"), http.StatusNotFound)
		return
	}

	httpRespSuccess(w, r, http.StatusOK, farms, &pagination)
}

func (c *controller) UpdateFarm(w http.ResponseWriter, r *http.Request) {
	farmID := mux.Vars(r)["id"]

	// read request body
	var reqBody entity.UpdateFarmRequest
	if err := json.NewDecoder(r.Body).Decode(&reqBody); err != nil {
		httpRespError(w, r, fmt.Errorf("Error Decode Request Body : %w", err), http.StatusInternalServerError)
		return
	}

	var farm entity.Farm
	farmRes, err := c.domain.GetFarmByID(farmID)
	if err != nil {
		httpRespError(w, r, fmt.Errorf("Error when get farm by id : %w", err), http.StatusInternalServerError)
		return
	} else if farmRes == nil {
		farm, err = c.domain.CreateFarm(entity.CreateFarmRequest{
			ID:          farmID,
			Name:        reqBody.Name,
			Description: reqBody.Description,
		})
		if err != nil {
			var mysqlError *mysql.MySQLError
			if errors.As(err, &mysqlError) {
				// check duplicate constraint
				if mysqlError.Number == entity.CodeMySQLDuplicateEntry {
					httpRespError(w, r, err, http.StatusConflict)
					return
				}
			}
			httpRespError(w, r, err, http.StatusInternalServerError)
			return
		}
	} else {
		// update
		farm = *farmRes
		farm.Name = reqBody.Name
		farm.Description = reqBody.Description
		farm.UpdatedAt = time.Now().UTC()

		tx := c.gorm.Save(&farm)
		if tx.Error != nil {
			httpRespError(w, r, tx.Error, http.StatusInternalServerError)
			return
		}
	}

	httpRespSuccess(w, r, http.StatusOK, farm, nil)
}

func (c *controller) DeleteFarm(w http.ResponseWriter, r *http.Request) {
	farmID := mux.Vars(r)["id"]

	var farm entity.Farm
	farmRes, err := c.domain.GetFarmByID(farmID)
	if err != nil {
		httpRespError(w, r, fmt.Errorf("Error when get farm by id : %w", err), http.StatusInternalServerError)
		return
	} else if farmRes == nil {
		httpRespError(w, r, fmt.Errorf("Farm Not Found !"), http.StatusBadRequest)
		return
	} else {
		farm = *farmRes
		if !farm.IsDeleted.Bool {
			// soft delete
			farm.IsDeleted = sql.NullBool{Bool: true, Valid: true}
			farm.DeletedAt = sql.NullTime{Time: time.Now().UTC(), Valid: true}
			c.gorm.Save(&farm)
		}
	}

	httpRespSuccess(w, r, http.StatusOK, nil, nil)
}
