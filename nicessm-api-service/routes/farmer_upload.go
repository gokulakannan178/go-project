package routes

import (
	"net/http"

	"github.com/gorilla/mux"
)

//FarmerLandUploadRoutes : ""
func (route *Route) FarmerLandUploadRoutes(r *mux.Router) {
	r.Handle("/farmerland/upload", Adapt(http.HandlerFunc(route.Handler.FarmerLandUploadExcel))).Methods("POST")
	r.Handle("/farmeraggregation/upload", Adapt(http.HandlerFunc(route.Handler.FarmerAggregationUploadExcel))).Methods("POST")
	r.Handle("/farmeraggregation/upload/names/v2", Adapt(http.HandlerFunc(route.Handler.FarmerAggregationUploadExcelWithNamesV2))).Methods("POST")
	r.Handle("/farmeraggregation/upload/names", Adapt(http.HandlerFunc(route.Handler.FarmerAggregationUploadExcelWithNames))).Methods("POST")
	r.Handle("/farmercaste/upload", Adapt(http.HandlerFunc(route.Handler.FarmerCasteUploadExcel))).Methods("POST")
	r.Handle("/farmerSoil/upload", Adapt(http.HandlerFunc(route.Handler.FarmerSoilUploadExcel))).Methods("POST")
	r.Handle("/farmercrop/upload", Adapt(http.HandlerFunc(route.Handler.FarmerCropUploadExcel))).Methods("POST")

}
