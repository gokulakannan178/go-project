package routes

import (
	"net/http"

	"github.com/gorilla/mux"
)

func (route *Route) FPORoutes(r *mux.Router) {
	// FPO
	r.Handle("/fpo", Adapt(http.HandlerFunc(route.Handler.SaveFPO))).Methods("POST")
	r.Handle("/fporegistration", Adapt(http.HandlerFunc(route.Handler.SaveFPORegistration))).Methods("POST")
	r.Handle("/fpo", Adapt(http.HandlerFunc(route.Handler.GetSingleFPO))).Methods("GET")
	r.Handle("/fpo", Adapt(http.HandlerFunc(route.Handler.UpdateFPO))).Methods("PUT")
	r.Handle("/fpo/status/enable", Adapt(http.HandlerFunc(route.Handler.EnableFPO))).Methods("PUT")
	r.Handle("/fpo/status/disable", Adapt(http.HandlerFunc(route.Handler.DisableFPO))).Methods("PUT")
	r.Handle("/fpo/status/delete", Adapt(http.HandlerFunc(route.Handler.DeleteFPO))).Methods("DELETE")
	r.Handle("/fpo/filter", Adapt(http.HandlerFunc(route.Handler.FilterFPO))).Methods("POST")
	r.Handle("/fpo/updatelocation", Adapt(http.HandlerFunc(route.Handler.FPOUpdateLocation))).Methods("PUT")
	r.Handle("/fpoupdation", Adapt(http.HandlerFunc(route.Handler.UpdateFPORegistration))).Methods("PUT")
	r.Handle("/fpo/masterreport", Adapt(http.HandlerFunc(route.Handler.FPOMasterReport))).Methods("POST")
	r.Handle("/fpo/monthReport", Adapt(http.HandlerFunc(route.Handler.FPOMonthReport))).Methods("POST")
	r.Handle("/fbo/nearby", Adapt(http.HandlerFunc(route.Handler.FBONearBy))).Methods("POST")
	r.Handle("/ulb/nearby/district", Adapt(http.HandlerFunc(route.Handler.ULBNearBy))).Methods("POST")
}

func (route *Route) FPOInventoryRoutes(r *mux.Router) {
	// FPO
	r.Handle("/fpo/inventory", Adapt(http.HandlerFunc(route.Handler.SaveFPOInventory))).Methods("POST")
	r.Handle("/fpo/inventory", Adapt(http.HandlerFunc(route.Handler.GetSingleFPOInventory))).Methods("GET")
	r.Handle("/fpo/inventory", Adapt(http.HandlerFunc(route.Handler.UpdateFPOInventory))).Methods("PUT")
	r.Handle("/fpo/inventory/status/enable", Adapt(http.HandlerFunc(route.Handler.EnableFPOInventory))).Methods("PUT")
	r.Handle("/fpo/inventory/status/disable", Adapt(http.HandlerFunc(route.Handler.DisableFPOInventory))).Methods("PUT")
	r.Handle("/fpo/inventory/status/delete", Adapt(http.HandlerFunc(route.Handler.DeleteFPOInventory))).Methods("DELETE")
	r.Handle("/fpo/inventory/quantityupdate", Adapt(http.HandlerFunc(route.Handler.FPOInventoryQuantityUpdate))).Methods("PUT")
	r.Handle("/fpo/inventory/priceupdate", Adapt(http.HandlerFunc(route.Handler.FPOInventoryPriceUpdate))).Methods("PUT")
	//r.Handle("/fpo/inventory/filter", Adapt(http.HandlerFunc(route.Handler.FilterFPO))).Methods("POST")
	//r.Handle("/fpo/inventory/updatelocation", Adapt(http.HandlerFunc(route.Handler.FPOUpdateLocation))).Methods("PUT")
}
