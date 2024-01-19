package routes

import (
	"net/http"

	"github.com/gorilla/mux"
)

//NutrientValueRoutes : ""
func (route *Route) NutrientValueRoutes(r *mux.Router) {
	r.Handle("/nutrientValue", Adapt(http.HandlerFunc(route.Handler.SaveNutrientValue))).Methods("POST")
	r.Handle("/nutrientValue", Adapt(http.HandlerFunc(route.Handler.GetSingleNutrientValue))).Methods("GET")
	r.Handle("/nutrientValue", Adapt(http.HandlerFunc(route.Handler.UpdateNutrientValue))).Methods("PUT")
	r.Handle("/nutrientValue/status/enable", Adapt(http.HandlerFunc(route.Handler.EnableNutrientValue))).Methods("PUT")
	r.Handle("/nutrientValue/status/disable", Adapt(http.HandlerFunc(route.Handler.DisableNutrientValue))).Methods("PUT")
	r.Handle("/nutrientValue/status/delete", Adapt(http.HandlerFunc(route.Handler.DeleteNutrientValue))).Methods("DELETE")
	r.Handle("/nutrientValue/filter", Adapt(http.HandlerFunc(route.Handler.FilterNutrientValue))).Methods("POST")
}
