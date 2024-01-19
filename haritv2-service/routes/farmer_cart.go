package routes

import (
	"net/http"

	"github.com/gorilla/mux"
)

func (route *Route) FarmerCartRoutes(r *mux.Router) {
	// FarmerCart
	r.Handle("/farmercart", Adapt(http.HandlerFunc(route.Handler.SaveFarmerCart))).Methods("POST")
	r.Handle("/farmercart", Adapt(http.HandlerFunc(route.Handler.GetSingleFarmerCart))).Methods("GET")
	r.Handle("/farmercart", Adapt(http.HandlerFunc(route.Handler.UpdateFarmerCart))).Methods("PUT")
	r.Handle("/farmercart/status/enable", Adapt(http.HandlerFunc(route.Handler.EnableFarmerCart))).Methods("PUT")
	r.Handle("/farmercart/status/disable", Adapt(http.HandlerFunc(route.Handler.DisableFarmerCart))).Methods("PUT")
	r.Handle("/farmercart/status/delete", Adapt(http.HandlerFunc(route.Handler.DeleteFarmerCart))).Methods("DELETE")
	r.Handle("/farmercart/filter", Adapt(http.HandlerFunc(route.Handler.FilterFarmerCart))).Methods("POST")

}
