package routes

import (
	"net/http"

	"github.com/gorilla/mux"
)

func (route *Route) SolidWasteUserChargeSubCategoryRoutes(r *mux.Router) {
	// SolidWasteUserChargeSubCategory
	r.Handle("/solidwasteuserchargesubcategory", Adapt(http.HandlerFunc(route.Handler.SaveSolidWasteUserChargeSubCategory))).Methods("POST")
	r.Handle("/solidwasteuserchargesubcategory", Adapt(http.HandlerFunc(route.Handler.GetSingleSolidWasteUserChargeSubCategory))).Methods("GET")
	r.Handle("/solidwasteuserchargesubcategory", Adapt(http.HandlerFunc(route.Handler.UpdateSolidWasteUserChargeSubCategory))).Methods("PUT")
	r.Handle("/solidwasteuserchargesubcategory/status/enable", Adapt(http.HandlerFunc(route.Handler.EnableSolidWasteUserChargeSubCategory))).Methods("PUT")
	r.Handle("/solidwasteuserchargesubcategory/status/disable", Adapt(http.HandlerFunc(route.Handler.DisableSolidWasteUserChargeSubCategory))).Methods("PUT")
	r.Handle("/solidwasteuserchargesubcategory/status/delete", Adapt(http.HandlerFunc(route.Handler.DeleteSolidWasteUserChargeSubCategory))).Methods("DELETE")
	r.Handle("/solidwasteuserchargesubcategory/filter", Adapt(http.HandlerFunc(route.Handler.FilterSolidWasteUserChargeSubCategory))).Methods("POST")
}
