package domain

import (
	"errors"

	"github.com/alvinatthariq/farmsvc-go/entity"
	"github.com/go-redis/redis"
)

func (d *domain) UpsertAPIStatistic(apiPath string, userAgent string) error {
	res := d.redisClient.IncrBy(apiPath, 1)
	if res.Err() != nil {
		return res.Err()
	}

	res = d.redisClient.PFAdd(apiPath+"ua", userAgent)
	if res.Err() != nil {
		return res.Err()
	}

	return nil
}

func (d *domain) GetAPIStatistic() (apiStatistics []entity.APIStatistic, err error) {
	for _, apiPath := range entity.APIPaths {
		res := d.redisClient.Get(apiPath)
		if res.Err() != nil {
			if !errors.Is(res.Err(), redis.Nil) {
				return apiStatistics, res.Err()
			}
		}

		count, err := res.Int64()
		if err != nil {
			if !errors.Is(err, redis.Nil) {
				return apiStatistics, err
			}
		}

		resUa := d.redisClient.PFCount(apiPath + "ua")
		if resUa.Err() != nil {
			if !errors.Is(resUa.Err(), redis.Nil) {
				return apiStatistics, resUa.Err()
			}
		}

		uniqueUserAgent := resUa.Val()

		apiStat := entity.APIStatistic{
			Path:            apiPath,
			Count:           count,
			UniqueUserAgent: uniqueUserAgent,
		}

		apiStatistics = append(apiStatistics, apiStat)
	}

	return apiStatistics, nil
}
