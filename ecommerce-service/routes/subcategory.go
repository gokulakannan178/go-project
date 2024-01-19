package routes

import (
	"net/http"

	"github.com/gorilla/mux"
)

func (route *Route) SubCategoryRoutes(r *mux.Router) {
	// SubCategory
	r.Handle("/subcategory", Adapt(http.HandlerFunc(route.Handler.SaveSubCategory))).Methods("POST")
	r.Handle("/subcategory", Adapt(http.HandlerFunc(route.Handler.GetSingleSubCategory))).Methods("GET")
	r.Handle("/subcategory", Adapt(http.HandlerFunc(route.Handler.UpdateSubCategory))).Methods("PUT")
	r.Handle("/subcategory/status/enable", Adapt(http.HandlerFunc(route.Handler.EnableSubCategory))).Methods("PUT")
	r.Handle("/subcategory/status/disable", Adapt(http.HandlerFunc(route.Handler.DisableSubCategory))).Methods("PUT")
	r.Handle("/subcategory/status/delete", Adapt(http.HandlerFunc(route.Handler.DeleteSubCategory))).Methods("DELETE")
	r.Handle("/subcategory/filter", Adapt(http.HandlerFunc(route.Handler.FilterSubCategory))).Methods("POST")
}
