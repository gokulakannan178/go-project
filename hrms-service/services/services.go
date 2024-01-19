package services

import (
	"hrms-services/config"
	"hrms-services/daos"
	"hrms-services/models"
	"hrms-services/redis"
	"hrms-services/shared"
)

//Service : ""
type Service struct {
	Daos         *daos.Daos
	Shared       *shared.Shared
	Redis        *redis.RedisCli
	ConfigReader *config.ViperConfigReader
	Cache        *models.CacheMemory
}

//IService : ""
type IService interface {
}

//GetService :""
func GetService(dao *daos.Daos, s *shared.Shared, Redis *redis.RedisCli, configReader *config.ViperConfigReader, cache *models.CacheMemory) *Service {
	return &Service{dao, s, Redis, configReader, cache}
}
