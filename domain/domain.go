package domain

import (
	"github.com/alvinatthariq/farmsvc-go/entity"

	"github.com/go-redis/redis"
	"gorm.io/gorm"
)

type domain struct {
	gorm        *gorm.DB
	redisClient *redis.Client
}

type DomainItf interface {
	// Farm
	CreateFarm(v entity.CreateFarmRequest) (farm entity.Farm, err error)
	GetFarmByID(farmID string) (farm *entity.Farm, err error)
	GetFarm() (farms []entity.Farm, err error)
	UpdateFarm(farmID string, v entity.UpdateFarmRequest) (farm entity.Farm, err error)
	DeleteFarmByID(farmID string) (err error)

	// Pond
	CreatePond(v entity.CreatePondRequest) (pond entity.Pond, err error)
	GetPondByID(pondID string) (pond *entity.Pond, err error)
	GetPond() (ponds []entity.Pond, err error)
	UpdatePond(pondID string, v entity.UpdatePondRequest) (pond entity.Pond, err error)
	DeletePondByID(pondID string) (err error)

	// API Statistic
	UpsertAPIStatistic(apiPath string, userAgent string) error
	GetAPIStatistic() (apiStatistics []entity.APIStatistic, err error)
}

func Init(gorm *gorm.DB, redisClient *redis.Client) DomainItf {
	return &domain{
		gorm:        gorm,
		redisClient: redisClient,
	}
}
