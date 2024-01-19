package services

import (
	"municipalproduct1-service/config"
	"municipalproduct1-service/daos"
	"municipalproduct1-service/redis"
	"municipalproduct1-service/shared"
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
