package controllers

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"strconv"

	"github.com/alvinatthariq/farmsvc-go/entity"

	"github.com/gorilla/mux"
)

func (c *controller) CreatePond(w http.ResponseWriter, r *http.Request) {
	// upsert api statistic
	c.domain.UpsertAPIStatistic(entity.APIPathPOSTPond, r.UserAgent())

	// parse request body
	var createPondRequest entity.CreatePondRequest
	if err := json.NewDecoder(r.Body).Decode(&createPondRequest); err != nil {
		httpRespError(w, r, fmt.Errorf("Error Decode Request Body : %w", err), http.StatusInternalServerError)
		return
	}

	pond, err := c.domain.CreatePond(createPondRequest)
	if err != nil {
		switch err {
		case entity.ErrorPondAlreadyExist:
			httpRespError(w, r, err, http.StatusConflict)
			return
		case
			entity.ErrorFarmNotFound,
			entity.ErrorFarmIDRequired,
			entity.ErrorFarmIDMaxLength,
			entity.ErrorPondIDRequired,
			entity.ErrorPondIDMaxLength,
			entity.ErrorPondNameRequired,
			entity.ErrorPondNameMaxLength,
			entity.ErrorPondDescriptionRequired,
			entity.ErrorPondDescriptionMaxLength:
			httpRespError(w, r, err, http.StatusBadRequest)
			return
		default:
			httpRespError(w, r, err, http.StatusInternalServerError)
			return
		}
	}

	httpRespSuccess(w, r, http.StatusCreated, pond)
}

func (c *controller) GetPondByID(w http.ResponseWriter, r *http.Request) {
	// upsert api statistic
	c.domain.UpsertAPIStatistic(entity.APIPathGETPondByID, r.UserAgent())

	pondID := mux.Vars(r)["id"]

	pondRes, err := c.domain.GetPondByID(pondID)
	if err != nil {
		httpRespError(w, r, fmt.Errorf("Error when get pond by id : %w", err), http.StatusInternalServerError)
		return
	} else if pondRes == nil {
		httpRespError(w, r, entity.ErrorPondNotFound, http.StatusNotFound)
		return
	}

	httpRespSuccess(w, r, http.StatusOK, *pondRes)
}

func (c *controller) GetPond(w http.ResponseWriter, r *http.Request) {
	// upsert api statistic
	c.domain.UpsertAPIStatistic(entity.APIPathGETPond, r.UserAgent())

	// get url query param
	urlVal := r.URL.Query()

	// limit
	limitStr := urlVal.Get("limit")
	limit, _ := strconv.Atoi(limitStr)
	if limit < 1 {
		limit = 10
	}

	// page
	pageStr := urlVal.Get("page")
	page, _ := strconv.Atoi(pageStr)

	param := entity.PondParam{
		ID:     urlVal.Get("id"),
		FarmID: urlVal.Get("farm_id"),
		Name:   urlVal.Get("name"),
		Page:   page,
		Limit:  limit,
	}

	ponds, err := c.domain.GetPond(param)
	if err != nil {
		httpRespError(w, r, err, http.StatusInternalServerError)
		return
	}

	if len(ponds) < 1 {
		httpRespError(w, r, fmt.Errorf("Pond Not Found !"), http.StatusNotFound)
		return
	}

	httpRespSuccess(w, r, http.StatusOK, ponds)
}

func (c *controller) UpdatePond(w http.ResponseWriter, r *http.Request) {
	// upsert api statistic
	c.domain.UpsertAPIStatistic(entity.APIPathPUTPondByID, r.UserAgent())

	pondID := mux.Vars(r)["id"]

	// read request body
	var reqBody entity.UpdatePondRequest
	if err := json.NewDecoder(r.Body).Decode(&reqBody); err != nil {
		httpRespError(w, r, fmt.Errorf("Error Decode Request Body : %w", err), http.StatusInternalServerError)
		return
	}

	pond, err := c.domain.UpdatePond(pondID, reqBody)
	if err != nil {
		switch err {
		case entity.ErrorPondAlreadyExist:
			httpRespError(w, r, err, http.StatusConflict)
			return
		case
			entity.ErrorFarmNotFound,
			entity.ErrorFarmIDRequired,
			entity.ErrorFarmIDMaxLength,
			entity.ErrorPondIDRequired,
			entity.ErrorPondIDMaxLength,
			entity.ErrorPondNameRequired,
			entity.ErrorPondNameMaxLength,
			entity.ErrorPondDescriptionRequired,
			entity.ErrorPondDescriptionMaxLength:
			httpRespError(w, r, err, http.StatusBadRequest)
			return
		default:
			httpRespError(w, r, err, http.StatusInternalServerError)
			return
		}
	}

	httpRespSuccess(w, r, http.StatusOK, pond)
}

func (c *controller) DeletePondByID(w http.ResponseWriter, r *http.Request) {
	// upsert api statistic
	c.domain.UpsertAPIStatistic(entity.APIPathDELETEPondByID, r.UserAgent())

	pondID := mux.Vars(r)["id"]

	err := c.domain.DeletePondByID(pondID)
	if err != nil {
		if errors.Is(err, entity.ErrorPondNotFound) {
			httpRespError(w, r, err, http.StatusBadRequest)
			return
		}
		httpRespError(w, r, fmt.Errorf("Error DeletePondByID : %w", err), http.StatusInternalServerError)
		return
	}

	httpRespSuccess(w, r, http.StatusOK, nil)
}
