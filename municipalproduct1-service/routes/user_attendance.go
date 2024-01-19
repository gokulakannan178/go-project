package routes

import (
	"net/http"

	"github.com/gorilla/mux"
)

//ShopRentRoutes : ""
func (route *Route) UserAttendanceRoutes(r *mux.Router) {
	r.Handle("attendance/punchin", Adapt(http.HandlerFunc(route.Handler.SavePunchIn))).Methods("POST")
	r.Handle("attendance/punchout", Adapt(http.HandlerFunc(route.Handler.SavePunchOut))).Methods("PUT")

}
