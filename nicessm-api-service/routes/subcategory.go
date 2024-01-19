package routes

import (
	"net/http"

	"github.com/gorilla/mux"
)

func (route *Route) SubCategoryRoutes(r *mux.Router) {
	r.Handle("/subCategory", Adapt(http.HandlerFunc(route.Handler.SaveSubCategory))).Methods("POST")
	r.Handle("/subCategory", Adapt(http.HandlerFunc(route.Handler.GetSingleSubCategory))).Methods("GET")
	r.Handle("/subCategory", Adapt(http.HandlerFunc(route.Handler.UpdateSubCategory))).Methods("PUT")
	r.Handle("/subCategory/status/enable", Adapt(http.HandlerFunc(route.Handler.EnableSubCategory))).Methods("PUT")
	r.Handle("/subCategory/status/disable", Adapt(http.HandlerFunc(route.Handler.DisableSubCategory))).Methods("PUT")
	r.Handle("/subCategory/status/delete", Adapt(http.HandlerFunc(route.Handler.DeleteSubCategory))).Methods("DELETE")
	r.Handle("/subCategory/filter", Adapt(http.HandlerFunc(route.Handler.FilterSubCategory))).Methods("POST")
}
