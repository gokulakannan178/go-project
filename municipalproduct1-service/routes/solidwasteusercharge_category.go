package routes

import (
	"net/http"

	"github.com/gorilla/mux"
)

func (route *Route) SolidWasteUserChargeCategoryRoutes(r *mux.Router) {
	// SolidWasteUserChargeCategory
	r.Handle("/solidwasteuserchargecategory", Adapt(http.HandlerFunc(route.Handler.SaveSolidWasteUserChargeCategory))).Methods("POST")
	r.Handle("/solidwasteuserchargecategory", Adapt(http.HandlerFunc(route.Handler.GetSingleSolidWasteUserChargeCategory))).Methods("GET")
	r.Handle("/solidwasteuserchargecategory", Adapt(http.HandlerFunc(route.Handler.UpdateSolidWasteUserChargeCategory))).Methods("PUT")
	r.Handle("/solidwasteuserchargecategory/status/enable", Adapt(http.HandlerFunc(route.Handler.EnableSolidWasteUserChargeCategory))).Methods("PUT")
	r.Handle("/solidwasteuserchargecategory/status/disable", Adapt(http.HandlerFunc(route.Handler.DisableSolidWasteUserChargeCategory))).Methods("PUT")
	r.Handle("/solidwasteuserchargecategory/status/delete", Adapt(http.HandlerFunc(route.Handler.DeleteSolidWasteUserChargeCategory))).Methods("DELETE")
	r.Handle("/solidwasteuserchargecategory/filter", Adapt(http.HandlerFunc(route.Handler.FilterSolidWasteUserChargeCategory))).Methods("POST")
}
