package routes

import (
	"net/http"

	"github.com/gorilla/mux"
)

//FarmerCropRoutes : ""
func (route *Route) FarmerCropRoutes(r *mux.Router) {
	r.Handle("/farmerCrop", Adapt(http.HandlerFunc(route.Handler.SaveFarmerCrop))).Methods("POST")
	r.Handle("/farmerCrop", Adapt(http.HandlerFunc(route.Handler.GetSingleFarmerCrop))).Methods("GET")
	r.Handle("/farmerCrop", Adapt(http.HandlerFunc(route.Handler.UpdateFarmerCrop))).Methods("PUT")
	r.Handle("/farmerCrop/status/enable", Adapt(http.HandlerFunc(route.Handler.EnableFarmerCrop))).Methods("PUT")
	r.Handle("/farmerCrop/status/disable", Adapt(http.HandlerFunc(route.Handler.DisableFarmerCrop))).Methods("PUT")
	r.Handle("/farmerCrop/status/delete", Adapt(http.HandlerFunc(route.Handler.DeleteFarmerCrop))).Methods("DELETE")
	r.Handle("/farmerCrop/filter", Adapt(http.HandlerFunc(route.Handler.FilterFarmerCrop))).Methods("POST")
	r.Handle("/farmerCropDone", Adapt(http.HandlerFunc(route.Handler.UpdateFarmerCropDone))).Methods("PUT")
}
