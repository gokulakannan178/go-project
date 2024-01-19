package routes

import (
	"net/http"

	"github.com/gorilla/mux"
)

func (route *Route) SolidWasteUserChargeRoutes(r *mux.Router) {
	// SolidWasteUserCharge
	r.Handle("/solidwasteusercharge", Adapt(http.HandlerFunc(route.Handler.SaveSolidWasteUserCharge))).Methods("POST")
	r.Handle("/solidwasteusercharge", Adapt(http.HandlerFunc(route.Handler.GetSingleSolidWasteUserCharge))).Methods("GET")
	r.Handle("/solidwasteusercharge", Adapt(http.HandlerFunc(route.Handler.UpdateSolidWasteUserCharge))).Methods("PUT")
	r.Handle("/solidwasteusercharge/status/enable", Adapt(http.HandlerFunc(route.Handler.EnableSolidWasteUserCharge))).Methods("PUT")
	r.Handle("/solidwasteusercharge/status/disable", Adapt(http.HandlerFunc(route.Handler.DisableSolidWasteUserCharge))).Methods("PUT")
	r.Handle("/solidwasteusercharge/status/delete", Adapt(http.HandlerFunc(route.Handler.DeleteSolidWasteUserCharge))).Methods("DELETE")
	r.Handle("/solidwasteusercharge/filter", Adapt(http.HandlerFunc(route.Handler.FilterSolidWasteUserCharge))).Methods("POST")
}
