package handlers

import (
	"lgf-ccc-service/config"
	"lgf-ccc-service/redis"
	service "lgf-ccc-service/services"
	"lgf-ccc-service/shared"
)

// Handler : ""
type Handler struct {
	Service      *service.Service
	Shared       *shared.Shared
	Redis        *redis.RedisCli
	ConfigReader *config.ViperConfigReader
}

// GetHandler :""
func GetHandler(service *service.Service, s *shared.Shared, Redis *redis.RedisCli, configReader *config.ViperConfigReader) *Handler {
	return &Handler{service, s, Redis, configReader}
}
