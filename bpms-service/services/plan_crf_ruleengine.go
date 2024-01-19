package services

import (
	"bpms-service/models"
	"errors"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
)

/*
Active
InspectionPending
InspectionCompleted
Approved
Completed
Rejected
*/

//PlanCRFAccept : ""
func (s *Service) PlanCRFAccept(ctx *models.Context, request *models.PlanCRFAccept) error {
	log.Println("transaction start")
	//Start Transaction
	if err := ctx.Session.StartTransaction(); err != nil {
		return err
	}
	defer ctx.Session.EndSession(ctx.CTX)
	if err := mongo.WithSession(ctx.CTX, ctx.Session, func(sc mongo.SessionContext) error {

		re, err := s.Daos.GetSinglePlanCRFRuleByScenario(ctx, request.Scenario)
		if err != nil {
			return errors.New("Not able to get RE - " + err.Error())
		}
		if re == nil {
			return errors.New("Re is null")
		}
		t := time.Now()
		request.PlanTimeline.On = &t
		request.PlanTimeline.Type = re.Log.Type
		request.PlanTimeline.TypeLabel = re.Log.Label
		m := make(map[string]interface{})
		m["status"] = re.To
		m["remarks"] = request.PlanTimeline.Remarks
		err = s.Daos.PlanCRFFlowUpdate(ctx, request.CRFID, m, request.PlanTimeline)
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

//PlanCRFPostInspectionAccept : ""
func (s *Service) PlanCRFPostInspectionAccept(ctx *models.Context, request *models.PlanCRFPostInspectionAccept) error {
	log.Println("transaction start")
	//Start Transaction
	if err := ctx.Session.StartTransaction(); err != nil {
		return err
	}
	defer ctx.Session.EndSession(ctx.CTX)
	if err := mongo.WithSession(ctx.CTX, ctx.Session, func(sc mongo.SessionContext) error {

		re, err := s.Daos.GetSinglePlanCRFRuleByScenario(ctx, request.Scenario)
		if err != nil {
			return errors.New("Not able to get RE - " + err.Error())
		}
		if re == nil {
			return errors.New("Re is null")
		}
		t := time.Now()
		request.PlanTimeline.On = &t
		request.PlanTimeline.Type = re.Log.Type
		request.PlanTimeline.TypeLabel = re.Log.Label
		m := make(map[string]interface{})
		m["status"] = re.To
		m["remarks"] = request.PlanTimeline.Remarks
		err = s.Daos.PlanCRFFlowUpdate(ctx, request.CRFID, m, request.PlanTimeline)
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

//PlanCRFPostInspectionReject : ""
func (s *Service) PlanCRFPostInspectionReject(ctx *models.Context, request *models.PlanCRFPostInspectionReject) error {
	log.Println("transaction start")
	//Start Transaction
	if err := ctx.Session.StartTransaction(); err != nil {
		return err
	}
	defer ctx.Session.EndSession(ctx.CTX)
	if err := mongo.WithSession(ctx.CTX, ctx.Session, func(sc mongo.SessionContext) error {

		re, err := s.Daos.GetSinglePlanCRFRuleByScenario(ctx, request.Scenario)
		if err != nil {
			return errors.New("Not able to get RE - " + err.Error())
		}
		if re == nil {
			return errors.New("Re is null")
		}
		t := time.Now()
		request.PlanTimeline.On = &t
		request.PlanTimeline.Type = re.Log.Type
		request.PlanTimeline.TypeLabel = re.Log.Label
		m := make(map[string]interface{})
		m["status"] = re.To
		m["remarks"] = request.PlanTimeline.Remarks
		err = s.Daos.PlanCRFFlowUpdate(ctx, request.CRFID, m, request.PlanTimeline)
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

//PlanCRFCertificateComplete : ""
func (s *Service) PlanCRFCertificateComplete(ctx *models.Context, request *models.PlanCRFCertificateComplete) error {
	log.Println("transaction start")
	//Start Transaction
	if err := ctx.Session.StartTransaction(); err != nil {
		return err
	}
	defer ctx.Session.EndSession(ctx.CTX)
	if err := mongo.WithSession(ctx.CTX, ctx.Session, func(sc mongo.SessionContext) error {

		re, err := s.Daos.GetSinglePlanCRFRuleByScenario(ctx, request.Scenario)
		if err != nil {
			return errors.New("Not able to get RE - " + err.Error())
		}
		if re == nil {
			return errors.New("Re is null")
		}
		t := time.Now()
		request.PlanTimeline.On = &t
		request.PlanTimeline.Type = re.Log.Type
		request.PlanTimeline.TypeLabel = re.Log.Label
		m := make(map[string]interface{})
		m["status"] = re.To
		m["remarks"] = request.PlanTimeline.Remarks
		err = s.Daos.PlanCRFFlowUpdate(ctx, request.CRFID, m, request.PlanTimeline)
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

//CommonCRFFlowUpdate : ""
func (s *Service) CommonCRFFlowUpdate(ctx *models.Context, CRFID string, scenario string, timeline *models.PlanTimeline) error {
	re, err := s.Daos.GetSinglePlanCRFRuleByScenario(ctx, scenario)
	if err != nil {
		return errors.New("Not able to get RE - " + err.Error())
	}
	if re == nil {
		return errors.New("Re is null")
	}
	t := time.Now()
	timeline.On = &t
	timeline.Type = re.Log.Type
	timeline.TypeLabel = re.Log.Label
	m := make(map[string]interface{})
	m["status"] = re.To
	m["remarks"] = timeline.Remarks
	err = s.Daos.PlanCRFFlowUpdate(ctx, CRFID, m, *timeline)
	if err != nil {
		return errors.New("Not able to update plan - " + err.Error())
	}
	return nil
}

//PlanCRFReapply : ""
func (s *Service) PlanCRFReapply(ctx *models.Context, request *models.PlanCRFReapply) error {
	log.Println("transaction start")
	//Start Transaction
	if err := ctx.Session.StartTransaction(); err != nil {
		return err
	}
	defer ctx.Session.EndSession(ctx.CTX)
	if err := mongo.WithSession(ctx.CTX, ctx.Session, func(sc mongo.SessionContext) error {

		re, err := s.Daos.GetSinglePlanCRFRuleByScenario(ctx, request.Scenario)
		if err != nil {
			return errors.New("Not able to get RE - " + err.Error())
		}
		if re == nil {
			return errors.New("Re is null")
		}
		t := time.Now()
		request.PlanTimeline.On = &t
		request.PlanTimeline.Type = re.Log.Type
		request.PlanTimeline.TypeLabel = re.Log.Label
		m := make(map[string]interface{})
		m["status"] = re.To
		m["remarks"] = request.PlanTimeline.Remarks
		err = s.Daos.PlanCRFFlowUpdate(ctx, request.CRFID, m, request.PlanTimeline)
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

//PlanCRFReject : ""
func (s *Service) PlanCRFReject(ctx *models.Context, request *models.PlanCRFReject) error {
	log.Println("transaction start")
	//Start Transaction
	if err := ctx.Session.StartTransaction(); err != nil {
		return err
	}
	defer ctx.Session.EndSession(ctx.CTX)
	if err := mongo.WithSession(ctx.CTX, ctx.Session, func(sc mongo.SessionContext) error {

		re, err := s.Daos.GetSinglePlanCRFRuleByScenario(ctx, request.Scenario)
		if err != nil {
			return errors.New("Not able to get RE - " + err.Error())
		}
		if re == nil {
			return errors.New("Re is null")
		}
		t := time.Now()
		request.PlanTimeline.On = &t
		request.PlanTimeline.Type = re.Log.Type
		request.PlanTimeline.TypeLabel = re.Log.Label
		m := make(map[string]interface{})
		m["status"] = re.To
		m["remarks"] = request.PlanTimeline.Remarks
		err = s.Daos.PlanCRFFlowUpdate(ctx, request.CRFID, m, request.PlanTimeline)
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
