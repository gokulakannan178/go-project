package routes

import (
	"net/http"

	"github.com/gorilla/mux"
)

//ProductConfigurationRoutes : ""
func (route *Route) ProductConfiguration(r *mux.Router) {
	r.Handle("/product/configuration", Adapt(http.HandlerFunc(route.Handler.SaveProductConfiguration))).Methods("POST")
	r.Handle("/product/configuration", Adapt(http.HandlerFunc(route.Handler.GetSingleProductConfiguration))).Methods("GET")
	r.Handle("/product/logo/b64", Adapt(http.HandlerFunc(route.Handler.GetProductLogo))).Methods("GET")
	r.Handle("/product/watermarklogo/b64", Adapt(http.HandlerFunc(route.Handler.GetWatermarkLogo))).Methods("GET")
	r.Handle("/product/configuration/filter", Adapt(http.HandlerFunc(route.Handler.FilterProductConfiguration))).Methods("POST")

}
