package handlers

import (
	"nicessm-api-service/config"
	"nicessm-api-service/models"
	"nicessm-api-service/redis"
	"nicessm-api-service/services"
	"nicessm-api-service/shared"
)

//Handler : ""
type Handler struct {
	Service      *services.Service
	Shared       *shared.Shared
	Redis        *redis.RedisCli
	ConfigReader *config.ViperConfigReader
	C            chan *models.WeatherDisseminationChennal
	Cache        *models.CacheMemory
}

//GetHandler :""
func GetHandler(service *services.Service, s *shared.Shared, Redis *redis.RedisCli, configReader *config.ViperConfigReader, C chan *models.WeatherDisseminationChennal, cache *models.CacheMemory) *Handler {
	return &Handler{service, s, Redis, configReader, C, cache}
}
