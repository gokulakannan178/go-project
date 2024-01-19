package routes

import (
	"net/http"

	"github.com/gorilla/mux"
)

func (route *Route) BulkPrintRoutes(r *mux.Router) {
	// BirthCertificate
	r.Handle("/bulkprint/property/getdetails", Adapt(http.HandlerFunc(route.Handler.BulkPrintGetDetailsForProperty))).Methods("POST")
	r.Handle("/bulkprint/property/receipts", Adapt(http.HandlerFunc(route.Handler.BulkPrintGetDetailsReceiptsForProperty))).Methods("POST")
	// TradeLicence
	r.Handle("/bulkprint/tradelicense/getdetails", Adapt(http.HandlerFunc(route.Handler.BulkPrintGetDetailsForTradelicense))).Methods("POST")
	r.Handle("/bulkprint/tradelicense/receipts", Adapt(http.HandlerFunc(route.Handler.BulkPrintGetDetailsReceiptsForTradelicense))).Methods("POST")

}
