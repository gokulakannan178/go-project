package services

import (
	"bpms-service/constants"
	"bpms-service/models"
	"errors"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
)

//SavePlanDepartmentApproval :""
func (s *Service) SavePlanDepartmentApproval(ctx *models.Context, planDepartmentApproval *models.PlanDepartmentApproval) error {
	log.Println("transaction start")
	//Start Transaction
	if err := ctx.Session.StartTransaction(); err != nil {
		return err
	}
	defer ctx.Session.EndSession(ctx.CTX)

	planDepartmentApproval.UniqueID = s.Daos.GetUniqueID(ctx, constants.COLLECTIONPLANDEPARTMENTAPPROVAL)
	planDepartmentApproval.Status = constants.PLANDEPTAPPROVALSTATUSACTIVE
	t := time.Now()
	created := models.Created{}
	created.On = &t
	created.By = constants.SYSTEM
	log.Println("b4 planDepartmentApproval.created")
	planDepartmentApproval.Created = created
	log.Println("b4 planDepartmentApproval.created")
	if err := mongo.WithSession(ctx.CTX, ctx.Session, func(sc mongo.SessionContext) error {
		dberr := s.Daos.SavePlanDepartmentApproval(ctx, planDepartmentApproval)
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

//SaveMultiplePlanDepartmentApproval :""
func (s *Service) SaveMultiplePlanDepartmentApproval(ctx *models.Context, planDepartmentApproval []models.PlanDepartmentApproval) error {
	log.Println("transaction start")
	//Start Transaction
	if err := ctx.Session.StartTransaction(); err != nil {
		return err
	}
	defer ctx.Session.EndSession(ctx.CTX)
	for k := range planDepartmentApproval {
		if planDepartmentApproval[k].UniqueID == "" {
			planDepartmentApproval[k].UniqueID = s.Daos.GetUniqueID(ctx, constants.COLLECTIONPLANDEPARTMENTAPPROVAL)
			planDepartmentApproval[k].Status = constants.PLANDEPTAPPROVALSTATUSACTIVE
			t := time.Now()
			created := models.Created{}
			created.On = &t
			created.By = constants.SYSTEM
			log.Println("b4 planDepartmentApproval.created")
			planDepartmentApproval[k].Created = created
			log.Println("b4 planDepartmentApproval.created")
		}
	}

	if err := mongo.WithSession(ctx.CTX, ctx.Session, func(sc mongo.SessionContext) error {
		dberr := s.Daos.SaveMultiplePlanDepartmentApproval(ctx, planDepartmentApproval)
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

//UpdatePlanDepartmentApproval : ""
func (s *Service) UpdatePlanDepartmentApproval(ctx *models.Context, planDepartmentApproval *models.PlanDepartmentApproval) error {
	if err := ctx.Session.StartTransaction(); err != nil {
		return err
	}
	defer ctx.Session.EndSession(ctx.CTX)
	if err := mongo.WithSession(ctx.CTX, ctx.Session, func(sc mongo.SessionContext) error {

		err := s.Daos.UpdatePlanDepartmentApproval(ctx, planDepartmentApproval)
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

//EnablePlanDepartmentApproval : ""
func (s *Service) EnablePlanDepartmentApproval(ctx *models.Context, UniqueID string) error {
	if err := ctx.Session.StartTransaction(); err != nil {
		return err
	}
	defer ctx.Session.EndSession(ctx.CTX)
	if err := mongo.WithSession(ctx.CTX, ctx.Session, func(sc mongo.SessionContext) error {

		err := s.Daos.EnablePlanDepartmentApproval(ctx, UniqueID)
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

//DisablePlanDepartmentApproval : ""
func (s *Service) DisablePlanDepartmentApproval(ctx *models.Context, UniqueID string) error {
	if err := ctx.Session.StartTransaction(); err != nil {
		return err
	}
	defer ctx.Session.EndSession(ctx.CTX)
	if err := mongo.WithSession(ctx.CTX, ctx.Session, func(sc mongo.SessionContext) error {

		err := s.Daos.DisablePlanDepartmentApproval(ctx, UniqueID)
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

//DeletePlanDepartmentApproval : ""
func (s *Service) DeletePlanDepartmentApproval(ctx *models.Context, UniqueID string) error {
	if err := ctx.Session.StartTransaction(); err != nil {
		return err
	}
	defer ctx.Session.EndSession(ctx.CTX)
	if err := mongo.WithSession(ctx.CTX, ctx.Session, func(sc mongo.SessionContext) error {

		err := s.Daos.DeletePlanDepartmentApproval(ctx, UniqueID)
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

//GetSinglePlanDepartmentApproval :""
func (s *Service) GetSinglePlanDepartmentApproval(ctx *models.Context, UniqueID string) (*models.RefPlanDepartmentApproval, error) {
	planDepartmentApproval, err := s.Daos.GetSinglePlanDepartmentApproval(ctx, UniqueID)
	if err != nil {
		return nil, err
	}
	return planDepartmentApproval, nil
}

//FilterPlanDepartmentApproval :""
func (s *Service) FilterPlanDepartmentApproval(ctx *models.Context, planDepartmentApprovalfilter *models.PlanDepartmentApprovalFilter, pagination *models.Pagination) (planDepartmentApproval []models.RefPlanDepartmentApproval, err error) {
	return s.Daos.FilterPlanDepartmentApproval(ctx, planDepartmentApprovalfilter, pagination)
}

//GetAPlanDeptsApproval : ""
func (s *Service) GetAPlanDeptsApproval(ctx *models.Context, deptID string) (*models.GetAPlanDeptsApproval, error) {
	return s.Daos.GetAPlanDeptsApproval(ctx, deptID)
}

//GetAPlanDeptsApprovalV2 : ""
func (s *Service) GetAPlanDeptsApprovalV2(ctx *models.Context, deptID string) (*models.GetAPlanDeptsApprovalV2, error) {
	data, err := s.Daos.GetAPlanDeptsApprovalV2(ctx, deptID)
	if err != nil {
		return nil, err
	}
	if data != nil {
		for k := range data.Departments {
			if data.Departments[k].PlanDepartmentApproval == nil {
				data.Departments[k].PlanDepartmentApproval = new(models.PlanDepartmentApproval)
				data.Departments[k].PlanDepartmentApproval.Check = "No"
			}
		}
	}
	return data, nil
}

//GetAPlanDeptsApprovalV3 : ""
func (s *Service) GetAPlanDeptsApprovalV3(ctx *models.Context, deptID string) (*models.GetAPlanDeptsApprovalV3, error) {
	data, err := s.Daos.GetAPlanDeptsApprovalV3(ctx, deptID)
	if err != nil {
		return nil, err
	}
	if data != nil {
		for k := range data.Departments {
			if data.Departments[k].PlanDepartmentApproval == nil {
				data.Departments[k].PlanDepartmentApproval = new(models.PlanDepartmentApproval)
				data.Departments[k].PlanDepartmentApproval.Check = "No"
			}
		}
	}
	return data, nil
}
