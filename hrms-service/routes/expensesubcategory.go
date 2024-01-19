package routes

import (
	"net/http"

	"github.com/gorilla/mux"
)

//ExpenseSubcategoryRoutes : ""
func (route *Route) ExpenseSubcategoryRoutes(r *mux.Router) {
	r.Handle("/expensesubcategory", Adapt(http.HandlerFunc(route.Handler.SaveExpenseSubcategory))).Methods("POST")
	r.Handle("/expensesubcategory", Adapt(http.HandlerFunc(route.Handler.GetSingleExpenseSubcategory))).Methods("GET")
	r.Handle("/expensesubcategory", Adapt(http.HandlerFunc(route.Handler.UpdateExpenseSubcategory))).Methods("PUT")
	r.Handle("/expensesubcategory/status/enable", Adapt(http.HandlerFunc(route.Handler.EnableExpenseSubcategory))).Methods("PUT")
	r.Handle("/expensesubcategory/status/disable", Adapt(http.HandlerFunc(route.Handler.DisableExpenseSubcategory))).Methods("PUT")
	r.Handle("/expensesubcategory/status/delete", Adapt(http.HandlerFunc(route.Handler.DeleteExpenseSubcategory))).Methods("DELETE")
	r.Handle("/expensesubcategory/filter", Adapt(http.HandlerFunc(route.Handler.FilterExpenseSubcategory))).Methods("POST")

}
