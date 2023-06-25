package domain

import (
	"github.com/alvinatthariq/farmsvc-go/entity"

	"gorm.io/gorm"
)

type domain struct {
	gorm *gorm.DB
}

type DomainItf interface {
	CreateFarm(v entity.CreateFarmRequest) (farm entity.Farm, err error)
	GetFarmByID(farmID string) (farm *entity.Farm, err error)
}

func Init(gorm *gorm.DB) DomainItf {
	return &domain{
		gorm: gorm,
	}
}
