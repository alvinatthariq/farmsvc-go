package entity

var APIPaths = []string{
	APIPathPOSTFarm,
	APIPathGETFarm,
	APIPathGETFarmByID,
	APIPathPUTFarmByID,
	APIPathDELETEFarmByID,
}

const (
	APIPathPOSTFarm       = "POST /v1/farm"
	APIPathGETFarm        = "GET /v1/farm"
	APIPathGETFarmByID    = "GET /v1/farm/{id}"
	APIPathPUTFarmByID    = "PUT /v1/farm/{id}"
	APIPathDELETEFarmByID = "DELETE /v1/farm/{id}"
)

type APIStatistic struct {
	Path            string `json:"path"`
	Count           int64  `json:"count"`
	UniqueUserAgent int64  `json:"unique_user_agent"`
}
