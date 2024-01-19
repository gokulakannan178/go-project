package routes

import (
	"net/http"

	"github.com/gorilla/mux"
)

// HolidaysRoutes
func (route *Route) HolidaysRoutes(r *mux.Router) {
	r.Handle("/holidays", Adapt(http.HandlerFunc(route.Handler.SaveHolidays))).Methods("POST")
	r.Handle("/holidays", Adapt(http.HandlerFunc(route.Handler.GetSingleHolidays))).Methods("GET")
	r.Handle("/holidays", Adapt(http.HandlerFunc(route.Handler.UpdateHolidays))).Methods("PUT")
	r.Handle("/holidays/status/enable", Adapt(http.HandlerFunc(route.Handler.EnableHolidays))).Methods("PUT")
	r.Handle("/holidays/status/disable", Adapt(http.HandlerFunc(route.Handler.DisableHolidays))).Methods("PUT")
	r.Handle("/holidays/status/delete", Adapt(http.HandlerFunc(route.Handler.DeleteHolidays))).Methods("DELETE")
	r.Handle("/holidays/filter", Adapt(http.HandlerFunc(route.Handler.FilterHolidays))).Methods("POST")
	r.Handle("/holidays/weeks", Adapt(http.HandlerFunc(route.Handler.GetHolidays))).Methods("PUT")
	r.Handle("/holidays/upload", Adapt(http.HandlerFunc(route.Handler.HolidayUpload))).Methods("PUT")

}
