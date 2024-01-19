package services

import (
	"context"
	"errors"
	"fmt"

	"hrms-services/app"
	"hrms-services/constants"
	"hrms-services/models"
	"log"
	"time"

	"github.com/360EntSecGroup-Skylar/excelize"
	"go.mongodb.org/mongo-driver/mongo"
)

//SavePayrollLog :""
func (s *Service) SavePayrollLog(ctx *models.Context, payrollLog *models.PayrollLog) error {
	log.Println("transaction start")
	//Start Transaction
	if err := ctx.Session.StartTransaction(); err != nil {
		return err
	}
	defer ctx.Session.EndSession(ctx.CTX)
	payrollLog.UniqueID = s.Daos.GetUniqueID(ctx, constants.COLLECTIONPAYROLLLOG)
	payrollLog.Status = constants.PAYROLLLOGSTATUSACTIVE
	t := time.Now()
	created := models.Created{}
	created.On = &t
	created.By = constants.SYSTEM
	log.Println("b4 PayrollLog.created")
	payrollLog.Created = &created
	log.Println("b4 PayrollLog.created")
	if err := mongo.WithSession(ctx.CTX, ctx.Session, func(sc mongo.SessionContext) error {
		dberr := s.Daos.SavePayrollLog(ctx, payrollLog)
		if dberr != nil {
			if err1 := ctx.Session.AbortTransaction(sc); err1 != nil {
				log.Println("err in abort")
				return errors.New("Transaction Aborted with error" + err1.Error())
			}
			log.Println("err in abort out")
			return errors.New("Transaction Aborted - " + dberr.Error())
		}
		return nil

	}); err != nil {
		return err
	}

	return nil
}

//UpdatePayrollLog : ""
func (s *Service) UpdatePayrollLog(ctx *models.Context, payrollLog *models.PayrollLog) error {
	if err := ctx.Session.StartTransaction(); err != nil {
		return err
	}
	defer ctx.Session.EndSession(ctx.CTX)
	if err := mongo.WithSession(ctx.CTX, ctx.Session, func(sc mongo.SessionContext) error {

		err := s.Daos.UpdatePayrollLog(ctx, payrollLog)
		if err != nil {
			if err = ctx.Session.AbortTransaction(sc); err != nil {
				return errors.New("Transaction Aborted with error" + err.Error())
			}
			return errors.New("Transaction Aborted - " + err.Error())
		}
		return nil

	}); err != nil {
		return err
	}
	return nil
}

//EnablePayrollLog : ""
func (s *Service) EnablePayrollLog(ctx *models.Context, UniqueID string) error {
	if err := ctx.Session.StartTransaction(); err != nil {
		return err
	}
	defer ctx.Session.EndSession(ctx.CTX)
	if err := mongo.WithSession(ctx.CTX, ctx.Session, func(sc mongo.SessionContext) error {

		err := s.Daos.EnablePayrollLog(ctx, UniqueID)
		if err != nil {
			if err = ctx.Session.AbortTransaction(sc); err != nil {
				return errors.New("Transaction Aborted with error" + err.Error())
			}
			return errors.New("Transaction Aborted - " + err.Error())
		}
		return nil

	}); err != nil {
		return err
	}
	return nil
}

//DisablePayrollLog : ""
func (s *Service) DisablePayrollLog(ctx *models.Context, UniqueID string) error {
	if err := ctx.Session.StartTransaction(); err != nil {
		return err
	}
	defer ctx.Session.EndSession(ctx.CTX)
	if err := mongo.WithSession(ctx.CTX, ctx.Session, func(sc mongo.SessionContext) error {

		err := s.Daos.DisablePayrollLog(ctx, UniqueID)
		if err != nil {
			if err = ctx.Session.AbortTransaction(sc); err != nil {
				return errors.New("Transaction Aborted with error" + err.Error())
			}
			return errors.New("Transaction Aborted - " + err.Error())
		}
		return nil

	}); err != nil {
		return err
	}
	return nil
}

//DeletePayrollLog : ""
func (s *Service) DeletePayrollLog(ctx *models.Context, UniqueID string) error {
	if err := ctx.Session.StartTransaction(); err != nil {
		return err
	}
	defer ctx.Session.EndSession(ctx.CTX)
	if err := mongo.WithSession(ctx.CTX, ctx.Session, func(sc mongo.SessionContext) error {

		err := s.Daos.DeletePayrollLog(ctx, UniqueID)
		if err != nil {
			if err = ctx.Session.AbortTransaction(sc); err != nil {
				return errors.New("Transaction Aborted with error" + err.Error())
			}
			return errors.New("Transaction Aborted - " + err.Error())
		}
		return nil

	}); err != nil {
		return err
	}
	return nil
}

//GetSinglePayrollLog :""
func (s *Service) GetSinglePayrollLog(ctx *models.Context, UniqueID string) (*models.RefPayrollLog, error) {
	payrollLog, err := s.Daos.GetSinglePayrollLog(ctx, UniqueID)
	if err != nil {
		return nil, err
	}
	return payrollLog, nil
}

//FilterPayrollLog :""
func (s *Service) FilterPayrollLog(ctx *models.Context, filter *models.FilterPayrollLog, pagination *models.Pagination) ([]models.RefPayrollLog, error) {
	defer ctx.Session.EndSession(ctx.CTX)
	err := s.PayrollLogDataAccess(ctx, filter)
	if err != nil {
		return nil, err
	}

	return s.Daos.FilterPayrollLog(ctx, filter, pagination)

}
func (s *Service) PayrollLogDataAccess(ctx *models.Context, filter *models.FilterPayrollLog) (err error) {
	if filter != nil {
		dataaccess, err := s.Daos.DataAccess(ctx, &filter.DataAccess)
		if err != nil {
			return err
		}
		if dataaccess != nil {
			if len(dataaccess.Organisation) > 0 {
				for _, v := range dataaccess.Organisation {
					filter.OrganisationID = append(filter.OrganisationID, v.UniqueID)
				}
			}

		}

	}
	return err
}
func (s *Service) GetEmployeeSalarySlip(ctx *models.Context, UniqueID *models.EmployeeSalarySlip) (*models.RefPayrollLog, error) {
	payrollLog, err := s.Daos.GetEmployeeSalarySlip(ctx, UniqueID)
	if err != nil {
		return nil, err
	}
	//	t := time.Now()
	if payrollLog != nil {
		filter := new(models.DayWiseAttendanceReportFilter)
		filter.EmployeeId = UniqueID.EmployeeId
		filter.StartDate = UniqueID.StartDate
		report, err := s.EmployeeDayWiseAttendanceReport(ctx, filter, nil)
		if err != nil {
			return nil, err
		}
		fmt.Println("report=======>", len(report))
		year := UniqueID.StartDate.Year()
		month := UniqueID.StartDate.Month()
		Yearofmonth := fmt.Sprintf("%v%v", year, int(month))
		paySlipId := UniqueID.EmployeeId + Yearofmonth + s.Daos.GetUniqueID(ctx, constants.COLLECTIONEMPLOYEEPAYSLIP)

		//	t.Month()=
		t1 := time.Date(UniqueID.StartDate.Year(), UniqueID.StartDate.Month()+1, 0, 23, 59, 59, 999999999, filter.StartDate.Location())
		//t2 := time.Date(t.Year(), t.Month(), 0, 0, 0, 0, 0, t.Location())
		if len(report) > 0 {
			fmt.Println("report[0].NoOfHolidays=======>", report[0].NoOfHolidays)
			days := t1.Day() - int(report[0].NoOfHolidays)
			lop := (payrollLog.NetAmount / float64(days)) * report[0].NoOfLOP
			PartialPay := (payrollLog.NetAmount / float64(days) / 2) * report[0].NoOfParticalPaid
			fmt.Println("no.of.days=======>", days)
			fmt.Println("report[0].NoOfLOP", report[0].NoOfLOP)
			fmt.Println("lop", lop)
			fmt.Println("PartialPay", PartialPay)
			payrollLog.Deduction = payrollLog.Deduction + lop + PartialPay
			payrollLog.Detections.Lop = lop + PartialPay
			fmt.Println("payrollLog.Deduction", payrollLog.Deduction)
			fmt.Println("payrollLog.NetAmount", payrollLog.NetAmount)
			payrollLog.NetAmount = payrollLog.NetAmount - lop - PartialPay
			fmt.Println("payrollLog.NetAmountafter", payrollLog.NetAmount)

		} else {
			days := t1.Day()
			fmt.Println("days", days)

			lop := payrollLog.NetAmount
			payrollLog.Deduction = payrollLog.Deduction + lop
			payrollLog.Detections.Lop = lop
			fmt.Println("payrollLog.Deduction", payrollLog.Deduction)
			fmt.Println("payrollLog.NetAmount", payrollLog.NetAmount)
			payrollLog.NetAmount = payrollLog.NetAmount - lop
			fmt.Println("payrollLog.NetAmountafter", payrollLog.NetAmount)
			//PartialPay := (payrollLog.NetAmount / float64(days) / 2)
		}
		payrollLog.AmountWord = s.Shared.Convert(int(payrollLog.NetAmount)) + "Rupess only"
		payrollLog.PayslipId = paySlipId

	} else {
		fmt.Println("PayrollNotFounded======>", UniqueID.EmployeeId)
		return nil, nil
	}
	return payrollLog, nil
}
func (s *Service) EmployeePayrollExcel(ctx *models.Context, UniqueID *models.EmployeeSalarySlip) (*excelize.File, error) {
	t := time.Now()
	data, err := s.GetEmployeeSalarySlip(ctx, UniqueID)
	if err != nil {
		return nil, err
	}
	filter := new(models.DayWiseAttendanceReportFilter)
	filter.EmployeeId = UniqueID.EmployeeId
	if UniqueID.StartDate != nil {
		filter.StartDate = UniqueID.StartDate
	}
	report, err := s.EmployeeDayWiseAttendanceReport(ctx, filter, nil)
	if err != nil {
		return nil, err
	}
	//	t.Month()=
	t1 := time.Date(UniqueID.StartDate.Year(), UniqueID.StartDate.Month()+1, 0, 23, 59, 59, 999999999, UniqueID.StartDate.Location())
	//t2 := time.Date(t.Year(), t.Month(), 0, 0, 0, 0, 0, t.Location())\
	days := 0
	if len(report) < 0 {

	}
	if report[0].NoOfHolidays > 0 {
		days = t1.Day() - int(report[0].NoOfHolidays)
	} else {
		days = t1.Day()
	}
	fmt.Println("no.of.days=======>", days)
	fmt.Println("NetAmount=======>", data.NetAmount)
	fmt.Println("Deduction=======>", data.Deduction)
	duration := time.Since(t)
	log.Println("query Time taken ===> ", duration.Minutes(), "m")
	t = time.Now()
	//fmt.Println("data=====>", data)
	excel := excelize.NewFile()
	sheet1 := "Salary Slip"
	index := excel.NewSheet(sheet1)
	rowNo := 1

	excel.SetActiveSheet(index)
	excel.MergeCell(sheet1, "B1", "E1")
	// excel.MergeCell(sheet1, "C1", "C3")
	// excel.MergeCell(sheet1, "C4", "C5")
	style1, err := excel.NewStyle(`{"fill":{"type":"pattern","color":["#FFFF00"],"pattern":1},"alignment":{"horizontal":"center","vertical":"center"},"font":{"bold":true}}`)
	if err != nil {
		fmt.Println(err)
	}
	style3, err := excel.NewStyle(`{"fill":{"type":"pattern","color":["#B6DDE8"],"pattern":1},"alignment":{"horizontal":"center","vertical":"center"},"font":{"bold":true}}`)
	if err != nil {
		fmt.Println(err)
	}
	style4, err := excel.NewStyle(`{"fill":{"type":"pattern","pattern":1},"alignment":{"horizontal":"center","vertical":"center"},"font":{"bold":true}}`)
	if err != nil {
		fmt.Println(err)
	}
	// documentUrl := s.Shared.Config.GetString(s.Shared.GetCmdArg(constants.ENV) + "." + constants.FILEURL)
	// if err := excel.AddPicture(sheet1, fmt.Sprintf("%v%v", "A", rowNo), documentUrl+"municipal/logo.png", `{"x_scale": 0.115, "y_scale": 0.0935}`); err != nil {
	// 	fmt.Println(err)
	// }
	// excel.SetCellStyle(sheet1, fmt.Sprintf("%v%v", "C", rowNo), fmt.Sprintf("%v%v", "C", rowNo), style3)
	// excel.SetCellValue(sheet1, fmt.Sprintf("%v%v", "C", rowNo), ctx.ProductConfig.Name)
	// title :=
	excel.SetCellStyle(sheet1, fmt.Sprintf("%v%v", "B", rowNo), fmt.Sprintf("%v%v", "E", rowNo), style3)
	excel.SetCellValue(sheet1, fmt.Sprintf("%v%v", "C", rowNo), fmt.Sprintf("%v", sheet1))
	//Kitchen := "3:04PM"
	rowNo++
	rowNo++

	//	var totalAmount float64
	if data != nil {
		columnNo1 := 'A'
		columnNoC := 'C'
		columnNoE := 'E'
		rowNo1 := 3
		//columnNo2 := 'D'
		columnNo1++
		excel.SetCellValue(sheet1, fmt.Sprintf("%c%v", columnNo1, rowNo), "Employee Code:")
		excel.SetCellValue(sheet1, fmt.Sprintf("%c%v", columnNoC, rowNo), data.Ref.EmployeeId.UniqueID)
		rowNo++
		excel.SetCellValue(sheet1, fmt.Sprintf("%c%v", columnNo1, rowNo), "Employee Name:")
		excel.SetCellValue(sheet1, fmt.Sprintf("%c%v", columnNoC, rowNo), data.Ref.EmployeeId.Name)
		rowNo++
		excel.SetCellValue(sheet1, fmt.Sprintf("%c%v", columnNo1, rowNo), "Designation:")
		excel.SetCellValue(sheet1, fmt.Sprintf("%c%v", columnNoC, rowNo), data.Ref.DesignationID.Name)
		rowNo++
		excel.SetCellValue(sheet1, fmt.Sprintf("%c%v", columnNo1, rowNo), "Bank Account No:")
		excel.SetCellValue(sheet1, fmt.Sprintf("%c%v", columnNoC, rowNo), data.Ref.Bank.AccountNumber)
		rowNo++
		excel.SetCellValue(sheet1, fmt.Sprintf("%c%v", columnNo1, rowNo), "UAN No:")
		excel.SetCellValue(sheet1, fmt.Sprintf("%c%v", columnNoC, rowNo), "")
		rowNo++
		excel.SetCellValue(sheet1, fmt.Sprintf("%c%v", columnNo1, rowNo), "ESIC No:")
		excel.SetCellValue(sheet1, fmt.Sprintf("%c%v", columnNoC, rowNo), "")
		rowNo++
		excel.SetCellValue(sheet1, fmt.Sprintf("%c%v", columnNo1, rowNo), "Total Working Day in Month:")
		excel.SetCellValue(sheet1, fmt.Sprintf("%c%v", columnNoC, rowNo), days)
		rowNo++

		excel.SetCellStyle(sheet1, fmt.Sprintf("%v%v", "B", rowNo), fmt.Sprintf("%v%v", "E", rowNo), style1)
		excel.SetCellValue(sheet1, fmt.Sprintf("%c%v", columnNo1, rowNo), "Gross salary per month")
		rowNo++
		excel.SetCellValue(sheet1, fmt.Sprintf("%c%v", columnNo1, rowNo), "Components In Salary(Earnings):")
		excel.SetCellValue(sheet1, fmt.Sprintf("%c%v", columnNoC, rowNo), "Amount (in Rs.)")
		rowNo++
		excel.SetCellValue(sheet1, fmt.Sprintf("%c%v", columnNo1, rowNo), "Basic Salary")
		excel.SetCellValue(sheet1, fmt.Sprintf("%c%v", columnNoC, rowNo), data.Earnings.BasicSalary)
		rowNo++
		excel.SetCellValue(sheet1, fmt.Sprintf("%c%v", columnNo1, rowNo), "HRA")
		excel.SetCellValue(sheet1, fmt.Sprintf("%c%v", columnNoC, rowNo), data.Earnings.Hra)
		rowNo++
		excel.SetCellValue(sheet1, fmt.Sprintf("%c%v", columnNo1, rowNo), "Conveyance allowances")
		excel.SetCellValue(sheet1, fmt.Sprintf("%c%v", columnNoC, rowNo), data.Earnings.ConveyanceAllowances)
		rowNo++
		excel.SetCellValue(sheet1, fmt.Sprintf("%c%v", columnNo1, rowNo), "Education Allowance")
		excel.SetCellValue(sheet1, fmt.Sprintf("%c%v", columnNoC, rowNo), data.Earnings.EducationAllowance)
		rowNo++
		excel.SetCellValue(sheet1, fmt.Sprintf("%c%v", columnNo1, rowNo), "Performance Allowance")
		excel.SetCellValue(sheet1, fmt.Sprintf("%c%v", columnNoC, rowNo), data.Earnings.PerformanceAllowance)
		rowNo++
		excel.SetCellStyle(sheet1, fmt.Sprintf("%v%v", "B", rowNo), fmt.Sprintf("%v%v", "C", rowNo), style4)
		excel.SetCellValue(sheet1, fmt.Sprintf("%c%v", columnNo1, rowNo), "Total Gross Salary")
		excel.SetCellValue(sheet1, fmt.Sprintf("%c%v", columnNoC, rowNo), data.GrossAmount)
		rowNo++
		excel.SetCellStyle(sheet1, fmt.Sprintf("%v%v", "B", rowNo), fmt.Sprintf("%v%v", "E", rowNo), style4)
		excel.SetCellValue(sheet1, fmt.Sprintf("%c%v", columnNo1, rowNo), "Net Salary (Gross-Total deductions)")
		excel.SetCellValue(sheet1, fmt.Sprintf("%c%v", columnNoE, rowNo), data.NetAmount)

		columnNo1++
		columnNo1++
		excel.SetCellValue(sheet1, fmt.Sprintf("%c%v", columnNo1, rowNo1), "Date Of Birth:")
		excel.SetCellValue(sheet1, fmt.Sprintf("%c%v", columnNoE, rowNo1), fmt.Sprintf("%v-%v-%v", data.Ref.EmployeeId.DOB.Day(), data.Ref.EmployeeId.DOB.Month(), data.Ref.EmployeeId.DOB.Year()))
		rowNo1++
		excel.SetCellValue(sheet1, fmt.Sprintf("%c%v", columnNo1, rowNo1), "Date Of Joining:")
		excel.SetCellValue(sheet1, fmt.Sprintf("%c%v", columnNoE, rowNo1), fmt.Sprintf("%v-%v-%v", data.Ref.EmployeeId.JoiningDate.Day(), data.Ref.EmployeeId.JoiningDate.Month(), data.Ref.EmployeeId.JoiningDate.Year()))
		rowNo1++
		excel.SetCellValue(sheet1, fmt.Sprintf("%c%v", columnNo1, rowNo1), "Location:")
		excel.SetCellValue(sheet1, fmt.Sprintf("%c%v", columnNoE, rowNo1), "")
		rowNo1++
		excel.SetCellValue(sheet1, fmt.Sprintf("%c%v", columnNo1, rowNo1), "Bank Name:")
		excel.SetCellValue(sheet1, fmt.Sprintf("%c%v", columnNoE, rowNo1), data.Ref.Bank.BankName)
		rowNo1++
		excel.SetCellValue(sheet1, fmt.Sprintf("%c%v", columnNo1, rowNo1), "IFSC Code:")
		excel.SetCellValue(sheet1, fmt.Sprintf("%c%v", columnNoE, rowNo1), data.Ref.Bank.IFSC)
		rowNo1++
		excel.SetCellValue(sheet1, fmt.Sprintf("%c%v", columnNo1, rowNo1), "PAN Code:")
		rowNo1++
		excel.SetCellValue(sheet1, fmt.Sprintf("%c%v", columnNo1, rowNo1), "Absent Days in a Month:")
		excel.SetCellValue(sheet1, fmt.Sprintf("%c%v", columnNoE, rowNo1), report[0].NoOfLOP+(report[0].NoOfParticalPaid/2))
		rowNo1++
		excel.SetCellValue(sheet1, fmt.Sprintf("%c%v", columnNo1, rowNo1), data.GrossAmount)
		rowNo1++
		excel.SetCellValue(sheet1, fmt.Sprintf("%c%v", columnNo1, rowNo1), "Components In salary (Deduction)")
		excel.SetCellValue(sheet1, fmt.Sprintf("%c%v", columnNoE, rowNo1), "Amount (in Rs.)")
		rowNo1++
		excel.SetCellValue(sheet1, fmt.Sprintf("%c%v", columnNo1, rowNo1), "PF contribution")
		excel.SetCellValue(sheet1, fmt.Sprintf("%c%v", columnNoE, rowNo1), data.Detections.PfContribution)
		rowNo1++
		excel.SetCellValue(sheet1, fmt.Sprintf("%c%v", columnNo1, rowNo1), "ESIC contribution")
		excel.SetCellValue(sheet1, fmt.Sprintf("%c%v", columnNoE, rowNo1), data.Detections.ESICContribution)
		rowNo1++
		excel.SetCellValue(sheet1, fmt.Sprintf("%c%v", columnNo1, rowNo1), "LOP")
		excel.SetCellValue(sheet1, fmt.Sprintf("%c%v", columnNoE, rowNo1), data.Detections.Lop)

		rowNo1++
		rowNo1++
		rowNo1++
		excel.SetCellStyle(sheet1, fmt.Sprintf("%v%v", "D", rowNo1), fmt.Sprintf("%v%v", "E", rowNo), style4)
		excel.SetCellValue(sheet1, fmt.Sprintf("%c%v", columnNo1, rowNo1), "Total Deduction")

		columnNo1++
		excel.SetCellValue(sheet1, fmt.Sprintf("%c%v", columnNo1, rowNo1), data.Deduction)
		rowNo++

	}

	duration = time.Since(t)
	log.Println("excel Time taken ===> ", duration.Minutes(), "m")
	return excel, nil

}
func (s *Service) PayrollLogList(ctx *models.Context, filter *models.FilterPayrollLog, pagination *models.Pagination) ([]models.RefPayrollLog, error) {
	defer ctx.Session.EndSession(ctx.CTX)
	err := s.PayrollLogDataAccess(ctx, filter)
	if err != nil {
		return nil, err
	}
	payrolllog, err := s.Daos.PayrollLogList(ctx, filter, pagination)
	if err != nil {
		return nil, err
	}
	for k, v := range payrolllog {
		//t := time.Now()
		reports := new(models.DayWiseAttendanceReportFilter)
		reports.EmployeeId = v.EmployeeId
		reports.StartDate = filter.StartDate
		report, err := s.EmployeeDayWiseAttendanceReport(ctx, reports, nil)
		if err != nil {
			return nil, err
		}
		//	t.Month()=
		if len(report) > 0 {
			t1 := time.Date(filter.StartDate.Year(), filter.StartDate.Month()+1, 0, 23, 59, 59, 999999999, filter.StartDate.Location())
			//t2 := time.Date(t.Year(), t.Month(), 0, 0, 0, 0, 0, t.Location())
			var days int
			var lop float64
			var PartialPay float64
			if int(report[0].NoOfHolidays) > 0 {
				days = t1.Day() - int(report[0].NoOfHolidays)
			} else {
				days = t1.Day()

			}
			fmt.Println("no.of.days=======>", days)
			fmt.Println("report[0].NoOfLOP", report[0].NoOfLOP)

			if report[0].NoOfLOP > 0 {
				if float64(days) > 0 {
					lop = (v.NetAmount / float64(days)) * report[0].NoOfLOP
				}
			}
			if report[0].NoOfParticalPaid > 0 {
				if float64(days) > 0 {
					PartialPay = (v.NetAmount / float64(days) / 2) * report[0].NoOfParticalPaid
				}
			}
			fmt.Println("lop", lop)
			fmt.Println("PartialPay", PartialPay)
			payrolllog[k].Deduction = payrolllog[k].Deduction + lop + PartialPay
			payrolllog[k].Detections.Lop = lop + PartialPay
			payrolllog[k].NetAmount = payrolllog[k].NetAmount - lop - PartialPay

		} else {
			fmt.Println("ddddd", v.EmployeeId)
			t1 := time.Date(filter.StartDate.Year(), filter.StartDate.Month()+1, 0, 23, 59, 59, 999999999, filter.StartDate.Location())

			var days int
			var lop float64
			var PartialPay float64
			days = t1.Day()
			lop = (v.NetAmount / float64(days)) * float64(days)

			payrolllog[k].Deduction = payrolllog[k].Deduction + lop + PartialPay
			payrolllog[k].Detections.Lop = lop + PartialPay
			payrolllog[k].NetAmount = payrolllog[k].NetAmount - lop - PartialPay
		}

	}
	return payrolllog, nil

}
func (s *Service) EmployeePayrollPdf(ctx *models.Context, UniqueID *models.EmployeeSalarySlip) ([]byte, error) {

	r := NewRequestPdf("")

	data, err := s.GetEmployeeSalarySlip(ctx, UniqueID)
	if err != nil {
		return nil, err
	}
	//fmt.Println(data)
	//productConfigUniqueID := "6176962a9dac3d102e979b54"
	//productConfig, err := s.Daos.GetactiveProductConfig(ctx, true)
	if err != nil {
		return nil, errors.New("Error in getting product config" + err.Error())
	}
	// year := UniqueID.StartDate.Year()
	// month := UniqueID.StartDate.Month()
	// Yearofmonth := fmt.Sprintf("%v%v", year, int(month))
	// paySlipId := UniqueID.EmployeeId + Yearofmonth + s.Daos.GetUniqueID(ctx, constants.COLLECTIONEMPLOYEEPAYSLIP)

	m := make(map[string]interface{})
	m2 := make(map[string]interface{})
	m["payslip"] = data
	m2["currentDateYear"] = UniqueID.StartDate.Year()
	m2["currentDateMonth"] = UniqueID.StartDate.Month()
	m2["paySlipId"] = data.PayslipId
	fmt.Println("currentDateYear======>", m2["currentDateYear"])
	fmt.Println("currentDateMonth======>", m2["currentDateMonth"])
	fmt.Println("data======>", m["payslip"])
	var pdfdata models.PDFData
	pdfdata.Data = m
	pdfdata.RefData = m2
	//pdfdata.Config = productConfig.ProductConfig

	templatePathStart := s.ConfigReader.GetString(s.Shared.GetCmdArg(constants.ENV) + "." + constants.TEMPLATELOC)
	//fmt.Println("this is yuva", templatePathStart)
	//html template path
	templatePath := templatePathStart + "salaryslip.html"
	err = r.ParseTemplate(templatePath, pdfdata)
	if err != nil {
		return nil, err
	}
	ok, file, err := r.GeneratePDFAsFile()
	if err != nil {
		return nil, err
	}
	//fmt.Println("{}byte======>", file)
	fmt.Println(ok, "pdf generated successfully")

	return file, nil
}
func (s *Service) EmployeePayrollPdfV2(ctx *models.Context, UniqueID *models.EmployeeSalarySlip) (string, error) {

	//r := NewRequestPdf("")

	data, err := s.GetEmployeeSalarySlip(ctx, UniqueID)
	if err != nil {
		return "", err
	}
	year := UniqueID.StartDate.Year()
	month := UniqueID.StartDate.Month()
	Yearofmonth := fmt.Sprintf("%v%v", year, int(month))
	paySlipId := UniqueID.EmployeeId + Yearofmonth + s.Daos.GetUniqueID(ctx, constants.COLLECTIONEMPLOYEEPAYSLIP)
	//	"LGF-0105_2022August01"
	Filename := paySlipId + ".pdf"
	m := make(map[string]interface{})
	m2 := make(map[string]interface{})
	m["payslip"] = data
	m2["currentDateYear"] = UniqueID.StartDate.Year()
	m2["currentDateMonth"] = UniqueID.StartDate.Month()
	m2["paySlipId"] = paySlipId
	var pdfdata models.PDFData
	pdfdata.Data = m
	pdfdata.RefData = m2
	templatePathStart := s.ConfigReader.GetString(s.Shared.GetCmdArg(constants.ENV) + "." + constants.TEMPLATELOC)
	templatePath := templatePathStart + "salaryslip.html"

	docStart := s.ConfigReader.GetString(s.Shared.GetCmdArg(constants.ENV) + "." + constants.DOCLOC)

	fileuri := docStart + "payslip" + "/" + Filename
	responseURI := constants.DEFAULTFILEURL + "payslip" + "/" + Filename
	fmt.Println("fileuri=", fileuri)
	url, err := s.generatePDFFromTemplateV2(pdfdata, fileuri, templatePath)
	if err != nil {
		log.Println("Payslip error", err)
		return "", err
	}
	t := time.Now()
	fmt.Println("Outttt=", url)
	fmt.Println("responseURI=", responseURI)
	employeePayslip := new(models.EmployeePayslip)
	employeePayslip.UniqueID = paySlipId
	employeePayslip.EmployeeId = UniqueID.EmployeeId
	employeePayslip.OrganisationId = data.OrganisationId
	employeePayslip.ResponseUrl = responseURI
	employeePayslip.PayslipId = paySlipId
	employeePayslip.YearOfMonth = Yearofmonth
	employeePayslip.Date = &t
	employeePayslip.FileUrl = fileuri
	created := models.Created{}
	created.On = &t
	created.By = constants.SYSTEM
	log.Println("b4 EmployeePayslip.created")
	employeePayslip.Created = &created
	employeePayslip.Status = constants.EMPLOYEEPAYSLIPSTATUSACTIVE
	err = s.Daos.SaveEmployeePayslip(ctx, employeePayslip)
	if err != nil {
		log.Println("SaveEmployeePayslip error", err)
		return "", err
	}

	return responseURI, nil
}
func (s *Service) EmployeePayslipCorn() {
	c := context.TODO()
	ctx := app.GetApp(c, s.Daos)
	defer ctx.Client.Disconnect(c)
	//Dissemination := new(models.Dissemination)
	employees, err := s.Daos.GetActiveEmployee(ctx)
	if err != nil {
		log.Println("dissemination not found" + err.Error())
	}
	for _, v := range employees {
		filter := new(models.EmployeeSalarySlip)
		filter.EmployeeId = v.UniqueID
		t := time.Now()
		filter.StartDate = &t
		data, err := s.GetEmployeeSalarySlip(ctx, filter)
		if err != nil {
			log.Println("Payslip error", err)
			return
		}
		if data != nil {

			year := filter.StartDate.Year()
			month := filter.StartDate.Month()
			Yearofmonth := fmt.Sprintf("%v%v", year, int(month))
			paySlipId := filter.EmployeeId + Yearofmonth + s.Daos.GetUniqueID(ctx, constants.COLLECTIONEMPLOYEEPAYSLIP)
			//	"LGF-0105_2022August01"
			Filename := paySlipId + ".pdf"
			m := make(map[string]interface{})
			m2 := make(map[string]interface{})
			m["payslip"] = data
			m2["currentDateYear"] = filter.StartDate.Year()
			m2["currentDateMonth"] = filter.StartDate.Month()
			m2["paySlipId"] = paySlipId
			var pdfdata models.PDFData
			pdfdata.Data = m
			pdfdata.RefData = m2
			templatePathStart := s.ConfigReader.GetString(s.Shared.GetCmdArg(constants.ENV) + "." + constants.TEMPLATELOC)
			templatePath := templatePathStart + "salaryslip.html"

			docStart := s.ConfigReader.GetString(s.Shared.GetCmdArg(constants.ENV) + "." + constants.DOCLOC)

			fileuri := docStart + "payslip" + "/" + Filename
			responseURI := constants.DEFAULTFILEURL + "payslip" + "/" + Filename
			fmt.Println("fileuri=", fileuri)
			url, err := s.generatePDFFromTemplateV2(pdfdata, fileuri, templatePath)
			if err != nil {
				log.Println("Payslip error", err)
				return
			}
			fmt.Println("Outttt=", url)
			fmt.Println("responseURI=", responseURI)
			employeePayslip := new(models.EmployeePayslip)
			employeePayslip.UniqueID = paySlipId
			employeePayslip.EmployeeId = v.UniqueID
			employeePayslip.OrganisationId = v.OrganisationID
			employeePayslip.ResponseUrl = responseURI
			employeePayslip.PayslipId = paySlipId
			employeePayslip.YearOfMonth = Yearofmonth
			employeePayslip.Date = &t
			employeePayslip.FileUrl = fileuri
			created := models.Created{}
			created.On = &t
			created.By = constants.SYSTEM
			log.Println("b4 EmployeePayslip.created")
			employeePayslip.Created = &created
			employeePayslip.Status = constants.EMPLOYEEPAYSLIPSTATUSACTIVE
			err = s.Daos.SaveEmployeePayslip(ctx, employeePayslip)
			if err != nil {
				log.Println("SaveEmployeePayslip error", err)
				return
			}
		}
	}

}
