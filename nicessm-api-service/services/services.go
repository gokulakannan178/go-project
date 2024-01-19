package services

import (
	"nicessm-api-service/config"
	"nicessm-api-service/daos"
	"nicessm-api-service/models"
	"nicessm-api-service/redis"
	"nicessm-api-service/shared"
)

//Service : ""
type Service struct {
	Daos         *daos.Daos
	Shared       *shared.Shared
	Redis        *redis.RedisCli
	ConfigReader *config.ViperConfigReader
	C            chan *models.WeatherDisseminationChennal
	Cache        *models.CacheMemory
}

//IService : ""
type IService interface {
}

//GetService :""
func GetService(dao *daos.Daos, s *shared.Shared, Redis *redis.RedisCli, configReader *config.ViperConfigReader, C chan *models.WeatherDisseminationChennal, cache *models.CacheMemory) *Service {
	return &Service{dao, s, Redis, configReader, C, cache}
}
