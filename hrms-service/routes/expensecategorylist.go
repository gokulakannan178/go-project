package routes

import (
	"net/http"

	"github.com/gorilla/mux"
)

//ExpenseCategoryListRoutes : ""
func (route *Route) ExpenseCategoryListRoutes(r *mux.Router) {
	r.Handle("/expensecategorylist", Adapt(http.HandlerFunc(route.Handler.SaveExpenseCategoryList))).Methods("POST")
	r.Handle("/expensecategorylist", Adapt(http.HandlerFunc(route.Handler.GetSingleExpenseCategoryList))).Methods("GET")
	r.Handle("/expensecategorylist", Adapt(http.HandlerFunc(route.Handler.UpdateExpenseCategoryList))).Methods("PUT")
	r.Handle("/expensecategorylist/status/enable", Adapt(http.HandlerFunc(route.Handler.EnableExpenseCategoryList))).Methods("PUT")
	r.Handle("/expensecategorylist/status/disable", Adapt(http.HandlerFunc(route.Handler.DisableExpenseCategoryList))).Methods("PUT")
	r.Handle("/expensecategorylist/status/delete", Adapt(http.HandlerFunc(route.Handler.DeleteExpenseCategoryList))).Methods("DELETE")
	r.Handle("/expensecategorylist/filter", Adapt(http.HandlerFunc(route.Handler.FilterExpenseCategoryList))).Methods("POST")

}
