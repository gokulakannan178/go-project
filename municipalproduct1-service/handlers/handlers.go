package handlers

import (
	"log"
	"municipalproduct1-service/app"
	"municipalproduct1-service/config"
	"municipalproduct1-service/constants"
	"municipalproduct1-service/models"
	"municipalproduct1-service/redis"
	"municipalproduct1-service/response"
	"municipalproduct1-service/services"
	"municipalproduct1-service/shared"
	"net/http"
)

//Handler : ""
type Handler struct {
	Service      *services.Service
	Shared       *shared.Shared
	Redis        *redis.RedisCli
	ConfigReader *config.ViperConfigReader
}

//GetHandler :""
func GetHandler(service *services.Service, s *shared.Shared, Redis *redis.RedisCli, configReader *config.ViperConfigReader) *Handler {
	return &Handler{service, s, Redis, configReader}
}

func (h *Handler) Authorization(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		var ctx *models.Context
		ctx = app.GetApp(r.Context(), h.Service.Daos)
		defer ctx.Client.Disconnect(r.Context())
		requestAgentType := r.Header.Get("requestAgentType")
		if requestAgentType == "web" {
			isLoggedIn := r.Header.Get("isLoggedIn")
			if isLoggedIn == "Yes" {
				userName := r.Header.Get("userName")
				authorization := r.Header.Get("Authorization")
				resUser, err := h.Service.GetSingleUser(ctx, userName)
				if err != nil {
					response.With401mV2(w, err.Error(), "")
					return
				}
				if resUser == nil {
					response.With401mV2(w, "unathourised user - user not found", "")
					log.Println("user not found - " + userName)
					return
				}
				if resUser.Token != authorization {
					response.With401mV2(w, "unathourised user - token missmatch", "")
					log.Println("user not found - " + userName)
					return
				}
				if resUser.IsForcedLogout == "Yes" {
					response.With401mV2(w, "unathourised user - forced logout", "")
					log.Println("user not found - " + userName)
					return
				}
				if resUser.Status == constants.USERSTATUSDISABLED || resUser.Status == constants.USERSTATUSDELETED {
					response.With401mV2(w, "unauthorised user - please contact administartor", "")
					return
				}
				if resUser.IsForcedLogout == "Yes" {
					response.With401mV2(w, "unauthorised user", "")
					return
				}
			}
		}

		if requestAgentType == "mobile" {
			isLoggedIn := r.Header.Get("isLoggedIn")
			if isLoggedIn == "Yes" {
				userName := r.Header.Get("userName")
				authorization := r.Header.Get("Authorization")
				resUser, err := h.Service.GetSingleUser(ctx, userName)
				if err != nil {
					response.With401mV2(w, err.Error(), constants.PLATFORMMOBILE)
					return
				}
				if resUser == nil {
					response.With401mV2(w, "unathourised user - user not found", constants.PLATFORMMOBILE)
					log.Println("user not found - " + userName)
					return
				}
				if resUser.MToken != authorization {
					response.With401mV2(w, "unathourised user - token missmatch", constants.PLATFORMMOBILE)
					log.Println("user not found - " + userName)
					return
				}
				if resUser.IsForcedLogout == "Yes" {
					response.With401mV2(w, "unathourised user - forced logout", constants.PLATFORMMOBILE)
					log.Println("user not found - " + userName)
					return
				}
				if resUser.Status == constants.USERSTATUSDISABLED || resUser.Status == constants.USERSTATUSDELETED {
					response.With401mV2(w, "unauthorised user - please contact administartor", constants.PLATFORMMOBILE)
					return
				}
				if resUser.IsForcedLogout == "Yes" {
					response.With401mV2(w, "unauthorised user", constants.PLATFORMMOBILE)
					return
				}
			}
		}

		next.ServeHTTP(w, r)

	})
}
