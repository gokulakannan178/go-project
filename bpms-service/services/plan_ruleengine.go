package services

import (
	"bpms-service/constants"
	"bpms-service/models"
	"errors"
	"fmt"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
)

//*Draft
//*Submitted
//PassedScrutiny
//PCP
//DeptApproval
//PendingCC
//Defective - PCPDefective, DeptDefective, CCDefective
//Cancelled - PCPCancelled, DeptCancelled, CCCancelled
//PendingPayment
//Approved

//PlanMakeFailScrutiny : ""
func (s *Service) PlanMakeFailScrutiny(ctx *models.Context, pfs *models.PlanFailScrutiny) error {
	log.Println("transaction start")
	//Start Transaction
	if err := ctx.Session.StartTransaction(); err != nil {
		return err
	}
	defer ctx.Session.EndSession(ctx.CTX)
	if err := mongo.WithSession(ctx.CTX, ctx.Session, func(sc mongo.SessionContext) error {

		re, err := s.Daos.GetSinglePlanRuleByScenario(ctx, pfs.Scenario)
		if err != nil {
			return errors.New("Not able to get RE - " + err.Error())
		}
		if re == nil {
			return errors.New("Re is null")
		}
		t := time.Now()
		pfs.PlanTimeline.On = &t
		pfs.PlanTimeline.Type = re.Log.Type
		pfs.PlanTimeline.TypeLabel = re.Log.Label
		m := make(map[string]interface{})
		m["status"] = re.To
		m["remarks"] = pfs.PlanTimeline.Remarks
		err = s.Daos.PlanFlowUpdate(ctx, pfs.PlanID, m, pfs.PlanTimeline)
		if err != nil {
			return errors.New("Not able to update plan - " + err.Error())
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

//PlanMakePassScrutiny : ""
func (s *Service) PlanMakePassScrutiny(ctx *models.Context, pps *models.PlanPassScrutiny) error {
	log.Println("transaction start")
	//Start Transaction
	if err := ctx.Session.StartTransaction(); err != nil {
		return err
	}
	defer ctx.Session.EndSession(ctx.CTX)
	if err := mongo.WithSession(ctx.CTX, ctx.Session, func(sc mongo.SessionContext) error {

		re, err := s.Daos.GetSinglePlanRuleByScenario(ctx, pps.Scenario)
		if err != nil {
			return errors.New("Not able to get RE - " + err.Error())
		}
		if re == nil {
			return errors.New("Re is null")
		}
		t := time.Now()
		pps.PlanTimeline.On = &t
		pps.PlanTimeline.Type = re.Log.Type
		pps.PlanTimeline.TypeLabel = re.Log.Label
		m := make(map[string]interface{})
		m["status"] = re.To
		m["remarks"] = pps.PlanTimeline.Remarks
		err = s.Daos.PlanFlowUpdate(ctx, pps.PlanID, m, pps.PlanTimeline)
		if err != nil {
			return errors.New("Not able to update plan - " + err.Error())
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

//ProceedPCP : ""
func (s *Service) ProceedPCP(ctx *models.Context, ppcp *models.ProceedPCP) error {
	log.Println("transaction start")
	//Start Transaction
	if err := ctx.Session.StartTransaction(); err != nil {
		return err
	}
	defer ctx.Session.EndSession(ctx.CTX)
	if err := mongo.WithSession(ctx.CTX, ctx.Session, func(sc mongo.SessionContext) error {

		re, err := s.Daos.GetSinglePlanRuleByScenario(ctx, ppcp.Scenario)
		if err != nil {
			return errors.New("Not able to get RE - " + err.Error())
		}
		if re == nil {
			return errors.New("Re is null")
		}
		t := time.Now()
		ppcp.PlanTimeline.On = &t
		ppcp.PlanTimeline.Type = re.Log.Type
		ppcp.PlanTimeline.TypeLabel = re.Log.Label
		m := make(map[string]interface{})
		m["status"] = re.To
		m["remarks"] = ppcp.PlanTimeline.Remarks
		err = s.Daos.PlanFlowUpdate(ctx, ppcp.PlanID, m, ppcp.PlanTimeline)
		if err != nil {
			return errors.New("Not able to update plan - " + err.Error())
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

//MakePCPDefective : ""
func (s *Service) MakePCPDefective(ctx *models.Context, mpcpd *models.MakePCPDefective) error {
	log.Println("transaction start")
	//Start Transaction
	if err := ctx.Session.StartTransaction(); err != nil {
		return err
	}
	defer ctx.Session.EndSession(ctx.CTX)
	if err := mongo.WithSession(ctx.CTX, ctx.Session, func(sc mongo.SessionContext) error {

		re, err := s.Daos.GetSinglePlanRuleByScenario(ctx, mpcpd.Scenario)
		if err != nil {
			return errors.New("Not able to get RE - " + err.Error())
		}
		if re == nil {
			return errors.New("Re is null")
		}
		t := time.Now()
		mpcpd.PlanTimeline.On = &t
		mpcpd.PlanTimeline.Type = re.Log.Type
		mpcpd.PlanTimeline.TypeLabel = re.Log.Label
		m := make(map[string]interface{})
		m["status"] = re.To
		m["remarks"] = mpcpd.PlanTimeline.Remarks
		err = s.Daos.PlanFlowUpdate(ctx, mpcpd.PlanID, m, mpcpd.PlanTimeline)
		if err != nil {
			return errors.New("Not able to update plan - " + err.Error())
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

//PCPAccept : ""
func (s *Service) PCPAccept(ctx *models.Context, data *models.PCPAccept) error {
	log.Println("transaction start")
	//Start Transaction
	if err := ctx.Session.StartTransaction(); err != nil {
		return err
	}
	defer ctx.Session.EndSession(ctx.CTX)
	if err := mongo.WithSession(ctx.CTX, ctx.Session, func(sc mongo.SessionContext) error {

		re, err := s.Daos.GetSinglePlanRuleByScenario(ctx, data.Scenario)
		if err != nil {
			return errors.New("Not able to get RE - " + err.Error())
		}
		if re == nil {
			return errors.New("Re is null")
		}
		t := time.Now()
		data.PlanTimeline.On = &t
		data.PlanTimeline.Type = re.Log.Type
		data.PlanTimeline.TypeLabel = re.Log.Label
		m := make(map[string]interface{})
		m["status"] = re.To
		m["remarks"] = data.PlanTimeline.Remarks
		err = s.Daos.PlanFlowUpdate(ctx, data.PlanID, m, data.PlanTimeline)
		if err != nil {
			return errors.New("Not able to update plan - " + err.Error())
		}
		/*****************/
		//Get Plan
		plan, err := s.Daos.GetSinglePlan(ctx, data.PlanID)
		if err != nil {
			return errors.New("Error in geting plan - " + err.Error())
		}
		if plan == nil {
			return errors.New("Nil Plan  ")
		}
		fmt.Println("GOT Pan")
		//Get Approvals
		// approvals, err := s.GetAPlanDeptsApprovalV2(ctx, plan.RegType)
		// if err != nil {
		// 	return errors.New("Error in geting approvals - " + err.Error())
		// }
		// if approvals == nil {
		// 	return errors.New("Nil Approvals  ")
		// }
		//Create CRF
		//	var crfs []models.CRF
		// if len(approvals.Departments) > 0 {
		// 	for _, v := range approvals.Departments {
		// 		if v.PlanDepartmentApproval != nil {
		// 			if v.PlanDepartmentApproval.Check == "Yes" {
		// 				var crf models.CRF
		// 				crf.UniqueID = s.Daos.GetUniqueID(ctx, constants.COLLECTIONCRF)
		// 				crf.PlanRegTypeID = plan.RegType
		// 				crf.DeptID = v.UniqueID
		// 				crf.PlanID = plan.UniqueID
		// 				crf.Status = constants.PLANCRFSTATUSINSPECTIONPENDING
		// 				crf.Created = models.Created{On: &t, By: "System"}
		// 				crfs = append(crfs, crf)
		// 			}
		// 		}

		// 	}

		// } else {
		// 	fmt.Println("No Departments")
		// }
		// if len(crfs) > 0 {
		// 	err = s.Daos.SaveCRF(ctx, crfs)
		// 	if err != nil {
		// 		return errors.New("Errors in creating CRF - " + err.Error())
		// 	}
		// } else {
		// 	fmt.Println("No CRFS")
		// }
		/*************/

		PlanAproval, err := s.Daos.GetSinglePlanDepartmentApprovalWithPlantype(ctx, plan.RegType)
		if err != nil {
			return err
		}
		for _, v := range PlanAproval {
			depart, err := s.Daos.GetSingleDepartmentWithDistrictAndDepartmentType(ctx, v.DepartmentTypeID, plan.Address.DistrictCode)
			if err != nil {
				return err
			}
			//var crf models.CRF
			crf := new(models.CRF)
			crf.UniqueID = s.Daos.GetUniqueID(ctx, constants.COLLECTIONCRF)
			crf.PlanRegTypeID = plan.RegType
			crf.DeptID = depart.UniqueID
			crf.PlanID = plan.UniqueID
			crf.Status = constants.PLANCRFSTATUSINSPECTIONPENDING
			crf.Created = models.Created{On: &t, By: "System"}
			err = s.Daos.SaveSingleCRF(ctx, crf)
			if err != nil {
				return err
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

//DeptApprovalAccept : ""
func (s *Service) DeptApprovalAccept(ctx *models.Context, data *models.DeptApprovalAccept) error {
	log.Println("transaction start")
	//Start Transaction
	if err := ctx.Session.StartTransaction(); err != nil {
		return err
	}
	defer ctx.Session.EndSession(ctx.CTX)
	if err := mongo.WithSession(ctx.CTX, ctx.Session, func(sc mongo.SessionContext) error {

		re, err := s.Daos.GetSinglePlanRuleByScenario(ctx, data.Scenario)
		if err != nil {
			return errors.New("Not able to get RE - " + err.Error())
		}
		if re == nil {
			return errors.New("Re is null")
		}
		t := time.Now()
		data.PlanTimeline.On = &t
		data.PlanTimeline.Type = re.Log.Type
		data.PlanTimeline.TypeLabel = re.Log.Label
		m := make(map[string]interface{})
		m["status"] = re.To
		m["remarks"] = data.PlanTimeline.Remarks
		err = s.Daos.PlanFlowUpdate(ctx, data.PlanID, m, data.PlanTimeline)
		if err != nil {
			return errors.New("Not able to update plan - " + err.Error())
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

//DeptApprovalReject : ""
func (s *Service) DeptApprovalReject(ctx *models.Context, data *models.DeptApprovalReject) error {
	log.Println("transaction start")
	//Start Transaction
	if err := ctx.Session.StartTransaction(); err != nil {
		return err
	}
	defer ctx.Session.EndSession(ctx.CTX)
	if err := mongo.WithSession(ctx.CTX, ctx.Session, func(sc mongo.SessionContext) error {

		re, err := s.Daos.GetSinglePlanRuleByScenario(ctx, data.Scenario)
		if err != nil {
			return errors.New("Not able to get RE - " + err.Error())
		}
		if re == nil {
			return errors.New("Re is null")
		}
		t := time.Now()
		data.PlanTimeline.On = &t
		data.PlanTimeline.Type = re.Log.Type
		data.PlanTimeline.TypeLabel = re.Log.Label
		m := make(map[string]interface{})
		m["status"] = re.To
		m["remarks"] = data.PlanTimeline.Remarks
		err = s.Daos.PlanFlowUpdate(ctx, data.PlanID, m, data.PlanTimeline)
		if err != nil {
			return errors.New("Not able to update plan - " + err.Error())
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

//CCAccept : ""
func (s *Service) CCAccept(ctx *models.Context, data *models.CCAccept) error {
	log.Println("transaction start")
	//Start Transaction
	if err := ctx.Session.StartTransaction(); err != nil {
		return err
	}
	defer ctx.Session.EndSession(ctx.CTX)
	if err := mongo.WithSession(ctx.CTX, ctx.Session, func(sc mongo.SessionContext) error {

		re, err := s.Daos.GetSinglePlanRuleByScenario(ctx, data.Scenario)
		if err != nil {
			return errors.New("Not able to get RE - " + err.Error())
		}
		if re == nil {
			return errors.New("Re is null")
		}
		t := time.Now()
		data.PlanTimeline.On = &t
		data.PlanTimeline.Type = re.Log.Type
		data.PlanTimeline.TypeLabel = re.Log.Label
		m := make(map[string]interface{})
		m["status"] = re.To
		m["remarks"] = data.PlanTimeline.Remarks
		err = s.Daos.PlanFlowUpdate(ctx, data.PlanID, m, data.PlanTimeline)
		if err != nil {
			return errors.New("Not able to update plan - " + err.Error())
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

//CCReject : ""
func (s *Service) CCReject(ctx *models.Context, data *models.CCReject) error {
	log.Println("transaction start")
	//Start Transaction
	if err := ctx.Session.StartTransaction(); err != nil {
		return err
	}
	defer ctx.Session.EndSession(ctx.CTX)
	if err := mongo.WithSession(ctx.CTX, ctx.Session, func(sc mongo.SessionContext) error {

		re, err := s.Daos.GetSinglePlanRuleByScenario(ctx, data.Scenario)
		if err != nil {
			return errors.New("Not able to get RE - " + err.Error())
		}
		if re == nil {
			return errors.New("Re is null")
		}
		t := time.Now()
		data.PlanTimeline.On = &t
		data.PlanTimeline.Type = re.Log.Type
		data.PlanTimeline.TypeLabel = re.Log.Label
		m := make(map[string]interface{})
		m["status"] = re.To
		m["remarks"] = data.PlanTimeline.Remarks
		err = s.Daos.PlanFlowUpdate(ctx, data.PlanID, m, data.PlanTimeline)
		if err != nil {
			return errors.New("Not able to update plan - " + err.Error())
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

//MakePayment : ""
func (s *Service) MakePayment(ctx *models.Context, data *models.MakePayment) error {
	log.Println("transaction start")
	//Start Transaction
	if err := ctx.Session.StartTransaction(); err != nil {
		return err
	}
	defer ctx.Session.EndSession(ctx.CTX)
	if err := mongo.WithSession(ctx.CTX, ctx.Session, func(sc mongo.SessionContext) error {

		re, err := s.Daos.GetSinglePlanRuleByScenario(ctx, data.Scenario)
		if err != nil {
			return errors.New("Not able to get RE - " + err.Error())
		}
		if re == nil {
			return errors.New("Re is null")
		}
		t := time.Now()
		data.PlanTimeline.On = &t
		data.PlanTimeline.Type = re.Log.Type
		data.PlanTimeline.TypeLabel = re.Log.Label
		m := make(map[string]interface{})
		m["status"] = re.To
		m["remarks"] = data.PlanTimeline.Remarks
		err = s.Daos.PlanFlowUpdate(ctx, data.PlanID, m, data.PlanTimeline)
		if err != nil {
			return errors.New("Not able to update plan - " + err.Error())
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

//ReapplyDefective : ""
func (s *Service) ReapplyDefective(ctx *models.Context, data *models.ReapplyDefective) error {
	log.Println("transaction start")
	//Start Transaction
	if err := ctx.Session.StartTransaction(); err != nil {
		return err
	}
	defer ctx.Session.EndSession(ctx.CTX)
	if err := mongo.WithSession(ctx.CTX, ctx.Session, func(sc mongo.SessionContext) error {
		refPlan, err := s.Daos.GetSinglePlan(ctx, data.PlanID)
		if err != nil {
			return errors.New("error in geting plan - " + err.Error())
		}
		if refPlan == nil {
			return errors.New("error in geting plan - ")
		}
		if refPlan.Status == "" {
			return errors.New("error in plan status- empty")
		}
		re, err := s.Daos.GetSinglePlanRuleByScenarioAndFromStatus(ctx, data.Scenario, refPlan.Status)
		if err != nil {
			return errors.New("Not able to get RE - " + err.Error())
		}
		if re == nil {
			return errors.New("Re is null")
		}
		t := time.Now()
		data.PlanTimeline.On = &t
		data.PlanTimeline.Type = re.Log.Type
		data.PlanTimeline.TypeLabel = re.Log.Label
		m := make(map[string]interface{})
		m["status"] = re.To
		m["remarks"] = data.PlanTimeline.Remarks
		err = s.Daos.PlanFlowUpdate(ctx, data.PlanID, m, data.PlanTimeline)
		if err != nil {
			return errors.New("Not able to update plan - " + err.Error())
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
