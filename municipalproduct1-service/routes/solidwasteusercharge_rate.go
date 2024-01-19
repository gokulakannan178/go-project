package routes

import (
	"net/http"

	"github.com/gorilla/mux"
)

func (route *Route) SolidWasteUserChargeRateRoutes(r *mux.Router) {
	// SolidWasteUserChargeRate
	r.Handle("/solidwasteuserchargerate", Adapt(http.HandlerFunc(route.Handler.SaveSolidWasteUserChargeRate))).Methods("POST")
	r.Handle("/solidwasteuserchargerate", Adapt(http.HandlerFunc(route.Handler.GetSingleSolidWasteUserChargeRate))).Methods("GET")
	r.Handle("/solidwasteuserchargerate", Adapt(http.HandlerFunc(route.Handler.UpdateSolidWasteUserChargeRate))).Methods("PUT")
	r.Handle("/solidwasteuserchargerate/status/enable", Adapt(http.HandlerFunc(route.Handler.EnableSolidWasteUserChargeRate))).Methods("PUT")
	r.Handle("/solidwasteuserchargerate/status/disable", Adapt(http.HandlerFunc(route.Handler.DisableSolidWasteUserChargeRate))).Methods("PUT")
	r.Handle("/solidwasteuserchargerate/status/delete", Adapt(http.HandlerFunc(route.Handler.DeleteSolidWasteUserChargeRate))).Methods("DELETE")
	r.Handle("/solidwasteuserchargerate/filter", Adapt(http.HandlerFunc(route.Handler.FilterSolidWasteUserChargeRate))).Methods("POST")
}
