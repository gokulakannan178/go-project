package routes

import (
	"net/http"

	"github.com/gorilla/mux"
)

//FarmerCropCalendarRoutes : ""
func (route *Route) FarmerCropCalendarRoutes(r *mux.Router) {
	r.Handle("/FarmerCropCalendar", Adapt(http.HandlerFunc(route.Handler.SaveFarmerCropCalendar))).Methods("POST")
	r.Handle("/FarmerCropCalendar", Adapt(http.HandlerFunc(route.Handler.GetSingleFarmerCropCalendar))).Methods("GET")
	r.Handle("/FarmerCropCalendar", Adapt(http.HandlerFunc(route.Handler.UpdateFarmerCropCalendar))).Methods("PUT")
	r.Handle("/FarmerCropCalendar/status/enable", Adapt(http.HandlerFunc(route.Handler.EnableFarmerCropCalendar))).Methods("PUT")
	r.Handle("/FarmerCropCalendar/status/disable", Adapt(http.HandlerFunc(route.Handler.DisableFarmerCropCalendar))).Methods("PUT")
	r.Handle("/FarmerCropCalendar/status/delete", Adapt(http.HandlerFunc(route.Handler.DeleteFarmerCropCalendar))).Methods("DELETE")
	r.Handle("/FarmerCropCalendar/filter", Adapt(http.HandlerFunc(route.Handler.FilterFarmerCropCalendar))).Methods("POST")

}
