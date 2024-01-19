package routes

import (
	"net/http"

	"github.com/gorilla/mux"
)

func (route *Route) DistrictweatheralertdissiminationRoutes(r *mux.Router) {
	r.Handle("/districtweatheralertdissimination", Adapt(http.HandlerFunc(route.Handler.SaveDistrictweatheralertdissimination))).Methods("POST")
	r.Handle("/districtweatheralertdissimination", Adapt(http.HandlerFunc(route.Handler.GetSingleDistrictweatheralertdissimination))).Methods("GET")
	r.Handle("/districtweatheralertdissimination", Adapt(http.HandlerFunc(route.Handler.UpdateDistrictweatheralertdissimination))).Methods("PUT")
	r.Handle("/districtweatheralertdissimination/status/enable", Adapt(http.HandlerFunc(route.Handler.EnableDistrictweatheralertdissimination))).Methods("PUT")
	r.Handle("/districtweatheralertdissimination/status/disable", Adapt(http.HandlerFunc(route.Handler.DisableDistrictweatheralertdissimination))).Methods("PUT")
	r.Handle("/districtweatheralertdissimination/status/delete", Adapt(http.HandlerFunc(route.Handler.DeleteDistrictweatheralertdissimination))).Methods("DELETE")
	r.Handle("/districtweatheralertdissimination/filter", Adapt(http.HandlerFunc(route.Handler.FilterDistrictweatheralertdissimination))).Methods("POST")
	r.Handle("/districtweatheralertdissimination/dissemination/sendnow", Adapt(http.HandlerFunc(route.Handler.SaveDistrictweatheralertdissiminationSendNow))).Methods("POST")
	r.Handle("/districtweatheralertdissimination/getfarmerusercount", Adapt(http.HandlerFunc(route.Handler.GetDistrictWeatherAlertFarmerUserCount))).Methods("GET")
	r.Handle("/districtweatheralertdissimination/report", Adapt(http.HandlerFunc(route.Handler.FilterDistrictweatheralertdissiminationReport))).Methods("POST")

}
