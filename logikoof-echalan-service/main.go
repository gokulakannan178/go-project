package main

import (
	"fmt"
	"log"
	"logikoof-echalan-service/config"
	"logikoof-echalan-service/constants"
	"logikoof-echalan-service/daos"
	"logikoof-echalan-service/handlers"
	"logikoof-echalan-service/middlewares"
	"logikoof-echalan-service/redis"
	"logikoof-echalan-service/routes"
	"logikoof-echalan-service/services"
	"logikoof-echalan-service/shared"
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
	r := mux.NewRouter()
	//commonRoute := rr.PathPrefix("/api").Subrouter()
	//r := commonRoute.NewRoute().Subrouter()
	r.Use(middlewares.Log)
	r.Use(middlewares.AllowCors)
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
	route.UserAuthRoutes(r)

	//Vehicle Routes
	route.VehicleRoutes(r)
	route.VehicleChallanRoutes(r)
	route.OffenceTypeRoutes(r)

	//Dashboard Routes
	route.DashboardRoutes(r)

	//Video Routes
	route.LiveVideoRoutes(r)
	route.OffenceVideoRoutes(r)

	log.Println("listening in port : ", sh.GetCmdArg(constants.PORT))
	log.Fatal(http.ListenAndServe("0.0.0.0:"+sh.GetCmdArg(constants.PORT), r))
}
