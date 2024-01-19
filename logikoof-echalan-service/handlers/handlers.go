package handlers

import (
	"logikoof-echalan-service/config"
	"logikoof-echalan-service/redis"
	"logikoof-echalan-service/services"
	"logikoof-echalan-service/shared"
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
