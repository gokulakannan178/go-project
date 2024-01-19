package handlers

import (
	"hrms-services/config"
	"hrms-services/models"
	"hrms-services/redis"
	"hrms-services/services"
	"hrms-services/shared"
)

//Handler : ""
type Handler struct {
	Service      *services.Service
	Shared       *shared.Shared
	Redis        *redis.RedisCli
	ConfigReader *config.ViperConfigReader
	Cache        *models.CacheMemory
}

//GetHandler :""
func GetHandler(service *services.Service, s *shared.Shared, Redis *redis.RedisCli, configReader *config.ViperConfigReader, Cache *models.CacheMemory) *Handler {
	return &Handler{service, s, Redis, configReader, Cache}
}
