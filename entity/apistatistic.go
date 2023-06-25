package entity

var APIPaths = []string{
	APIPathPOSTFarm,
	APIPathGETFarm,
	APIPathGETFarmByID,
	APIPathPUTFarmByID,
	APIPathDELETEFarmByID,
	APIPathPOSTPond,
	APIPathGETPond,
	APIPathGETPondByID,
	APIPathPUTPondByID,
	APIPathDELETEPondByID,
}

const (
	APIPathPOSTFarm       = "POST /v1/farm"
	APIPathGETFarm        = "GET /v1/farm"
	APIPathGETFarmByID    = "GET /v1/farm/{id}"
	APIPathPUTFarmByID    = "PUT /v1/farm/{id}"
	APIPathDELETEFarmByID = "DELETE /v1/farm/{id}"
	APIPathPOSTPond       = "POST /v1/pond"
	APIPathGETPond        = "GET /v1/pond"
	APIPathGETPondByID    = "GET /v1/pond/{id}"
	APIPathPUTPondByID    = "PUT /v1/pond/{id}"
	APIPathDELETEPondByID = "DELETE /v1/pond/{id}"
)

type APIStatistic struct {
	Path            string `json:"path"`
	Count           int64  `json:"count"`
	UniqueUserAgent int64  `json:"unique_user_agent"`
}
