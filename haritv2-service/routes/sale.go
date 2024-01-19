package routes

import (
	"net/http"

	"github.com/gorilla/mux"
)

func (route *Route) SaleRoutes(r *mux.Router) {
	r.Handle("/save/sale", Adapt(http.HandlerFunc(route.Handler.SaveSale))).Methods("POST")
	r.Handle("/sale/filter", Adapt(http.HandlerFunc(route.Handler.FilterSale))).Methods("POST")

}
