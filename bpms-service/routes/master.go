package routes

import (
	"net/http"

	"github.com/gorilla/mux"
)

//RoadTypeRoutes : ""
func (route *Route) RoadTypeRoutes(r *mux.Router) {
	r.Handle("/roadtype", Adapt(http.HandlerFunc(route.Handler.SaveRoadType))).Methods("POST")
	r.Handle("/roadtype", Adapt(http.HandlerFunc(route.Handler.GetSingleRoadType))).Methods("GET")
	r.Handle("/roadtype", Adapt(http.HandlerFunc(route.Handler.UpdateRoadType))).Methods("PUT")
	r.Handle("/roadtype/status/enable", Adapt(http.HandlerFunc(route.Handler.EnableRoadType))).Methods("PUT")
	r.Handle("/roadtype/status/disable", Adapt(http.HandlerFunc(route.Handler.DisableRoadType))).Methods("PUT")
	r.Handle("/roadtype/status/delete", Adapt(http.HandlerFunc(route.Handler.DeleteRoadType))).Methods("DELETE")
	r.Handle("/roadtype/filter", Adapt(http.HandlerFunc(route.Handler.FilterRoadType))).Methods("POST")
}

//OccupancyTypeRoutes : ""
func (route *Route) OccupancyTypeRoutes(r *mux.Router) {
	r.Handle("/occupancytype", Adapt(http.HandlerFunc(route.Handler.SaveOccupancyType))).Methods("POST")
	r.Handle("/occupancytype", Adapt(http.HandlerFunc(route.Handler.GetSingleOccupancyType))).Methods("GET")
	r.Handle("/occupancytype", Adapt(http.HandlerFunc(route.Handler.UpdateOccupancyType))).Methods("PUT")
	r.Handle("/occupancytype/status/enable", Adapt(http.HandlerFunc(route.Handler.EnableOccupancyType))).Methods("PUT")
	r.Handle("/occupancytype/status/disable", Adapt(http.HandlerFunc(route.Handler.DisableOccupancyType))).Methods("PUT")
	r.Handle("/occupancytype/status/delete", Adapt(http.HandlerFunc(route.Handler.DeleteOccupancyType))).Methods("DELETE")
	r.Handle("/occupancytype/filter", Adapt(http.HandlerFunc(route.Handler.FilterOccupancyType))).Methods("POST")
}

//RoofTypeRoutes : ""
func (route *Route) RoofTypeRoutes(r *mux.Router) {
	r.Handle("/rooftype", Adapt(http.HandlerFunc(route.Handler.SaveRoofType))).Methods("POST")
	r.Handle("/rooftype", Adapt(http.HandlerFunc(route.Handler.GetSingleRoofType))).Methods("GET")
	r.Handle("/rooftype", Adapt(http.HandlerFunc(route.Handler.UpdateRoofType))).Methods("PUT")
	r.Handle("/rooftype/status/enable", Adapt(http.HandlerFunc(route.Handler.EnableRoofType))).Methods("PUT")
	r.Handle("/rooftype/status/disable", Adapt(http.HandlerFunc(route.Handler.DisableRoofType))).Methods("PUT")
	r.Handle("/rooftype/status/delete", Adapt(http.HandlerFunc(route.Handler.DeleteRoofType))).Methods("DELETE")
	r.Handle("/rooftype/filter", Adapt(http.HandlerFunc(route.Handler.FilterRoofType))).Methods("POST")
}

//AmenitiesRoutes : ""
func (route *Route) AmenitiesRoutes(r *mux.Router) {
	r.Handle("/amenities", Adapt(http.HandlerFunc(route.Handler.SaveAmenities))).Methods("POST")
	r.Handle("/amenities", Adapt(http.HandlerFunc(route.Handler.GetSingleAmenities))).Methods("GET")
	r.Handle("/amenities", Adapt(http.HandlerFunc(route.Handler.UpdateAmenities))).Methods("PUT")
	r.Handle("/amenities/status/enable", Adapt(http.HandlerFunc(route.Handler.EnableAmenities))).Methods("PUT")
	r.Handle("/amenities/status/disable", Adapt(http.HandlerFunc(route.Handler.DisableAmenities))).Methods("PUT")
	r.Handle("/amenities/status/delete", Adapt(http.HandlerFunc(route.Handler.DeleteAmenities))).Methods("DELETE")
	r.Handle("/amenities/filter", Adapt(http.HandlerFunc(route.Handler.FilterAmenities))).Methods("POST")
}
