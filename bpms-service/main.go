package main

import (
	"bpms-service/config"
	"bpms-service/constants"
	"bpms-service/daos"
	"bpms-service/handlers"
	"bpms-service/middlewares"
	"bpms-service/redis"
	"bpms-service/routes"
	"bpms-service/services"
	"bpms-service/shared"
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
)

func main() {
	argsWithoutProg := os.Args[1:]
	config := config.Config()
	sh := shared.NewShared(shared.SplitCmdArguments(argsWithoutProg), config)
	redisConn := redis.Connect(config)
	db := daos.GetDaos(sh, redisConn, config)
	ser := services.GetService(db, sh, redisConn, config)
	han := handlers.GetHandler(ser, sh, redisConn, config)
	route := routes.GetRoute(han, sh, redisConn, config)
	rr := mux.NewRouter()
	commonRoute := rr.PathPrefix("/api").Subrouter()
	r := commonRoute.NewRoute().Subrouter()
	r.Use(middlewares.Log)
	r.Use(middlewares.AllowCors)
	r.Methods("OPTIONS").HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE, PATCH")
		w.Header().Set("Access-Control-Allow-Headers", "Accept, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, Access-Control-Request-Headers, Access-Control-Request-Method, Connection, Host, Origin, User-Agent, Referer, Cache-Control, X-header")
		w.WriteHeader(http.StatusNoContent)
		return
	})
	//Common Routes
	route.CommonRoutes(r)
	route.FileRoutes(r)

	//Geo Location Routes
	route.StateRoutes(r)
	route.DistrictRoutes(r)
	route.VillageRoutes(r)
	route.ZoneRoutes(r)
	route.WardRoutes(r)

	//bpms routes
	route.PreregistrationRoutes(r)
	route.ApplicantRoutes(r)
	route.OTPRoutes(r)
	route.ULBRoutes(r)

	//Master Routes
	route.ApplicantTypeRoutes(r)
	route.EducationTypeRoutes(r)
	route.RoadTypeRoutes(r)
	route.OccupancyTypeRoutes(r)
	route.RoofTypeRoutes(r)
	route.AmenitiesRoutes(r)

	//ACL Routes
	route.ModuleRoutes(r)
	route.MenuRoutes(r)
	route.TabRoutes(r)
	route.FeatureRoutes(r)
	route.ACLMasterUserRoutes(r)
	route.ACLAccess(r)

	//Department Routes
	route.DepartmentRoutes(r)
	route.DeptChecklistRoutes(r)
	route.DepartmentTypeRoutes(r)

	//User Routes
	route.UserRoutes(r)
	route.UserTypeRoutes(r)
	route.UserAuthRoutes(r)

	//Plan Routes
	route.PlanRoutes(r)
	route.PlanRegistrationTypeRoutes(r)
	route.PlanDepartmentApprovalRoutes(r)
	route.PlanDepartmentFlowRoutes(r)
	route.PlanReqDocumentRoutes(r)
	route.PlanDocumentRoutes(r)

	//Plan CRF Routed
	route.PlanCRFRoutes(r)

	log.Println("listening in port : ", sh.GetCmdArg(constants.PORT))
	log.Fatal(http.ListenAndServe("0.0.0.0:"+sh.GetCmdArg(constants.PORT), r))
}
