package routes

import (
	"net/http"

	"github.com/gorilla/mux"
)

func (route *Route) WeatherAlertDissiminationRoutes(r *mux.Router) {
	r.Handle("/weatheralertdissimination", Adapt(http.HandlerFunc(route.Handler.SaveStateWeatherAlertDissimination))).Methods("POST")
	r.Handle("/weatheralertdissimination", Adapt(http.HandlerFunc(route.Handler.GetSingleStateWeatherAlertDissimination))).Methods("GET")
	r.Handle("/weatheralertdissimination", Adapt(http.HandlerFunc(route.Handler.UpdateStateWeatherAlertDissimination))).Methods("PUT")
	r.Handle("/weatheralertdissimination/status/enable", Adapt(http.HandlerFunc(route.Handler.EnableStateWeatherAlertDissimination))).Methods("PUT")
	r.Handle("/weatheralertdissimination/status/disable", Adapt(http.HandlerFunc(route.Handler.DisableStateWeatherAlertDissimination))).Methods("PUT")
	r.Handle("/weatheralertdissimination/status/delete", Adapt(http.HandlerFunc(route.Handler.DeleteStateWeatherAlertDissimination))).Methods("DELETE")
	r.Handle("/weatheralertdissimination/filter", Adapt(http.HandlerFunc(route.Handler.FilterStateWeatherAlertDissimination))).Methods("POST")
	r.Handle("/weatheralert/dissemination/sendnow", Adapt(http.HandlerFunc(route.Handler.SaveStateWeatherAlertDissiminationSendNow))).Methods("POST")
	r.Handle("/weatheralert/getfarmerusercount", Adapt(http.HandlerFunc(route.Handler.GetWeatherAlertFarmerUserCount))).Methods("GET")
	r.Handle("/weatheralertdissimination/report", Adapt(http.HandlerFunc(route.Handler.FilterStateWeatherAlertDissiminationReport))).Methods("POST")

}
