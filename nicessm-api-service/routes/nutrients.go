package routes

import (
	"net/http"

	"github.com/gorilla/mux"
)

func (route *Route) NutrientsRoutes(r *mux.Router) {
	r.Handle("/nutrients", Adapt(http.HandlerFunc(route.Handler.SaveNutrients))).Methods("POST")
	r.Handle("/nutrients", Adapt(http.HandlerFunc(route.Handler.GetSingleNutrients))).Methods("GET")
	r.Handle("/nutrients", Adapt(http.HandlerFunc(route.Handler.UpdateNutrients))).Methods("PUT")
	r.Handle("/nutrients/status/enable", Adapt(http.HandlerFunc(route.Handler.EnableNutrients))).Methods("PUT")
	r.Handle("/nutrients/status/disable", Adapt(http.HandlerFunc(route.Handler.DisableNutrients))).Methods("PUT")
	r.Handle("/nutrients/status/delete", Adapt(http.HandlerFunc(route.Handler.DeleteNutrients))).Methods("DELETE")
	r.Handle("/nutrients/filter", Adapt(http.HandlerFunc(route.Handler.FilterNutrients))).Methods("POST")
}
