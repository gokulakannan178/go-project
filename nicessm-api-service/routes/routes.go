package routes

import (
	"fmt"
	"net/http"
	"nicessm-api-service/config"
	"nicessm-api-service/handlers"
	"nicessm-api-service/middlewares"
	"nicessm-api-service/models"
	"nicessm-api-service/redis"
	"nicessm-api-service/shared"

	"github.com/gorilla/mux"
)

//Route : ""
type Route struct {
	Handler      *handlers.Handler
	Shared       *shared.Shared
	Redis        *redis.RedisCli
	ConfigReader *config.ViperConfigReader
	Cache        *models.CacheMemory
}

//GetRoute : ""
func GetRoute(handler *handlers.Handler, s *shared.Shared, Redis *redis.RedisCli, configReader *config.ViperConfigReader, cache *models.CacheMemory) *Route {
	return &Route{handler, s, Redis, configReader, cache}
}

// Adapt :
func Adapt(h http.Handler, adapters ...middlewares.Adapter) http.Handler {
	if len(adapters) == 0 {
		return h
	}
	return adapters[0](Adapt(h, adapters[1:cap(adapters)]...))
}

//CommonRoutes : " All common routes come here"
//Log
//Updated by Solomon Arumuhan (solomon@logikoof.com) on 07-Mar-2022 - new route '/common/uniqueness'
//Updated by  Gokulkannan (Gokulkannan.M@logikoof.com) on 10-Mar-2022 - new route '/common/uniquenesscheck'
func (rout *Route) CommonRoutes(r *mux.Router) {
	//CommonRoutes Default Route
	r.Handle("/", Adapt(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("Handled By Basic Handler updated")
	})))
	//Added by Solomon Arumuhan (solomon@logikoof.com) on 07-Mar-2022
	r.Handle("/common/uniqueness", Adapt(http.HandlerFunc(rout.Handler.ChkCommonUniqueness))).Methods("GET")
	r.Handle("/common/uniquenesscheck", Adapt(http.HandlerFunc(rout.Handler.ChkCommonUniquenessWithoutRegex))).Methods("GET")

}
