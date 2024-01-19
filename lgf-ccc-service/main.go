package main

import (
	"lgf-ccc-service/config"
	"lgf-ccc-service/constants"
	"lgf-ccc-service/daos"
	"lgf-ccc-service/handlers"
	"lgf-ccc-service/middlewares"
	"lgf-ccc-service/redis"
	"lgf-ccc-service/routes"
	service "lgf-ccc-service/services"
	"lgf-ccc-service/shared"
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
	ser := service.GetService(db, sh, redisConn, config)
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
	route.CircleRoutes(r)
	route.SectorRoutes(r)

	//User Routes
	route.OrganisationRoutes(r)
	route.UserRoutes(r)
	route.UserAuthRoutes(r)
	route.UserTypeRoutes(r)
	route.UserLocationRoutes(r)
	route.ConsumerRoutes(r)
	route.DriverRoutes(r)

	//ACL Routes
	route.ModuleRoutes(r)
	route.MenuRoutes(r)
	route.TabRoutes(r)
	route.FeatureRoutes(r)
	route.ACLMasterUserRoutes(r)
	route.ACLAccess(r)

	//Task Routes
	route.TaskRoutes(r)
	//PropertiesRoutes
	route.PropertiesRoutes(r)
	// DepartmentRoutes
	route.DepartmentRoutes(r)

	//BranchRoutes
	route.BranchRoutes(r)

	//BeatRoutes
	route.BeatRoutes(r)

	//ChecklistRoutes
	route.ChecklistRoutes(r)

	//CitizenGrevience Routes
	route.CitizenGrevienceRoutes(r)

	//RouteMaster Routes
	route.RouteMasterRoutes(r)

	//BeatMaster Routes
	route.BeatMasterRoutes(r)

	//ChecklistMaster Routes
	route.ChecklistMasterRoutes(r)

	//WorkScheduleRoutes
	route.WorkScheduleRoutes(r)

	//DesignationRoutes
	route.DesignationRoutes(r)

	//DashboardRoutes
	route.DashboardRoutes(r)

	//HouseVisitedRoutes
	route.HouseVisitedRoutes(r)
	//WasteCollectedRoutes
	route.WasteCollectedRoutes(r)
	//MySurveyRoutes
	route.MySurveyRoutes(r)
	//ServiceRequest
	route.ServiceRequestRoutes(r)

	//DumpHistoryRoutes
	route.DumpSiteRoutes(r)
	route.DumpHistoryRoutes(r)

	//VechileRoutes
	route.VechileRoutes(r)
	route.VehicleTypeRoutes(r)
	route.VehicleLogRoutes(r)
	route.FuelHistoryRoutes(r)
	route.VehicleInsuranceRoutes(r)

	//NoticePolicyRoutes
	route.NoticePolicyRoutes(r)

	//SmsLogRoutes
	route.SmsLogRoutes(r)

	//EmailLogRoutes
	route.EmailLogRoutes(r)
	//RoadTypeRoutes
	route.RoadTypeRoutes(r)
	//ApptokenRoutes
	route.ApptokenRoutes(r)

	//ser.DaywiseAttendance()

	//EmployeeDocumentsRoutes
	//route.EmployeeDocumentsRoutes(r)

	//AssetRoutes
	//route.AssetRoutes(r)

	//AssetTypeRoutes

	//UserLocationTrackerRoutes
	route.UserLocationTrackerRoutes(r)

	//EmployeeExperienceRoutes

	//NotificationLogRoutes
	route.NotificationLogRoutes(r)

	//PropertyTypeRoutes
	route.PropertyTypeRoutes(r)
	route.IdentityTypeRoutes(r)
	route.EmployeeShiftRoutes(r)

	// //EmployeeDeductionMasterRoutes
	// route.EmployeeDeductionMasterRoutes(r)

	// //EmployeeEarningRoutes
	// route.EmployeeEarningRoutes(r)

	// //EmployeeDeductionRoutes
	// route.EmployeeDeductionRoutes(r)

	//DemoUserRoutes
	route.DemoUserRoutes(r)
	route.ReportRoutes(r)
	route.VehicleLocationRoutes(r)
	route.WardWiseDumpHistoryRoutes(r)
	route.CircleWiseDumpHistoryRoutes(r)
	route.CircleWiseHouseVisitedRoutes(r)
	route.WardWiseHouseVisitedRoutes(r)
	route.VehicleLocationRoutes(r)
	route.VehicleTripRoutes(r)

	//attendance
	route.AttendanceRoutes(r)
	route.AttendanceLogRoutes(r)

	//Beat Routes
	route.HelperBeatRoutes(r)
	route.AreaAssignLogRoutes(r)
	//ser.DailyAttendanceCorn()

	//EmployeeAssetsRoutes

	// c := cron.New()
	// c.AddFunc("30 06 * * *", ser.DailyAttendanceCorn)

	//c.Start()
	log.Println("listening in port : ", sh.GetCmdArg(constants.PORT))
	log.Fatal(http.ListenAndServe("0.0.0.0:"+sh.GetCmdArg(constants.PORT), r))
}
