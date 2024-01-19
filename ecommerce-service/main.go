package main

import (
	"ecommerce-service/config"
	"ecommerce-service/constants"
	"ecommerce-service/daos"
	"ecommerce-service/handlers"
	"ecommerce-service/middlewares"
	"ecommerce-service/redis"
	"ecommerce-service/routes"
	"ecommerce-service/services"
	"ecommerce-service/shared"
	"fmt"
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
	route.ZoneRoutes(r)
	route.WardRoutes(r)

	//User Routes
	route.OrganisationRoutes(r)
	route.UserRoutes(r)
	route.UserAuthRoutes(nonau)
	route.UserTypeRoutes(r)

	//ACL Routes
	route.ModuleRoutes(r)
	route.MenuRoutes(r)
	route.TabRoutes(r)
	route.FeatureRoutes(r)
	route.ACLMasterUserRoutes(r)
	route.ACLAccess(r)

	// Inventory Routes
	route.InventoryRoutes(r)

	// Product Routes
	route.ProductRoutes(r)

	// Product Variant Routes
	route.ProductVariantRoutes(r)

	// Product Variant Type Routes
	route.ProductVariantTypeRoutes(r)
	// ProductVariantMeshRoutes
	route.ProductVariantMeshRoutes(r)
	// Product Register Routes
	route.ProductRegisterRoutes(r)
	// Vendor Routes
	route.VendorRoutes(r)

	// Vendor Info Routes
	route.VendorInfoRoutes(r)

	// Category Routes
	route.CategoryRoutes(r)

	// SubCategory Routess
	route.SubCategoryRoutes(r)

	// CartRoutes
	route.CartRoutes(r)

	//Customer
	route.CustomerRoutes(r)
	//CustomerAuthRoutes
	route.CustomerAuthRoutes(r)
	//wallet
	route.WalletRoutes(r)
	//walletLog
	route.WalletLogRoutes(r)
	//BasicLogics
	route.ModuleLogic(r)

	//OrderRoutes
	route.OrderRoutes(r)
	//PaymentModeRoutes
	route.PaymentModeRoutes(r)
	//OrderPaymentRoutes
	route.OrderPaymentRoutes(r)
	//ScenarioRoutes
	route.ScenarioRoutes(r)
	//VendorDashboardRoutes
	route.VendorDashboardRoutes(r)

	http.DefaultClient.Timeout = time.Minute * 10
	log.Println("listening in port : ", sh.GetCmdArg(constants.PORT))
	log.Fatal(http.ListenAndServe("0.0.0.0:"+sh.GetCmdArg(constants.PORT), rr))
}
