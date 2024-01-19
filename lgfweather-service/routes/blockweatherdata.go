package routes

import (
	"net/http"

	"github.com/gorilla/mux"
)

func (route *Route) BlockWeatherDataRoutes(r *mux.Router) {
	r.Handle("/blockWeatherData", Adapt(http.HandlerFunc(route.Handler.SaveBlockWeatherData))).Methods("POST")
	r.Handle("/blockWeatherData", Adapt(http.HandlerFunc(route.Handler.GetSingleBlockWeatherData))).Methods("GET")
	r.Handle("/blockWeatherData", Adapt(http.HandlerFunc(route.Handler.UpdateBlockWeatherData))).Methods("PUT")
	r.Handle("/blockWeatherData/status/enable", Adapt(http.HandlerFunc(route.Handler.EnableBlockWeatherData))).Methods("PUT")
	r.Handle("/blockWeatherData/status/disable", Adapt(http.HandlerFunc(route.Handler.DisableBlockWeatherData))).Methods("PUT")
	r.Handle("/blockWeatherData/status/delete", Adapt(http.HandlerFunc(route.Handler.DeleteBlockWeatherData))).Methods("DELETE")
	r.Handle("/blockWeatherData/filter", Adapt(http.HandlerFunc(route.Handler.FilterBlockWeatherData))).Methods("POST")
	r.Handle("/blockWeatherData/Block", Adapt(http.HandlerFunc(route.Handler.SaveBlockWeatherDataWithOpenWebsite))).Methods("GET")
	r.Handle("/blockWeatherData/currentdate", Adapt(http.HandlerFunc(route.Handler.GetSingleBlockWeatherDataWithCurrentDate))).Methods("GET")
	r.Handle("/weatherdata/byblock/state/district", Adapt(http.HandlerFunc(route.Handler.GetBlockWeatherDataByBlockId))).Methods("GET")
}
