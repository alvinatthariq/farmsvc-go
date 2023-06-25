package controllers

import "net/http"

func (c *controller) GetAPIStatistic(w http.ResponseWriter, r *http.Request) {
	apiStatistics, err := c.domain.GetAPIStatistic()
	if err != nil {
		httpRespError(w, r, err, http.StatusInternalServerError)
		return
	}

	httpRespSuccess(w, r, http.StatusOK, apiStatistics, nil)
}
