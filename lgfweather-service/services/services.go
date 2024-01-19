package services

import (
	"lgfweather-service/config"
	"lgfweather-service/daos"
	"lgfweather-service/redis"
	"lgfweather-service/shared"
)

//Service : ""
type Service struct {
	Daos         *daos.Daos
	Shared       *shared.Shared
	Redis        *redis.RedisCli
	ConfigReader *config.ViperConfigReader
}

//IService : ""
type IService interface {
}

//GetService :""
func GetService(dao *daos.Daos, s *shared.Shared, Redis *redis.RedisCli, configReader *config.ViperConfigReader) *Service {
	return &Service{dao, s, Redis, configReader}
}
