package routes

import (
	"net/http"

	"github.com/gorilla/mux"
)

//LandCropRoutes : ""
func (route *Route) LandCropRoutes(r *mux.Router) {
	r.Handle("/landCrop", Adapt(http.HandlerFunc(route.Handler.SaveLandCrop))).Methods("POST")
	r.Handle("/landCrop", Adapt(http.HandlerFunc(route.Handler.GetSingleLandCrop))).Methods("GET")
	r.Handle("/landCrop", Adapt(http.HandlerFunc(route.Handler.UpdateLandCrop))).Methods("PUT")
	r.Handle("/landCrop/status/enable", Adapt(http.HandlerFunc(route.Handler.EnableLandCrop))).Methods("PUT")
	r.Handle("/landCrop/status/disable", Adapt(http.HandlerFunc(route.Handler.DisableLandCrop))).Methods("PUT")
	r.Handle("/landCrop/status/delete", Adapt(http.HandlerFunc(route.Handler.DeleteLandCrop))).Methods("DELETE")
	r.Handle("/landCrop/filter", Adapt(http.HandlerFunc(route.Handler.FilterLandCrop))).Methods("POST")
}
