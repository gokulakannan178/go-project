package routes

import (
	"net/http"

	"github.com/gorilla/mux"
)

func (route *Route) ProductVariantMeshRoutes(r *mux.Router) {
	// ProductVariantMesh
	r.Handle("/productvariantmesh", Adapt(http.HandlerFunc(route.Handler.SaveProductVariantMesh))).Methods("POST")
	r.Handle("/productvariantmesh", Adapt(http.HandlerFunc(route.Handler.GetSingleProductVariantMesh))).Methods("GET")
	r.Handle("/productvariantmesh", Adapt(http.HandlerFunc(route.Handler.UpdateProductVariantMesh))).Methods("PUT")
	r.Handle("/productvariantmesh/status/enable", Adapt(http.HandlerFunc(route.Handler.EnableProductVariantMesh))).Methods("PUT")
	r.Handle("/productvariantmesh/status/disable", Adapt(http.HandlerFunc(route.Handler.DisableProductVariantMesh))).Methods("PUT")
	r.Handle("/productvariantmesh/status/delete", Adapt(http.HandlerFunc(route.Handler.DeleteProductVariantMesh))).Methods("DELETE")
	r.Handle("/productvariantmesh/filter", Adapt(http.HandlerFunc(route.Handler.FilterProductVariantMesh))).Methods("POST")
}
