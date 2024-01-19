package routes

import (
	"net/http"

	"github.com/gorilla/mux"
)

//ModuleLogic : ""
func (route *Route) ModuleLogic(r *mux.Router) {
	r.Handle("/logic", Adapt(http.HandlerFunc(route.Handler.Logic))).Methods("POST")

}
