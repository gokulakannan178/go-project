package routes

import (
	"net/http"

	"github.com/gorilla/mux"
)

//ContentDisseminationRoutes : ""
func (route *Route) ContentDisseminationRoutes(r *mux.Router) {
	r.Handle("/contentdissemination", Adapt(http.HandlerFunc(route.Handler.GetContentDisseminationUserAndFarmer))).Methods("GET")
	r.Handle("/contentdissemination/farmeruser/count", Adapt(http.HandlerFunc(route.Handler.GetContentDisseminationUserAndFarmerCount))).Methods("POST")

}
