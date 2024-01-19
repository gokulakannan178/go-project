package routes

import (
	"net/http"

	"github.com/gorilla/mux"
)

func (route *Route) BoringChargesRoutes(r *mux.Router) {
	r.Handle("/boringcharges", Adapt(http.HandlerFunc(route.Handler.SaveBoringCharges))).Methods("POST")
	r.Handle("/boringcharges", Adapt(http.HandlerFunc(route.Handler.GetSingleBoringCharges))).Methods("GET")
	r.Handle("/boringcharges", Adapt(http.HandlerFunc(route.Handler.UpdateBoringCharges))).Methods("PUT")
	r.Handle("/boringcharges/status/enable", Adapt(http.HandlerFunc(route.Handler.EnableBoringCharges))).Methods("PUT")
	r.Handle("/boringcharges/status/disable", Adapt(http.HandlerFunc(route.Handler.DisableBoringCharges))).Methods("PUT")
	r.Handle("/boringcharges/status/delete", Adapt(http.HandlerFunc(route.Handler.DeleteBoringCharges))).Methods("DELETE")
	r.Handle("/boringcharges/filter", Adapt(http.HandlerFunc(route.Handler.FilterBoringCharges))).Methods("POST")
}
