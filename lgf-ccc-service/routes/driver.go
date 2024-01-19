package routes

import (
	"net/http"

	"github.com/gorilla/mux"
)

// DriverRoutes
func (route *Route) DriverRoutes(r *mux.Router) {
	r.Handle("/driver", Adapt(http.HandlerFunc(route.Handler.SaveDriverDetails))).Methods("POST")
	r.Handle("/driver", Adapt(http.HandlerFunc(route.Handler.GetSingleDriverDetails))).Methods("GET")
	r.Handle("/driver", Adapt(http.HandlerFunc(route.Handler.UpdateDriverDetails))).Methods("PUT")
	r.Handle("/driver/status/enable", Adapt(http.HandlerFunc(route.Handler.EnableDriverDetails))).Methods("PUT")
	r.Handle("/driver/status/disable", Adapt(http.HandlerFunc(route.Handler.DisableDriverDetails))).Methods("PUT")
	r.Handle("/driver/status/delete", Adapt(http.HandlerFunc(route.Handler.DeleteDriverDetails))).Methods("DELETE")
	r.Handle("/driver/filter", Adapt(http.HandlerFunc(route.Handler.FilterDriverDetails))).Methods("POST")
	//r.Handle("/Driver/assign", Adapt(http.HandlerFunc(route.Handler.DriverAssign))).Methods("POST")
	//r.Handle("/Driver/revoke", Adapt(http.HandlerFunc(route.Handler.RevokeDriver))).Methods("PUT")

}
