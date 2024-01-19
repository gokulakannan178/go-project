package routes

import (
	"net/http"

	"github.com/gorilla/mux"
)

func (route *Route) DeliverSaleRoutes(r *mux.Router) {
	// Deliver Sale
	r.Handle("/order/deliver", Adapt(http.HandlerFunc(route.Handler.CreateFPOPurchaseULBSale))).Methods("PUT")
	r.Handle("/order/placeanddeliver", Adapt(http.HandlerFunc(route.Handler.PlaceAndDeliverOrder))).Methods("PUT")
}
