package routes

import (
	"net/http"

	"github.com/gorilla/mux"
)

//AssetLogRoutes : ""
func (route *Route) AssetLogRoutes(r *mux.Router) {
	r.Handle("/assetlog", Adapt(http.HandlerFunc(route.Handler.SaveAssetLog))).Methods("POST")
	r.Handle("/assetlog", Adapt(http.HandlerFunc(route.Handler.GetSingleAssetLog))).Methods("GET")
	r.Handle("/assetlog", Adapt(http.HandlerFunc(route.Handler.UpdateAssetLog))).Methods("PUT")
	r.Handle("/assetlog/status/enable", Adapt(http.HandlerFunc(route.Handler.EnableAssetLog))).Methods("PUT")
	r.Handle("/assetlog/status/disable", Adapt(http.HandlerFunc(route.Handler.DisableAssetLog))).Methods("PUT")
	r.Handle("/assetlog/status/delete", Adapt(http.HandlerFunc(route.Handler.DeleteAssetLog))).Methods("DELETE")
	r.Handle("/assetlog/filter", Adapt(http.HandlerFunc(route.Handler.FilterAssetLog))).Methods("POST")

}
