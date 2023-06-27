package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/alvinatthariq/farmsvc-go/entity"
)

func httpRespError(w http.ResponseWriter, r *http.Request, err error, statusCode int) {
	statusStr := http.StatusText(statusCode)

	jsonErrResp := &entity.HTTPEmptyResp{
		Meta: entity.Meta{
			Path:       r.URL.String(),
			StatusCode: statusCode,
			Status:     statusStr,
			Message:    fmt.Sprintf("%s %s [%d] %s", r.Method, r.URL.RequestURI(), statusCode, statusStr),
			Error:      err.Error(),
			Timestamp:  time.Now().Format(time.RFC3339),
		},
	}

	raw, err := json.Marshal(jsonErrResp)
	if err != nil {
		statusCode = http.StatusInternalServerError
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)
	_, _ = w.Write(raw)
}

func httpRespSuccess(w http.ResponseWriter, r *http.Request, statusCode int, resp interface{}) {
	meta := entity.Meta{
		Path:       r.URL.String(),
		StatusCode: statusCode,
		Status:     http.StatusText(statusCode),
		Message:    fmt.Sprintf("%s %s [%d] %s", r.Method, r.URL.RequestURI(), statusCode, http.StatusText(statusCode)),
		Timestamp:  time.Now().Format(time.RFC3339),
	}

	var (
		raw []byte
		err error
	)
	switch data := resp.(type) {
	case nil:
		httpResp := &entity.HTTPEmptyResp{
			Meta: meta,
		}
		raw, err = json.Marshal(httpResp)
		if err != nil {
			statusCode = http.StatusInternalServerError
		}
	case entity.Farm:
		httpResp := &entity.HTTPFarmResp{
			Meta: meta,
			Data: entity.HTTPFarmData{
				Farm: data,
			},
		}
		raw, err = json.Marshal(httpResp)
		if err != nil {
			statusCode = http.StatusInternalServerError
		}
	case []entity.Farm:
		httpResp := &entity.HTTPFarmsResp{
			Meta: meta,
			Data: entity.HTTPFarmsData{
				Farms: data,
			},
		}
		raw, err = json.Marshal(httpResp)
		if err != nil {
			statusCode = http.StatusInternalServerError
		}
	case entity.Pond:
		httpResp := &entity.HTTPPondResp{
			Meta: meta,
			Data: entity.HTTPPondData{
				Pond: data,
			},
		}
		raw, err = json.Marshal(httpResp)
		if err != nil {
			statusCode = http.StatusInternalServerError
		}
	case []entity.Pond:
		httpResp := &entity.HTTPPondsResp{
			Meta: meta,
			Data: entity.HTTPPondsData{
				Ponds: data,
			},
		}
		raw, err = json.Marshal(httpResp)
		if err != nil {
			statusCode = http.StatusInternalServerError
		}
	case []entity.APIStatistic:
		httpResp := &entity.HTTPAPIStatisticsResp{
			Meta: meta,
			Data: entity.HTTPAPIStatisticsData{
				APIStatistics: data,
			},
		}
		raw, err = json.Marshal(httpResp)
		if err != nil {
			statusCode = http.StatusInternalServerError
		}

	default:
		httpRespError(w, r, fmt.Errorf("cannot cast type of %+v", data), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)
	_, _ = w.Write(raw)
}
