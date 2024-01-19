package routes

import (
	"fmt"
	"haritv2-service/config"
	"haritv2-service/handlers"
	"haritv2-service/middlewares"
	"haritv2-service/redis"
	"haritv2-service/shared"
	"net/http"

	"github.com/gorilla/mux"
)

//Route : ""
type Route struct {
	Handler      *handlers.Handler
	Shared       *shared.Shared
	Redis        *redis.RedisCli
	ConfigReader *config.ViperConfigReader
}

//GetRoute : ""
func GetRoute(handler *handlers.Handler, s *shared.Shared, Redis *redis.RedisCli, configReader *config.ViperConfigReader) *Route {
	return &Route{handler, s, Redis, configReader}
}

// Adapt :
func Adapt(h http.Handler, adapters ...middlewares.Adapter) http.Handler {
	if len(adapters) == 0 {
		return h
	}
	return adapters[0](Adapt(h, adapters[1:cap(adapters)]...))
}

//CommonRoutes : ""
func (rout *Route) CommonRoutes(r *mux.Router) {
	//CommonRoutes Default Route
	r.Handle("/", Adapt(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("Handled By Basic Handler updated")
	})))

}
