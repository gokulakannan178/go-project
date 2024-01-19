package routes

import (
	"net/http"

	"github.com/gorilla/mux"
)

//LandCropCalendarRoutes : ""
func (route *Route) LandCropCalendarRoutes(r *mux.Router) {
	r.Handle("/landCropCalendar", Adapt(http.HandlerFunc(route.Handler.SaveLandCropCalendar))).Methods("POST")
	r.Handle("/landCropCalendar", Adapt(http.HandlerFunc(route.Handler.GetSingleLandCropCalendar))).Methods("GET")
	r.Handle("/landCropCalendar", Adapt(http.HandlerFunc(route.Handler.UpdateLandCropCalendar))).Methods("PUT")
	r.Handle("/landCropCalendar/status/enable", Adapt(http.HandlerFunc(route.Handler.EnableLandCropCalendar))).Methods("PUT")
	r.Handle("/landCropCalendar/status/disable", Adapt(http.HandlerFunc(route.Handler.DisableLandCropCalendar))).Methods("PUT")
	r.Handle("/landCropCalendar/status/delete", Adapt(http.HandlerFunc(route.Handler.DeleteLandCropCalendar))).Methods("DELETE")
	r.Handle("/landCropCalendar/filter", Adapt(http.HandlerFunc(route.Handler.FilterLandCropCalendar))).Methods("POST")
}
