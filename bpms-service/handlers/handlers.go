package handlers

import (
	"bpms-service/config"
	"bpms-service/redis"
	"bpms-service/services"
	"bpms-service/shared"
)

//Handler : ""
type Handler struct {
	Service      *services.Service
	Shared       *shared.Shared
	Redis        *redis.RedisCli
	ConfigReader *config.ViperConfigReader
}

//GetHandler :""
func GetHandler(service *services.Service, s *shared.Shared, Redis *redis.RedisCli, configReader *config.ViperConfigReader) *Handler {
	return &Handler{service, s, Redis, configReader}
}
