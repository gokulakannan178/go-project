package routes

import (
	"net/http"

	"github.com/gorilla/mux"
)

func (route *Route) VillageWeatherDataRoutes(r *mux.Router) {
	r.Handle("/villageWeatherData", Adapt(http.HandlerFunc(route.Handler.SaveVillageWeatherData))).Methods("POST")
	r.Handle("/villageWeatherData", Adapt(http.HandlerFunc(route.Handler.GetSingleVillageWeatherData))).Methods("GET")
	r.Handle("/villageWeatherData", Adapt(http.HandlerFunc(route.Handler.UpdateVillageWeatherData))).Methods("PUT")
	r.Handle("/villageWeatherData/status/enable", Adapt(http.HandlerFunc(route.Handler.EnableVillageWeatherData))).Methods("PUT")
	r.Handle("/villageWeatherData/status/disable", Adapt(http.HandlerFunc(route.Handler.DisableVillageWeatherData))).Methods("PUT")
	r.Handle("/villageWeatherData/status/delete", Adapt(http.HandlerFunc(route.Handler.DeleteVillageWeatherData))).Methods("DELETE")
	r.Handle("/villageWeatherData/filter", Adapt(http.HandlerFunc(route.Handler.FilterVillageWeatherData))).Methods("POST")
	r.Handle("/villageWeatherData/Village", Adapt(http.HandlerFunc(route.Handler.SaveVillageWeatherDataWithOpenWebsite))).Methods("GET")
	r.Handle("/villageWeatherData/currentdate", Adapt(http.HandlerFunc(route.Handler.GetSingleVillageWeatherDataWithCurrentDate))).Methods("GET")
}
