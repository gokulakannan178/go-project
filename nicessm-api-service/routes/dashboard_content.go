package routes

import (
	"net/http"

	"github.com/gorilla/mux"
)

func (route *Route) DashboardContentRoutes(r *mux.Router) {
	r.Handle("/dashboard/content/sms/count", Adapt(http.HandlerFunc(route.Handler.DashboardContentSmsCount))).Methods("POST")
	r.Handle("/dashboard/content/voice/count", Adapt(http.HandlerFunc(route.Handler.DashboardContentVoiceCount))).Methods("POST")
	r.Handle("/dashboard/content/video/count", Adapt(http.HandlerFunc(route.Handler.DashboardContentVideoCount))).Methods("POST")
	r.Handle("/dashboard/content/poster/count", Adapt(http.HandlerFunc(route.Handler.DashboardContentPosterCount))).Methods("POST")
	r.Handle("/dashboard/content/docment/count", Adapt(http.HandlerFunc(route.Handler.DashboardContentDocmentCount))).Methods("POST")
	r.Handle("/content/report/daywisedemand", Adapt(http.HandlerFunc(route.Handler.DayWiseContentDemandChart))).Methods("POST")

}
