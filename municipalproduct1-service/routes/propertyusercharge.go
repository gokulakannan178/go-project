package routes

import (
	"net/http"

	"github.com/gorilla/mux"
)

// PropertyUserChargeRoutes : ""
func (route *Route) PropertyUserChargeRoutes(r *mux.Router) {
	r.Handle("/property/usercharge", Adapt(http.HandlerFunc(route.Handler.SavePropertyUserCharge))).Methods("POST")
	r.Handle("/propertyusercharge", Adapt(http.HandlerFunc(route.Handler.GetSinglePropertyUserCharge))).Methods("GET")
	r.Handle("/property/usercharge/update", Adapt(http.HandlerFunc(route.Handler.UpdatePropertyUserCharge))).Methods("PUT")
	r.Handle("/property/usercharge/verify", Adapt(http.HandlerFunc(route.Handler.VerifyPropertyUserCharge))).Methods("PUT")
	r.Handle("/property/usercharge/reject", Adapt(http.HandlerFunc(route.Handler.RejectPropertyUserCharge))).Methods("PUT")
	r.Handle("/property/usercharge/status/delete", Adapt(http.HandlerFunc(route.Handler.DeletePropertyUserCharge))).Methods("PUT")
	r.Handle("/propertyusercharge/filter", Adapt(http.HandlerFunc(route.Handler.FilterPropertyUserCharge))).Methods("POST")
}
