package services

import (
	"errors"
	"fmt"

	"hrms-services/constants"
	"hrms-services/models"
	"log"
	"time"

	"github.com/360EntSecGroup-Skylar/excelize"
	"go.mongodb.org/mongo-driver/mongo"
)

//SavePayroll :""
func (s *Service) SavePayroll(ctx *models.Context, payroll *models.Payroll) error {
	log.Println("transaction start")
	//Start Transaction
	if err := ctx.Session.StartTransaction(); err != nil {
		return err
	}
	defer ctx.Session.EndSession(ctx.CTX)
	payroll.UniqueID = s.Daos.GetUniqueID(ctx, constants.COLLECTIONPAYROLL)
	payroll.Status = constants.PAYROLLSTATUSACTIVE
	payroll.CTC = payroll.GrossAmount * 12
	t := time.Now()
	created := models.Created{}
	created.On = &t
	created.By = constants.SYSTEM
	log.Println("b4 Payroll.created")
	payroll.Created = &created
	log.Println("b4 Payroll.created")
	if err := mongo.WithSession(ctx.CTX, ctx.Session, func(sc mongo.SessionContext) error {
		dberr := s.Daos.SavePayroll(ctx, payroll)
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

//UpdatePayroll : ""
func (s *Service) UpdatePayroll(ctx *models.Context, payroll *models.Payroll) error {
	if err := ctx.Session.StartTransaction(); err != nil {
		return err
	}
	defer ctx.Session.EndSession(ctx.CTX)
	if err := mongo.WithSession(ctx.CTX, ctx.Session, func(sc mongo.SessionContext) error {

		err := s.Daos.UpdatePayroll(ctx, payroll)
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

//EnablePayroll : ""
func (s *Service) EnablePayroll(ctx *models.Context, UniqueID string) error {
	if err := ctx.Session.StartTransaction(); err != nil {
		return err
	}
	defer ctx.Session.EndSession(ctx.CTX)
	if err := mongo.WithSession(ctx.CTX, ctx.Session, func(sc mongo.SessionContext) error {

		err := s.Daos.EnablePayroll(ctx, UniqueID)
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

//DisablePayroll : ""
func (s *Service) DisablePayroll(ctx *models.Context, UniqueID string) error {
	if err := ctx.Session.StartTransaction(); err != nil {
		return err
	}
	defer ctx.Session.EndSession(ctx.CTX)
	if err := mongo.WithSession(ctx.CTX, ctx.Session, func(sc mongo.SessionContext) error {

		err := s.Daos.DisablePayroll(ctx, UniqueID)
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

//DeletePayroll : ""
func (s *Service) DeletePayroll(ctx *models.Context, UniqueID string) error {
	if err := ctx.Session.StartTransaction(); err != nil {
		return err
	}
	defer ctx.Session.EndSession(ctx.CTX)
	if err := mongo.WithSession(ctx.CTX, ctx.Session, func(sc mongo.SessionContext) error {

		err := s.Daos.DeletePayroll(ctx, UniqueID)
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

//GetSinglePayroll :""
func (s *Service) GetSinglePayroll(ctx *models.Context, UniqueID string) (*models.RefPayroll, error) {
	payroll, err := s.Daos.GetSinglePayroll(ctx, UniqueID)
	if err != nil {
		return nil, err
	}
	return payroll, nil
}

//FilterPayroll :""
func (s *Service) FilterPayroll(ctx *models.Context, filter *models.FilterPayroll, pagination *models.Pagination) ([]models.RefPayroll, error) {
	defer ctx.Session.EndSession(ctx.CTX)
	if filter != nil {
		dataaccess, err := s.Daos.DataAccess(ctx, &filter.DataAccess)
		if err != nil {
			return nil, err
		}
		if dataaccess != nil {
			if len(dataaccess.Organisation) > 0 {
				for _, v := range dataaccess.Organisation {
					filter.OrganisationID = append(filter.OrganisationID, v.UniqueID)
				}
			}

		}

	}
	return s.Daos.FilterPayroll(ctx, filter, pagination)

}
func (s *Service) SavePayrollWithEmployee(ctx *models.Context, payroll *models.Payroll) error {
	if err := ctx.Session.StartTransaction(); err != nil {
		return err
	}
	payroll.UniqueID = s.Daos.GetUniqueID(ctx, constants.COLLECTIONPAYROLL)
	payroll.Status = constants.PAYROLLSTATUSACTIVE
	payroll.CTC = payroll.GrossAmount * 12
	t := time.Now()
	created := models.Created{}
	created.On = &t
	created.By = constants.SYSTEM
	log.Println("b4 Payroll.created")
	payroll.Created = &created
	payroll.Date = &t
	log.Println("b4 Payroll.created")
	defer ctx.Session.EndSession(ctx.CTX)
	if err := mongo.WithSession(ctx.CTX, ctx.Session, func(sc mongo.SessionContext) error {
		err := s.Daos.SavePayrollWithUpsert(ctx, payroll)
		if err != nil {
			return err
		}
		payrollog, err := s.Daos.GetSinglePayrollLogWithEmployeeId(ctx, payroll.EmployeeId)
		if err != nil {
			return err
		}
		if payrollog != nil {

			if payrollog.StartDate.Month() == payroll.Date.Month() && payrollog.StartDate.Year() == payroll.Date.Year() {
				UpdateEmployeePayroll := new(models.PayrollLog)
				UpdateEmployeePayroll.OrganisationId = payroll.OrganisationId
				UpdateEmployeePayroll.UniqueID = payrollog.UniqueID
				UpdateEmployeePayroll.EmployeeId = payroll.EmployeeId
				UpdateEmployeePayroll.Method = payroll.Method
				if payroll.SalaryConfigId != "" {
					UpdateEmployeePayroll.SalaryConfigId = payroll.SalaryConfigId
				}
				UpdateEmployeePayroll.CTC = payroll.CTC
				UpdateEmployeePayroll.NetAmount = payroll.NetSalary
				UpdateEmployeePayroll.GrossAmount = payroll.GrossAmount
				UpdateEmployeePayroll.Deduction = payroll.TotalDeduction
				UpdateEmployeePayroll.Earnings = payroll.Earnings
				UpdateEmployeePayroll.Detections = payroll.Detections
				t := time.Now()
				t = time.Date(t.Year(), t.Month(), t.Day(), 0, 0, 0, 0, t.Location())

				UpdateEmployeePayroll.StartDate = &t
				UpdateEmployeePayroll.Status = constants.EMPLOYEEPAYROLLSTATUSACTIVE
				created := models.Created{}
				created.On = &t
				created.By = constants.SYSTEM
				log.Println("b4 EmployeePayrolllog.created")
				UpdateEmployeePayroll.Created = &created
				err := s.Daos.UpdatePayrollLog(ctx, UpdateEmployeePayroll)
				if err != nil {
					return err
				}
			} else {
				err := s.Daos.ArchivedPayrollLog(ctx, payrollog.UniqueID)
				if err != nil {
					return err
				}
				NewEmployeePayroll := new(models.PayrollLog)
				NewEmployeePayroll.OrganisationId = payroll.OrganisationId
				NewEmployeePayroll.UniqueID = s.Daos.GetUniqueID(ctx, constants.COLLECTIONPAYROLLLOG)
				NewEmployeePayroll.EmployeeId = payroll.EmployeeId
				NewEmployeePayroll.CTC = payroll.CTC
				NewEmployeePayroll.NetAmount = payroll.NetSalary
				NewEmployeePayroll.GrossAmount = payroll.GrossAmount
				NewEmployeePayroll.Deduction = payroll.TotalDeduction
				NewEmployeePayroll.Method = payroll.Method
				if payroll.SalaryConfigId != "" {
					NewEmployeePayroll.SalaryConfigId = payroll.SalaryConfigId
				}
				NewEmployeePayroll.Earnings = payroll.Earnings
				NewEmployeePayroll.Detections = payroll.Detections
				t := time.Now()
				t = time.Date(t.Year(), t.Month(), t.Day(), 0, 0, 0, 0, t.Location())

				NewEmployeePayroll.StartDate = &t
				NewEmployeePayroll.Status = constants.EMPLOYEEPAYROLLSTATUSACTIVE
				created := models.Created{}
				created.On = &t
				created.By = constants.SYSTEM
				log.Println("b4 EmployeePayrolllog.created")
				NewEmployeePayroll.Created = &created
				debrr := s.Daos.SavePayrollLog(ctx, NewEmployeePayroll)
				if debrr != nil {
					return debrr
				}
			}
		} else {
			NewEmployeePayroll := new(models.PayrollLog)
			NewEmployeePayroll.OrganisationId = payroll.OrganisationId
			NewEmployeePayroll.UniqueID = s.Daos.GetUniqueID(ctx, constants.COLLECTIONPAYROLLLOG)
			NewEmployeePayroll.EmployeeId = payroll.EmployeeId
			NewEmployeePayroll.CTC = payroll.CTC
			NewEmployeePayroll.NetAmount = payroll.NetSalary
			NewEmployeePayroll.GrossAmount = payroll.GrossAmount
			NewEmployeePayroll.Deduction = payroll.TotalDeduction
			NewEmployeePayroll.Method = payroll.Method
			if payroll.SalaryConfigId != "" {
				NewEmployeePayroll.SalaryConfigId = payroll.SalaryConfigId
			}
			NewEmployeePayroll.Earnings = payroll.Earnings
			NewEmployeePayroll.Detections = payroll.Detections
			t := time.Now()
			t = time.Date(t.Year(), t.Month(), t.Day(), 0, 0, 0, 0, t.Location())

			NewEmployeePayroll.StartDate = &t
			NewEmployeePayroll.Status = constants.EMPLOYEEPAYROLLSTATUSACTIVE
			created := models.Created{}
			created.On = &t
			created.By = constants.SYSTEM
			log.Println("b4 EmployeePayrolllog.created")
			NewEmployeePayroll.Created = &created
			debrr := s.Daos.SavePayrollLog(ctx, NewEmployeePayroll)
			if debrr != nil {
				return debrr
			}
		}

		if err := sc.CommitTransaction(sc); err != nil {
			return errors.New("Not able to commit - " + err.Error())
		}
		return nil
	}); err != nil {
		if err1 := ctx.Session.AbortTransaction(ctx.CTX); err1 != nil {
			log.Println("err in abort")
			return errors.New("Transaction Abort with error" + err1.Error())
		}
		return err
	}
	return nil
}
func (s *Service) GetSinglePayrollWithEmployee(ctx *models.Context, UniqueID string) (*models.RefPayroll, error) {
	payroll, err := s.Daos.GetSinglePayrollWithEmployee(ctx, UniqueID)
	if err != nil {
		return nil, err
	}
	return payroll, nil
}
func (s *Service) AttendanceExcel(ctx *models.Context, UniqueID string) (*excelize.File, error) {
	data, err := s.GetSinglePayrollWithEmployee(ctx, UniqueID)
	if err != nil {
		return nil, err
	}
	excel := excelize.NewFile()
	sheet1 := "Salary Slip"
	rowNo := 1
	index := excel.NewSheet(sheet1)
	excel.SetActiveSheet(index)
	excel.MergeCell(sheet1, "A1", "B5")
	excel.MergeCell(sheet1, "C1", "AF3")
	excel.MergeCell(sheet1, "C4", "AF5")
	excel.MergeCell(sheet1, "A6", "AF6")
	style3, err := excel.NewStyle(`{"fill":{"type":"pattern","color":["#B6DDE8"],"pattern":1},"alignment":{"horizontal":"center","vertical":"center"},"font":{"bold":true}}`)
	if err != nil {
		fmt.Println(err)
	}
	style1, err := excel.NewStyle(`{"fill":{"type":"pattern","color":["#FFDC6D"],"pattern":1},"alignment":{"horizontal":"center","vertical":"center"},"font":{"bold":true}}`)
	if err != nil {
		fmt.Println(err)
	}
	// style4, err := excel.NewStyle(`{"fill":{"type":"pattern","color":["#FFDC6D"],"pattern":1},"alignment":{"horizontal":"right","vertical":"center"},"font":{"bold":true}}`)
	// if err != nil {
	// 	fmt.Println(err)
	// }
	// style2, err := excel.NewStyle(`{"fill":{"type":"pattern","color":["#B6DDE8"],"pattern":1},"font":{"bold":true}}`)
	// if err != nil {
	// 	fmt.Println(err)
	// }
	excel.SetCellStyle(sheet1, fmt.Sprintf("%v%v", "C", rowNo), fmt.Sprintf("%v%v", "C", rowNo), style3)
	// documentUrl := s.Shared.Config.GetString(s.Shared.GetCmdArg(constants.ENV) + "." + constants.FILEURL)
	// if err := excel.AddPicture(sheet1, fmt.Sprintf("%v%v", "A", rowNo), documentUrl+"municipal/logo.png", `{"x_scale": 0.115, "y_scale": 0.0935}`); err != nil {
	// 	fmt.Println(err)
	// }
	//excel.SetCellValue(sheet1, fmt.Sprintf("%v%v", "C", rowNo), s.Shared.Config.GetString(s.Shared.GetCmdArg(constants.ENV)+"."+constants.LOCATIONNAME))
	rowNo++
	rowNo++
	rowNo++

	excel.SetCellStyle(sheet1, fmt.Sprintf("%v%v", "C", rowNo), fmt.Sprintf("%v%v", "C", rowNo), style3)
	excel.SetCellValue(sheet1, fmt.Sprintf("%v%v", "C", rowNo), "Salary Slip ")
	rowNo++
	rowNo++
	// t := time.Now()
	// toDate := t.Format("02-January-2006")

	// reportFromMsg := "Report Generated on" + " " + toDate

	// excel.SetCellStyle(sheet1, fmt.Sprintf("%v%v", "A", rowNo), fmt.Sprintf("%v%v", "A", rowNo), style3)
	// excel.SetCellValue(sheet1, fmt.Sprintf("%v%v", "A", rowNo), reportFromMsg)
	rowNo++

	if data != nil {
		columnNo1 := 'B'
		columnNo2 := 'D'

		excel.SetCellValue(sheet1, fmt.Sprintf("%c%v", columnNo1, rowNo), "Employee Code:")
		columnNo1++
		excel.SetCellValue(sheet1, fmt.Sprintf("%c%v", columnNo1, rowNo), "Employee Name:")
		columnNo1++
		excel.SetCellValue(sheet1, fmt.Sprintf("%c%v", columnNo1, rowNo), "Designation:")
		columnNo1++
		excel.SetCellValue(sheet1, fmt.Sprintf("%c%v", columnNo1, rowNo), "Bank Account No:")
		columnNo1++
		excel.SetCellValue(sheet1, fmt.Sprintf("%c%v", columnNo1, rowNo), "UAN No:")
		columnNo1++
		excel.SetCellValue(sheet1, fmt.Sprintf("%c%v", columnNo1, rowNo), "ESIC No:")
		columnNo1++
		excel.SetCellValue(sheet1, fmt.Sprintf("%c%v", columnNo1, rowNo), "Total Working Day in Month:")
		columnNo1++
		excel.SetCellValue(sheet1, fmt.Sprintf("%c%v", columnNo1, rowNo), "Absent Days In Month:")
		columnNo1++
		excel.SetCellValue(sheet1, fmt.Sprintf("%c%v", columnNo2, rowNo), "Date Of Birth:")
		columnNo2++
		excel.SetCellValue(sheet1, fmt.Sprintf("%c%v", columnNo2, rowNo), "Date Of Joining:")
		columnNo2++
		excel.SetCellValue(sheet1, fmt.Sprintf("%c%v", columnNo2, rowNo), "Location:")
		columnNo2++
		excel.SetCellValue(sheet1, fmt.Sprintf("%c%v", columnNo2, rowNo), "Bank Name:")
		columnNo2++
		excel.SetCellValue(sheet1, fmt.Sprintf("%c%v", columnNo2, rowNo), "IFSC Code:")
		columnNo2++
		excel.SetCellValue(sheet1, fmt.Sprintf("%c%v", columnNo2, rowNo), "PAN Code:")
		columnNo2++
		excel.SetCellValue(sheet1, fmt.Sprintf("%c%v", columnNo2, rowNo), "Balance Leaves in Previous Month:")
		columnNo2++
		excel.SetCellValue(sheet1, fmt.Sprintf("%c%v", columnNo2, rowNo), "Balance Leaves in Current Month:")
		columnNo2++
		excel.SetCellValue(sheet1, fmt.Sprintf("%c%v", "B", rowNo), "Components In Salary(Earnings):")
		rowNo++
		excel.SetCellValue(sheet1, fmt.Sprintf("%c%v", "C", rowNo), "Amount(in Rs):")
		rowNo++
		excel.SetCellValue(sheet1, fmt.Sprintf("%c%v", "D", rowNo), "Components In Salary(Deduction):")
		rowNo++
		excel.SetCellValue(sheet1, fmt.Sprintf("%c%v", "E", rowNo), "Amount(in Rs):")
		columnNo1++
		excel.SetCellValue(sheet1, fmt.Sprintf("%c%v", columnNo1, rowNo), "Basic Salary:")
		columnNo1++
		excel.SetCellValue(sheet1, fmt.Sprintf("%c%v", columnNo1, rowNo), "HRA:")
		columnNo1++
		excel.SetCellValue(sheet1, fmt.Sprintf("%c%v", columnNo1, rowNo), "Conveyance Allowances:")
		columnNo1++
		excel.SetCellValue(sheet1, fmt.Sprintf("%c%v", columnNo1, rowNo), "Educational Allowances:")
		columnNo1++
		excel.SetCellValue(sheet1, fmt.Sprintf("%c%v", columnNo1, rowNo), "Performance Allowances:")
		columnNo1++
		excel.SetCellValue(sheet1, fmt.Sprintf("%c%v", columnNo1, rowNo), "Total Gross Salary:")
		columnNo2++
		excel.SetCellValue(sheet1, fmt.Sprintf("%c%v", columnNo2, rowNo), "PF Contribution:")
		columnNo1++
		excel.SetCellValue(sheet1, fmt.Sprintf("%c%v", columnNo2, rowNo), "ESIC Contribution:")
		columnNo2++
		columnNo2++
		columnNo2++
		excel.SetCellValue(sheet1, fmt.Sprintf("%c%v", columnNo2, rowNo), "Total Deduction:")
		columnNo1++
		// var ch1 rune
		// var ch2 rune
		// ch1 = 'B'
		// ch2 = 'B'
		// column := ""
		excel.SetCellStyle(sheet1, fmt.Sprintf("%v%v", "B", rowNo), fmt.Sprintf("%v%v", "E", rowNo), style1)
		// for _, v := range data[0]. {

		// 	if columnNo1 <= 'Z' {
		// 		column = fmt.Sprintf("%c", columnNo1)
		// 		day := fmt.Sprintf("%v", v.ID)
		// 		excel.SetCellValue(sheet1, fmt.Sprintf("%v%v", column, rowNo), day)
		// 	} else {
		// 		CH1 := fmt.Sprintf("%c", ch1)
		// 		CH2 := fmt.Sprintf("%c", ch2)
		// 		fmt.Println("ch1===>", CH1)
		// 		fmt.Println("ch2===>", CH2)
		// 		column = fmt.Sprintf("%c%c", ch1, ch2)
		// 		day := fmt.Sprintf("%v", v.ID)
		// 		excel.SetCellValue(sheet1, fmt.Sprintf("%v%v", column, rowNo), day)
		// 	}
		// 	if columnNo1 > 'Z' {
		// 		ch2++
		// 	}
		// 	columnNo1++
		// }
		rowNo++
		// for i, v2 := range data {
		// 	columnNo2 := 'A'
		// 	excel.SetCellValue(sheet1, fmt.Sprintf("%c%v", columnNo2, rowNo), i+1)
		// 	columnNo2++

		// 	excel.SetCellValue(sheet1, fmt.Sprintf("%c%v", columnNo2, rowNo), v2.Name)
		// 	columnNo2++

		// 	ch3 := 'A'
		// 	ch4 := 'A'
		// 	column2 := ""
		// 	for _, v3 := range v2.Payments {
		// 		if columnNo2 <= 'Z' {
		// 			column2 = fmt.Sprintf("%c", columnNo2)
		// 			collection := fmt.Sprintf("%v", v3.TotalCollection)
		// 			excel.SetCellValue(sheet1, fmt.Sprintf("%v%v", column2, rowNo), collection)
		// 		} else {
		// 			column2 = fmt.Sprintf("%c%c", ch3, ch4)
		// 			collection := fmt.Sprintf("%v", v3.TotalCollection)
		// 			excel.SetCellValue(sheet1, fmt.Sprintf("%v%v", column2, rowNo), collection)

		// 		}
		// 		if columnNo2 > 'Z' {
		// 			ch4++
		// 		}
		// 		columnNo2++

		// 	}
		// 	rowNo++
		// }
	}
	return excel, nil
}
