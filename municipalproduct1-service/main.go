package main

import (
	"fmt"
	"log"
	"municipalproduct1-service/config"
	"municipalproduct1-service/constants"
	"municipalproduct1-service/daos"
	"municipalproduct1-service/handlers"
	"municipalproduct1-service/middlewares"
	"municipalproduct1-service/redis"
	"municipalproduct1-service/routes"
	"municipalproduct1-service/services"
	"municipalproduct1-service/shared"
	"net/http"
	"os"
	"time"

	"github.com/gorilla/mux"
	"github.com/robfig/cron/v3"
)

func main() {
	argsWithoutProg := os.Args[1:]
	config := config.Config()
	sh := shared.NewShared(shared.SplitCmdArguments(argsWithoutProg), config)
	redisConn := redis.Connect(config)
	db := daos.GetDaos(sh, redisConn, config)
	ser := services.GetService(db, sh, redisConn, config)
	han := handlers.GetHandler(ser, sh, redisConn, config)
	fmt.Println(argsWithoutProg)
	if sh.GetCmdArg("--FORSCRIPT") == "Yes" {
		//Place your script here
		ser.PropertyDemandSummaryCalc()
		return
	}
	route := routes.GetRoute(han, sh, redisConn, config)
	r := mux.NewRouter()
	r.Use(middlewares.Log)
	r.Use(middlewares.AllowCors)
	r.Methods("OPTIONS").HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("options called")
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE, PATCH")
		w.Header().Set("Access-Control-Allow-Headers", "Accept, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, Access-Control-Request-Headers, Access-Control-Request-Method, Connection, Host, Origin, User-Agent, Referer, Cache-Control, X-header, requestAgentType, isLoggedIn, userName")
		w.WriteHeader(http.StatusNoContent)
		return
	})
	r.Use(han.Authorization)

	//Over all demand calc
	// ser.PropertyCalc()
	// return

	// Save Demand in property collection
	//err := ser.SavePropertyDemandForAll()
	//fmt.Println(err)
	// return

	// //Save Collection in propertypayment collection
	// err = ser.CalcualteTotalCollectionForAllPayments()
	// fmt.Println(err)
	// return

	// // Save Collection in property collection
	// err = ser.SavePropertyCollectionForAll()
	// fmt.Println(err)
	// return

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
	//Property Master Routes
	route.ConstructionTypeRoutes(r)
	route.FloorTypeRoutes(r)
	route.MunicipalTypeRoutes(r)
	route.OccupancyTypeRoutes(r)
	route.RoadTypeRoutes(r)
	route.UsageTypeRoutes(r)
	route.PropertyTypeRoutes(r)
	route.FinancialYearRoutes(r)
	route.VacantLandRateRoutes(r)
	route.FloorRatableAreaRoutes(r)
	route.AVRRoutes(r)
	route.NonResidentialUsageFactorRoutes(r)
	route.OwnershipRoutes(r)
	route.PropertyTaxRoutes(r)
	route.RebateRoutes(r)
	route.PenaltyRoutes(r)
	route.ResidentialTypeRoutes(r)
	route.OtherChargesRoutes(r)

	//Property Master - SAF Setup Routes
	route.AVRRangeRoutes(r)
	route.PropertyOtherTaxRoutes(r)
	//Property Routes
	route.PropertyRoutes(r)
	route.PropertyTaxCalculation(r)

	//ProductConfiguration Routes
	route.ProductConfiguration(r)

	//BankDeposit Routes
	route.CreatedBankDeposit(r)

	//User Routes
	route.OrganisationRoutes(r)
	route.UserRoutes(r)
	route.UserAuthRoutes(r)
	route.UserTypeRoutes(r)
	route.UserLocationRoutes(r)
	route.ConsumerRoutes(r)
	//legacy
	route.LegacyPropertyRoutes(r)
	//contactsus
	route.ContactUs(r)
	//paymentgateway
	route.PaymentGatewayConfigurationRoutes(r)

	//Payment Gateway - Paytm
	route.PaymentGatewayPaytmRoutes(r)

	// MobileTowerRoutes
	route.MobileTowerRoutes(r)
	route.MobileTowerDemandRoutes(r)
	route.MobileTowerPaymentRoutes(r)
	route.MobileTowerPaymentReportsRoutes(r)
	route.MobileTowerRegistrationPaymentRoutes(r)
	route.MobileTowerLegacyRegistrationPaymentRoutes(r)
	route.MobileTowerRegistrationTaxRoutes(r)

	// CitizenGraviansRoutes
	route.CitizenGrievanceRoutes(r)
	// Reassessment Request Routes
	route.ReassessmentRequestRoutes(r)
	// Property Penalty
	route.PropertyPenaltyRoutes(r)
	//LetterUpload
	route.LetterUpload(r)
	//LetterGenerate
	route.LetterGenerate(r)

	//ACL Routes
	route.ModuleRoutes(r)
	route.MenuRoutes(r)
	route.TabRoutes(r)
	route.FeatureRoutes(r)
	route.ACLMasterUserRoutes(r)
	route.ACLAccess(r)

	//GSTRateMaster
	route.GSTRateMasterRoutes(r)

	//LeaseRent
	route.LeaseRentRoutes(r)
	route.LeaseRentShopCategoryRoutes(r)
	route.LeaseRentShopSubCategoryRoutes(r)

	//LeaseRentRateMaster
	route.LeaseRentRateMasterRoutes(r)

	//ShopRent
	route.ShopRentShopCategoryRoutes(r)
	route.ShopRentShopSubCategoryRoutes(r)
	route.ShopRentRoutes(r)
	route.ShopRentDemandRoutes(r)
	route.ShopRentPaymenRoutes(r)
	route.ShopRentCalcRoutes(r)

	//ShopRentRateMaster
	route.ShopRentRateMasterRoutes(r)

	//PropertyVisitLogRemarkType
	route.PropertyVisitLogRemarkType(r)
	//propertyVisitLog
	route.PropertyVisitLog(r)

	// User Attendance
	route.UserAttendanceRoutes(r)

	// Lease Routes
	route.LeaseRoutes(r)

	// Lease Day Wise
	route.LeaseDayWiseRoutes(r)

	// MobileTowerRegistrationRateMaster
	route.MobileTowerRegistrationRateMaster(r)

	// Mobile Tower
	route.MobileTowerDashboardRoutes(r)

	// Mobile Tower Day Wise
	route.MobileTowerDashBoardDayWiseRoutes(r)

	// Dashboard Property
	route.DashboardPropertyRoutes(r)

	// Dashboard Property Day Wise
	route.DashboardPropertyDayWiseRoutes(r)

	// Shop Rent Dashboard
	route.ShopRentDashboardRoutes(r)

	// Shop Rent DashBoard DayWise
	route.ShopRentDashboardDayWiseRoutes(r)

	// Trade Licence
	route.TradeLicenseDashboardRoutes(r)
	route.TradeLicenseDemandRoutes(r)
	route.TradeLicensePaymenRoutes(r)

	// Trade Licence Day Wise
	route.TradeLicenseDashboardDayWiseRoutes(r)

	// User Charge
	route.UserChargeDashboardRoutes(r)

	// User Charge Day Wise
	route.UserChargeDayWiseDashboardRoutes(r)

	// Water Bill
	route.WaterBillDashboardRoutes(r)

	// Water Bill Day Wise
	route.WaterBillDayWiseDashboardRoutes(r)

	// Monthly Target
	route.MonthlyTargetRoutes(r)

	// User Ward Access
	route.UserWardAccessRoutes(r)

	// User Zone Access
	route.UserZoneAccessRoutes(r)

	//TcDashBoard
	route.TcDashboardRoutes(r)

	// Trade License Routes
	route.TradeLicenseRoutes(r)

	// Trade License Business TypeRoutes
	route.TradeLicenseBusinessTypeRoutes(r)

	// Trade License Category TypeRoutes
	route.TradeLicenseCategoryTypeRoutes(r)

	// Trade License Rate Master
	route.TradeLicenseRateMasterRoutes(r)

	// PmTarget Routes
	route.PmTargetRoutes(r)

	// PmAchievement Routes
	route.PmAchievementRoutes(r)

	// Property Wallet Routes
	route.PropertyWalletRoutes(r)

	// Property WalletLog Routes
	route.PropertyWalletLogRoutes(r)

	// PropertyPartPaymentPaymenRoutes
	route.PropertyPartPaymentPaymentRoutes(r)
	// OverallPropertyDemand Routes
	route.OverallPropertyDemandRoutes(r)
	// OverallPropertyDemand Routes
	route.OverallPropertyDemandRoutes(r)
	// Job Routes
	route.JobRoutes(r)
	// JobLogRoutes
	route.JobLogRoutes(r)
	// Major Update Routes
	route.MajorUpdateRoutes(r)
	//CronRoutes
	route.CronRoutes(r)

	// ShopRentMonthlyDemand Routes
	route.ShopRentMonthlyDemandRoutes(r)
	route.PropertyReportRoutes(r)
	route.ShopRentReportRoute(r)
	route.TradeLicenseReportRoute(r)
	route.MobileTowerReportRoute(r)

	// PropertyDemandLogRoutes
	route.PropertyDemandLogRoutes(r)
	//CitizenGraviansLogRoutes
	route.CitizenGraviansLogRoutes(r)
	// TicketRoutes
	route.TicketRoutes(r)
	// TicketUserRoutes
	route.TicketUserRoutes(r)
	// TicketCommentRoutes
	route.TicketCommentRoutes(r)
	//DocumentListRoutes
	route.DocumentListRoutes(r)
	//PropertyDocumentRoutes
	route.PropertyDocumentRoutes(r)
	route.CronLogRoutes(r)
	//ShoprentReassessmentRequestRoutes
	route.ShoprentReassessmentRequestRoutes(r)
	// TradeLicenseReassessmentRequestRoutes
	route.TradeLicenseReassessmentRequestRoutes(r)
	// MobileTowerReassessmentRequestRoutes
	route.MobileTowerReassessmentRequestRoutes(r)
	//PropertyRequiredDocumentRoutes
	route.PropertyRequiredDocumentRoutes(r)
	// UserLocationTrackerRoutes
	route.UserLocationTrackerRoutes(r)
	// UserLocationLogRoutes
	route.UserLocationLogRoutes(r)
	// TradeLicenseRebateRoutes
	route.TradeLicenseRebateRoutes(r)
	//TradeLicensePart2Routes
	route.TradeLicenseDemandPart2Routes(r)
	// TradeLicensePaymentPart2Routes
	route.TradeLicensePaymentsPart2Routes(r)
	//SaveStoredDemandRoute
	route.SaveStoredDemandRoute(r)
	// Property Payment Mode Change Routes
	route.PropertyPaymentModeChangeRoutes(r)
	route.UserchargePaymentModeChangeRoutes(r)
	route.TradelicensePaymentModeChangeRoutes(r)
	route.ShoprentPaymentModeChangeRoutes(r)
	// Estimated Property Demand Routes
	route.EstimatedPropertyDemandRoutes(r)
	// Property Mutation Request Routes
	route.PropertyMutationRequestRoutes(r)
	//Property Delete Request Routes
	route.PropertyDeleteRequestRoutes(r)
	// Boring Charges Routes
	route.BoringChargesRoutes(r)

	//SolidWasteUserCharge
	route.SolidWasteUserChargeRoutes(r)
	route.SolidWasteUserChargeRateRoutes(r)
	route.SolidWasteUserChargeCategoryRoutes(r)
	route.SolidWasteUserChargeSubCategoryRoutes(r)
	route.SolidWasteUserChargeMonthlyDemandRoutes(r)

	//PropertyOtherDemandRoutes
	route.PropertyOtherDemandRoutes(r)
	// HDFCPaymentGatewayRoutes
	route.HDFCPaymentGatewayRoutes(r)

	//WaterTax
	route.WaterTaxArvRoutes(r)
	route.WaterTaxConnectionTypeRoutes(r)

	//BirthCertificateRoutes
	route.BirthCertificateRoutes(r)
	route.HospitalRoutes(r)

	//propertyFixedArv
	route.PropertyFixedArvRoutes(r)
	route.PropertyFixedArvLogRoutes(r)
	route.SolidWasteReassessmentRequestRoutes(r)

	// Parking Penalty Routes
	route.ParkingPenaltyRoutes(r)
	route.TradeLicenseDeleteRequestRoutes(r)
	route.ShopRentDeleteRequestRoutes(r)
	route.MobileTowerDeleteRequestRoutes(r)
	route.OurServiceRoutes(r)

	//usercharge
	route.UserChargeCategoryRoutes(r)
	route.PropertyUserChargeRoutes(r)
	route.UserChargeRateMasterRoutes(r)
	route.UserChargeMonthlyDemandRoutes(r)
	route.UserChargeUpdateLogRoutes(r)

	//BulkPrintRoutes
	route.BulkPrintRoutes(r)

	route.UserChargeLogRoutes(r)
	route.PropertyPayeeNameChangeRoutes(r)
	route.UserchargePayeeNameChangeRoutes(r)
	route.TradelicensePayeeNameChangeRoutes(r)
	route.ShoprentPayeeNameChangeRoutes(r)
	// ser.PropertyOverallCollectionCron()
	// ser.PropertyTodayCollectionCron()
	// ser.PropertyOverallDemandCron()
	//ser.PropertyTodayDemandCron()
	// return

	// ser.UpdateTotalDeand()
	// ser.PropertyCalcCron()
	// ser.ShopRentCalcCron()
	// return
	// return
	// ser.PropertyCalc()
	c := cron.New()
	// c.AddFunc("* * * * * *", func() { fmt.Println("Every hour on the half hour") })
	// c.AddFunc("19 11-13,20-23 * * *", func() {
	// 	fmt.Println("On the half hour of 3-6am, 8-11pm")
	// })
	PROPERTYDEMANDCRON := sh.Config.GetString(sh.GetCmdArg(constants.ENV) + "." + constants.PROPERTYDEMANDCRON)
	if PROPERTYDEMANDCRON != "" {
		c.AddFunc(PROPERTYDEMANDCRON, ser.PropertyCalc)

	}

	PROPERTYDASHBOARDCRON := sh.Config.GetString(sh.GetCmdArg(constants.ENV) + "." + constants.PROPERTYDASHBOARDCRON)
	if PROPERTYDASHBOARDCRON != "" {
		c.AddFunc(PROPERTYDASHBOARDCRON, ser.PropertyTodayDemandCron)
		c.AddFunc(PROPERTYDASHBOARDCRON, ser.PropertyOverallDemandCron)
		c.AddFunc(PROPERTYDASHBOARDCRON, ser.PropertyTodayCollectionCron)
		c.AddFunc(PROPERTYDASHBOARDCRON, ser.PropertyOverallCollectionCron)
		c.AddFunc(PROPERTYDASHBOARDCRON, ser.PropertyMonthDemandCron)
		c.AddFunc(PROPERTYDASHBOARDCRON, ser.PropertyMonthCollectionCron)
		c.AddFunc(PROPERTYDASHBOARDCRON, ser.PropertyYearDemandCron)
		c.AddFunc(PROPERTYDASHBOARDCRON, ser.PropertyYearCollectionCron)

	}
	// ser.StoredPeropertyDemandCron()
	// //@midnight
	//c.Start()
	// ser.FilterPaymentSummery()
	http.DefaultClient.Timeout = time.Minute * 10
	log.Println("listening in port : ", sh.GetCmdArg(constants.PORT))
	log.Fatal(http.ListenAndServe("0.0.0.0:"+sh.GetCmdArg(constants.PORT), r))
}
