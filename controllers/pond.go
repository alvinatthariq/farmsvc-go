package controllers

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"

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
		if errors.Is(err, entity.ErrorPondAlreadyExist) {
			httpRespError(w, r, err, http.StatusConflict)
			return
		} else if errors.Is(err, entity.ErrorFarmNotFound) {
			httpRespError(w, r, err, http.StatusBadRequest)
			return
		}

		httpRespError(w, r, err, http.StatusInternalServerError)
		return
	}

	httpRespSuccess(w, r, http.StatusCreated, pond, nil)
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

	httpRespSuccess(w, r, http.StatusOK, *pondRes, nil)
}

func (c *controller) GetPond(w http.ResponseWriter, r *http.Request) {
	// upsert api statistic
	c.domain.UpsertAPIStatistic(entity.APIPathGETPond, r.UserAgent())

	ponds, err := c.domain.GetPond()
	if err != nil {
		httpRespError(w, r, err, http.StatusInternalServerError)
		return
	}

	if len(ponds) < 1 {
		httpRespError(w, r, fmt.Errorf("Pond Not Found !"), http.StatusNotFound)
		return
	}

	httpRespSuccess(w, r, http.StatusOK, ponds, nil)
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
		if errors.Is(err, entity.ErrorFarmNotFound) {
			httpRespError(w, r, err, http.StatusBadRequest)
			return
		}

		httpRespError(w, r, fmt.Errorf("Error UpdatePond : %w", err), http.StatusInternalServerError)
		return
	}

	httpRespSuccess(w, r, http.StatusOK, pond, nil)
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

	httpRespSuccess(w, r, http.StatusOK, nil, nil)
}
