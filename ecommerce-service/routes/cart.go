package routes

import (
	"net/http"

	"github.com/gorilla/mux"
)

func (route *Route) CartRoutes(r *mux.Router) {
	r.Handle("/cart/add", Adapt(http.HandlerFunc(route.Handler.AddCart))).Methods("POST")
	r.Handle("/cart", Adapt(http.HandlerFunc(route.Handler.GetSingleCart))).Methods("GET")
	r.Handle("/cart", Adapt(http.HandlerFunc(route.Handler.UpdateCart))).Methods("PUT")
	r.Handle("/cart/status/enable", Adapt(http.HandlerFunc(route.Handler.EnableCart))).Methods("PUT")
	r.Handle("/cart/status/disable", Adapt(http.HandlerFunc(route.Handler.DisableCart))).Methods("PUT")
	r.Handle("/cart/status/delete", Adapt(http.HandlerFunc(route.Handler.DeleteCart))).Methods("DELETE")
	r.Handle("/cart/filter", Adapt(http.HandlerFunc(route.Handler.FilterCart))).Methods("POST")
	r.Handle("/cart/updatecartitemquanity", Adapt(http.HandlerFunc(route.Handler.UpdateCartItemQuanity))).Methods("PUT")

}
