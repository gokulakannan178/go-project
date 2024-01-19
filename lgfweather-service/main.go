package main

import (
	"fmt"
	"lgfweather-service/config"
	"lgfweather-service/constants"
	"lgfweather-service/daos"
	"lgfweather-service/handlers"
	"lgfweather-service/middlewares"
	"lgfweather-service/redis"
	"lgfweather-service/routes"
	"lgfweather-service/services"
	"lgfweather-service/shared"
	"log"
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
	route := routes.GetRoute(han, sh, redisConn, config)
	rr := mux.NewRouter()
	commonRoute := rr.PathPrefix("/api").Subrouter()
	commonRoute.Use(middlewares.Log)
	commonRoute.Use(middlewares.AllowCors)
	r := commonRoute.NewRoute().Subrouter()
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

	//Common Routes
	route.CommonRoutes(r)
	route.FileRoutes(r)

	//Geo Location Routes
	route.StateRoutes(r)
	route.DistrictRoutes(r)
	route.VillageRoutes(r)
	route.BlockRoutes(r)
	route.GramPanchayatRoutes(r)
	route.DivisionRoutes(r)

	//User Routes
	route.OrganisationRoutes(r)
	route.UserRoutes(r)
	route.UserAuthRoutes(nonau)
	route.UserTypeRoutes(r)

	//ProductConfig
	route.ProductconfigsRoutes(r)

	//ACL Routes
	route.ModuleRoutes(r)
	route.MenuRoutes(r)
	route.TabRoutes(r)
	route.FeatureRoutes(r)
	route.ACLMasterUserRoutes(r)
	route.ACLAccess(r)

	//Popup Notification
	route.PopupNotificationRoutes(r)

	//Apptoken
	route.ApptokenRoutes(r)

	// Notification Routes
	route.NotificationRoutes(r)

	//StateWeatherDataRoutes
	route.StateWeatherDataRoutes(r)
	//DistrictWeatherDataRoutes
	route.DistrictWeatherDataRoutes(r)
	//BlockWeatherDataRoutes
	route.BlockWeatherDataRoutes(r)
	//GramPanchayatWeatherDataRoutes
	route.GramPanchayatWeatherDataRoutes(r)
	//VillageWeatherDataRoutes
	route.VillageWeatherDataRoutes(r)

	//WeatherAlertTypeRoutes
	route.WeatherAlertTypeRoutes(r)
	//WeatherparameterRoutes
	route.WeatherparameterRoutes(r)
	//MonthSeasonRoutes
	route.MonthSeasonRoutes(r)
	//ser.SaveWeatherDataCron()
	//ser.SaveDistrictWeatherDataCron()
	//ser.SaveBlockWeatherDataCron()
	//ser.SaveGramPanchayatWeatherDataCron()
	//ser.SaveVillageWeatherDataCron()
	//ser.LoadIMDDistrictWeatherV2()
	//ser.LoadIMDDistrictWeatherWithState()
	//ser.LoadIMDBlockWeatherWithState()
	c := cron.New()
	//c.AddFunc("@every 30m", ser.FarmerExcelCron)
	//ser.SendWeatheralertCron()
	// c.AddFunc("@midnight", ser.FarmerExcelCron)
	// c.AddFunc("@earlymorning", ser.SendLaterCron)
	//c.AddFunc("30 06 * * *", ser.SendLaterCron)
	c.AddFunc("30 06 * * *", ser.SaveStateWeatherDataCron)
	c.AddFunc("30 06 * * *", ser.SaveDistrictWeatherDataCron)
	c.AddFunc("30 05 * * *", ser.SaveBlockWeatherDataCron)
	c.AddFunc("30 06 * * *", ser.SaveGramPanchayatWeatherDataCron)
	c.AddFunc("30 06 * * *", ser.SaveVillageWeatherDataCron)
	//  c.AddFunc("30 05 * * *", ser.SendStateWeatherAlertCron)

	c.Start()
	//ser.LoadIMDDistrictWeather()
	http.DefaultClient.Timeout = time.Minute * 10
	log.Println("listening in port : ", sh.GetCmdArg(constants.PORT))
	log.Fatal(http.ListenAndServe("0.0.0.0:"+sh.GetCmdArg(constants.PORT), rr))
}
