package main

import (
	"fmt"
	"log"
	"net/http"
	"nicessm-api-service/config"
	"nicessm-api-service/constants"
	"nicessm-api-service/daos"
	"nicessm-api-service/handlers"
	"nicessm-api-service/middlewares"
	"nicessm-api-service/models"
	"nicessm-api-service/redis"
	"nicessm-api-service/routes"
	"nicessm-api-service/services"
	"nicessm-api-service/shared"
	"os"
	"time"

	"github.com/gorilla/mux"
	"github.com/robfig/cron/v3"
)

func main() {
	ch := make(chan *models.WeatherDisseminationChennal)
	go SendSMS(ch)
	argsWithoutProg := os.Args[1:]
	config := config.Config()
	sh := shared.NewShared(shared.SplitCmdArguments(argsWithoutProg), config)
	redisConn := redis.Connect(config)
	cache := new(models.CacheMemory)
	cache.Data = make(map[string]models.CacheData)
	db := daos.GetDaos(sh, redisConn, config, cache)
	ser := services.GetService(db, sh, redisConn, config, ch, cache)
	han := handlers.GetHandler(ser, sh, redisConn, config, ch, cache)
	route := routes.GetRoute(han, sh, redisConn, config, cache)
	rr := mux.NewRouter()
	commonRoute := rr.PathPrefix("/api").Subrouter()
	// fileRoute := rr.PathPrefix("/").Subrouter()
	//UIRoute := rr.PathPrefix("/ui").Subrouter()
	rr.Use(middlewares.Log)
	rr.Use(middlewares.AllowCors)
	r := commonRoute.NewRoute().Subrouter()
	// fr := fileRoute.NewRoute().Subrouter()
	//r.Use(middlewares.JWT)
	nonau := commonRoute.NewRoute().Subrouter()
	r.Methods("OPTIONS").HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("options called")
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE, PATCH")
		w.Header().Set("Access-Control-Allow-Headers", "Accept, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, Access-Control-Request-Headers, Access-Control-Request-Method, Connection, Host, Origin, User-Agent, Referer, Cache-Control, X-header")
		w.WriteHeader(http.StatusNoContent)
		return
	})
	//UI Routes
	//route.UIRoutes(UIRoute)
	//Common Routes
	route.CommonRoutes(r)
	route.FileRoutes(r)

	//Geo Location Routes
	route.StateRoutes(r)
	route.DistrictRoutes(r)
	route.VillageRoutes(r)
	route.BlockRoutes(r)
	route.GramPanchayatRoutes(r)

	//User Routes
	route.OrganisationRoutes(r)
	route.UserOrganisationRoutes(r)
	route.UserRoutes(r)
	route.UserAuthRoutes(nonau)
	route.UserTypeRoutes(r)

	//User Language
	route.LanguageRoutes(r)

	//User Cropseason
	route.CropseasonRoutes(r)

	//User Insect
	route.InsectRoutes(r)
	//user AgroEcologicalZone
	route.AgroEcologicalZoneRoutes(r)
	//User Market
	route.MarketRoutes(r)

	//User Aidlocation
	route.AidlocationRoutes(r)

	//User ProductConfig
	route.ProductConfigRoutes(r)

	//ACL Routes
	route.ModuleRoutes(r)
	route.MenuRoutes(r)
	route.TabRoutes(r)
	route.FeatureRoutes(r)
	route.ACLMasterUserRoutes(r)
	route.ACLAccess(r)
	//SoilTypeRoutes
	route.SoilTypeRoutes(r)
	//AssetRoutes
	route.AssetRoutes(r)
	//CastRoutes
	route.CastRoutes(r)

	// Project Routes
	route.ProjectRoutes(r)
	// ProjectKnowledgeDomain Routes
	route.ProjectKnowledgeDomainRoutes(r)
	// KnowledgeDomain Routes
	route.KnowledgeDomainRoutes(r)
	//SubDomainRoutes
	route.SubDomainRoutes(r)
	//TopicRoutes
	route.TopicRoutes(r)
	//TopicRoutes
	route.SubTopicRoutes(r)
	//FarmerRoutes
	route.FarmerRoutes(r)
	//FarmerLandRoutes
	route.FarmerLandRoutes(r)
	//FeedBackRoutes
	route.FeedBackRoutes(r)
	//query
	route.QueryRoutes(r)
	//NutrientsRoutes
	route.NutrientsRoutes(r)
	//ClusterRoutes
	route.ClusterRoutes(r)
	//CommonLandRoutes
	route.CommonLandRoutes(r)
	//LiveStockVaccinationRoutes
	route.LiveStockVaccinationRoutes(r)
	//BlockCropRoutes
	route.BlockCropRoutes(r)
	//StateLiveStockRoutes
	route.StateLiveStockRoutes(r)

	// ProjectState Routes
	route.ProjectStateRoutes(r)
	// ProjectUser Routes
	route.ProjectUserRoutes(r)
	// AidCategoryRoutes
	route.AidCategoryRoutes(r)
	// ContentRoutes
	route.ContentRoutes(r)
	// ContentcommentRoutes
	route.ContentCommentRoutes(r)
	// ContentTranslationRoutes
	route.ContentTranslationRoutes(r)
	// ProjectFarmer Routes
	route.ProjectFarmerRoutes(r)
	// ProjectPartner Routes
	route.ProjectPartnerRoutes(r)
	// DiseaseRoutes
	route.DiseaseRoutes(r)
	// BannedItemRoutes
	route.BannedItemRoutes(r)
	// Vaccine Routes
	route.VaccineRoutes(r)
	// CommodityCategory Routes
	route.CommodityCategoryRoutes(r)
	// CommodityFunction Routes
	route.CommodityFunctionRoutes(r)
	// Commodity Routes
	route.CommodityRoutes(r)
	// CommodityStage Routes
	route.CommodityStageRoutes(r)
	// CommodityVariety Routes
	route.CommodityVarietyRoutes(r)
	// CommoditySubVariety Routes
	route.CommoditySubVarietyRoutes(r)
	// DistrictWeatherDataRoutes
	route.DistrictWeatherDataRoutes(r)
	// SelfRegisterRoutes
	route.SelfRegisterRoutes(r)
	// FarmerLiveStockRoutes
	route.FarmerLiveStockRoutes(r)
	// FarmerCropRoutes
	route.FarmerCropRoutes(r)
	//  NutrientValueRoutes
	route.NutrientValueRoutes(r)
	// FarmerSoilDataRoutes
	route.FarmerSoilDataRoutes(r)
	// FarmerLandUploadRoutes
	route.FarmerLandUploadRoutes(r)
	// LandCropRoutes
	route.LandCropRoutes(r)
	// FarmerCropCalendarRoutes
	route.FarmerCropCalendarRoutes(r)
	// LandCropCalendarItemRoutes
	route.LandCropCalendarRoutes(r)
	//DisseminationRoutes
	route.DisseminationRoutes(r)
	// ContentManagerRoutes
	route.ContentManagerRoutes(r)

	route.QueryReportRoutes(r)
	// UserReportRoutes
	route.UserReportRoutes(r)
	//FarmerReportRoutes
	route.FarmerReportRoutes(r)
	//ContentReportRoutes
	route.ContentReportRoutes(r)
	//EmailLogRoutes
	route.EmailLogRoutes(r)
	//SmsLogRoutes
	route.SmsLogRoutes(r)
	//NotificationLogRoutes
	route.NotificationLogRoutes(r)
	//WhatsappLogRoutes
	route.WhatsappLogRoutes(r)
	//TelegramLogRoutes
	route.TelegramLogRoutes(r)
	// ContentDisseminationRoutes
	route.ContentDisseminationRoutes(r)
	// DashboardContentRoutes
	route.DashboardContentRoutes(r)
	// DashboardUserCountRoutes
	route.DashboardUserCountRoutes(r)
	// DashboardFarmerCountRoutes
	route.DashboardFarmerCountRoutes(r)
	// DashboardQueryCountRoutes
	route.DashboardQueryCountRoutes(r)
	//DealerRoutes
	route.DealerRoutes(r)
	//DealerRegistrationRoutes
	route.DealerRegistrationRoutes(r)
	//COLLECTIONSUBCATEGORY
	route.SubCategoryRoutes(r)
	//ProductRoutes
	route.ProductRoutes(r)
	//CategoryRoutes
	route.CategoryRoutes(r)
	//OrderRoutes
	route.OrderRoutes(r)
	//OrderPaymentsRoutes
	route.OrderPaymentsRoutes(r)
	//FarmerAuthRoutes
	route.FarmerAuthRoutes(r)
	//ApptokenRoutes
	route.ApptokenRoutes(r)
	//MonthSeasonRoutes
	route.MonthSeasonRoutes(r)
	//SeasonRoutes
	route.SeasonRoutes(r)
	//WeatherparameterRoutes
	route.WeatherparameterRoutes(r)
	//WeatherMasterRoutes
	route.StateWeatherAlertMasterRoutes(r)
	//OnePageAttachmentRoutes
	route.OnePageAttachmentRoutes(r)
	//WeatherAlertRoutes
	route.StateWeatherAlertRoutes(r)
	//StateWeatherDataRoutes
	route.StateWeatherDataRoutes(r)
	//WeatherAlertTypeRoutes
	route.WeatherAlertTypeRoutes(r)
	//CommunicationCreditRoutes
	route.CommunicationCreditRoutes(r)
	//CommunicationCreditRoutesLog
	route.CommunicationCreditLogRoutes(r)
	//WeatherAlertDissiminationRoutes
	route.WeatherAlertDissiminationRoutes(r)
	//DistrictWeatherAlertMasterRoutes
	route.DistrictWeatherAlertMasterRoutes(r)
	//DistrictWeatherAlertRoutes
	route.DistrictWeatherAlertRoutes(r)
	//DistrictWeatherAlertNotInRangeRoutes
	route.DistrictWeatherAlertNotInRangeRoutes(r)
	//DistrictweatheralertdissiminationRoutes
	route.DistrictweatheralertdissiminationRoutes(r)
	//WeatherAlertNotInRangeRoutes
	route.WeatherAlertNotInRangeRoutes(r)
	//ContentCountLogRoutes
	route.ContentCountLogRoutes(r)
	//UserLoginLogRoutes
	route.UserLoginLogRoutes(r)
	//ContentViewLogRoutes
	route.ContentViewLogRoutes(r)
	//UserAclRoutes
	route.UserAclRoutes(r)
	//CompendiumRoutes
	route.CompendiumRoutes(r)
	//CommonLanguageTranslationsRoutes
	route.CommonLanguageTranslationsRoutes(r)
	//LanguageTranslationRoutes
	route.LanguageTranslationRoutes(r)
	//TodayAdvisoryRoutes
	route.TodayAdvisoryRoutes(r)
	//ser.SaveWeatherDataCron()
	//ser.SaveDistrictWeatherDataCron()

	c := cron.New()
	//c.AddFunc("@every 30m", ser.FarmerExcelCron)
	//ser.SendWeatheralertCron()
	// c.AddFunc("@midnight", ser.FarmerExcelCron)
	// c.AddFunc("@earlymorning", ser.SendLaterCron)
	c.AddFunc("30 06 * * *", ser.SendLaterCron)
	c.AddFunc("30 06 * * *", ser.SaveWeatherDataCron)
	c.AddFunc("30 06 * * *", ser.SaveDistrictWeatherDataCron)
	c.AddFunc("30 05 * * *", ser.SendStateWeatherAlertCron)

	c.Start()

	http.DefaultClient.Timeout = time.Minute * 10
	log.Println("listening in port : ", sh.GetCmdArg(constants.PORT))
	log.Fatal(http.ListenAndServe("0.0.0.0:"+sh.GetCmdArg(constants.PORT), rr))
}

func SendSMS(c chan *models.WeatherDisseminationChennal) {
	for {
		fmt.Println("waiting for  data")
		fmt.Println("data ==", <-c)
		weather := <-c
		fmt.Println(weather)
		time.Sleep(3600 * time.Second)
		fmt.Println("COMPLETED")
	}
}
