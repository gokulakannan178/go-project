package routes

import (
	"net/http"

	"github.com/gorilla/mux"
)

//PropertyRoutes : ""
func (route *Route) LegacyPropertyRoutes(r *mux.Router) {
	r.Handle("/legacy", Adapt(http.HandlerFunc(route.Handler.SaveLegacy))).Methods("POST")
	r.Handle("/legacy", Adapt(http.HandlerFunc(route.Handler.GetLegacyForAProperty))).Methods("GET")
	r.Handle("/legacy", Adapt(http.HandlerFunc(route.Handler.UpdateLegacyForAProperty))).Methods("PUT")
	r.Handle("/legacy/financialyears", Adapt(http.HandlerFunc(route.Handler.GetFinancialYearsForLegacyPayments))).Methods("GET")
	r.Handle("/legacy/financialyears/newsaf", Adapt(http.HandlerFunc(route.Handler.GetReqFinancialYearForLegacy))).Methods("POST")
	r.Handle("/legacy/v2", Adapt(http.HandlerFunc(route.Handler.GetLegacyForAPropertyV2))).Methods("GET")
	r.Handle("/legacy/filter", Adapt(http.HandlerFunc(route.Handler.FilterLegacy))).Methods("POST")

}
