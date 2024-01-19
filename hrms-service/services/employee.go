package services

import (
	"errors"
	"fmt"
	"mime/multipart"

	"hrms-services/constants"
	"hrms-services/models"
	"log"
	"time"

	"github.com/360EntSecGroup-Skylar/excelize"
	"go.mongodb.org/mongo-driver/mongo"
)

//Employee Login
// func (s *Service) EmployeeLogin(ctx *models.Context, EmployeeLogin *models.EmployeeLogin) (string, bool, error) {
// 	data, err := s.Daos.GetSingleEmployee(ctx, EmployeeLogin.UserName)
// 	if err != nil {
// 		fmt.Println(err)
// 		return "dal err", false, err
// 	}
// 	if ok := data.Password == EmployeeLogin.PassWord; !ok {
// 		log.Println("Data password ==>", data.Password)
// 		log.Println("login password ==>", EmployeeLogin.PassWord)
// 		return "Passs false", false, nil
// 	}
// 	if data.Status == constants.USERSTATUSINIT {
// 		return "", false, errors.New("Awaiting Activation")
// 	}

// 	return "", true, nil
// }

//SaveEmployee :""
func (s *Service) SaveEmployee(ctx *models.Context, employee *models.Employee) error {
	log.Println("transaction start")
	//Start Transaction
	if err := ctx.Session.StartTransaction(); err != nil {
		return err
	}
	defer ctx.Session.EndSession(ctx.CTX)
	employee.UniqueID = s.Daos.GetUniqueID(ctx, constants.COLLECTIONEMPLOYEE)
	employee.UserName = employee.UniqueID
	employee.Status = constants.EMPLOYEESTATUSONBORADING
	t := time.Now()
	created := models.Created{}
	created.On = &t
	created.By = constants.SYSTEM
	log.Println("b4 Employee.created")
	employee.Created = &created
	employee.Type = constants.USERTYPEEMPLOYEE
	log.Println("b4 Employee.created")
	if err := mongo.WithSession(ctx.CTX, ctx.Session, func(sc mongo.SessionContext) error {
		//Employee
		dberr := s.Daos.SaveEmployee(ctx, employee)
		if dberr != nil {
			return dberr
		}
		//User
		user := new(models.User)
		user.UserName = employee.UniqueID
		user.Password = "#nature32" //Default Password
		user.UniqueID = s.Daos.GetUniqueID(ctx, constants.COLLECTIONUSER)
		user.Name = employee.Name
		user.LastName = employee.LastName
		user.Gender = employee.Gender
		user.Mobile = employee.Mobile
		user.Email = employee.Email
		user.OfficialEmail = employee.OfficialEmail
		user.DOB = employee.DOB
		user.JoiningDate = employee.JoiningDate
		user.OrganisationID = employee.OrganisationID
		user.EmployeeId = employee.UniqueID
		user.Status = constants.EMPLOYEESTATUSREJECT
		user.Role = employee.Role
		user.Type = "Employee"
		user.OrganisationID = employee.OrganisationID
		dberr = s.Daos.SaveUser(ctx, user)
		if dberr != nil {
			return dberr
		}
		//EmployeeLog
		employeeLog := new(models.EmployeeLog)
		employeeLog.Name = employee.Name
		employeeLog.UniqueID = s.Daos.GetUniqueID(ctx, constants.COLLECTIONJOBTIMELINE)
		employeeLog.OrganisationId = employee.OrganisationID
		employeeLog.DepartmentId = employee.DepartmentID
		employeeLog.BranchId = employee.BranchID
		employeeLog.DesignationId = employee.DepartmentID
		employeeLog.Desc = constants.EMPLOYEESTATUSONBORADINGDESC
		employeeLog.Action.UserID = employee.UniqueID
		employeeLog.Action.UserType = employee.Role
		employeeLog.Action.Date = t
		employeeLog.EmployeeId = employee.UniqueID
		employeeLog.Status = constants.EMPLOYEESTATUSONBORADING
		employeeLog.Remark = employee.Remark
		dberr = s.Daos.SaveEmployeeLog(ctx, employeeLog)
		if dberr != nil {
			return dberr
		}

		//jobtimeline
		jobtimeline := new(models.JobTimeline)
		jobtimeline.State = constants.EMPLOYEESTATUSONBORADING
		jobtimeline.UniqueID = s.Daos.GetUniqueID(ctx, constants.COLLECTIONJOBTIMELINE)
		jobtimeline.StartDate = &t
		jobtimeline.Status = constants.EMPLOYEESTATUSACTIVESTAGE
		jobtimeline.OrganisationId = employee.OrganisationID
		jobtimeline.DepartmentId = employee.DepartmentID
		jobtimeline.BranchId = employee.BranchID
		jobtimeline.DesignationId = employee.DepartmentID
		jobtimeline.Assigned.UserType = employee.Role
		jobtimeline.Assigned.Date = t
		jobtimeline.EmployeeId = employee.UniqueID
		jobtimeline.Remark = employee.Remark
		dberr = s.Daos.SaveJobTimeline(ctx, jobtimeline)
		if dberr != nil {
			return dberr
		}
		if err := ctx.Session.CommitTransaction(sc); err != nil {
			return errors.New("Not able to commit - " + err.Error())
		}
		return nil
	}); err != nil {
		log.Println("Transaction start aborting")
		if abortError := ctx.Session.AbortTransaction(ctx.CTX); abortError != nil {
			return errors.New("Error while aborting transaction" + abortError.Error())
		}
		log.Println("Transaction aborting completed successfully")
		return err
	}

	return nil
}
func (s *Service) SaveEmployeeWithoutTransaction(ctx *models.Context, employee *models.Employee) error {

	employee.UniqueID = s.Daos.GetUniqueID(ctx, constants.COLLECTIONEMPLOYEE)
	employee.UserName = employee.UniqueID
	employee.Status = constants.EMPLOYEESTATUSONBORADING
	t := time.Now()
	created := models.Created{}
	created.On = &t
	created.By = constants.SYSTEM
	log.Println("b4 Employee.created")
	employee.Created = &created
	employee.Type = constants.USERTYPEEMPLOYEE
	log.Println("b4 Employee.created")
	dberr := s.Daos.SaveEmployee(ctx, employee)
	if dberr != nil {
		return dberr
	}
	//User
	user := new(models.User)
	user.UserName = employee.UniqueID
	user.Password = "#nature32" //Default Password
	user.UniqueID = s.Daos.GetUniqueID(ctx, constants.COLLECTIONUSER)
	user.Name = employee.Name
	user.LastName = employee.LastName
	user.Gender = employee.Gender
	user.Mobile = employee.Mobile
	user.Email = employee.Email
	user.OfficialEmail = employee.OfficialEmail
	user.DOB = employee.DOB
	user.JoiningDate = employee.JoiningDate
	user.OrganisationID = employee.OrganisationID
	user.EmployeeId = employee.UniqueID
	user.Status = constants.EMPLOYEESTATUSONBORADING
	user.Role = employee.Role
	user.Type = employee.Type
	dberr = s.Daos.SaveUser(ctx, user)
	if dberr != nil {
		return dberr
	}
	//EmployeeLog
	employeeLog := new(models.EmployeeLog)
	employeeLog.Name = employee.Name
	employeeLog.UniqueID = s.Daos.GetUniqueID(ctx, constants.COLLECTIONJOBTIMELINE)
	employeeLog.OrganisationId = employee.OrganisationID
	employeeLog.DepartmentId = employee.DepartmentID
	employeeLog.BranchId = employee.BranchID
	employeeLog.DesignationId = employee.DepartmentID
	employeeLog.Desc = constants.EMPLOYEESTATUSONBORADINGDESC
	employeeLog.Action.UserID = employee.UniqueID
	employeeLog.Action.UserType = employee.Role
	employeeLog.Action.Date = t
	employeeLog.EmployeeId = employee.UniqueID
	employeeLog.Status = constants.EMPLOYEESTATUSONBORADING
	employeeLog.Remark = employee.Remark
	dberr = s.Daos.SaveEmployeeLog(ctx, employeeLog)
	if dberr != nil {
		return dberr
	}

	//jobtimeline
	jobtimeline := new(models.JobTimeline)
	jobtimeline.State = constants.EMPLOYEESTATUSONBORADING
	jobtimeline.UniqueID = s.Daos.GetUniqueID(ctx, constants.COLLECTIONJOBTIMELINE)
	jobtimeline.StartDate = &t
	jobtimeline.Status = constants.EMPLOYEESTATUSACTIVESTAGE
	jobtimeline.OrganisationId = employee.OrganisationID
	jobtimeline.DepartmentId = employee.DepartmentID
	jobtimeline.BranchId = employee.BranchID
	jobtimeline.DesignationId = employee.DepartmentID
	jobtimeline.Assigned.UserType = employee.Role
	jobtimeline.Assigned.Date = t
	jobtimeline.EmployeeId = employee.UniqueID
	jobtimeline.Remark = employee.Remark
	dberr = s.Daos.SaveJobTimeline(ctx, jobtimeline)
	if dberr != nil {
		return dberr
	}

	return nil
}

//UpdateEmployee : ""
func (s *Service) UpdateEmployee(ctx *models.Context, employee *models.Employee) error {
	if err := ctx.Session.StartTransaction(); err != nil {
		return err
	}
	defer ctx.Session.EndSession(ctx.CTX)
	if err := mongo.WithSession(ctx.CTX, ctx.Session, func(sc mongo.SessionContext) error {

		err := s.Daos.UpdateEmployee(ctx, employee)
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

//EnableEmployee : ""
func (s *Service) EnableEmployee(ctx *models.Context, UniqueID string) error {
	if err := ctx.Session.StartTransaction(); err != nil {
		return err
	}
	defer ctx.Session.EndSession(ctx.CTX)
	if err := mongo.WithSession(ctx.CTX, ctx.Session, func(sc mongo.SessionContext) error {

		err := s.Daos.EnableEmployee(ctx, UniqueID)
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

//DisableEmployee : ""
func (s *Service) DisableEmployee(ctx *models.Context, UniqueID string) error {
	if err := ctx.Session.StartTransaction(); err != nil {
		return err
	}
	defer ctx.Session.EndSession(ctx.CTX)
	if err := mongo.WithSession(ctx.CTX, ctx.Session, func(sc mongo.SessionContext) error {

		err := s.Daos.DisableEmployee(ctx, UniqueID)
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

//DeleteEmployee : ""
func (s *Service) DeleteEmployee(ctx *models.Context, UniqueID string) error {
	if err := ctx.Session.StartTransaction(); err != nil {
		return err
	}
	defer ctx.Session.EndSession(ctx.CTX)
	if err := mongo.WithSession(ctx.CTX, ctx.Session, func(sc mongo.SessionContext) error {

		err := s.Daos.DeleteEmployee(ctx, UniqueID)
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

//GetSingleEmployee :""
func (s *Service) GetSingleEmployee(ctx *models.Context, UniqueID string) (*models.RefEmployee, error) {
	Employee, err := s.Daos.GetSingleEmployee(ctx, UniqueID)
	if err != nil {
		return nil, err
	}
	return Employee, nil
}

//FilterEmployee :""
func (s *Service) FilterEmployee(ctx *models.Context, filter *models.FilterEmployee, pagination *models.Pagination) ([]models.RefEmployee, error) {
	defer ctx.Session.EndSession(ctx.CTX)
	err := s.EmployeeDataAccess(ctx, filter)
	if err != nil {
		return nil, err
	}
	return s.Daos.FilterEmployee(ctx, filter, pagination)

}
func (s *Service) EmployeeDataAccess(ctx *models.Context, filter *models.FilterEmployee) (err error) {
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

//EmployeeReject : ""
func (s *Service) EmployeeReject(ctx *models.Context, employeereject *models.EmployeeMoveToReject) error {

	userdata, err := s.Daos.GetSingleUserbyemployeeid(ctx, employeereject.EmployeeID)
	if err != nil {
		fmt.Println(err)
		return err
	}

	employee, err := s.Daos.GetSingleEmployee(ctx, employeereject.EmployeeID)
	if err != nil {
		fmt.Println(err)
		return err
	}

	datajobtimeline, err := s.Daos.GetSingleJobTimelineemployeeid(ctx, employeereject.EmployeeID, constants.EMPLOYEESTATUSACTIVESTAGE)
	if err != nil {
		fmt.Println(err)
		return err
	}

	if err := ctx.Session.StartTransaction(); err != nil {
		return err
	}
	defer ctx.Session.EndSession(ctx.CTX)

	if err := mongo.WithSession(ctx.CTX, ctx.Session, func(sc mongo.SessionContext) error {
		t := time.Now()
		//update Employee data
		employee.RejectDate = &t
		employee.Status = constants.EMPLOYEESTATUSREJECT
		employee.Remark = employeereject.Remark
		dberr := s.Daos.EmployeeReject(ctx, &employee.Employee)
		if dberr != nil {
			return dberr
		}
		//JobTimeLine update field here
		datajobtimeline.EndDate = &t
		datajobtimeline.Status = constants.EMPLOYEESTATUSCOMPLETEDSTAGE
		dberr = s.Daos.UpdateJobTimeline(ctx, &datajobtimeline.JobTimeline)
		if dberr != nil {
			return dberr
		}

		//User
		userdata.Status = constants.EMPLOYEESTATUSREJECT
		userdata.OrganisationID = employee.OrganisationID
		userdata.DepartmentID = employee.DepartmentID
		userdata.BranchID = employee.BranchID
		userdata.DesignationID = employee.DepartmentID
		dberr = s.Daos.UpdateUserbyemployeeId(ctx, &userdata.User)
		if dberr != nil {
			return dberr
		}

		//EmployeeLog
		employeeLog := new(models.EmployeeLog)
		employeeLog.Name = employee.Name
		employeeLog.UniqueID = s.Daos.GetUniqueID(ctx, constants.COLLECTIONJOBTIMELINE)
		employeeLog.OrganisationId = employee.OrganisationID
		employeeLog.DepartmentId = employee.DepartmentID
		employeeLog.BranchId = employee.BranchID
		employeeLog.DesignationId = employee.DepartmentID
		employeeLog.Desc = constants.EMPLOYEESTATUSREJECTDESC
		employeeLog.Action.UserID = employeereject.ByID
		employeeLog.Action.UserType = employeereject.ByType
		employeeLog.Action.Date = t
		employeeLog.EmployeeId = employee.UniqueID
		employeeLog.Status = constants.EMPLOYEESTATUSREJECT
		employeeLog.Remark = employeereject.Remark
		dberr = s.Daos.SaveEmployeeLog(ctx, employeeLog)
		if dberr != nil {
			return dberr
		}

		//jobtimeline
		jobtimeline := new(models.JobTimeline)
		jobtimeline.State = constants.EMPLOYEESTATUSREJECTDESC
		jobtimeline.UniqueID = s.Daos.GetUniqueID(ctx, constants.COLLECTIONJOBTIMELINE)
		jobtimeline.StartDate = &t
		jobtimeline.Status = constants.EMPLOYEESTATUSACTIVESTAGE
		jobtimeline.OrganisationId = employee.OrganisationID
		jobtimeline.DepartmentId = employee.DepartmentID
		jobtimeline.BranchId = employee.BranchID
		jobtimeline.DesignationId = employee.DepartmentID
		jobtimeline.Assigned.UserID = employeereject.By
		jobtimeline.Assigned.UserType = employeereject.ByType
		jobtimeline.Assigned.Date = t
		jobtimeline.EmployeeId = employee.UniqueID
		jobtimeline.Remark = employeereject.Remark
		dberr = s.Daos.SaveJobTimeline(ctx, jobtimeline)
		if dberr != nil {
			return dberr
		}

		if err := ctx.Session.CommitTransaction(sc); err != nil {
			return errors.New("Not able to commit - " + err.Error())
		}
		return nil
	}); err != nil {
		log.Println("Transaction start aborting")
		if abortError := ctx.Session.AbortTransaction(ctx.CTX); abortError != nil {
			return errors.New("Error while aborting transaction" + abortError.Error())
		}
		log.Println("Transaction aborting completed successfully")
		return err
	}

	return nil
}

//EmployeeOnboarding : ""
func (s *Service) EmployeeOnboarding(ctx *models.Context, employeeonboarding *models.EmployeeMoveToOnboarding) error {

	userdata, err := s.Daos.GetSingleUserbyemployeeid(ctx, employeeonboarding.EmployeeID)
	if err != nil {
		fmt.Println(err)
		return err
	}
	employee, err := s.Daos.GetSingleEmployee(ctx, employeeonboarding.EmployeeID)
	if err != nil {
		fmt.Println(err)
		return err
	}

	datajobtimeline, err := s.Daos.GetSingleJobTimelineemployeeid(ctx, employeeonboarding.EmployeeID, constants.EMPLOYEESTATUSACTIVESTAGE)
	if err != nil {
		fmt.Println(err)
		return err
	}

	if err := ctx.Session.StartTransaction(); err != nil {
		return err
	}
	defer ctx.Session.EndSession(ctx.CTX)

	if err := mongo.WithSession(ctx.CTX, ctx.Session, func(sc mongo.SessionContext) error {
		t := time.Now()
		//update Employee data
		employee.ReOnboardingDate = &t
		employee.Status = constants.EMPLOYEESTATUSONBORADING
		employee.Remark = employeeonboarding.Remark
		dberr := s.Daos.EmployeeOnboarding(ctx, &employee.Employee)
		if dberr != nil {
			return dberr
		}
		//JobTimeLine update field here
		datajobtimeline.EndDate = &t
		datajobtimeline.Status = constants.EMPLOYEESTATUSCOMPLETEDSTAGE
		dberr = s.Daos.UpdateJobTimeline(ctx, &datajobtimeline.JobTimeline)
		if dberr != nil {
			return dberr
		}

		//User
		userdata.Status = constants.EMPLOYEESTATUSONBORADING
		userdata.OrganisationID = employee.OrganisationID
		userdata.DepartmentID = employee.DepartmentID
		userdata.BranchID = employee.BranchID
		userdata.DesignationID = employee.DepartmentID
		dberr = s.Daos.UpdateUserbyemployeeId(ctx, &userdata.User)
		if dberr != nil {
			return dberr
		}

		//EmployeeLog
		employeeLog := new(models.EmployeeLog)
		employeeLog.Name = employee.Name
		employeeLog.UniqueID = s.Daos.GetUniqueID(ctx, constants.COLLECTIONJOBTIMELINE)
		employeeLog.OrganisationId = employee.OrganisationID
		employeeLog.DepartmentId = employee.DepartmentID
		employeeLog.BranchId = employee.BranchID
		employeeLog.DesignationId = employee.DepartmentID
		employeeLog.Desc = constants.EMPLOYEESTATUSONBORADINGDESC
		employeeLog.Action.UserID = employeeonboarding.ByID
		employeeLog.Action.UserType = employeeonboarding.ByType
		employeeLog.Action.Date = t
		employeeLog.EmployeeId = employee.UniqueID
		employeeLog.Status = constants.EMPLOYEESTATUSONBORADING
		employeeLog.Remark = employeeonboarding.Remark
		dberr = s.Daos.SaveEmployeeLog(ctx, employeeLog)
		if dberr != nil {
			return dberr
		}

		//jobtimeline
		jobtimeline := new(models.JobTimeline)
		jobtimeline.State = constants.EMPLOYEESTATUSONBORADINGDESC
		jobtimeline.UniqueID = s.Daos.GetUniqueID(ctx, constants.COLLECTIONJOBTIMELINE)
		jobtimeline.StartDate = &t
		jobtimeline.Status = constants.EMPLOYEESTATUSACTIVESTAGE
		jobtimeline.OrganisationId = employee.OrganisationID
		jobtimeline.DepartmentId = employee.DepartmentID
		jobtimeline.BranchId = employee.BranchID
		jobtimeline.DesignationId = employee.DepartmentID
		jobtimeline.Assigned.UserID = employeeonboarding.By
		jobtimeline.Assigned.UserType = employeeonboarding.ByType
		jobtimeline.Assigned.Date = t
		jobtimeline.EmployeeId = employee.UniqueID
		jobtimeline.Remark = employeeonboarding.Remark
		dberr = s.Daos.SaveJobTimeline(ctx, jobtimeline)
		if dberr != nil {
			return dberr
		}

		if err := ctx.Session.CommitTransaction(sc); err != nil {
			return errors.New("Not able to commit - " + err.Error())
		}
		return nil
	}); err != nil {
		log.Println("Transaction start aborting")
		if abortError := ctx.Session.AbortTransaction(ctx.CTX); abortError != nil {
			return errors.New("Error while aborting transaction" + abortError.Error())
		}
		log.Println("Transaction aborting completed successfully")
		return err
	}

	return nil
}

//EmployeeProbationary : ""
func (s *Service) EmployeeProbationary(ctx *models.Context, employeeprobationary *models.EmployeeMoveToProbationary) error {
	probationarydata, err := s.Daos.GetSingleProbationary(ctx, employeeprobationary.ProbationaryID)
	if err != nil {
		fmt.Println(err)
		return err
	}
	userdata, err := s.Daos.GetSingleUserbyemployeeid(ctx, employeeprobationary.EmployeeID)
	if err != nil {
		fmt.Println(err)
		return err
	}
	datajobtimeline, err := s.Daos.GetSingleJobTimelineemployeeid(ctx, employeeprobationary.EmployeeID, constants.EMPLOYEESTATUSACTIVESTAGE)
	if err != nil {
		fmt.Println(err)
		return err
	}
	employee, err := s.Daos.GetSingleEmployee(ctx, employeeprobationary.EmployeeID)
	if err != nil {
		fmt.Println(err)
		return err
	}

	if err := ctx.Session.StartTransaction(); err != nil {
		return err
	}
	defer ctx.Session.EndSession(ctx.CTX)

	if err := mongo.WithSession(ctx.CTX, ctx.Session, func(sc mongo.SessionContext) error {
		//update Employee data
		currentDate := time.Now()
		employee.ProbationaryendDate = currentDate.AddDate(0, 0, probationarydata.ProbationaryDays)
		employee.Remark = employeeprobationary.Remark
		employee.Status = constants.EMPLOYEESTATUSPROBATIONARY
		employee.NoticeID = employeeprobationary.NoticeID
		employee.DepartmentID = employeeprobationary.DepartmentID
		employee.DesignationID = employeeprobationary.DesignationID
		employee.BranchID = employeeprobationary.BranchID
		employee.WorkScheduleID = employeeprobationary.WorkScheduleID
		employee.LeavePolicyID = employeeprobationary.LeavePolicyID
		employee.DocumentPolicyID = employeeprobationary.DocumentPolicyID
		employee.ProbationaryID = employeeprobationary.ProbationaryID
		employee.OfficialEmail = employeeprobationary.OfficialEmail
		employee.LineManager = employeeprobationary.LineManager
		// employee.PayrollPolicyId = employeeprobationary.PayrollPolicyId
		dberr := s.Daos.EmployeeProbationary(ctx, &employee.Employee)
		if dberr != nil {
			return dberr
		}

		t := time.Now()
		//JobTimeLine update field here
		datajobtimeline.EndDate = &t
		datajobtimeline.Status = constants.EMPLOYEESTATUSCOMPLETEDSTAGE

		dberr = s.Daos.UpdateJobTimeline(ctx, &datajobtimeline.JobTimeline)
		if dberr != nil {
			return dberr
		}

		//User
		userdata.Status = constants.EMPLOYEESTATUSPROBATIONARY
		userdata.OrganisationID = employee.OrganisationID
		userdata.DepartmentID = employee.DepartmentID
		userdata.BranchID = employee.BranchID
		userdata.DesignationID = employee.DesignationID
		// userdata.p = employeeprobationary.ProbationaryID

		dberr = s.Daos.UpdateUserbyemployeeId(ctx, &userdata.User)
		if dberr != nil {
			return dberr
		}

		//fmt.Println("EmployeeLog here start transaction")
		//EmployeeLog
		employeeLog := new(models.EmployeeLog)
		employeeLog.Name = employee.Name
		employeeLog.UniqueID = s.Daos.GetUniqueID(ctx, constants.COLLECTIONJOBTIMELINE)
		employeeLog.OrganisationId = employee.OrganisationID
		employeeLog.DepartmentId = employee.DepartmentID
		employeeLog.BranchId = employee.BranchID
		employeeLog.DesignationId = employee.DepartmentID
		employeeLog.Desc = constants.EMPLOYEESTATUSPROBATIONARYDESC
		employeeLog.Action.UserID = employeeprobationary.ByID
		employeeLog.Action.UserType = employeeprobationary.ByType
		employeeLog.Action.Date = t
		employeeLog.EmployeeId = employee.UniqueID
		employeeLog.Status = constants.EMPLOYEESTATUSPROBATIONARY
		employeeLog.Remark = employeeprobationary.Remark
		dberr = s.Daos.SaveEmployeeLog(ctx, employeeLog)
		if dberr != nil {
			return dberr
		}

		//fmt.Println("jobtimeline here start transaction")
		//jobtimeline
		jobtimeline := new(models.JobTimeline)
		jobtimeline.State = constants.EMPLOYEESTATUSPROBATIONARY
		jobtimeline.UniqueID = s.Daos.GetUniqueID(ctx, constants.COLLECTIONJOBTIMELINE)
		jobtimeline.StartDate = &t
		jobtimeline.Status = constants.EMPLOYEESTATUSACTIVESTAGE
		jobtimeline.OrganisationId = employee.OrganisationID
		jobtimeline.DepartmentId = employee.DepartmentID
		jobtimeline.BranchId = employee.BranchID
		jobtimeline.DesignationId = employee.DepartmentID
		jobtimeline.Assigned.UserID = employeeprobationary.By
		jobtimeline.Assigned.UserType = employeeprobationary.ByType
		jobtimeline.Assigned.Date = t
		jobtimeline.EmployeeId = employee.UniqueID
		jobtimeline.Remark = employeeprobationary.Remark
		dberr = s.Daos.SaveJobTimeline(ctx, jobtimeline)
		if dberr != nil {
			return dberr
		}
		leave, err := s.Daos.GetSingleLeavePolicy(ctx, employeeprobationary.LeavePolicyID)
		if err != nil {
			return err
		}
		employeeleave := new(models.EmployeeLeave)
		for k, v := range leave.LeaveMaster {
			employeeleave.OrganisationId = employee.OrganisationID
			employeeleave.EmployeeId = employeeprobationary.EmployeeID
			fmt.Println("Leavetype===", v.UniqueID)
			fmt.Println("Leavetype K===", leave.LeaveMaster[k].UniqueID)
			employeeleave.LeaveType = v.UniqueID
			employeeleave.Value = int64(v.Value)
			employeeleave.UniqueID = s.Daos.GetUniqueID(ctx, constants.COLLECTIONEMPLOYEELEAVE)
			employeeleave.Status = constants.EMPLOYEELEAVESTATUSACTIVE
			t := time.Now()
			created := models.Created{}
			created.On = &t
			created.By = constants.SYSTEM
			log.Println("b4 EmployeeLeave.created")
			employeeleave.Created = &created
			log.Println("b4 EmployeeLeave.created")
			err := s.Daos.SaveEmployeeLeave(ctx, employeeleave)
			if err != nil {
				return err
			}
			employeeleavelog := new(models.EmployeeLeaveLog)
			employeeleavelog.UniqueID = s.Daos.GetUniqueID(ctx, constants.COLLECTIONEMPLOYEELEAVELOG)
			employeeleavelog.EmployeeId = employeeleave.EmployeeId
			employeeleavelog.LeaveType = employeeleave.LeaveType
			employeeleavelog.Value = employeeleave.Value
			employeeleavelog.Date = &t
			employeeleavelog.CreateBy = "System"
			employeeleavelog.CreateDate = &t
			employeeleavelog.Status = constants.EMPLOYEELEAVESTATUSACTIVE
			employeeleavelog.Created = &created
			employeeleavelog.Remarks = "System to Probationary"
			err = s.Daos.SaveEmployeeLeaveLog(ctx, employeeleavelog)
			if err != nil {
				return err
			}
		}
		url := s.ConfigReader.GetString(s.Shared.GetCmdArg(constants.ENV) + "." + constants.APKURL)

		msg := fmt.Sprintf(constants.COMMONTEMPLATE, employee.Name, employee.OfficialEmail, url, employee.UniqueID)
		//	msg:=fmt.Sprintf("D")
		var sendmailto []string
		fmt.Println("Employee email======>", employee.Email)
		sendmailto = append(sendmailto, employee.Email)
		err = s.SendEmail("Welcome To Logikoof", sendmailto, msg)
		if err != nil {
			return errors.New("email Sending Error - " + err.Error())
		}
		if err == nil {
			emaillog := new(models.EmailLog)
			to2 := models.ToEmailLog{}
			to2.Email = employee.Email
			to2.Name = employee.UserName
			to2.UserName = employee.UserName
			to2.UserType = "Employee"
			t := time.Now()
			emaillog.SentDate = &t
			emaillog.IsJob = false
			emaillog.Message = msg
			emaillog.SentFor = "login"
			emaillog.Status = "Active"
			emaillog.To = to2
			err = s.Daos.SaveEmailLog(ctx, emaillog)
			if err != nil {
				return errors.New("login email not save")
			}
		}
		if err := ctx.Session.CommitTransaction(sc); err != nil {
			return errors.New("Not able to commit - " + err.Error())
		}
		return nil
	}); err != nil {
		log.Println("Transaction start aborting")
		if abortError := ctx.Session.AbortTransaction(ctx.CTX); abortError != nil {
			return errors.New("Error while aborting transaction" + abortError.Error())
		}
		log.Println("Transaction aborting completed successfully")
		return err
	}

	return nil
}

//EmployeeActive : ""
func (s *Service) EmployeeActive(ctx *models.Context, employeeactive *models.EmployeeMoveToActive) error {

	userdata, err := s.Daos.GetSingleUserbyemployeeid(ctx, employeeactive.EmployeeID)
	if err != nil {
		fmt.Println(err)
		return err
	}

	datajobtimeline, err := s.Daos.GetSingleJobTimelineemployeeid(ctx, employeeactive.EmployeeID, constants.EMPLOYEESTATUSACTIVESTAGE)
	if err != nil {
		fmt.Println(err)
		return err
	}
	employee, err := s.Daos.GetSingleEmployee(ctx, employeeactive.EmployeeID)
	if err != nil {
		fmt.Println(err)
		return err
	}

	if err := ctx.Session.StartTransaction(); err != nil {
		return err
	}
	defer ctx.Session.EndSession(ctx.CTX)

	if err := mongo.WithSession(ctx.CTX, ctx.Session, func(sc mongo.SessionContext) error {
		t := time.Now()
		//employee
		employee.Status = constants.EMPLOYEESTATUSACTIVEEMPLOYEE
		employee.ConfirmationDate = &t
		employee.Remark = employeeactive.Remark
		dberr := s.Daos.EmployeeActive(ctx, &employee.Employee)
		if dberr != nil {
			return dberr
		}

		//JobTimeLine update field here
		datajobtimeline.EndDate = &t
		datajobtimeline.Status = constants.EMPLOYEESTATUSCOMPLETEDSTAGE
		dberr = s.Daos.UpdateJobTimeline(ctx, &datajobtimeline.JobTimeline)
		if dberr != nil {
			return dberr
		}

		//User
		userdata.Status = constants.EMPLOYEESTATUSACTIVEEMPLOYEE
		userdata.OrganisationID = employee.OrganisationID
		userdata.DepartmentID = employee.DepartmentID
		userdata.BranchID = employee.BranchID
		userdata.DesignationID = employee.DepartmentID
		dberr = s.Daos.UpdateUserbyemployeeId(ctx, &userdata.User)
		if dberr != nil {
			return dberr
		}

		//EmployeeLog
		employeeLog := new(models.EmployeeLog)
		employeeLog.Name = employee.Name
		employeeLog.UniqueID = s.Daos.GetUniqueID(ctx, constants.COLLECTIONJOBTIMELINE)
		employeeLog.OrganisationId = employee.OrganisationID
		employeeLog.DepartmentId = employee.DepartmentID
		employeeLog.BranchId = employee.BranchID
		employeeLog.DesignationId = employee.DepartmentID
		employeeLog.Desc = constants.EMPLOYEESTATUSACTIVEEMPLOYEEDESC
		employeeLog.Action.UserID = employeeactive.ByID
		employeeLog.Action.UserType = employeeactive.ByType
		employeeLog.Action.Date = t
		employeeLog.EmployeeId = employee.UniqueID
		employeeLog.Status = constants.EMPLOYEESTATUSACTIVEEMPLOYEE
		employeeLog.Remark = employee.Remark
		dberr = s.Daos.SaveEmployeeLog(ctx, employeeLog)
		if dberr != nil {
			return dberr
		}

		//jobtimeline
		jobtimeline := new(models.JobTimeline)
		jobtimeline.State = constants.EMPLOYEESTATUSACTIVEEMPLOYEE
		jobtimeline.UniqueID = s.Daos.GetUniqueID(ctx, constants.COLLECTIONJOBTIMELINE)
		jobtimeline.StartDate = &t
		jobtimeline.Status = constants.EMPLOYEESTATUSACTIVESTAGE
		jobtimeline.OrganisationId = employee.OrganisationID
		jobtimeline.DepartmentId = employee.DepartmentID
		jobtimeline.BranchId = employee.BranchID
		jobtimeline.DesignationId = employee.DepartmentID
		jobtimeline.Assigned.UserID = employeeactive.By
		jobtimeline.Assigned.UserType = employeeactive.ByType
		jobtimeline.Assigned.Date = t
		jobtimeline.EmployeeId = employee.UniqueID
		jobtimeline.Remark = employee.Remark
		dberr = s.Daos.SaveJobTimeline(ctx, jobtimeline)
		if dberr != nil {
			return dberr
		}

		if err := ctx.Session.CommitTransaction(sc); err != nil {
			return errors.New("Not able to commit - " + err.Error())
		}
		return nil
	}); err != nil {
		log.Println("Transaction start aborting")
		if abortError := ctx.Session.AbortTransaction(ctx.CTX); abortError != nil {
			return errors.New("Error while aborting transaction" + abortError.Error())
		}
		log.Println("Transaction aborting completed successfully")
		return err
	}

	return nil
}

//EmployeeBench : ""
func (s *Service) EmployeeBench(ctx *models.Context, employeebench *models.EmployeeMoveToBench) error {

	if err := ctx.Session.StartTransaction(); err != nil {
		return err
	}
	defer ctx.Session.EndSession(ctx.CTX)

	userdata, err := s.Daos.GetSingleUserbyemployeeid(ctx, employeebench.EmployeeID)
	if err != nil {
		fmt.Println(err)
		return err
	}

	datajobtimeline, err := s.Daos.GetSingleJobTimelineemployeeid(ctx, employeebench.EmployeeID, constants.EMPLOYEESTATUSACTIVESTAGE)
	if err != nil {
		fmt.Println(err)
		return err
	}

	employee, err := s.Daos.GetSingleEmployee(ctx, employeebench.EmployeeID)
	if err != nil {
		fmt.Println(err)
		return err
	}

	if err := mongo.WithSession(ctx.CTX, ctx.Session, func(sc mongo.SessionContext) error {

		t := time.Now()
		//Employee
		employee.Status = constants.EMPLOYEESTATUSBENCH
		employee.BenchDate = &t
		employee.Remark = employeebench.Remark
		emp := s.Daos.EmployeeBench(ctx, &employee.Employee)
		if emp != nil {
			return emp
		}
		//JobTimeLine update field here
		datajobtimeline.EndDate = &t
		datajobtimeline.Status = constants.EMPLOYEESTATUSCOMPLETEDSTAGE
		dberr := s.Daos.UpdateJobTimeline(ctx, &datajobtimeline.JobTimeline)
		if dberr != nil {
			return dberr
		}
		//User

		userdata.Status = constants.EMPLOYEESTATUSBENCH
		userdata.OrganisationID = employee.OrganisationID
		userdata.DepartmentID = employee.DepartmentID
		userdata.BranchID = employee.BranchID
		userdata.DesignationID = employee.DepartmentID
		dberr = s.Daos.UpdateUserbyemployeeId(ctx, &userdata.User)
		if dberr != nil {
			return dberr
		}

		//EmployeeLog
		employeeLog := new(models.EmployeeLog)
		employeeLog.Name = employee.Name
		employeeLog.UniqueID = s.Daos.GetUniqueID(ctx, constants.COLLECTIONJOBTIMELINE)
		employeeLog.OrganisationId = employee.OrganisationID
		employeeLog.DepartmentId = employee.DepartmentID
		employeeLog.BranchId = employee.BranchID
		employeeLog.DesignationId = employee.DepartmentID
		employeeLog.Desc = constants.EMPLOYEESTATUSBENCHDESC
		employeeLog.Action.UserID = employeebench.ByID
		employeeLog.Action.UserType = employeebench.ByType
		employeeLog.Action.Date = t
		employeeLog.EmployeeId = employee.UniqueID
		employeeLog.Status = constants.EMPLOYEESTATUSBENCH
		employeeLog.Remark = employee.Remark
		dberr = s.Daos.SaveEmployeeLog(ctx, employeeLog)
		if dberr != nil {
			return dberr
		}

		//jobtimeline
		jobtimeline := new(models.JobTimeline)
		jobtimeline.State = constants.EMPLOYEESTATUSBENCH
		jobtimeline.UniqueID = s.Daos.GetUniqueID(ctx, constants.COLLECTIONJOBTIMELINE)
		jobtimeline.StartDate = &t
		jobtimeline.Status = constants.EMPLOYEESTATUSACTIVESTAGE
		jobtimeline.OrganisationId = employee.OrganisationID
		jobtimeline.DepartmentId = employee.DepartmentID
		jobtimeline.BranchId = employee.BranchID
		jobtimeline.DesignationId = employee.DepartmentID
		jobtimeline.Assigned.UserID = employeebench.ByID
		jobtimeline.Assigned.UserType = employeebench.ByType
		jobtimeline.Assigned.Date = t
		jobtimeline.EmployeeId = employee.UniqueID
		jobtimeline.Remark = employee.Remark
		dberr = s.Daos.SaveJobTimeline(ctx, jobtimeline)
		if dberr != nil {
			return dberr
		}

		if err := ctx.Session.CommitTransaction(sc); err != nil {
			return errors.New("Not able to commit - " + err.Error())
		}
		return nil
	}); err != nil {
		log.Println("Transaction start aborting")
		if abortError := ctx.Session.AbortTransaction(ctx.CTX); abortError != nil {
			return errors.New("Error while aborting transaction" + abortError.Error())
		}
		log.Println("Transaction aborting completed successfully")
		return err
	}

	return nil
}

//EmployeeNotice : ""
func (s *Service) EmployeeNotice(ctx *models.Context, employeenotice *models.EmployeeMoveToNotice) error {

	if err := ctx.Session.StartTransaction(); err != nil {
		return err
	}
	defer ctx.Session.EndSession(ctx.CTX)

	noticedata, err := s.Daos.GetSingleNoticePolicy(ctx, employeenotice.NoticeID)
	if err != nil {
		fmt.Println(err)
		return err
	}

	userdata, err := s.Daos.GetSingleUserbyemployeeid(ctx, employeenotice.EmployeeID)
	if err != nil {
		fmt.Println(err)
		return err
	}

	employee, err := s.Daos.GetSingleEmployee(ctx, employeenotice.EmployeeID)
	if err != nil {
		fmt.Println(err)
		return err
	}

	datajobtimeline, err := s.Daos.GetSingleJobTimelineemployeeid(ctx, employee.UniqueID, constants.EMPLOYEESTATUSACTIVESTAGE)
	if err != nil {
		fmt.Println(err)
		return err
	}

	if err := mongo.WithSession(ctx.CTX, ctx.Session, func(sc mongo.SessionContext) error {

		t := time.Now()
		//Employee
		employee.Status = constants.EMPLOYEESTATUSNOTICE
		currenttimevalue := t
		employee.NoticeendDate = currenttimevalue.AddDate(0, 0, noticedata.NoticeDays)
		employee.Remark = employeenotice.Remark
		emp := s.Daos.EmployeeNotice(ctx, &employee.Employee)
		if emp != nil {
			return emp
		}

		//JobTimeLine update field here
		datajobtimeline.EndDate = &t
		datajobtimeline.Status = constants.EMPLOYEESTATUSCOMPLETEDSTAGE
		dberr := s.Daos.UpdateJobTimeline(ctx, &datajobtimeline.JobTimeline)
		if dberr != nil {
			return dberr
		}

		//User
		userdata.Status = constants.EMPLOYEESTATUSNOTICE
		userdata.OrganisationID = employee.OrganisationID
		userdata.DepartmentID = employee.DepartmentID
		userdata.BranchID = employee.BranchID
		userdata.DesignationID = employee.DepartmentID
		dberr = s.Daos.UpdateUserbyemployeeId(ctx, &userdata.User)
		if dberr != nil {
			return dberr
		}

		//EmployeeLog
		employeeLog := new(models.EmployeeLog)
		employeeLog.Name = employee.Name
		employeeLog.UniqueID = s.Daos.GetUniqueID(ctx, constants.COLLECTIONJOBTIMELINE)
		employeeLog.OrganisationId = employee.OrganisationID
		employeeLog.DepartmentId = employee.DepartmentID
		employeeLog.BranchId = employee.BranchID
		employeeLog.DesignationId = employee.DepartmentID
		employeeLog.Desc = constants.EMPLOYEESTATUSNOTICEDESC
		employeeLog.Action.UserID = employee.UniqueID
		employeeLog.Action.UserID = employeenotice.ByID
		employeeLog.Action.UserType = employeenotice.ByType
		employeeLog.Action.Date = t
		employeeLog.EmployeeId = employee.UniqueID
		employeeLog.Status = constants.EMPLOYEESTATUSNOTICE
		employeeLog.Remark = employee.Remark
		dberr = s.Daos.SaveEmployeeLog(ctx, employeeLog)
		if dberr != nil {
			return dberr
		}

		//jobtimeline
		jobtimeline := new(models.JobTimeline)
		jobtimeline.State = constants.EMPLOYEESTATUSNOTICE
		jobtimeline.UniqueID = s.Daos.GetUniqueID(ctx, constants.COLLECTIONJOBTIMELINE)
		jobtimeline.StartDate = &t
		jobtimeline.Status = constants.EMPLOYEESTATUSACTIVESTAGE
		jobtimeline.OrganisationId = employee.OrganisationID
		jobtimeline.DepartmentId = employee.DepartmentID
		jobtimeline.BranchId = employee.BranchID
		jobtimeline.DesignationId = employee.DepartmentID
		jobtimeline.Assigned.UserID = employeenotice.ByID
		jobtimeline.Assigned.UserType = employeenotice.ByType
		jobtimeline.Assigned.Date = t
		jobtimeline.EmployeeId = employee.UniqueID
		jobtimeline.Remark = employee.Remark
		dberr = s.Daos.SaveJobTimeline(ctx, jobtimeline)
		if dberr != nil {
			return dberr
		}

		if err := ctx.Session.CommitTransaction(sc); err != nil {
			return errors.New("Not able to commit - " + err.Error())
		}
		return nil
	}); err != nil {
		log.Println("Transaction start aborting")
		if abortError := ctx.Session.AbortTransaction(ctx.CTX); abortError != nil {
			return errors.New("Error while aborting transaction" + abortError.Error())
		}
		log.Println("Transaction aborting completed successfully")
		return err
	}

	return nil
}

//EmployeeOffboard : ""
func (s *Service) EmployeeOffboard(ctx *models.Context, employeeoffboard *models.EmployeeMoveToOffboard) error {

	if err := ctx.Session.StartTransaction(); err != nil {
		return err
	}
	defer ctx.Session.EndSession(ctx.CTX)

	userdata, err := s.Daos.GetSingleUserbyemployeeid(ctx, employeeoffboard.EmployeeID)
	if err != nil {
		fmt.Println(err)
		return err
	}

	datajobtimeline, err := s.Daos.GetSingleJobTimelineemployeeid(ctx, employeeoffboard.EmployeeID, constants.EMPLOYEESTATUSACTIVESTAGE)
	if err != nil {
		fmt.Println(err)
		return err
	}

	employee, err := s.Daos.GetSingleEmployee(ctx, employeeoffboard.EmployeeID)
	if err != nil {
		fmt.Println(err)
		return err
	}

	if err := mongo.WithSession(ctx.CTX, ctx.Session, func(sc mongo.SessionContext) error {

		t := time.Now()
		//Employee
		employee.Status = constants.EMPLOYEESTATUSOFFBOARD
		employee.OffBoardDate = &t
		employee.Remark = employeeoffboard.Remark
		emp := s.Daos.EmployeeOffboard(ctx, &employee.Employee)
		if emp != nil {
			return emp
		}
		//JobTimeLine update field here
		datajobtimeline.EndDate = &t
		datajobtimeline.Status = constants.EMPLOYEESTATUSCOMPLETEDSTAGE
		dberr := s.Daos.UpdateJobTimeline(ctx, &datajobtimeline.JobTimeline)
		if dberr != nil {
			return dberr
		}

		//User
		userdata.Status = constants.EMPLOYEESTATUSOFFBOARD
		userdata.OrganisationID = employee.OrganisationID
		userdata.DepartmentID = employee.DepartmentID
		userdata.BranchID = employee.BranchID
		userdata.DesignationID = employee.DepartmentID

		dberr = s.Daos.UpdateUserbyemployeeId(ctx, &userdata.User)
		if dberr != nil {
			return dberr
		}

		//EmployeeLog
		employeeLog := new(models.EmployeeLog)
		employeeLog.Name = employee.Name
		employeeLog.UniqueID = s.Daos.GetUniqueID(ctx, constants.COLLECTIONJOBTIMELINE)
		employeeLog.OrganisationId = employee.OrganisationID
		employeeLog.DepartmentId = employee.DepartmentID
		employeeLog.BranchId = employee.BranchID
		employeeLog.DesignationId = employee.DepartmentID
		employeeLog.Desc = constants.EMPLOYEESTATUSOFFBOARDDECS
		employeeLog.Action.UserID = employeeoffboard.ByID
		employeeLog.Action.UserType = employeeoffboard.ByType
		employeeLog.Action.Date = t
		employeeLog.EmployeeId = employee.UniqueID
		employeeLog.Status = constants.EMPLOYEESTATUSOFFBOARD
		employeeLog.Remark = employee.Remark
		dberr = s.Daos.SaveEmployeeLog(ctx, employeeLog)
		if dberr != nil {
			return dberr
		}

		//jobtimeline
		jobtimeline := new(models.JobTimeline)
		jobtimeline.State = constants.EMPLOYEESTATUSOFFBOARD
		jobtimeline.UniqueID = s.Daos.GetUniqueID(ctx, constants.COLLECTIONJOBTIMELINE)
		jobtimeline.StartDate = &t
		jobtimeline.Status = constants.EMPLOYEESTATUSACTIVESTAGE
		jobtimeline.OrganisationId = employee.OrganisationID
		jobtimeline.DepartmentId = employee.DepartmentID
		jobtimeline.BranchId = employee.BranchID
		jobtimeline.DesignationId = employee.DepartmentID
		jobtimeline.Assigned.UserID = employeeoffboard.ByID
		jobtimeline.Assigned.UserType = employeeoffboard.ByType
		jobtimeline.Assigned.Date = t
		jobtimeline.EmployeeId = employee.UniqueID
		jobtimeline.Remark = employee.Remark
		dberr = s.Daos.SaveJobTimeline(ctx, jobtimeline)
		if dberr != nil {
			return dberr
		}

		if err := ctx.Session.CommitTransaction(sc); err != nil {
			return errors.New("Not able to commit - " + err.Error())
		}
		return nil
	}); err != nil {
		log.Println("Transaction start aborting")
		if abortError := ctx.Session.AbortTransaction(ctx.CTX); abortError != nil {
			return errors.New("Error while aborting transaction" + abortError.Error())
		}
		log.Println("Transaction aborting completed successfully")
		return err
	}

	return nil
}

//EmployeeRelieve : ""
func (s *Service) EmployeeRelieve(ctx *models.Context, employeerelieve *models.EmployeeMoveToRelieve) error {

	if err := ctx.Session.StartTransaction(); err != nil {
		return err
	}
	defer ctx.Session.EndSession(ctx.CTX)

	userdata, err := s.Daos.GetSingleUserbyemployeeid(ctx, employeerelieve.EmployeeID)
	if err != nil {
		fmt.Println(err)
		return err
	}

	datajobtimeline, err := s.Daos.GetSingleJobTimelineemployeeid(ctx, employeerelieve.EmployeeID, constants.EMPLOYEESTATUSACTIVESTAGE)
	if err != nil {
		fmt.Println(err)
		return err
	}
	employee, err := s.Daos.GetSingleEmployee(ctx, employeerelieve.EmployeeID)
	if err != nil {
		fmt.Println(err)
		return err
	}

	if err := mongo.WithSession(ctx.CTX, ctx.Session, func(sc mongo.SessionContext) error {
		t := time.Now()
		//Employee
		employee.Status = constants.EMPLOYEESTATUSRELIEVE
		employee.RelieveDate = &t
		employee.Remark = employeerelieve.Remark
		emp := s.Daos.EmployeeRelieve(ctx, &employee.Employee)
		if emp != nil {
			return emp
		}

		//JobTimeLine update field here
		datajobtimeline.EndDate = &t
		datajobtimeline.Status = constants.EMPLOYEESTATUSCOMPLETEDSTAGE
		dberr := s.Daos.UpdateJobTimeline(ctx, &datajobtimeline.JobTimeline)
		if dberr != nil {
			return dberr
		}

		//User
		userdata.Status = constants.EMPLOYEESTATUSRELIEVE
		userdata.OrganisationID = employee.OrganisationID
		userdata.DepartmentID = employee.DepartmentID
		userdata.BranchID = employee.BranchID
		userdata.DesignationID = employee.DepartmentID
		dberr = s.Daos.UpdateUserbyemployeeId(ctx, &userdata.User)
		if dberr != nil {
			return dberr
		}

		//EmployeeLog
		employeeLog := new(models.EmployeeLog)
		employeeLog.Name = employee.Name
		employeeLog.UniqueID = s.Daos.GetUniqueID(ctx, constants.COLLECTIONJOBTIMELINE)
		employeeLog.OrganisationId = employee.OrganisationID
		employeeLog.DepartmentId = employee.DepartmentID
		employeeLog.BranchId = employee.BranchID
		employeeLog.DesignationId = employee.DepartmentID
		employeeLog.Desc = constants.EMPLOYEESTATUSRELIEVEDECS
		employeeLog.Action.UserID = employee.UniqueID
		employeeLog.Action.UserType = employee.Role
		employeeLog.Action.Date = t
		employeeLog.EmployeeId = employee.UniqueID
		employeeLog.Status = constants.EMPLOYEESTATUSRELIEVE
		employeeLog.Remark = employee.Remark
		dberr = s.Daos.SaveEmployeeLog(ctx, employeeLog)
		if dberr != nil {
			return dberr
		}

		//jobtimeline
		jobtimeline := new(models.JobTimeline)
		jobtimeline.State = constants.EMPLOYEESTATUSRELIEVE
		jobtimeline.UniqueID = s.Daos.GetUniqueID(ctx, constants.COLLECTIONJOBTIMELINE)
		jobtimeline.StartDate = &t
		jobtimeline.Status = constants.EMPLOYEESTATUSACTIVESTAGE
		jobtimeline.OrganisationId = employee.OrganisationID
		jobtimeline.DepartmentId = employee.DepartmentID
		jobtimeline.BranchId = employee.BranchID
		jobtimeline.DesignationId = employee.DepartmentID
		jobtimeline.Assigned.UserType = employee.Role
		jobtimeline.Assigned.Date = t
		jobtimeline.EmployeeId = employee.UniqueID
		jobtimeline.Remark = employee.Remark
		dberr = s.Daos.SaveJobTimeline(ctx, jobtimeline)
		if dberr != nil {
			return dberr
		}

		if err := ctx.Session.CommitTransaction(sc); err != nil {
			return errors.New("Not able to commit - " + err.Error())
		}
		return nil
	}); err != nil {
		log.Println("Transaction start aborting")
		if abortError := ctx.Session.AbortTransaction(ctx.CTX); abortError != nil {
			return errors.New("Error while aborting transaction" + abortError.Error())
		}
		log.Println("Transaction aborting completed successfully")
		return err
	}

	return nil
}
func (s *Service) UpdateEmployeeBioData(ctx *models.Context, employee *models.UpdateBioData, UniqueID string) error {
	if err := ctx.Session.StartTransaction(); err != nil {
		return err
	}
	defer ctx.Session.EndSession(ctx.CTX)
	if err := mongo.WithSession(ctx.CTX, ctx.Session, func(sc mongo.SessionContext) error {

		err := s.Daos.UpdateEmployeeBioData(ctx, employee, UniqueID)
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
func (s *Service) UpdateEmployeeEmergencyContact(ctx *models.Context, employee *models.UpdateEmergencyContact, UniqueID string) error {
	if err := ctx.Session.StartTransaction(); err != nil {
		return err
	}
	defer ctx.Session.EndSession(ctx.CTX)
	if err := mongo.WithSession(ctx.CTX, ctx.Session, func(sc mongo.SessionContext) error {

		err := s.Daos.UpdateEmployeeEmergencyContact(ctx, employee, UniqueID)
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

func (s *Service) UpdateEmployeePersonalInformation(ctx *models.Context, employee *models.PersonalInformation, UniqueID string) error {
	if err := ctx.Session.StartTransaction(); err != nil {
		return err
	}
	defer ctx.Session.EndSession(ctx.CTX)
	if err := mongo.WithSession(ctx.CTX, ctx.Session, func(sc mongo.SessionContext) error {

		err := s.Daos.UpdateEmployeePersonalInformation(ctx, employee, UniqueID)
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
func (s *Service) EmployeeDayWiseAttendanceReport(ctx *models.Context, filter *models.DayWiseAttendanceReportFilter, pagination *models.Pagination) ([]models.EmployeeDayWiseAttendanceReport, error) {
	err := s.EmployeeDayWiseAttendanceAccess(ctx, filter)
	if err != nil {
		return nil, err
	}
	defer ctx.Session.EndSession(ctx.CTX)
	employereport, err := s.Daos.EmployeeDayWiseAttendanceReport(ctx, filter, pagination)
	if err != nil {
		return nil, err
	}
	for k, v := range employereport {
		var overtime float64
		for k2, v2 := range v.Days {
			overtime = overtime + v2.Attendance.OverTime
			v.Days[k2].Dates = int64(k2) + 1
			date := time.Date(filter.StartDate.Year(), filter.StartDate.Month(), k2+1, 0, 0, 0, 0, filter.StartDate.Location())
			v.Days[k2].Date = &date
			///fmt.Println("Payroll===>", v2.Attendance.PayRoll)
			if v2.Attendance.PayRoll == constants.ATTENDANCESTATUSLOP {
				employereport[k].NoOfLOP = employereport[k].NoOfLOP + 1
			}
			if v2.Attendance.PayRoll == constants.ATTENDANCESTATUSPARTIALPAY {
				employereport[k].NoOfParticalPaid = employereport[k].NoOfParticalPaid + 1
			}
			if v2.Attendance.PayRoll == constants.ATTENDANCESTATUSPAID {
				employereport[k].NoOfPaid = employereport[k].NoOfPaid + 1
			}
			if v2.Attendance.PayRoll == constants.ATTENDANCESTATUSHOLIDAY {
				employereport[k].NoOfHolidays = employereport[k].NoOfHolidays + 1
			}

			if v2.Attendance.TotalWorkingHoursStr == "" {
				v.Days[k2].Attendance.TotalWorkingHoursStr = fmt.Sprintf("%vh%vmins", 0, 0)
			}
			if v2.Attendance.Status == constants.ATTENDANCESTATUSPENDING {
				employereport[k].PendingStatus = employereport[k].PendingStatus + 1
			}
		}
		hours := overtime / 60
		Minutes := int64(overtime) % 60
		employereport[k].Overtime = fmt.Sprintf("%vh%vmins", int64(hours), Minutes)
		fmt.Println("Employee====Dayes", v.Name, len(v.Days))
	}

	return employereport, nil

}
func (s *Service) EmployeeDayWiseAttendanceAccess(ctx *models.Context, filter *models.DayWiseAttendanceReportFilter) (err error) {
	if filter != nil {
		dataaccess, err := s.Daos.DataAccess(ctx, &filter.DataAccess)
		if err != nil {
			return err
		}
		if dataaccess != nil {
			if len(dataaccess.Organisation) > 0 {
				for _, v := range dataaccess.Organisation {
					filter.OrganisationId = append(filter.OrganisationId, v.UniqueID)
				}
			}
			fmt.Println(" dataaccess.SuperAdmin===>", dataaccess.SuperAdmin)
			if dataaccess.SuperAdmin {
				filter.Manager = ""
			}

		}

	}
	return err
}
func (s *Service) UpdateEmployeeProfileImage(ctx *models.Context, employee *models.UpdateBioData, UniqueID string) error {
	if err := ctx.Session.StartTransaction(); err != nil {
		return err
	}
	defer ctx.Session.EndSession(ctx.CTX)
	if err := mongo.WithSession(ctx.CTX, ctx.Session, func(sc mongo.SessionContext) error {

		dberr := s.Daos.UpdateEmployeeProfileImage(ctx, employee, UniqueID)
		if dberr != nil {
			return dberr
		}

		if err := ctx.Session.CommitTransaction(sc); err != nil {
			return errors.New("Not able to commit - " + err.Error())
		}
		err := s.Daos.UpdateUserProfileImage(ctx, employee, UniqueID)
		if err != nil {
			return err
		}

		return nil

	}); err != nil {
		log.Println("Transaction start aborting")
		if abortError := ctx.Session.AbortTransaction(ctx.CTX); abortError != nil {
			return errors.New("Error while aborting transaction" + abortError.Error())
		}
		log.Println("Transaction aborting completed successfully")
		return err
	}
	return nil
}
func (s *Service) EmployeeDayWiseAttendanceReportExcel(ctx *models.Context, filter *models.DayWiseAttendanceReportFilter) (*excelize.File, error) {
	t := time.Now()
	data, err := s.EmployeeDayWiseAttendanceReport(ctx, filter, nil)
	if err != nil {
		return nil, err
	}
	duration := time.Since(t)
	log.Println("query Time taken ===> ", duration.Minutes(), "m")
	t = time.Now()
	// var k []string
	// for _, v := range data {
	// 	if len(v.Days) > 0 {
	// 		switch len(v.Days) {
	// 		case 31:
	// 			k = append(k, "C", "D", "E", "F", "G", "H", "I", "J", "K", "L", "M", "N", "O", "P", "Q", "R", "S", "T", "U", "V", "W", "X", "Y", "Z", "AA", "AB", "AC", "AD", "AE", "AF", "AG")
	// 		case 30:
	// 			k = append(k, "C", "D", "E", "F", "G", "H", "I", "J", "K", "L", "M", "N", "O", "P", "Q", "R", "S", "T", "U", "V", "W", "X", "Y", "Z", "AA", "AB", "AC", "AD", "AE", "AF")
	// 		case 28:
	// 			k = append(k, "C", "D", "E", "F", "G", "H", "I", "J", "K", "L", "M", "N", "O", "P", "Q", "R", "S", "T", "U", "V", "W", "X", "Y", "Z", "AA", "AB", "AC", "AD")
	// 		default:
	// 			k = append(k, "C", "D", "E", "F", "G", "H", "I", "J", "K", "L", "M", "N", "O", "P", "Q", "R", "S", "T", "U", "V", "W", "X", "Y", "Z", "AA", "AB", "AC", "AD", "AE")
	// 		}
	// 	}
	// }
	// n := len(k)
	excel := excelize.NewFile()
	sheet1 := "EmployeesDayWiseReport"
	index := excel.NewSheet(sheet1)
	excel.SetActiveSheet(index)
	rowNo := 1
	style1, err := excel.NewStyle(`{"fill":{"type":"pattern","color":["#FFDC6D"],"pattern":1},"alignment":{"horizontal":"center","vertical":"center"},"font":{"bold":true}}`)
	if err != nil {
		fmt.Println(err)
	}
	style3, err := excel.NewStyle(`{"fill":{"type":"pattern","color":["#B6DDE8"],"pattern":1},"alignment":{"horizontal":"center","vertical":"center"},"font":{"bold":true}}`)
	if err != nil {
		fmt.Println(err)
	}
	Lop, err := excel.NewStyle(`{"fill":{"type":"pattern","color":["#FF0000"],"pattern":1},"alignment":{"horizontal":"center","vertical":"center"},"font":{"bold":false}}`)
	if err != nil {
		fmt.Println(err)
	}
	paid, err := excel.NewStyle(`{"fill":{"type":"pattern","color":["#00FF00"],"pattern":1},"alignment":{"horizontal":"center","vertical":"center"},"font":{"bold":false}}`)
	if err != nil {
		fmt.Println(err)
	}
	particalpaid, err := excel.NewStyle(`{"fill":{"type":"pattern","color":["#FFFFE0"],"pattern":1},"alignment":{"horizontal":"center","vertical":"center"},"font":{"bold":false}}`)
	if err != nil {
		fmt.Println(err)
	}
	holiday, err := excel.NewStyle(`{"fill":{"type":"pattern","color":["#81C3B4"],"pattern":1},"alignment":{"horizontal":"center","vertical":"center"},"font":{"bold":false}}`)
	if err != nil {
		fmt.Println(err)
	}
	// excel.MergeCell(sheet1, "C1", "C3")
	// excel.MergeCell(sheet1, "C4", "C5")

	// documentUrl := s.Shared.Config.GetString(s.Shared.GetCmdArg(constants.ENV) + "." + constants.FILEURL)
	// if err := excel.AddPicture(sheet1, fmt.Sprintf("%v%v", "A", rowNo), documentUrl+"municipal/logo.png", `{"x_scale": 0.115, "y_scale": 0.0935}`); err != nil {
	// 	fmt.Println(err)
	// }
	// excel.SetCellStyle(sheet1, fmt.Sprintf("%v%v", "C", rowNo), fmt.Sprintf("%v%v", "C", rowNo), style3)
	// excel.SetCellValue(sheet1, fmt.Sprintf("%v%v", "C", rowNo), ctx.ProductConfig.Name)
	// title :=

	//Kitchen := "22-08-2022"
	rowNo++
	rowNo++
	if len(data) > 0 {
		ch := 'A'
		excel.SetCellValue(sheet1, fmt.Sprintf("%c%v", ch, rowNo), "S.No")
		ch++
		excel.SetCellValue(sheet1, fmt.Sprintf("%c%v", ch, rowNo), "EmployeeName")
		ch++
		excel.SetCellValue(sheet1, fmt.Sprintf("%c%v", ch, rowNo), "EmployeeId")
		ch++
		excel.SetCellValue(sheet1, fmt.Sprintf("%c%v", ch, rowNo), "Designation")
		ch++
		excel.SetCellValue(sheet1, fmt.Sprintf("%c%v", ch, rowNo), "Branch")
		ch++

		var ch1 rune
		var ch2 rune
		ch1 = 'A'
		ch2 = 'A'
		column := ""
		for _, v := range data[0].Days {
			rowNo = 3
			if ch < 'Z' {
				column = fmt.Sprintf("%c", ch)
				day := fmt.Sprintf("%v%v%v", v.Date.Day(), v.Date.Month(), v.Date.Year())
				excel.SetCellValue(sheet1, fmt.Sprintf("%v%v", column, rowNo), day)

			} else if ch == 'Z' {
				column = fmt.Sprintf("%c", ch)
				day := fmt.Sprintf("%v%v%v", v.Date.Day(), v.Date.Month(), v.Date.Year())
				excel.SetCellValue(sheet1, fmt.Sprintf("%v%v", column, rowNo), day)
			}
			if ch > 'Z' {
				CH1 := fmt.Sprintf("%c", ch1)
				CH2 := fmt.Sprintf("%c", ch2)
				fmt.Println("ch1===>", CH1)
				fmt.Println("ch2===>", CH2)
				column = fmt.Sprintf("%c%c", ch1, ch2)
				day := fmt.Sprintf("%v%v%v", v.Date.Day(), v.Date.Month(), v.Date.Year())
				excel.SetCellValue(sheet1, fmt.Sprintf("%v%v", column, rowNo), day)

				//CH := fmt.Sprintf("%c", ch)
				fmt.Println("column===>", column)
			}
			//CH := fmt.Sprintf("%c", ch)
			fmt.Println("column===>", column)
			if ch > 'Z' {
				ch2++
			}
			ch++
		}
		CH1 := fmt.Sprintf("%c", ch1)
		CH2 := fmt.Sprintf("%c", ch2)
		fmt.Println("ch1===>", CH1)
		fmt.Println("ch2===>", CH2)
		column = fmt.Sprintf("%c%c", ch1, ch2)
		excel.SetCellValue(sheet1, fmt.Sprintf("%v%v", column, rowNo), "No.Of.Paid")
		ch2++
		CH1 = fmt.Sprintf("%c", ch1)
		CH2 = fmt.Sprintf("%c", ch2)
		fmt.Println("ch1===>", CH1)
		fmt.Println("ch2===>", CH2)
		column = fmt.Sprintf("%c%c", ch1, ch2)
		excel.SetCellValue(sheet1, fmt.Sprintf("%v%v", column, rowNo), "No.Of.PartialPay")
		ch2++
		CH1 = fmt.Sprintf("%c", ch1)
		CH2 = fmt.Sprintf("%c", ch2)
		fmt.Println("ch1===>", CH1)
		fmt.Println("ch2===>", CH2)
		column = fmt.Sprintf("%c%c", ch1, ch2)
		excel.SetCellValue(sheet1, fmt.Sprintf("%v%v", column, rowNo), "No.Of.Lop")
		ch2++
		CH1 = fmt.Sprintf("%c", ch1)
		CH2 = fmt.Sprintf("%c", ch2)
		fmt.Println("ch1===>", CH1)
		fmt.Println("ch2===>", CH2)
		column = fmt.Sprintf("%c%c", ch1, ch2)
		excel.SetCellValue(sheet1, fmt.Sprintf("%v%v", column, rowNo), "No.Of.Holidays")
		excel.SetCellStyle(sheet1, fmt.Sprintf("%v%v", "A", rowNo), fmt.Sprintf("%v%v", column, rowNo), style1)

		column1 := fmt.Sprintf("%v%v", column, 1)
		excel.MergeCell(sheet1, "A1", column1)
		excel.SetCellStyle(sheet1, fmt.Sprintf("%v%v", "A", 1), fmt.Sprintf("%v%v", column, 1), style3)
		excel.SetCellValue(sheet1, fmt.Sprintf("%v%v", "C", 1), fmt.Sprintf("%v-%v%v", sheet1, filter.StartDate.Month(), filter.StartDate.Year()))

	}
	rowNo++
	rowNo++
	//	var totalAmount float64
	for k, v := range data {
		ch := 'A'
		excel.SetCellValue(sheet1, fmt.Sprintf("%c%v", ch, rowNo), k+1)
		ch++
		excel.SetCellValue(sheet1, fmt.Sprintf("%c%v", ch, rowNo), v.Name)
		ch++
		excel.SetCellValue(sheet1, fmt.Sprintf("%c%v", ch, rowNo), v.UserName)
		ch++
		excel.SetCellValue(sheet1, fmt.Sprintf("%c%v", ch, rowNo), v.Ref.DesignationID.Name)
		ch++
		excel.SetCellValue(sheet1, fmt.Sprintf("%c%v", ch, rowNo), v.Ref.BranchID.Name)
		ch++
		var ch1 rune
		var ch2 rune
		ch1 = 'A'
		ch2 = 'A'
		column := ""
		for _, v2 := range v.Days {
			if ch < 'Z' {
				column = fmt.Sprintf("%c", ch)
				excel.SetCellValue(sheet1, fmt.Sprintf("%v%v", column, rowNo), v2.Attendance.PayRoll)
				if v2.Attendance.PayRoll == "LOP" {
					excel.SetCellStyle(sheet1, fmt.Sprintf("%v%v", column, rowNo), fmt.Sprintf("%v%v", column, rowNo), Lop)
				} else if v2.Attendance.PayRoll == "Paid" {
					excel.SetCellStyle(sheet1, fmt.Sprintf("%v%v", column, rowNo), fmt.Sprintf("%v%v", column, rowNo), paid)
				} else if v2.Attendance.PayRoll == "PartialPay" {
					excel.SetCellStyle(sheet1, fmt.Sprintf("%v%v", column, rowNo), fmt.Sprintf("%v%v", column, rowNo), particalpaid)
				} else {
					excel.SetCellStyle(sheet1, fmt.Sprintf("%v%v", column, rowNo), fmt.Sprintf("%v%v", column, rowNo), holiday)
				}

			} else if ch == 'Z' {
				column = fmt.Sprintf("%c", ch)
				excel.SetCellValue(sheet1, fmt.Sprintf("%v%v", column, rowNo), v2.Attendance.PayRoll)
				if v2.Attendance.PayRoll == "LOP" {
					excel.SetCellStyle(sheet1, fmt.Sprintf("%v%v", column, rowNo), fmt.Sprintf("%v%v", column, rowNo), Lop)
				} else if v2.Attendance.PayRoll == "Paid" {
					excel.SetCellStyle(sheet1, fmt.Sprintf("%v%v", column, rowNo), fmt.Sprintf("%v%v", column, rowNo), paid)
				} else if v2.Attendance.PayRoll == "PartialPay" {
					excel.SetCellStyle(sheet1, fmt.Sprintf("%v%v", column, rowNo), fmt.Sprintf("%v%v", column, rowNo), particalpaid)
				} else {
					excel.SetCellStyle(sheet1, fmt.Sprintf("%v%v", column, rowNo), fmt.Sprintf("%v%v", column, rowNo), holiday)
				}
			}
			if ch > 'Z' {
				CH1 := fmt.Sprintf("%c", ch1)
				CH2 := fmt.Sprintf("%c", ch2)
				fmt.Println("ch1===>", CH1)
				fmt.Println("ch2===>", CH2)
				column = fmt.Sprintf("%c%c", ch1, ch2)
				excel.SetCellValue(sheet1, fmt.Sprintf("%v%v", column, rowNo), v2.Attendance.PayRoll)
				if v2.Attendance.PayRoll == "LOP" {
					excel.SetCellStyle(sheet1, fmt.Sprintf("%v%v", column, rowNo), fmt.Sprintf("%v%v", column, rowNo), Lop)
				} else if v2.Attendance.PayRoll == "Paid" {
					excel.SetCellStyle(sheet1, fmt.Sprintf("%v%v", column, rowNo), fmt.Sprintf("%v%v", column, rowNo), paid)
				} else if v2.Attendance.PayRoll == "PartialPay" {
					excel.SetCellStyle(sheet1, fmt.Sprintf("%v%v", column, rowNo), fmt.Sprintf("%v%v", column, rowNo), particalpaid)
				} else {
					excel.SetCellStyle(sheet1, fmt.Sprintf("%v%v", column, rowNo), fmt.Sprintf("%v%v", column, rowNo), holiday)
				}
				//CH := fmt.Sprintf("%c", ch)
				if ch > 'Z' {
					ch2++
				}

			}
			fmt.Println("column Row===>", column)

			ch++
		}
		CH1 := fmt.Sprintf("%c", ch1)
		CH2 := fmt.Sprintf("%c", ch2)
		fmt.Println("ch1===>", CH1)
		fmt.Println("ch2===>", CH2)
		column = fmt.Sprintf("%c%c", ch1, ch2)
		excel.SetCellValue(sheet1, fmt.Sprintf("%v%v", column, rowNo), v.NoOfPaid)
		ch2++
		CH1 = fmt.Sprintf("%c", ch1)
		CH2 = fmt.Sprintf("%c", ch2)
		fmt.Println("ch1===>", CH1)
		fmt.Println("ch2===>", CH2)
		column = fmt.Sprintf("%c%c", ch1, ch2)
		excel.SetCellValue(sheet1, fmt.Sprintf("%v%v", column, rowNo), v.NoOfParticalPaid)
		ch2++
		CH1 = fmt.Sprintf("%c", ch1)
		CH2 = fmt.Sprintf("%c", ch2)
		fmt.Println("ch1===>", CH1)
		fmt.Println("ch2===>", CH2)
		column = fmt.Sprintf("%c%c", ch1, ch2)
		excel.SetCellValue(sheet1, fmt.Sprintf("%v%v", column, rowNo), v.NoOfLOP)
		ch2++
		CH1 = fmt.Sprintf("%c", ch1)
		CH2 = fmt.Sprintf("%c", ch2)
		fmt.Println("ch1===>", CH1)
		fmt.Println("ch2===>", CH2)
		column = fmt.Sprintf("%c%c", ch1, ch2)
		excel.SetCellValue(sheet1, fmt.Sprintf("%v%v", column, rowNo), v.NoOfHolidays)
		rowNo++
	}
	//excel.MergeCell(sheet1, "A", "D")
	//excel.SetCellStyle(sheet1, fmt.Sprintf("%v%v", "A", rowNo), fmt.Sprintf("%v%v", "A", rowNo), style1)
	//excel.SetCellStyle(sheet1, fmt.Sprintf("%v%v", "B", rowNo), fmt.Sprintf("%v%v", "B", rowNo), style1)
	// excel.SetCellStyle(sheet1, fmt.Sprintf("%v%v", "C", rowNo), fmt.Sprintf("%v%v", "C", rowNo), style1)
	// //	excel.SetCellValue(sheet1, fmt.Sprintf("%v%v", "C", rowNo), fmt.Sprintf(" %.2f", totalAmount))

	duration = time.Since(t)
	log.Println("excel Time taken ===> ", duration.Minutes(), "m")
	return excel, nil

}
func (s *Service) EmployeeUploadExcel(ctx *models.Context, file multipart.File) []models.EmployeeUploadError {
	log.Println("transaction start")
	//Start Transaction
	// orgRefMap := make(map[string]primitive.ObjectID)
	// projectRefMap := make(map[string]primitive.ObjectID)
	// stateRefMap := make(map[string]primitive.ObjectID)
	// districtRefMap := make(map[string]primitive.ObjectID)
	// blockRefMap := make(map[string]primitive.ObjectID)
	// grampRefMap := make(map[string]primitive.ObjectID)
	// villageRefMap := make(map[string]primitive.ObjectID)
	var errs []models.EmployeeUploadError
	var employeeerr models.EmployeeUploadError
	if err := ctx.Session.StartTransaction(); err != nil {
		employeeerr.Error = err.Error()
		errs = append(errs, employeeerr)
		return errs
	}
	defer ctx.Session.EndSession(ctx.CTX)
	const (
		MAXCOLUMN              = 11
		OMITROWS               = 0
		ORGANISATIONCOLUMN     = 0
		FIRSTNAMECOLUMN        = 1
		LASTNAMECOLUMN         = 2
		FATHERNAMECOLUMN       = 3
		DOBCOLUMN              = 4
		GENDERCOLUMN           = 5
		MOBILENOCOLUMN         = 6
		PERSONALEMAILCOLUMN    = 7
		OFFICIALEMAILCOLUMN    = 8
		ONBOARDINGPOLICYCOLUMN = 9
		JOININGDATECOLUMN      = 10
		CREATEDDATELAYOUT      = "02-January-2006"
	)
	if err := mongo.WithSession(ctx.CTX, ctx.Session, func(sc mongo.SessionContext) error {
		fmt.Println("started reading file")
		f, err := excelize.OpenReader(file)
		if err != nil {
			return err
		}
		employees := make([]models.Employee, 0)
		rows := f.GetRows("Sheet1")
		//var errors []string
		fmt.Println("started looping")
		for rowIndex, row := range rows {
			fmt.Println("row no === ", rowIndex)
			if rowIndex <= OMITROWS {
				continue
			}
			if len(row) < MAXCOLUMN {
				return errors.New("Excel is not upto the format")
			}
			employee := new(models.Employee)
			employeevaildation, _ := s.Daos.GetsingleEmployeeWithMobileNumber(ctx, row[MOBILENOCOLUMN])
			if employeevaildation != nil {
				employeeerr.Name = row[FIRSTNAMECOLUMN]
				employeeerr.UserName = employeevaildation.UniqueID
				employeeerr.MobileNumber = row[MOBILENOCOLUMN]
				employeeerr.Error = "Employee MobileNumber Already Registered"
				errs = append(errs, employeeerr)
				continue
			} else {
				employee.Mobile = row[MOBILENOCOLUMN]
			}
			organisation, _ := s.Daos.GetSingleActiveOrganisationWithName(ctx, row[ORGANISATIONCOLUMN])
			if organisation == nil {
				employeeerr.Name = row[FIRSTNAMECOLUMN]
				employeeerr.MobileNumber = row[MOBILENOCOLUMN]
				employeeerr.Error = "organisation Not Found"
				errs = append(errs, employeeerr)
				continue
			}
			if organisation != nil {
				employee.OrganisationID = organisation.UniqueID
			}
			employee.FatherName = row[FATHERNAMECOLUMN]
			employee.Name = row[FIRSTNAMECOLUMN] + row[LASTNAMECOLUMN]
			employee.Gender = row[GENDERCOLUMN]
			employee.OfficialEmail = row[OFFICIALEMAILCOLUMN]
			employee.Email = row[PERSONALEMAILCOLUMN]
			onboardingpolicy, _ := s.Daos.GetSingleActiveOnboardingPolicyWithName(ctx, row[ONBOARDINGPOLICYCOLUMN])
			if onboardingpolicy == nil {
				employeeerr.Name = row[FIRSTNAMECOLUMN]
				employeeerr.MobileNumber = row[MOBILENOCOLUMN]
				employeeerr.Error = "onboardingpolicy Not Found"
				errs = append(errs, employeeerr)
				continue
			}
			if onboardingpolicy != nil {
				employee.OnboardingpolicyID = onboardingpolicy.UniqueID
			}
			if row[DOBCOLUMN] != "" {
				layout := CREATEDDATELAYOUT
				t, err := time.Parse(layout, row[DOBCOLUMN])
				if err != nil {
					return err
				}
				employee.DOB = &t
			}
			if row[JOININGDATECOLUMN] != "" {
				layout := CREATEDDATELAYOUT
				t, err := time.Parse(layout, row[JOININGDATECOLUMN])
				if err != nil {
					return err
				}
				employee.JoiningDate = &t
			}
			err = s.SaveEmployeeWithoutTransaction(ctx, employee)
			if err != nil {
				employeeerr.Name = row[FIRSTNAMECOLUMN]
				employeeerr.MobileNumber = row[MOBILENOCOLUMN]
				employeeerr.Error = err.Error()
				errs = append(errs, employeeerr)
				continue
			}
			employees = append(employees, *employee)
			if err == nil {
				employeeerr.Name = row[FIRSTNAMECOLUMN]
				employeeerr.UserName = employee.UniqueID
				employeeerr.MobileNumber = row[MOBILENOCOLUMN]
				employeeerr.Error = "Sucess"
				errs = append(errs, employeeerr)
				continue
			}
		}
		fmt.Println("no.of.employee==>", len(employees))
		return nil

	}); err != nil {
		employeeerr.Error = err.Error()
		errs = append(errs, employeeerr)
		return errs
	}
	return errs
}
func (s *Service) DashboardEmployeeCount(ctx *models.Context, filter *models.FilterEmployee) ([]models.DashboardEmployeeCount, error) {
	defer ctx.Session.EndSession(ctx.CTX)
	err := s.EmployeeDataAccess(ctx, filter)
	if err != nil {
		return nil, err
	}
	return s.Daos.DashboardEmployeeCount(ctx, filter)

}

// func (s *Service) EmployeeTree() {
// 	c := context.TODO()
// 	ctx := app.GetApp(c, s.Daos)
// 	employeetree := new(models.EmployeeTree)

// 	employee, err := s.Daos.GetParetentEmployee(ctx)
// 	if err != nil {
// 		fmt.Println(err)
// 	}
// 	fmt.Println("GetParetentEmployee==>", employee)
// 	childemployee, err := s.Daos.GetChildEmployee(ctx, employee.UniqueID)
// 	if err != nil {
// 		fmt.Println(err)
// 	}
// 	employeetree.Employee = employee
// 	employeetree.Child = childemployee
// 	fmt.Println("employee=====>", employeetree)
// }
func (s *Service) FindChild(ctx *models.Context) (*models.EmployeeTree, error) {
	//fmt.Println("mmmm==========>", tree)
	var tree *models.EmployeeTree
	tree, err := s.Daos.GetParetentEmployee(ctx)
	if err != nil {
		return nil, err
	}
	fmt.Println("pare====>", tree)
	if tree != nil {
		childemployee, err := s.Daos.GetChildEmployee(ctx, tree.Employee.UniqueID)
		if err != nil {
			return nil, err
		}
		if childemployee != nil {
			tree.Child = childemployee
		}
	}

	return tree, nil
}

func (s *Service) GetLineManagerEmployee(ctx *models.Context, UniqueID string) ([]models.LineManagerEmployee, error) {
	Employee, err := s.Daos.GetLineManagerEmployee(ctx, UniqueID)
	if err != nil {
		return nil, err
	}
	return Employee, nil
}
func (s *Service) FindChildwithcache(ctx *models.Context, tree *models.EmployeeTreev2) (*models.EmployeeTreev2, error) {
	//fmt.Println("mmmm==========>", tree)
	//var tree *models.EmployeeTreev2
	if tree.Employee == nil {
		trees, err := s.Daos.GetParetentEmployee(ctx)
		if err != nil {
			return nil, err
		}
		tree.Employee = trees.Employee
		fmt.Println("pare====>", tree)
	}
	//var childemployee models.LineManagerEmployee
	if tree != nil {
		childemployee, err := s.Daos.GetChildEmployeeWithtree(ctx, tree.Employee.UniqueID)
		if err != nil {
			return nil, err
		}
		if childemployee != nil {
			tree.Child = childemployee
		}
	}
	for _, v := range tree.Child {

		key := fmt.Sprintf("%v_%v", tree.Employee.Name, tree.Employee.UniqueID)
		err := s.Redis.SetValue(key, v, 1111)
		if err != nil {
			return nil, err
		}
		_, err = s.FindChildwithcache(ctx, &v)
		if err != nil {
			return nil, err
		}
		var fam []interface{}

		value := s.Redis.GetValue(key)
		fam = append(fam, value)
		fmt.Println("fam-====>", fam)

	}

	return tree, nil
}
func (s *Service) GetallEmployee(ctx *models.Context) ([]models.AllEmployees, error) {
	defer ctx.Session.EndSession(ctx.CTX)
	employees, err := s.Daos.GetallEmployee(ctx)
	if err != nil {
		return nil, err
	}
	for k, v := range employees {
		for _, v2 := range v.Employee {
			linemanger, err := s.Daos.GetSingleEmployee(ctx, v2.LineManager)
			if err != nil {
				return nil, err
			}
			if linemanger != nil {
				employees[k].Employee = append(employees[k].Employee, linemanger.Employee)
			}

		}
	}
	return employees, nil
}
func (s *Service) GetOrg(ctx *models.Context, employee []models.AllEmployees) (map[string][]models.Employee, error) {
	relations := make(map[string][]models.Employee)

	for _, relation := range employee {
		if len(relation.Employee) == 2 {
			child, parent := relation.Employee[0], relation.Employee[1].UniqueID
			relations[parent] = append(relations[parent], child)

		} else {
			child := relation.Employee[0]
			//var parent models.Employee
			relations[""] = append(relations[""], child)
		}
	}

	return relations, nil
}
func (s *Service) GetOrgChart(ctx *models.Context, employess []models.Employee, partent map[string][]models.Employee) []models.Orgchart {
	org := make([]models.Orgchart, len(employess))
	fmt.Println("len(employess)====>", len(employess))
	for i, id := range employess {
		var c models.Orgchart
		c.Employee = employess[i]
		if c.Employee.ProfileImg != "" {
			URL := s.ConfigReader.GetString(s.Shared.GetCmdArg(constants.ENV) + "." + constants.BASEURL)
			c.Employee.Image = URL + c.Employee.ProfileImg
		} else {
			if c.Employee.Gender == "Male" {
				c.Employee.Image = "assets/img/profile_male.png"
			} else if c.Employee.Gender == "Female" {
				c.Employee.Image = "assets/img/profile_female.png"
			} else {
				c.Employee.Image = "assets/img/profiles/avatar-05.jpg"
			}
		}

		fmt.Println("c.Employee.Image====>", c.Employee.Image)
		fmt.Println("Orgchart====>", c.Employee.Name)
		if childIDs, ok := partent[id.UniqueID]; ok {
			for _, k := range childIDs {
				fmt.Println("Child====>", k.Name)
			}
			c.Child = s.GetOrgChart(ctx, childIDs, partent)

		}
		org[i] = c

	}
	fmt.Println("org====>", org)

	return org
}
func (s *Service) GetAllOrgChart(ctx *models.Context) ([]models.Orgchart, error) {
	defer ctx.Session.EndSession(ctx.CTX)
	var org []models.Orgchart
	employee, err := s.GetallEmployee(ctx)
	if err != nil {
		return nil, err
	}
	partent, err := s.GetOrg(ctx, employee)
	if err != nil {
		return nil, err
	}

	org = s.GetOrgChart(ctx, partent[""], partent)
	if err != nil {
		return nil, err
	}

	return org, nil
}

func (s *Service) GetEmployeeLinemanagerCheck(ctx *models.Context, UniqueID string) (*models.EmployeeLinemanagerCheck, error) {
	Employee, err := s.Daos.GetEmployeeLinemanagerCheck(ctx, UniqueID)
	if err != nil {
		return nil, err
	}
	getemployee, err := s.Daos.GetSingleUser(ctx, UniqueID)
	if err != nil {
		return nil, err
	}
	if getemployee != nil {
		if getemployee.Type == constants.USERTYPESUPERADMIN {
			Employee.LineManager = true
		}
		if getemployee.Type == constants.USERTYPEHR {
			Employee.LineManager = true
		}
		// if getemployee.Type == constants.USERTYPE {
		// 	Employee.LineManager = true
		// }
	}
	return Employee, nil
}
func (s *Service) EmployeeUploadExcelV2(ctx *models.Context, file multipart.File) []models.EmployeeUploadError {
	log.Println("transaction start")
	//Start Transaction
	// orgRefMap := make(map[string]primitive.ObjectID)
	// projectRefMap := make(map[string]primitive.ObjectID)
	// stateRefMap := make(map[string]primitive.ObjectID)
	// districtRefMap := make(map[string]primitive.ObjectID)
	// blockRefMap := make(map[string]primitive.ObjectID)
	// grampRefMap := make(map[string]primitive.ObjectID)
	// villageRefMap := make(map[string]primitive.ObjectID)
	var errs []models.EmployeeUploadError
	var employeeerr models.EmployeeUploadError
	if err := ctx.Session.StartTransaction(); err != nil {
		employeeerr.Error = err.Error()
		errs = append(errs, employeeerr)
		return errs
	}
	defer ctx.Session.EndSession(ctx.CTX)
	const (
		MAXCOLUMN              = 20
		OMITROWS               = 0
		ORGANISATIONCOLUMN     = 0
		FIRSTNAMECOLUMN        = 1
		LASTNAMECOLUMN         = 2
		FATHERNAMECOLUMN       = 3
		DOBCOLUMN              = 4
		GENDERCOLUMN           = 5
		MOBILENOCOLUMN         = 6
		PERSONALEMAILCOLUMN    = 7
		OFFICIALEMAILCOLUMN    = 8
		ONBOARDINGPOLICYCOLUMN = 9
		WORKPOLICYCOLUMN       = 10
		LEAVEPOLICYCOLUMN      = 11
		DOCUMENTPOLICYCOLUMN   = 12
		PROBATIONOLICYCOLUMN   = 13
		NOTICEPOLICYCOLUMN     = 14
		BRANCHCOLUMN           = 15
		DEPARTMENTCOLUMN       = 16
		DESIGNATIONCOLUMN      = 17
		LINEMANAGERCOLUMN      = 18
		JOININGDATECOLUMN      = 19
		LOGINIDCOLUMN          = 20
		CREATEDDATELAYOUT      = "02-January-2006"
	)
	if err := mongo.WithSession(ctx.CTX, ctx.Session, func(sc mongo.SessionContext) error {
		fmt.Println("started reading file")
		f, err := excelize.OpenReader(file)
		if err != nil {
			return err
		}
		employees := make([]models.Employee, 0)
		rows := f.GetRows("Sheet1")
		//var errors []string
		fmt.Println("started looping")
		for rowIndex, row := range rows {
			fmt.Println("row no === ", rowIndex)
			if rowIndex <= OMITROWS {
				continue
			}
			if len(row) < MAXCOLUMN {
				return errors.New("Excel is not upto the format")
			}
			employee := new(models.Employee)
			employeevaildation, _ := s.Daos.GetsingleEmployeeWithMobileNumber(ctx, row[MOBILENOCOLUMN])
			if employeevaildation != nil {
				employeeerr.Name = row[FIRSTNAMECOLUMN]
				employeeerr.UserName = employeevaildation.UniqueID
				employeeerr.MobileNumber = row[MOBILENOCOLUMN]
				employeeerr.Error = "Employee MobileNumber Already Registered"
				errs = append(errs, employeeerr)
				continue
			} else {
				employee.Mobile = row[MOBILENOCOLUMN]
			}
			organisation, _ := s.Daos.GetSingleActiveOrganisationWithName(ctx, row[ORGANISATIONCOLUMN])
			if organisation == nil {
				employeeerr.Name = row[FIRSTNAMECOLUMN]
				employeeerr.UserName = ""
				employeeerr.MobileNumber = row[MOBILENOCOLUMN]
				employeeerr.Error = "organisation Not Found"
				errs = append(errs, employeeerr)
				continue
			}
			if organisation != nil {
				employee.OrganisationID = organisation.UniqueID
			}
			employee.FatherName = row[FATHERNAMECOLUMN]
			employee.Name = row[FIRSTNAMECOLUMN] + row[LASTNAMECOLUMN]
			employee.Gender = row[GENDERCOLUMN]
			employee.OfficialEmail = row[OFFICIALEMAILCOLUMN]
			employee.Email = row[PERSONALEMAILCOLUMN]
			onboardingpolicy, _ := s.Daos.GetSingleActiveOnboardingPolicyWithName(ctx, row[ONBOARDINGPOLICYCOLUMN])
			if onboardingpolicy == nil {
				employeeerr.Name = row[FIRSTNAMECOLUMN]
				employeeerr.UserName = ""
				employeeerr.MobileNumber = row[MOBILENOCOLUMN]
				employeeerr.Error = "onboardingpolicy Not Found"
				errs = append(errs, employeeerr)
				continue
			}
			if onboardingpolicy != nil {
				employee.OnboardingpolicyID = onboardingpolicy.UniqueID
			}
			WorkSchedulepolicy, err := s.Daos.GetSingleWorkScheduleActiveWithName(ctx, row[WORKPOLICYCOLUMN])
			if WorkSchedulepolicy == nil {
				employeeerr.Name = row[FIRSTNAMECOLUMN]
				employeeerr.UserName = ""
				employeeerr.MobileNumber = row[MOBILENOCOLUMN]
				employeeerr.Error = "WorkSchedulepolicy Not Found"
				errs = append(errs, employeeerr)
				continue
			}
			if WorkSchedulepolicy != nil {
				employee.WorkScheduleID = WorkSchedulepolicy.UniqueID
			}
			LeavePolicy, err := s.Daos.GetSingleLeavePolicyWithActiveName(ctx, row[LEAVEPOLICYCOLUMN])
			if LeavePolicy == nil {
				employeeerr.Name = row[FIRSTNAMECOLUMN]
				employeeerr.UserName = ""
				employeeerr.MobileNumber = row[MOBILENOCOLUMN]
				employeeerr.Error = "LeavePolicy Not Found"
				errs = append(errs, employeeerr)
				continue
			}
			if LeavePolicy != nil {
				employee.LeavePolicyID = LeavePolicy.UniqueID
			}

			fmt.Println("DOCUMENTPOLICYCOLUMN", row[DOCUMENTPOLICYCOLUMN])
			DocumentPolicy, err := s.Daos.GetSingleDocumentPolicyWithActiveName(ctx, row[DOCUMENTPOLICYCOLUMN])
			if DocumentPolicy == nil {
				employeeerr.Name = row[FIRSTNAMECOLUMN]
				employeeerr.UserName = ""
				employeeerr.MobileNumber = row[MOBILENOCOLUMN]
				employeeerr.Error = "DocumentPolicy Not Found"
				errs = append(errs, employeeerr)
				continue
			}
			if DocumentPolicy != nil {
				employee.DocumentPolicyID = DocumentPolicy.UniqueID
			}
			ProbationaryID, err := s.Daos.GetSingleProbationaryWithActiveName(ctx, row[PROBATIONOLICYCOLUMN])
			if ProbationaryID == nil {
				employeeerr.Name = row[FIRSTNAMECOLUMN]
				employeeerr.UserName = ""
				employeeerr.MobileNumber = row[MOBILENOCOLUMN]
				employeeerr.Error = "ProbationaryPolicy Not Found"
				errs = append(errs, employeeerr)
				continue
			}
			if ProbationaryID != nil {
				employee.ProbationaryID = ProbationaryID.UniqueID
			}
			NoticeID, err := s.Daos.GetSingleNoticePolicyActiveWithName(ctx, row[NOTICEPOLICYCOLUMN])
			if NoticeID == nil {
				employeeerr.Name = row[FIRSTNAMECOLUMN]
				employeeerr.UserName = ""
				employeeerr.MobileNumber = row[MOBILENOCOLUMN]
				employeeerr.Error = "NoticePolicy Not Found"
				errs = append(errs, employeeerr)
				continue
			}
			if NoticeID != nil {
				employee.NoticeID = NoticeID.UniqueID
			}
			BranchID, err := s.Daos.GetSingleBranchActiveWithName(ctx, row[BRANCHCOLUMN])
			if BranchID == nil {
				employeeerr.Name = row[FIRSTNAMECOLUMN]
				employeeerr.UserName = ""
				employeeerr.MobileNumber = row[MOBILENOCOLUMN]
				employeeerr.Error = "Branch Not Found"
				errs = append(errs, employeeerr)
				continue
			}
			if BranchID != nil {
				employee.BranchID = BranchID.UniqueID
			}
			Department, err := s.Daos.GetSingleDepartmentActivewithName(ctx, row[DEPARTMENTCOLUMN])
			if Department == nil {
				employeeerr.Name = row[FIRSTNAMECOLUMN]
				employeeerr.UserName = ""
				employeeerr.MobileNumber = row[MOBILENOCOLUMN]
				employeeerr.Error = "Department Not Found"
				errs = append(errs, employeeerr)
				continue
			}
			if Department != nil {
				employee.DepartmentID = Department.UniqueID
			}
			fmt.Println("LINEMANAGERCOLUMN==========>", row[LINEMANAGERCOLUMN])
			LineManager, err := s.Daos.GetSingleEmployeeActiveWithName(ctx, row[LINEMANAGERCOLUMN])
			if LineManager == nil {
				employeeerr.Name = row[FIRSTNAMECOLUMN]
				employeeerr.UserName = ""
				employeeerr.MobileNumber = row[MOBILENOCOLUMN]
				employeeerr.Error = "LineManager Not Found"
				errs = append(errs, employeeerr)
				continue
			}
			if LineManager != nil {
				employee.LineManager = LineManager.UserName
			}
			DesignationID, err := s.Daos.GetSingleDesignationActiveWithName(ctx, row[DESIGNATIONCOLUMN])
			if DesignationID == nil {
				employeeerr.Name = row[FIRSTNAMECOLUMN]
				employeeerr.UserName = ""
				employeeerr.MobileNumber = row[MOBILENOCOLUMN]
				employeeerr.Error = "Designation Not Found"
				errs = append(errs, employeeerr)
				continue
			}
			if DesignationID != nil {
				employee.DesignationID = DesignationID.UniqueID
			}

			if row[DOBCOLUMN] != "" {
				layout := CREATEDDATELAYOUT
				t, err := time.Parse(layout, row[DOBCOLUMN])
				if err != nil {
					return err
				}
				employee.DOB = &t
			}
			if row[JOININGDATECOLUMN] != "" {
				layout := CREATEDDATELAYOUT
				t, err := time.Parse(layout, row[JOININGDATECOLUMN])
				if err != nil {
					return err
				}
				employee.JoiningDate = &t
			}
			employee.LoginId = row[LOGINIDCOLUMN]

			err = s.EmployeeProbationaryWithoutTranction(ctx, employee)
			if err != nil {
				employeeerr.Name = row[FIRSTNAMECOLUMN]
				employeeerr.UserName = ""
				employeeerr.MobileNumber = row[MOBILENOCOLUMN]
				employeeerr.Error = err.Error()
				errs = append(errs, employeeerr)
				continue
			}
			employees = append(employees, *employee)
			if err == nil {
				employeeerr.Name = row[FIRSTNAMECOLUMN]
				employeeerr.UserName = ""
				employeeerr.UserName = employee.UniqueID
				employeeerr.MobileNumber = row[MOBILENOCOLUMN]
				employeeerr.Error = "Sucess"
				errs = append(errs, employeeerr)
				continue
			}
		}
		fmt.Println("no.of.employee==>", len(employees))
		return nil

	}); err != nil {
		employeeerr.Error = err.Error()
		errs = append(errs, employeeerr)
		return errs
	}
	return errs
}

//EmployeeProbationaryWithoutTranction : ""
func (s *Service) EmployeeProbationaryWithoutTranction(ctx *models.Context, employee *models.Employee) error {
	employee.UniqueID = s.Daos.GetUniqueID(ctx, constants.COLLECTIONEMPLOYEE)
	employee.UserName = employee.UniqueID
	employee.Status = constants.EMPLOYEESTATUSACTIVEEMPLOYEE
	t := time.Now()
	created := models.Created{}
	created.On = &t
	created.By = constants.SYSTEM
	log.Println("b4 Employee.created")
	employee.Created = &created
	employee.Type = constants.USERTYPEEMPLOYEE
	log.Println("b4 Employee.created")

	//update Employee data
	dberr := s.Daos.SaveEmployee(ctx, employee)
	if dberr != nil {
		return dberr
	}
	user := new(models.User)
	user.UserName = employee.UniqueID
	user.Password = "#nature32" //Default Password
	user.UniqueID = s.Daos.GetUniqueID(ctx, constants.COLLECTIONUSER)
	user.Name = employee.Name
	user.LastName = employee.LastName
	user.Gender = employee.Gender
	user.Mobile = employee.Mobile
	user.Email = employee.Email
	user.OfficialEmail = employee.OfficialEmail
	user.DOB = employee.DOB
	user.JoiningDate = employee.JoiningDate
	user.OrganisationID = employee.OrganisationID
	user.EmployeeId = employee.UniqueID
	user.Status = constants.EMPLOYEESTATUSACTIVEEMPLOYEE
	user.Role = employee.Role
	user.Type = "Employee"
	user.OrganisationID = employee.OrganisationID
	user.DepartmentID = employee.DepartmentID
	user.DesignationID = employee.DesignationID
	user.BranchID = employee.BranchID
	dberr = s.Daos.SaveUser(ctx, user)
	if dberr != nil {
		return dberr
	}
	//t = time.Now()
	//EmployeeLog
	employeeLog := new(models.EmployeeLog)
	employeeLog.Name = employee.Name
	employeeLog.UniqueID = s.Daos.GetUniqueID(ctx, constants.COLLECTIONJOBTIMELINE)
	employeeLog.OrganisationId = employee.OrganisationID
	employeeLog.DepartmentId = employee.DepartmentID
	employeeLog.BranchId = employee.BranchID
	employeeLog.DesignationId = employee.DepartmentID
	employeeLog.Desc = constants.EMPLOYEESTATUSACTIVEEMPLOYEEDESC
	employeeLog.Action.UserID = employee.UniqueID
	employeeLog.Action.UserType = employee.Role
	employeeLog.Action.Date = t
	employeeLog.EmployeeId = employee.UniqueID
	employeeLog.Status = constants.EMPLOYEESTATUSACTIVE
	employeeLog.Remark = employee.Remark
	dberr = s.Daos.SaveEmployeeLog(ctx, employeeLog)
	if dberr != nil {
		return dberr
	}
	//jobtimeline
	jobtimeline := new(models.JobTimeline)
	jobtimeline.State = constants.EMPLOYEESTATUSACTIVE
	jobtimeline.UniqueID = s.Daos.GetUniqueID(ctx, constants.COLLECTIONJOBTIMELINE)
	jobtimeline.StartDate = &t
	jobtimeline.Status = constants.EMPLOYEESTATUSACTIVESTAGE
	jobtimeline.OrganisationId = employee.OrganisationID
	jobtimeline.DepartmentId = employee.DepartmentID
	jobtimeline.BranchId = employee.BranchID
	jobtimeline.DesignationId = employee.DepartmentID
	jobtimeline.Assigned.UserType = employee.Role
	jobtimeline.Assigned.Date = t
	jobtimeline.EmployeeId = employee.UniqueID
	jobtimeline.Remark = employee.Remark
	dberr = s.Daos.SaveJobTimeline(ctx, jobtimeline)
	if dberr != nil {
		return dberr
	}

	leave, err := s.Daos.GetSingleLeavePolicy(ctx, employee.LeavePolicyID)
	if err != nil {
		return err
	}
	employeeleave := new(models.EmployeeLeave)
	for k, v := range leave.LeaveMaster {
		employeeleave.OrganisationId = employee.OrganisationID
		employeeleave.EmployeeId = employee.UserName
		fmt.Println("Leavetype===", v.UniqueID)
		fmt.Println("Leavetype K===", leave.LeaveMaster[k].UniqueID)
		employeeleave.LeaveType = v.UniqueID
		employeeleave.Value = int64(v.Value)
		employeeleave.UniqueID = s.Daos.GetUniqueID(ctx, constants.COLLECTIONEMPLOYEELEAVE)
		employeeleave.Status = constants.EMPLOYEELEAVESTATUSACTIVE
		t := time.Now()
		created := models.Created{}
		created.On = &t
		created.By = constants.SYSTEM
		log.Println("b4 EmployeeLeave.created")
		employeeleave.Created = &created
		log.Println("b4 EmployeeLeave.created")
		err := s.Daos.SaveEmployeeLeave(ctx, employeeleave)
		if err != nil {
			return err
		}
		employeeleavelog := new(models.EmployeeLeaveLog)
		employeeleavelog.UniqueID = s.Daos.GetUniqueID(ctx, constants.COLLECTIONEMPLOYEELEAVELOG)
		employeeleavelog.EmployeeId = employeeleave.EmployeeId
		employeeleavelog.LeaveType = employeeleave.LeaveType
		employeeleavelog.Value = employeeleave.Value
		employeeleavelog.Date = &t
		employeeleavelog.CreateBy = "System"
		employeeleavelog.CreateDate = &t
		employeeleavelog.Status = constants.EMPLOYEELEAVESTATUSACTIVE
		employeeleavelog.Created = &created
		err = s.Daos.SaveEmployeeLeaveLog(ctx, employeeleavelog)
		if err != nil {
			return err
		}
	}
	url := s.ConfigReader.GetString(s.Shared.GetCmdArg(constants.ENV) + "." + constants.APKURL)

	msg := fmt.Sprintf(constants.COMMONTEMPLATE, employee.Name, employee.OfficialEmail, url, employee.UniqueID, employee.Mobile)
	//	msg:=fmt.Sprintf("D")
	var sendmailto []string
	fmt.Println("Employee email======>", employee.Email)
	sendmailto = append(sendmailto, employee.Email)
	err = s.SendEmail("Welcome To Logikoof", sendmailto, msg)
	if err != nil {
		return errors.New("email Sending Error - " + err.Error())
	}
	if err == nil {
		emaillog := new(models.EmailLog)
		to2 := models.ToEmailLog{}
		to2.Email = employee.Email
		to2.Name = employee.UserName
		to2.UserName = employee.UserName
		to2.UserType = "Employee"
		t := time.Now()
		emaillog.SentDate = &t
		emaillog.IsJob = false
		emaillog.Message = msg
		emaillog.SentFor = "login"
		emaillog.Status = "Active"
		emaillog.To = to2
		err = s.Daos.SaveEmailLog(ctx, emaillog)
		if err != nil {
			return errors.New("login email not save")
		}
	}

	return nil
}
func (s *Service) EmployeeUpdateLoginId(ctx *models.Context, file multipart.File) []models.EmployeeUploadError {
	log.Println("transaction start")
	//Start Transaction
	// orgRefMap := make(map[string]primitive.ObjectID)
	// projectRefMap := make(map[string]primitive.ObjectID)
	// stateRefMap := make(map[string]primitive.ObjectID)
	// districtRefMap := make(map[string]primitive.ObjectID)
	// blockRefMap := make(map[string]primitive.ObjectID)
	// grampRefMap := make(map[string]primitive.ObjectID)
	// villageRefMap := make(map[string]primitive.ObjectID)
	var errs []models.EmployeeUploadError
	var employeeerr models.EmployeeUploadError
	duplicatecheck := make(map[string]string)
	if err := ctx.Session.StartTransaction(); err != nil {
		employeeerr.Error = err.Error()
		errs = append(errs, employeeerr)
		return errs
	}
	defer ctx.Session.EndSession(ctx.CTX)
	const (
		MAXCOLUMN      = 2
		OMITROWS       = 0
		SNO            = 0
		USERNAMECOLUNM = 1
		LOGINIDCOLUMN  = 2
	)
	if err := mongo.WithSession(ctx.CTX, ctx.Session, func(sc mongo.SessionContext) error {
		fmt.Println("started reading file")
		f, err := excelize.OpenReader(file)
		if err != nil {
			return err
		}
		rows := f.GetRows("Sheet1")
		//var errors []string
		fmt.Println("started looping")
		for rowIndex, row := range rows {
			fmt.Println("row no === ", rowIndex)
			if rowIndex <= OMITROWS {
				continue
			}
			if len(row) < MAXCOLUMN {
				return errors.New("Excel is not upto the format")
			}
			if row[USERNAMECOLUNM] == "" {
				employeeerr.UserName = row[USERNAMECOLUNM]
				employeeerr.SNo = row[SNO]
				employeeerr.Error = "Please Enter UserName"
				errs = append(errs, employeeerr)
				continue
			}
			if row[LOGINIDCOLUMN] == "" {
				employeeerr.UserName = row[USERNAMECOLUNM]
				employeeerr.SNo = row[SNO]
				employeeerr.Error = "Please Enter LoginId"
				errs = append(errs, employeeerr)
				continue
			}
			if duplicatecheck[row[USERNAMECOLUNM]] == "Yes" {
				employeeerr.UserName = row[USERNAMECOLUNM]
				employeeerr.SNo = row[SNO]
				employeeerr.Error = "UserName Duplicate Found"
				errs = append(errs, employeeerr)
				continue
			}
			loginidCheck, err := s.Daos.GetSingleEmployeeWithLoginId(ctx, row[LOGINIDCOLUMN])
			if err != nil {
				return err
			}
			if loginidCheck != nil {
				employeeerr.SNo = row[SNO]
				employeeerr.UserName = loginidCheck.UserName
				employeeerr.Error = "LoginId Duplicate Found"
				errs = append(errs, employeeerr)
				continue
			}
			employees, err := s.Daos.GetSingleEmployeeWithUserName(ctx, row[USERNAMECOLUNM])
			if err != nil {
				return err
			}
			if employees == nil {
				employeeerr.UserName = row[USERNAMECOLUNM]
				employeeerr.Error = "Employee UserName Not Found"
				errs = append(errs, employeeerr)
				continue
			}
			if employees.LoginId == row[LOGINIDCOLUMN] {
				employeeerr.UserName = row[USERNAMECOLUNM]
				employeeerr.Error = "LoginId Already Update"
				errs = append(errs, employeeerr)
				continue
			}
			err = s.Daos.UpdateEmployeeLoginId(ctx, row[USERNAMECOLUNM], row[LOGINIDCOLUMN])
			if err != nil {
				return err
			}
			err = s.Daos.UpdateUserLoginId(ctx, row[USERNAMECOLUNM], row[LOGINIDCOLUMN])
			if err != nil {
				return err
			}
			duplicatecheck[row[USERNAMECOLUNM]] = "Yes"
			employeeerr.UserName = row[USERNAMECOLUNM]
			employeeerr.Error = "Scuessfully Update"
			errs = append(errs, employeeerr)
			continue
		}
		return nil

	}); err != nil {
		employeeerr.Error = err.Error()
		errs = append(errs, employeeerr)
		return errs
	}
	return errs
}
