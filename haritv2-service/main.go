package main

import (
	"fmt"
	"haritv2-service/config"
	"haritv2-service/constants"
	"haritv2-service/daos"
	"haritv2-service/handlers"
	"haritv2-service/middlewares"
	"haritv2-service/redis"
	"haritv2-service/routes"
	"haritv2-service/services"
	"haritv2-service/shared"
	"log"
	"net/http"
	"os"
	"time"

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
	route.ULBRoutes(r)
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

	// FPO Routes
	route.FPORoutes(r)
	route.FPOInventoryRoutes(r)
	//CART
	route.CartRoutes(r)
	//THINGSTOKNOW
	route.ThingsToKnowRoutes(r)

	// TodayTips Routes
	route.TodayTipsRoutes(r)

	// Farmer Routes
	route.FarmerRoutes(r)
	route.CustomerRoutes(r)

	//  Farmer Cart Routes
	route.FarmerCartRoutes(r)

	// GST Routes
	route.GSTRoutes(r)

	// Product Routes
	route.ProductRoutes(r)

	//Order Routes
	route.OrderRoutes(r)

	// Product Category Routes
	route.ProductCategoryRoutes(r)

	//Popup Notification
	route.PopupNotificationRoutes(r)
	//Advertisement
	route.AdvertisementRoutes(r)
	// Payment Routes
	route.PaymentRoutes(r)

	//Apptoken
	route.ApptokenRoutes(r)

	// PackageType  Routes
	route.PkgTypeRoutes(r)

	// DeliverSale Routes
	route.DeliverSaleRoutes(r)
	// Notification Routes
	route.NotificationRoutes(r)
	// CustomNotification Routes
	route.CustomNotificationRoutes(r)
	// NotifyRoutes
	route.NotifyRoutes(r)
	// 	route.SaleRoutes(r)
	route.SaleRoutes(r)
	// 	route.RoleRoutes(r)
	route.RoleRoutes(r)
	// Cron Routes
	route.CronRoutes(r)
	//SelfConsumptionRoutes
	route.SelfConsumptionRoutes(r)
	//ConsumerRegistrationRoutes
	route.CustomerRegistrationRoutes(r)

	http.DefaultClient.Timeout = time.Minute * 10
	log.Println("listening in port : ", sh.GetCmdArg(constants.PORT))
	log.Fatal(http.ListenAndServe("0.0.0.0:"+sh.GetCmdArg(constants.PORT), rr))
}
