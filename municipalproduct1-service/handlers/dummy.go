package handlers

import (
	"municipalproduct1-service/models"
	"municipalproduct1-service/response"
	"net/http"
)

//AssetAPIForInterview : ""
func (h *Handler) AssetAPIForInterview(w http.ResponseWriter, r *http.Request) {
	platform := r.URL.Query().Get("platform")
	data := new(models.AssetTestModelForInterview)

	data.InhandBalance = 80000
	data.Todaystarget = 100000
	data.Name = "Ravi Kumar"
	data.TodayCollection = 67599
	data.YesterdayCollection = 89897
	data.ProjectWise.TotalCollection = 89989
	data.ProjectWise.TotalDemand = 76767
	data.Arrear.ArrearCollection = 676
	data.Arrear.ArrearDemand = 7878
	data.UserTypeID = "Tax Collector"
	data.Achieved = true
	x := make(map[string]float64)
	x["total Demand"] = 2772.2

	y := make(map[string]float64)
	y["total Collection"] = 3332.89

	z := make(map[string]float64)
	z["total Survey"] = 4564.99

	s := make(map[string]float64)
	s["collections Of houses"] = 9865.56

	data.Array = append(data.Array, x, y, z, s)

	m := make(map[string]interface{})
	m["data"] = data
	response.With200V2(w, "Success", m, platform)
}
