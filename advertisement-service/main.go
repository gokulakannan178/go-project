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

	//User Routes
	route.OrganisationRoutes(r)
	route.UserRoutes(r)
	route.UserAuthRoutes(nonau)
	route.UserTypeRoutes(r)

	//Customer
	route.CustomerRoutes(r)
	//CustomerAuthRoutes
	route.CustomerAuthRoutes(r)
	//ProjectRoutes
	route.ProjectRoutes(r)
	//AdvertisementRoutes
	route.AdvertisementRoutes(r)

	http.DefaultClient.Timeout = time.Minute * 10
	log.Println("listening in port : ", sh.GetCmdArg(constants.PORT))
	log.Fatal(http.ListenAndServe("0.0.0.0:"+sh.GetCmdArg(constants.PORT), rr))
}
