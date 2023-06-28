package controllers

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"

	"github.com/alvinatthariq/farmsvc-go/entity"

	"github.com/gorilla/mux"
)

func (c *controller) CreateFarm(w http.ResponseWriter, r *http.Request) {
	// upsert api statistic
	c.domain.UpsertAPIStatistic(entity.APIPathPOSTFarm, r.UserAgent())

	// parse request body
	var createFarmRequest entity.CreateFarmRequest
	if err := json.NewDecoder(r.Body).Decode(&createFarmRequest); err != nil {
		httpRespError(w, r, fmt.Errorf("Error Decode Request Body : %w", err), http.StatusInternalServerError)
		return
	}

	farm, err := c.domain.CreateFarm(createFarmRequest)
	if err != nil {
		switch err {
		case entity.ErrorFarmAlreadyExist:
			httpRespError(w, r, err, http.StatusConflict)
			return
		case
			entity.ErrorFarmIDRequired,
			entity.ErrorFarmIDMaxLength,
			entity.ErrorFarmNameRequired,
			entity.ErrorFarmNameMaxLength,
			entity.ErrorFarmDescriptionRequired,
			entity.ErrorFarmDescriptionMaxLength:
			httpRespError(w, r, err, http.StatusBadRequest)
			return
		default:
			httpRespError(w, r, err, http.StatusInternalServerError)
			return
		}
	}

	httpRespSuccess(w, r, http.StatusCreated, farm)
}

func (c *controller) GetFarmByID(w http.ResponseWriter, r *http.Request) {
	// upsert api statistic
	c.domain.UpsertAPIStatistic(entity.APIPathGETFarmByID, r.UserAgent())

	farmID := mux.Vars(r)["id"]

	farmRes, err := c.domain.GetFarmByID(farmID)
	if err != nil {
		httpRespError(w, r, fmt.Errorf("Error when get farm by id : %w", err), http.StatusInternalServerError)
		return
	} else if farmRes == nil {
		httpRespError(w, r, entity.ErrorFarmNotFound, http.StatusNotFound)
		return
	}

	httpRespSuccess(w, r, http.StatusOK, *farmRes)
}

func (c *controller) GetFarm(w http.ResponseWriter, r *http.Request) {
	// upsert api statistic
	c.domain.UpsertAPIStatistic(entity.APIPathGETFarm, r.UserAgent())

	// get url query param
	urlVal := r.URL.Query()

	param := entity.FarmParam{
		ID:   urlVal.Get("id"),
		Name: urlVal.Get("name"),
	}

	farms, err := c.domain.GetFarm(param)
	if err != nil {
		httpRespError(w, r, err, http.StatusInternalServerError)
		return
	}

	if len(farms) < 1 {
		httpRespError(w, r, fmt.Errorf("Farm Not Found !"), http.StatusNotFound)
		return
	}

	httpRespSuccess(w, r, http.StatusOK, farms)
}

func (c *controller) UpdateFarm(w http.ResponseWriter, r *http.Request) {
	// upsert api statistic
	c.domain.UpsertAPIStatistic(entity.APIPathPUTFarmByID, r.UserAgent())

	farmID := mux.Vars(r)["id"]

	// read request body
	var reqBody entity.UpdateFarmRequest
	if err := json.NewDecoder(r.Body).Decode(&reqBody); err != nil {
		httpRespError(w, r, fmt.Errorf("Error Decode Request Body : %w", err), http.StatusInternalServerError)
		return
	}

	farm, err := c.domain.UpdateFarm(farmID, reqBody)
	if err != nil {
		switch err {
		case entity.ErrorFarmAlreadyExist:
			httpRespError(w, r, err, http.StatusConflict)
			return
		case
			entity.ErrorFarmIDRequired,
			entity.ErrorFarmIDMaxLength,
			entity.ErrorFarmNameRequired,
			entity.ErrorFarmNameMaxLength,
			entity.ErrorFarmDescriptionRequired,
			entity.ErrorFarmDescriptionMaxLength:
			httpRespError(w, r, err, http.StatusBadRequest)
			return
		default:
			httpRespError(w, r, err, http.StatusInternalServerError)
			return
		}
	}

	httpRespSuccess(w, r, http.StatusOK, farm)
}

func (c *controller) DeleteFarmByID(w http.ResponseWriter, r *http.Request) {
	// upsert api statistic
	c.domain.UpsertAPIStatistic(entity.APIPathDELETEFarmByID, r.UserAgent())

	farmID := mux.Vars(r)["id"]

	err := c.domain.DeleteFarmByID(farmID)
	if err != nil {
		if errors.Is(err, entity.ErrorFarmNotFound) {
			httpRespError(w, r, err, http.StatusBadRequest)
			return
		}
		httpRespError(w, r, fmt.Errorf("Error DeleteFarmByID : %w", err), http.StatusInternalServerError)
		return
	}

	httpRespSuccess(w, r, http.StatusOK, nil)
}
