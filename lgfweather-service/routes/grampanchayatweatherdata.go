package routes

import (
	"net/http"

	"github.com/gorilla/mux"
)

func (route *Route) GramPanchayatWeatherDataRoutes(r *mux.Router) {
	r.Handle("/gramPanchayatWeatherData", Adapt(http.HandlerFunc(route.Handler.SaveGramPanchayatWeatherData))).Methods("POST")
	r.Handle("/gramPanchayatWeatherData", Adapt(http.HandlerFunc(route.Handler.GetSingleGramPanchayatWeatherData))).Methods("GET")
	r.Handle("/gramPanchayatWeatherData", Adapt(http.HandlerFunc(route.Handler.UpdateGramPanchayatWeatherData))).Methods("PUT")
	r.Handle("/gramPanchayatWeatherData/status/enable", Adapt(http.HandlerFunc(route.Handler.EnableGramPanchayatWeatherData))).Methods("PUT")
	r.Handle("/gramPanchayatWeatherData/status/disable", Adapt(http.HandlerFunc(route.Handler.DisableGramPanchayatWeatherData))).Methods("PUT")
	r.Handle("/gramPanchayatWeatherData/status/delete", Adapt(http.HandlerFunc(route.Handler.DeleteGramPanchayatWeatherData))).Methods("DELETE")
	r.Handle("/gramPanchayatWeatherData/filter", Adapt(http.HandlerFunc(route.Handler.FilterGramPanchayatWeatherData))).Methods("POST")
	r.Handle("/gramPanchayatWeatherData/GramPanchayat", Adapt(http.HandlerFunc(route.Handler.SaveGramPanchayatWeatherDataWithOpenWebsite))).Methods("GET")
	r.Handle("/gramPanchayatWeatherData/currentdate", Adapt(http.HandlerFunc(route.Handler.GetSingleGramPanchayatWeatherDataWithCurrentDate))).Methods("GET")
}
