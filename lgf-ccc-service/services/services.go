package service

import (
	"lgf-ccc-service/config"
	"lgf-ccc-service/daos"
	"lgf-ccc-service/redis"
	"lgf-ccc-service/shared"
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
