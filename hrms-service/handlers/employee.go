package handlers

import (
	"encoding/json"
	"fmt"
	"hrms-services/app"
	"hrms-services/models"
	"hrms-services/response"
	"log"
	"net/http"
	"strconv"
)

//Employee Login
// func (h *Handler) EmployeeLogin(w http.ResponseWriter, r *http.Request) {
// 	platform := r.URL.Query().Get("platform")
// 	var ctx *models.Context
// 	ctx = app.GetApp(r.Context(), h.Service.Daos)
// 	euser := new(models.EmployeeLogin)
// 	err := json.NewDecoder(r.Body).Decode(&euser)
// 	defer r.Body.Close()
// 	if err != nil {
// 		response.With400V2(w, constants.RESPONSEINVALIDDATA+" "+err.Error(), platform)
// 		return
// 	}

// 	token, stat, err := h.Service.EmployeeLogin(ctx, euser)
// 	log.Println("stat ==>", stat)
// 	//	log.Println("err ==>", err.Error())
// 	log.Println("TOKEN==>", token)
// 	if err != nil {
// 		if err.Error() == constants.NOTFOUND {
// 			response.With403mV2(w, "Invalid User", platform)
// 			return
// 		}
// 		response.With500mV2(w, err.Error(), platform)
// 		return
// 	}
// 	if !stat {
// 		response.With403mV2(w, "Invalid Username or Password", platform)
// 		return
// 	}
// 	respUser, err := h.Service.GetSingleUser(ctx, euser.UserName)
// 	if err != nil {
// 		log.Println("err=>", err.Error())
// 	}
// 	m := make(map[string]interface{})
// 	m["token"] = token
// 	m["user"] = respUser
// 	// m["role"] = role
// 	response.With200V2(w, "Success", m, platform)
// }

//SaveEmployee : ""
func (h *Handler) SaveEmployee(w http.ResponseWriter, r *http.Request) {
	platform := r.URL.Query().Get("platform")
	employee := new(models.Employee)
	ctx := app.GetApp(r.Context(), h.Service.Daos)
	defer ctx.Client.Disconnect(r.Context())
	err := json.NewDecoder(r.Body).Decode(&employee)
	if err != nil {
		response.With400V2(w, err.Error(), platform)
		return
	}
	defer r.Body.Close()

	err = h.Service.SaveEmployee(ctx, employee)
	if err != nil {
		response.With500mV2(w, "failed - "+err.Error(), platform)
		return
	}
	m := make(map[string]interface{})
	m["employee"] = employee
	response.With200V2(w, "Success", m, platform)
}

//UpdateEmployee :""
func (h *Handler) UpdateEmployee(w http.ResponseWriter, r *http.Request) {

	platform := r.URL.Query().Get("platform")

	employee := new(models.Employee)
	var ctx *models.Context
	ctx = app.GetApp(r.Context(), h.Service.Daos)
	defer ctx.Client.Disconnect(r.Context())

	err := json.NewDecoder(r.Body).Decode(&employee)
	defer r.Body.Close()
	if err != nil {
		response.With400V2(w, err.Error(), platform)
		return
	}
	if employee.UniqueID == "" {
		response.With400V2(w, "id is missing", platform)
	}
	err = h.Service.UpdateEmployee(ctx, employee)
	if err != nil {
		response.With500mV2(w, "failed - "+err.Error(), platform)
		return
	}
	m := make(map[string]interface{})
	m["data"] = "success"
	response.With200V2(w, "Success", m, platform)
}

//EnableEmployee : ""
func (h *Handler) EnableEmployee(w http.ResponseWriter, r *http.Request) {

	platform := r.URL.Query().Get("platform")
	UniqueID := r.URL.Query().Get("id")

	if UniqueID == "" {
		response.With400V2(w, "id is missing", platform)
	}

	var ctx *models.Context
	ctx = app.GetApp(r.Context(), h.Service.Daos)
	defer ctx.Client.Disconnect(r.Context())

	err := h.Service.EnableEmployee(ctx, UniqueID)
	if err != nil {
		response.With500mV2(w, "failed - "+err.Error(), platform)
		return
	}
	m := make(map[string]interface{})
	m["data"] = "success"
	response.With200V2(w, "Success", m, platform)
}

//DisableEmployee : ""
func (h *Handler) DisableEmployee(w http.ResponseWriter, r *http.Request) {

	platform := r.URL.Query().Get("platform")

	UniqueID := r.URL.Query().Get("id")

	if UniqueID == "" {
		response.With400V2(w, "id is missing", platform)
	}
	var ctx *models.Context
	ctx = app.GetApp(r.Context(), h.Service.Daos)
	defer ctx.Client.Disconnect(r.Context())

	err := h.Service.DisableEmployee(ctx, UniqueID)
	if err != nil {
		response.With500mV2(w, "failed - "+err.Error(), platform)
		return
	}
	m := make(map[string]interface{})
	m["data"] = "success"
	response.With200V2(w, "Success", m, platform)
}

//DeleteEmployee : ""
func (h *Handler) DeleteEmployee(w http.ResponseWriter, r *http.Request) {

	platform := r.URL.Query().Get("platform")

	UniqueID := r.URL.Query().Get("id")

	if UniqueID == "" {
		response.With400V2(w, "id is missing", platform)
	}

	var ctx *models.Context
	ctx = app.GetApp(r.Context(), h.Service.Daos)
	defer ctx.Client.Disconnect(r.Context())

	err := h.Service.DeleteEmployee(ctx, UniqueID)
	if err != nil {
		response.With500mV2(w, "failed - "+err.Error(), platform)
		return
	}
	m := make(map[string]interface{})
	m["data"] = "success"
	response.With200V2(w, "Success", m, platform)
}

//GetSingleEmployee :""
func (h *Handler) GetSingleEmployee(w http.ResponseWriter, r *http.Request) {

	platform := r.URL.Query().Get("platform")
	UniqueID := r.URL.Query().Get("id")

	if UniqueID == "" {
		response.With400V2(w, "id is missing", platform)
	}

	Employee := new(models.RefEmployee)
	var ctx *models.Context
	ctx = app.GetApp(r.Context(), h.Service.Daos)
	defer ctx.Client.Disconnect(r.Context())

	Employee, err := h.Service.GetSingleEmployee(ctx, UniqueID)
	if err != nil {
		if err.Error() == "mongo: no documents in result" {
			response.With500mV2(w, "failed no data for this id", platform)
			return
		}
		response.With500mV2(w, "failed - "+err.Error(), platform)
		return
	}
	m := make(map[string]interface{})
	m["data"] = Employee
	response.With200V2(w, "Success", m, platform)
}

//FilterEmployee : ""
func (h *Handler) FilterEmployee(w http.ResponseWriter, r *http.Request) {

	platform := r.URL.Query().Get("platform")
	var employee *models.FilterEmployee
	var ctx *models.Context
	ctx = app.GetApp(r.Context(), h.Service.Daos)
	defer ctx.Client.Disconnect(r.Context())
	pageNo := r.URL.Query().Get("pageno")
	Limit := r.URL.Query().Get("limit")

	var pagination *models.Pagination
	if pageNo != "no" {
		pagination = new(models.Pagination)
		if pagination.PageNum = 1; pageNo != "" {
			page, err := strconv.Atoi(pageNo)
			if pagination.PageNum = 1; err == nil {
				pagination.PageNum = page
			}
		}
		if pagination.Limit = 10; Limit != "" {
			limit, err := strconv.Atoi(Limit)
			if pagination.Limit = 10; err == nil {
				pagination.Limit = limit
			}
		}
	}
	err := json.NewDecoder(r.Body).Decode(&employee)
	defer r.Body.Close()
	if err != nil {
		response.With400V2(w, err.Error(), platform)
		return
	}

	var employees []models.RefEmployee
	log.Println(pagination)
	employees, err = h.Service.FilterEmployee(ctx, employee, pagination)
	if err != nil {
		response.With500mV2(w, "failed - "+err.Error(), platform)
		return
	}
	m := make(map[string]interface{})

	if len(employees) > 0 {
		m["data"] = employees
	} else {
		res := make([]models.Employee, 0)
		m["data"] = res
	}
	if pagination != nil {
		if pagination.PageNum > 0 {
			m["pagination"] = pagination
		}
	}

	response.With200V2(w, "Success", m, platform)
}

//EmployeeReject :""
func (h *Handler) EmployeeReject(w http.ResponseWriter, r *http.Request) {

	platform := r.URL.Query().Get("platform")

	employeereject := new(models.EmployeeMoveToReject)
	var ctx *models.Context
	ctx = app.GetApp(r.Context(), h.Service.Daos)
	defer ctx.Client.Disconnect(r.Context())

	err := json.NewDecoder(r.Body).Decode(&employeereject)
	defer r.Body.Close()
	if err != nil {
		response.With400V2(w, err.Error(), platform)
		return
	}
	if employeereject.EmployeeID == "" {
		response.With400V2(w, "id is missing", platform)
	}
	err = h.Service.EmployeeReject(ctx, employeereject)
	if err != nil {
		response.With500mV2(w, "failed - "+err.Error(), platform)
		return
	}
	m := make(map[string]interface{})
	m["data"] = "success"
	response.With200V2(w, "Success", m, platform)
}

//EmployeeReject :""
func (h *Handler) EmployeeOnboarding(w http.ResponseWriter, r *http.Request) {

	platform := r.URL.Query().Get("platform")

	employeeonboarding := new(models.EmployeeMoveToOnboarding)
	var ctx *models.Context
	ctx = app.GetApp(r.Context(), h.Service.Daos)
	defer ctx.Client.Disconnect(r.Context())

	err := json.NewDecoder(r.Body).Decode(&employeeonboarding)
	defer r.Body.Close()
	if err != nil {
		response.With400V2(w, err.Error(), platform)
		return
	}
	if employeeonboarding.EmployeeID == "" {
		response.With400V2(w, "id is missing", platform)
	}
	err = h.Service.EmployeeOnboarding(ctx, employeeonboarding)
	if err != nil {
		response.With500mV2(w, "failed - "+err.Error(), platform)
		return
	}
	m := make(map[string]interface{})
	m["data"] = "success"
	response.With200V2(w, "Success", m, platform)
}

//EmployeeProbationary :""
func (h *Handler) EmployeeProbationary(w http.ResponseWriter, r *http.Request) {

	platform := r.URL.Query().Get("platform")

	employeeprobationary := new(models.EmployeeMoveToProbationary)
	var ctx *models.Context
	ctx = app.GetApp(r.Context(), h.Service.Daos)
	defer ctx.Client.Disconnect(r.Context())

	err := json.NewDecoder(r.Body).Decode(&employeeprobationary)
	defer r.Body.Close()
	if err != nil {
		response.With400V2(w, err.Error(), platform)
		return
	}
	if employeeprobationary.EmployeeID == "" {
		response.With400V2(w, "id is missing", platform)
	}
	err = h.Service.EmployeeProbationary(ctx, employeeprobationary)
	if err != nil {
		response.With500mV2(w, "failed - "+err.Error(), platform)
		return
	}
	m := make(map[string]interface{})
	m["data"] = "success"
	response.With200V2(w, "Success", m, platform)
}

//EmployeeActive :""
func (h *Handler) EmployeeActive(w http.ResponseWriter, r *http.Request) {

	platform := r.URL.Query().Get("platform")

	employeeactive := new(models.EmployeeMoveToActive)
	var ctx *models.Context
	ctx = app.GetApp(r.Context(), h.Service.Daos)
	defer ctx.Client.Disconnect(r.Context())

	err := json.NewDecoder(r.Body).Decode(&employeeactive)
	defer r.Body.Close()
	if err != nil {
		response.With400V2(w, err.Error(), platform)
		return
	}
	if employeeactive.EmployeeID == "" {
		response.With400V2(w, "id is missing", platform)
	}
	err = h.Service.EmployeeActive(ctx, employeeactive)
	if err != nil {
		response.With500mV2(w, "failed - "+err.Error(), platform)
		return
	}
	m := make(map[string]interface{})
	m["data"] = "success"
	response.With200V2(w, "Success", m, platform)
}

//EmployeeBench :""
func (h *Handler) EmployeeBench(w http.ResponseWriter, r *http.Request) {

	platform := r.URL.Query().Get("platform")

	employeebench := new(models.EmployeeMoveToBench)
	var ctx *models.Context
	ctx = app.GetApp(r.Context(), h.Service.Daos)
	defer ctx.Client.Disconnect(r.Context())

	err := json.NewDecoder(r.Body).Decode(&employeebench)
	defer r.Body.Close()
	if err != nil {
		response.With400V2(w, err.Error(), platform)
		return
	}
	if employeebench.EmployeeID == "" {
		response.With400V2(w, "id is missing", platform)
	}
	err = h.Service.EmployeeBench(ctx, employeebench)
	if err != nil {
		response.With500mV2(w, "failed - "+err.Error(), platform)
		return
	}
	m := make(map[string]interface{})
	m["data"] = "success"
	response.With200V2(w, "Success", m, platform)
}

//EmployeeNotice :""
func (h *Handler) EmployeeNotice(w http.ResponseWriter, r *http.Request) {

	platform := r.URL.Query().Get("platform")

	employeenotice := new(models.EmployeeMoveToNotice)
	var ctx *models.Context
	ctx = app.GetApp(r.Context(), h.Service.Daos)
	defer ctx.Client.Disconnect(r.Context())

	err := json.NewDecoder(r.Body).Decode(&employeenotice)
	defer r.Body.Close()
	if err != nil {
		response.With400V2(w, err.Error(), platform)
		return
	}
	if employeenotice.EmployeeID == "" {
		response.With400V2(w, "id is missing", platform)
	}
	err = h.Service.EmployeeNotice(ctx, employeenotice)
	if err != nil {
		response.With500mV2(w, "failed - "+err.Error(), platform)
		return
	}
	m := make(map[string]interface{})
	m["data"] = "success"
	response.With200V2(w, "Success", m, platform)
}

//EmployeeOffboard :""
func (h *Handler) EmployeeOffboard(w http.ResponseWriter, r *http.Request) {

	platform := r.URL.Query().Get("platform")

	employeeoffboard := new(models.EmployeeMoveToOffboard)
	var ctx *models.Context
	ctx = app.GetApp(r.Context(), h.Service.Daos)
	defer ctx.Client.Disconnect(r.Context())

	err := json.NewDecoder(r.Body).Decode(&employeeoffboard)
	defer r.Body.Close()
	if err != nil {
		response.With400V2(w, err.Error(), platform)
		return
	}
	if employeeoffboard.EmployeeID == "" {
		response.With400V2(w, "id is missing", platform)
	}
	err = h.Service.EmployeeOffboard(ctx, employeeoffboard)
	if err != nil {
		response.With500mV2(w, "failed - "+err.Error(), platform)
		return
	}
	m := make(map[string]interface{})
	m["data"] = "success"
	response.With200V2(w, "Success", m, platform)
}

//EmployeeRelieve :""
func (h *Handler) EmployeeRelieve(w http.ResponseWriter, r *http.Request) {

	platform := r.URL.Query().Get("platform")

	employeerelieve := new(models.EmployeeMoveToRelieve)
	var ctx *models.Context
	ctx = app.GetApp(r.Context(), h.Service.Daos)
	defer ctx.Client.Disconnect(r.Context())

	err := json.NewDecoder(r.Body).Decode(&employeerelieve)
	defer r.Body.Close()
	if err != nil {
		response.With400V2(w, err.Error(), platform)
		return
	}
	if employeerelieve.EmployeeID == "" {
		response.With400V2(w, "id is missing", platform)
	}
	err = h.Service.EmployeeRelieve(ctx, employeerelieve)
	if err != nil {
		response.With500mV2(w, "failed - "+err.Error(), platform)
		return
	}
	m := make(map[string]interface{})
	m["data"] = "success"
	response.With200V2(w, "Success", m, platform)
}

func (h *Handler) UpdateEmployeeBioData(w http.ResponseWriter, r *http.Request) {

	platform := r.URL.Query().Get("platform")

	employee := new(models.UpdateBioData)
	var ctx *models.Context
	ctx = app.GetApp(r.Context(), h.Service.Daos)
	defer ctx.Client.Disconnect(r.Context())
	UniqueID := r.URL.Query().Get("id")

	err := json.NewDecoder(r.Body).Decode(&employee)
	defer r.Body.Close()
	if err != nil {
		response.With400V2(w, err.Error(), platform)
		return
	}
	if UniqueID == "" {
		response.With400V2(w, "id is missing", platform)
	}
	err = h.Service.UpdateEmployeeBioData(ctx, employee, UniqueID)
	if err != nil {
		response.With500mV2(w, "failed - "+err.Error(), platform)
		return
	}
	m := make(map[string]interface{})
	m["data"] = "success"
	response.With200V2(w, "Success", m, platform)
}
func (h *Handler) UpdateEmployeeEmergencyContact(w http.ResponseWriter, r *http.Request) {

	platform := r.URL.Query().Get("platform")

	employee := new(models.UpdateEmergencyContact)
	var ctx *models.Context
	ctx = app.GetApp(r.Context(), h.Service.Daos)
	defer ctx.Client.Disconnect(r.Context())
	UniqueID := r.URL.Query().Get("id")

	err := json.NewDecoder(r.Body).Decode(&employee)
	defer r.Body.Close()
	if err != nil {
		response.With400V2(w, err.Error(), platform)
		return
	}
	if UniqueID == "" {
		response.With400V2(w, "id is missing", platform)
	}
	err = h.Service.UpdateEmployeeEmergencyContact(ctx, employee, UniqueID)
	if err != nil {
		response.With500mV2(w, "failed - "+err.Error(), platform)
		return
	}
	m := make(map[string]interface{})
	m["data"] = "success"
	response.With200V2(w, "Success", m, platform)
}
func (h *Handler) UpdateEmployeePersonalInformation(w http.ResponseWriter, r *http.Request) {

	platform := r.URL.Query().Get("platform")

	employee := new(models.PersonalInformation)
	var ctx *models.Context
	ctx = app.GetApp(r.Context(), h.Service.Daos)
	defer ctx.Client.Disconnect(r.Context())
	UniqueID := r.URL.Query().Get("id")

	err := json.NewDecoder(r.Body).Decode(&employee)
	defer r.Body.Close()
	if err != nil {
		response.With400V2(w, err.Error(), platform)
		return
	}
	if UniqueID == "" {
		response.With400V2(w, "id is missing", platform)
	}
	err = h.Service.UpdateEmployeePersonalInformation(ctx, employee, UniqueID)
	if err != nil {
		response.With500mV2(w, "failed - "+err.Error(), platform)
		return
	}
	m := make(map[string]interface{})
	m["data"] = "success"
	response.With200V2(w, "Success", m, platform)
}
func (h *Handler) EmployeeDayWiseAttendanceReport(w http.ResponseWriter, r *http.Request) {

	platform := r.URL.Query().Get("platform")
	var filter *models.DayWiseAttendanceReportFilter
	var ctx *models.Context
	ctx = app.GetApp(r.Context(), h.Service.Daos)
	pageNo := r.URL.Query().Get("pageno")
	Limit := r.URL.Query().Get("limit")

	var pagination *models.Pagination
	if pageNo != "no" {
		pagination = new(models.Pagination)
		if pagination.PageNum = 1; pageNo != "" {
			page, err := strconv.Atoi(pageNo)
			if pagination.PageNum = 1; err == nil {
				pagination.PageNum = page
			}
		}
		if pagination.Limit = 10; Limit != "" {
			limit, err := strconv.Atoi(Limit)
			if pagination.Limit = 10; err == nil {
				pagination.Limit = limit
			}
		}
	}
	err := json.NewDecoder(r.Body).Decode(&filter)
	defer r.Body.Close()
	if err != nil {
		response.With400V2(w, err.Error(), platform)
		return
	}
	resType := r.URL.Query().Get("resType")
	if resType == "reportexcel" {
		file, err := h.Service.EmployeeDayWiseAttendanceReportExcel(ctx, filter)
		if err != nil {
			response.With500mV2(w, "failed - "+err.Error(), platform)
			return
		}
		w.Header().Set("Content-Type", "application/octet-stream")
		w.Header().Set("Content-Disposition", "attachment; filename=EmployeesDaywiseReport.xlsx")
		w.Header().Set("ocntent-Transfer-Encoding", "binary")
		file.Write(w)
		return
	}

	var attendance []models.EmployeeDayWiseAttendanceReport

	attendance, err = h.Service.EmployeeDayWiseAttendanceReport(ctx, filter, pagination)
	if err != nil {
		response.With500mV2(w, "failed - "+err.Error(), platform)
		return
	}
	m := make(map[string]interface{})

	if len(attendance) > 0 {
		m["employee"] = attendance
	} else {
		res := make([]models.Attendance, 0)
		m["employee"] = res
	}
	if pagination != nil {
		if pagination.PageNum > 0 {
			m["pagination"] = pagination
		}
	}

	response.With200V2(w, "Success", m, platform)
}
func (h *Handler) UpdateEmployeeProfileImage(w http.ResponseWriter, r *http.Request) {

	platform := r.URL.Query().Get("platform")

	employee := new(models.UpdateBioData)
	var ctx *models.Context
	ctx = app.GetApp(r.Context(), h.Service.Daos)
	defer ctx.Client.Disconnect(r.Context())
	UniqueID := r.URL.Query().Get("id")

	err := json.NewDecoder(r.Body).Decode(&employee)
	defer r.Body.Close()
	if err != nil {
		response.With400V2(w, err.Error(), platform)
		return
	}
	if UniqueID == "" {
		response.With400V2(w, "id is missing", platform)
	}
	err = h.Service.UpdateEmployeeProfileImage(ctx, employee, UniqueID)
	if err != nil {
		response.With500mV2(w, "failed - "+err.Error(), platform)
		return
	}
	m := make(map[string]interface{})
	m["data"] = "success"
	response.With200V2(w, "Success", m, platform)
}
func (h *Handler) EmployeeUpload(w http.ResponseWriter, r *http.Request) {
	platform := r.URL.Query().Get("platform")
	file, _, err := r.FormFile("file")
	if err != nil {
		fmt.Fprintf(w, err.Error())
		return
	}
	defer file.Close()
	var ctx *models.Context
	ctx = app.GetApp(r.Context(), h.Service.Daos)
	defer ctx.Client.Disconnect(r.Context())
	defer r.Body.Close()

	farmer := h.Service.EmployeeUploadExcel(ctx, file)
	if err != nil {
		response.With500mV2(w, "failed - "+err.Error(), platform)
		return
	}
	m := make(map[string]interface{})
	m["EmployeeUpload"] = farmer
	response.With200V2(w, "Success", m, platform)
}
func (h *Handler) EmployeeUploadV2(w http.ResponseWriter, r *http.Request) {
	platform := r.URL.Query().Get("platform")
	file, _, err := r.FormFile("file")
	if err != nil {
		fmt.Fprintf(w, err.Error())
		return
	}
	defer file.Close()
	var ctx *models.Context
	ctx = app.GetApp(r.Context(), h.Service.Daos)
	defer ctx.Client.Disconnect(r.Context())
	defer r.Body.Close()

	farmer := h.Service.EmployeeUploadExcelV2(ctx, file)
	if err != nil {
		response.With500mV2(w, "failed - "+err.Error(), platform)
		return
	}
	m := make(map[string]interface{})
	m["EmployeeUpload"] = farmer
	response.With200V2(w, "Success", m, platform)
}
func (h *Handler) DashboardEmployeeCount(w http.ResponseWriter, r *http.Request) {

	platform := r.URL.Query().Get("platform")
	var employee *models.FilterEmployee
	var ctx *models.Context
	ctx = app.GetApp(r.Context(), h.Service.Daos)
	defer ctx.Client.Disconnect(r.Context())

	err := json.NewDecoder(r.Body).Decode(&employee)
	defer r.Body.Close()
	if err != nil {
		response.With400V2(w, err.Error(), platform)
		return
	}
	var employees []models.DashboardEmployeeCount
	employees, err = h.Service.DashboardEmployeeCount(ctx, employee)
	if err != nil {
		response.With500mV2(w, "failed - "+err.Error(), platform)
		return
	}
	m := make(map[string]interface{})

	if len(employees) > 0 {
		m["data"] = employees
	} else {
		res := make([]models.Employee, 0)
		m["data"] = res
	}
	response.With200V2(w, "Success", m, platform)
}
func (h *Handler) GetEmployeeChild(w http.ResponseWriter, r *http.Request) {

	platform := r.URL.Query().Get("platform")
	var ctx *models.Context
	ctx = app.GetApp(r.Context(), h.Service.Daos)
	defer ctx.Client.Disconnect(r.Context())
	//Employee := new(models.EmployeeTreev2)
	Employees, err := h.Service.FindChild(ctx)
	if err != nil {
		if err.Error() == "mongo: no documents in result" {
			response.With500mV2(w, "failed no data for this id", platform)
			return
		}
		response.With500mV2(w, "failed - "+err.Error(), platform)
		return
	}
	m := make(map[string]interface{})
	m["data"] = Employees
	response.With200V2(w, "Success", m, platform)
}
func (h *Handler) GetOrgChart(w http.ResponseWriter, r *http.Request) {

	platform := r.URL.Query().Get("platform")
	var ctx *models.Context
	ctx = app.GetApp(r.Context(), h.Service.Daos)
	defer ctx.Client.Disconnect(r.Context())
	//Employee := new(models.EmployeeTreev2)
	Employees, err := h.Service.GetAllOrgChart(ctx)
	if err != nil {
		if err.Error() == "mongo: no documents in result" {
			response.With500mV2(w, "failed no data for this id", platform)
			return
		}
		response.With500mV2(w, "failed - "+err.Error(), platform)
		return
	}
	m := make(map[string]interface{})
	m["data"] = Employees
	response.With200V2(w, "Success", m, platform)
}
func (h *Handler) GetLineManagerEmployee(w http.ResponseWriter, r *http.Request) {

	platform := r.URL.Query().Get("platform")
	UniqueID := r.URL.Query().Get("id")

	// if UniqueID == "" {
	// 	response.With400V2(w, "id is missing", platform)
	// }

	//	Employee := new(models.EmployeeTree)
	var ctx *models.Context
	ctx = app.GetApp(r.Context(), h.Service.Daos)
	defer ctx.Client.Disconnect(r.Context())

	Employees, err := h.Service.GetLineManagerEmployee(ctx, UniqueID)
	if err != nil {
		if err.Error() == "mongo: no documents in result" {
			response.With500mV2(w, "failed no data for this id", platform)
			return
		}
		response.With500mV2(w, "failed - "+err.Error(), platform)
		return
	}
	m := make(map[string]interface{})
	m["data"] = Employees
	response.With200V2(w, "Success", m, platform)
}
func (h *Handler) GetEmployeeLinemanagerCheck(w http.ResponseWriter, r *http.Request) {

	platform := r.URL.Query().Get("platform")
	UniqueID := r.URL.Query().Get("id")

	if UniqueID == "" {
		response.With400V2(w, "id is missing", platform)
	}

	//	Employee := new(models.EmployeeTree)
	var ctx *models.Context
	ctx = app.GetApp(r.Context(), h.Service.Daos)
	defer ctx.Client.Disconnect(r.Context())

	Employees, err := h.Service.GetEmployeeLinemanagerCheck(ctx, UniqueID)
	if err != nil {
		if err.Error() == "mongo: no documents in result" {
			response.With500mV2(w, "failed no data for this id", platform)
			return
		}
		response.With500mV2(w, "failed - "+err.Error(), platform)
		return
	}
	m := make(map[string]interface{})
	m["data"] = Employees
	response.With200V2(w, "Success", m, platform)
}
func (h *Handler) EmployeeUpdateLoginId(w http.ResponseWriter, r *http.Request) {
	platform := r.URL.Query().Get("platform")
	file, _, err := r.FormFile("file")
	if err != nil {
		fmt.Fprintf(w, err.Error())
		return
	}
	defer file.Close()
	var ctx *models.Context
	ctx = app.GetApp(r.Context(), h.Service.Daos)
	defer ctx.Client.Disconnect(r.Context())
	defer r.Body.Close()

	farmer := h.Service.EmployeeUpdateLoginId(ctx, file)
	if err != nil {
		response.With500mV2(w, "failed - "+err.Error(), platform)
		return
	}
	m := make(map[string]interface{})
	m["EmployeeUpload"] = farmer
	response.With200V2(w, "Success", m, platform)
}
