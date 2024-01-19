package main

import (
	"hrms-services/config"
	"hrms-services/constants"
	"hrms-services/daos"
	"hrms-services/handlers"
	"hrms-services/middlewares"
	"hrms-services/models"
	"hrms-services/redis"
	"hrms-services/routes"
	"hrms-services/services"
	"hrms-services/shared"
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
	"github.com/robfig/cron/v3"
)

func main() {
	argsWithoutProg := os.Args[1:]
	config := config.Config()
	sh := shared.NewShared(shared.SplitCmdArguments(argsWithoutProg), config)
	redisConn := redis.Connect(config)
	cache := new(models.CacheMemory)
	db := daos.GetDaos(sh, redisConn, config)
	ser := services.GetService(db, sh, redisConn, config, cache)
	han := handlers.GetHandler(ser, sh, redisConn, config, cache)
	route := routes.GetRoute(han, sh, redisConn, config, cache)
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
	//Master Routes
	route.HonorifficRoutes(r)
	route.RelationRoutes(r)

	//User Routes
	route.OrganisationRoutes(r)
	route.UserRoutes(r)
	route.UserAuthRoutes(r)
	route.UserTypeRoutes(r)
	route.UserLocationRoutes(r)
	route.ConsumerRoutes(r)

	//ACL Routes
	route.ModuleRoutes(r)
	route.MenuRoutes(r)
	route.TabRoutes(r)
	route.FeatureRoutes(r)
	route.ACLMasterUserRoutes(r)
	route.ACLAccess(r)
	route.Project(r)

	//Task Routes
	route.TaskRoutes(r)

	//Attendance Routes
	route.AttendanceRoutes(r)

	// DepartmentRoutes
	route.DepartmentRoutes(r)

	//BranchRoutes
	route.BranchRoutes(r)

	//EmployeeRoutes
	route.EmployeeRoutes(r)

	//EmployeeJobRoutes
	route.EmployeeJobRoutes(r)

	//EmployeeHistoryRoutes
	route.EmployeeHistoryRoutes(r)

	//EmergencyContactRoutes
	route.EmergencyContactRoutes(r)

	//DocumentTypeRoutes
	route.DocumentTypeRoutes(r)

	//DocumentScenarioRoutes
	route.DocumentScenarioRoutes(r)

	//DocumentMuxMasterRoutes
	route.DocumentMuxMasterRoutes(r)

	//WorkScheduleRoutes
	route.WorkScheduleRoutes(r)

	//BankInformationRoutes
	route.BankInformationRoutes(r)

	//JobTimelineRoutes
	route.JobTimelineRoutes(r)

	//ProbationaryRoutes
	route.ProbationaryRoutes(r)

	//DesignationRoutes
	route.DesignationRoutes(r)

	//DashboardRoutes
	route.DashboardRoutes(r)

	//EmployeeReportRoutes
	route.EmployeeReportRoutes(r)

	//LeaveMasterRoutes
	route.LeaveMasterRoutes(r)

	//LeavePolicyRoutes
	route.LeavePolicyRoutes(r)

	//PolicyRuleRoutes
	route.PolicyRuleRoutes(r)

	//EmployeeLogRoutes
	route.EmployeeLogRoutes(r)

	//EmployeeLeaveRoutes
	route.EmployeeLeaveRoutes(r)

	//OffboardingCheckListRoutes
	route.OffboardingCheckListRoutes(r)

	//OffboardingCheckListMasterRoutes
	route.OffboardingCheckListMasterRoutes(r)

	//OffboardingPolicyRoutes
	route.OffboardingPolicyRoutes(r)

	//NoticePolicyRoutes
	route.NoticePolicyRoutes(r)

	//OnboardingCheckListRoutes
	route.OnboardingCheckListRoutes(r)

	//OnboardingCheckListMasterRoutes
	route.OnboardingCheckListMasterRoutes(r)

	//OnboardingPolicyRoutes
	route.OnboardingPolicyRoutes(r)

	//AttendanceLogRoutes
	route.AttendanceLogRoutes(r)

	//EmployeeOnboardingCheckListRoutes
	route.EmployeeOnboardingCheckListRoutes(r)

	//EmployeeOffboardingCheckListRoutes
	route.EmployeeOffboardingCheckListRoutes(r)

	//DayOfWeekRoutes
	route.DayOfWeekRoutes(r)

	//SmsLogRoutes
	route.SmsLogRoutes(r)

	//EmailLogRoutes
	route.EmailLogRoutes(r)

	//ApptokenRoutes
	route.ApptokenRoutes(r)

	//NewsRoutes
	route.NewsRoutes(r)

	//BillClaimRoutes
	route.BillClaimRoutes(r)

	//ser.DaywiseAttendance()

	//NewsLikeRoutes
	route.NewsLikeRoutes(r)

	//NewsCommentRoutes
	route.NewsCommentRoutes(r)

	//DocumentPolicyRoutes
	route.DocumentPolicyRoutes(r)

	//DocumentMaterRoutes
	route.DocumentMaterRoutes(r)

	//DocumentPolicyDocumentsRoutes
	route.DocumentPolicyDocumentsRoutes(r)

	//AssetPolicyRoutes
	route.AssetPolicyRoutes(r)

	//AssetMasterRoutes
	route.AssetMasterRoutes(r)

	//AssetPolicyAssetsRoutes
	route.AssetPolicyAssetsRoutes(r)

	//EmployeeDocumentsRoutes
	route.EmployeeDocumentsRoutes(r)

	//HolidaysRoutes
	route.HolidaysRoutes(r)

	//AssetRoutes
	route.AssetRoutes(r)

	//AssetTypeRoutes
	route.AssetTypeRoutes(r)

	//AssetTypePropertysRoutes
	route.AssetTypePropertysRoutes(r)

	//UserLocationTrackerRoutes
	route.UserLocationTrackerRoutes(r)

	//EmployeeExperienceRoutes
	route.EmployeeExperienceRoutes(r)

	//EmployeeEducationRoutes
	route.EmployeeEducationRoutes(r)

	//NotificationLogRoutes
	route.NotificationLogRoutes(r)

	//EmployeeFamilyMembersRoutes
	route.EmployeeFamilyMembersRoutes(r)

	//EmployeeTimeOffRoutes
	route.EmployeeTimeOffRoutes(r)

	//EmployeeLeaveLogRoutes
	route.EmployeeLeaveLogRoutes(r)

	//SalaryRoutes
	route.SalaryRoutes(r)

	//PaymentRoutes
	route.PaymentRoutes(r)

	//EmployeeAttendanceCalendarRoutes
	route.EmployeeAttendanceCalendarRoutes(r)

	//EmployeeEarningMasterRoutes
	route.EmployeeEarningMasterRoutes(r)

	//EmployeeDeductionMasterRoutes
	route.EmployeeDeductionMasterRoutes(r)

	//EmployeeEarningRoutes
	route.EmployeeEarningRoutes(r)

	//EmployeeDeductionRoutes
	route.EmployeeDeductionRoutes(r)

	//UserAclRoutes
	route.UserAclRoutes(r)
	//EmployeePayrollRoutes
	route.EmployeePayrollRoutes(r)

	//PayrollPolicyRoutes
	route.PayrollPolicyRoutes(r)
	//PayrollPolicyEarningRoutes
	route.PayrollPolicyEarningRoutes(r)
	//PayrollPolicyDetectionRoutes
	route.PayrollPolicyDetectionRoutes(r)
	//DemoUserRoutes
	route.DemoUserRoutes(r)
	//\\ser.DailyAttendanceCorn()
	//EmployeePayslipRoutes
	route.EmployeePayslipRoutes(r)
	//PayrollRoutes
	route.PayrollRoutes(r)
	//PayrollRoutes
	route.PayrollLogRoutes(r)
	//EmployeeAssetsRoutes
	route.EmployeeAssetsRoutes(r)
	//AssetLogRoutes
	route.AssetLogRoutes(r)
	//EmployeeSalaryRoutes
	route.EmployeeSalaryRoutes(r)
	//ser.EmployeeTree()
	//SalaryConfigRoutes
	route.SalaryConfigRoutes(r)
	//BillclaimConfigRoutes
	route.BillclaimConfigRoutes(r)
	//ExpenseCategoryRoutes
	route.ExpenseCategoryRoutes(r)
	//ExpenseSubCategoryRoutes
	route.ExpenseSubcategoryRoutes(r)
	//ExpenseCategorylistRoutes
	route.ExpenseCategoryListRoutes(r)
	//GradeRoutes
	route.GradeRoutes(r)
	//BillclaimLevelsRoutes
	route.BillclaimLevelsRoutes(r)
	//ser.EmployeePayslipCorn()
	c := cron.New()
	c.AddFunc("30 03 * * *", ser.DailyAttendanceCorn)

	c.Start()
	log.Println("listening in port : ", sh.GetCmdArg(constants.PORT))
	log.Fatal(http.ListenAndServe("0.0.0.0:"+sh.GetCmdArg(constants.PORT), r))
}
