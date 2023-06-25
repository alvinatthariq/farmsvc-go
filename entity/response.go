package entity

type HTTPFarmResp struct {
	Meta Meta         `json:"meta"`
	Data HTTPFarmData `json:"data"`
}

type HTTPFarmData struct {
	Farm Farm `json:"farm"`
}

type HTTPFarmsResp struct {
	Meta       Meta          `json:"meta"`
	Data       HTTPFarmsData `json:"data"`
	Pagination *Pagination   `json:"pagination,omitempty"`
}

type HTTPFarmsData struct {
	Farms []Farm `json:"farms"`
}

type Meta struct {
	Path       string `json:"path"`
	StatusCode int    `json:"status_code"`
	Status     string `json:"status"`
	Message    string `json:"message"`
	Timestamp  string `json:"timestamp"`
	Error      string `json:"error,omitempty"`
}

type HTTPEmptyResp struct {
	Meta Meta `json:"metadata"`
}

type HTTPAPIStatisticsResp struct {
	Meta Meta                  `json:"meta"`
	Data HTTPAPIStatisticsData `json:"data"`
}

type HTTPAPIStatisticsData struct {
	APIStatistics []APIStatistic `json:"api_statistics"`
}

type HTTPPondResp struct {
	Meta Meta         `json:"meta"`
	Data HTTPPondData `json:"data"`
}

type HTTPPondData struct {
	Pond Pond `json:"pond"`
}

type HTTPPondsResp struct {
	Meta       Meta          `json:"meta"`
	Data       HTTPPondsData `json:"data"`
	Pagination *Pagination   `json:"pagination,omitempty"`
}

type HTTPPondsData struct {
	Ponds []Pond `json:"ponds"`
}
