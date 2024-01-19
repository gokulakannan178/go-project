package service

import (
	"errors"
	"lgf-ccc-service/constants"
	"lgf-ccc-service/models"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
)

// SaveServiceRequest : ""
func (s *Service) SaveServiceRequest(ctx *models.Context, serviceRequest *models.ServiceRequest) error {
	log.Println("transaction start")
	// Start Transaction
	if err := ctx.Session.StartTransaction(); err != nil {
		return err
	}
	defer ctx.Session.EndSession(ctx.CTX)
	serviceRequest.UniqueID = s.Daos.GetUniqueID(ctx, constants.COLLECTIONSERVICEREQUEST)
	serviceRequest.Status = constants.SERVICEREQUESTSTATUSINIT
	t := time.Now()
	created := models.CreatedV2{}
	created.On = &t
	created.By = constants.SYSTEM
	log.Println("b4 ServiceRequest.created")
	serviceRequest.Created = &created
	serviceRequest.RequstedDate = &t
	log.Println("b4 ServiceRequest.created")
	if err := mongo.WithSession(ctx.CTX, ctx.Session, func(sc mongo.SessionContext) error {
		dberr := s.Daos.SaveServiceRequest(ctx, serviceRequest)
		if dberr != nil {
			return dberr
		}

		// for _, v := range ServiceRequest.LeaveMasterId {
		// 	policyrule := new(models.PolicyRule)
		// 	leavemaster, err := s.Daos.GetSingleLeaveMasterWithActive(ctx, v, constants.ServiceRequestSTATUSACTIVE)
		// 	if err != nil {
		// 		fmt.Println(err)
		// 		return err
		// 	}
		// 	if leavemaster == nil {
		// 		return errors.New("Leave master id is not available" + err.Error())
		// 	}
		// 	fmt.Println("leavemaster", leavemaster)
		// 	policyrule.ServiceRequestID = ServiceRequest.UniqueID
		// 	policyrule.LeaveMasterID = v
		// 	policyrule.Name = ServiceRequest.Name
		// 	//policyrule.Value = leavemaster.Value
		// 	err = s.SavePolicyRuleWithoutTransaction(ctx, policyrule)
		// 	if err != nil {
		// 		return err
		// 	}
		// }

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

// GetSingleServiceRequest : ""
func (s *Service) GetSingleServiceRequest(ctx *models.Context, UniqueID string) (*models.RefServiceRequest, error) {
	ServiceRequest, err := s.Daos.GetSingleServiceRequest(ctx, UniqueID)
	if err != nil {
		return nil, err
	}
	return ServiceRequest, nil
}

//UpdateServiceRequest : ""
func (s *Service) UpdateServiceRequest(ctx *models.Context, serviceRequest *models.ServiceRequest) error {
	if err := ctx.Session.StartTransaction(); err != nil {
		return err
	}
	defer ctx.Session.EndSession(ctx.CTX)
	if err := mongo.WithSession(ctx.CTX, ctx.Session, func(sc mongo.SessionContext) error {
		// if len(ServiceRequest.LeaveMasterId) > 0 {
		// 	err := s.Daos.PolicyRuleRemoveNotPresentValue(ctx, ServiceRequest.UniqueID, ServiceRequest.LeaveMasterId)
		// 	if err != nil {
		// 		return err
		// 	}
		// }
		// err := s.Daos.PolicyRuleUpsert(ctx, ServiceRequest.UniqueID, ServiceRequest.LeaveMasterId, ServiceRequest.Name)
		// if err != nil {
		// 	return err
		// }

		//fmt.Println("error==>", err)

		err := s.Daos.UpdateServiceRequest(ctx, serviceRequest)
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

// EnableServiceRequest : ""
func (s *Service) EnableServiceRequest(ctx *models.Context, uniqueID string) error {

	if err := ctx.Session.StartTransaction(); err != nil {
		return err
	}
	defer ctx.Session.EndSession(ctx.CTX)
	if err := mongo.WithSession(ctx.CTX, ctx.Session, func(sc mongo.SessionContext) error {
		dberr := s.Daos.EnableServiceRequest(ctx, uniqueID)
		if dberr != nil {
			return dberr
		}
		if err := sc.CommitTransaction(sc); err != nil {
			return errors.New("Not able to commit - " + err.Error())
		}
		return nil
	}); err != nil {
		if err1 := ctx.Session.AbortTransaction(ctx.CTX); err1 != nil {
			log.Println("err in abort")
			return errors.New("Transaction Aborted with error" + err1.Error())
		}
		return err
	}

	return nil
}

// DisableServiceRequest : ""
func (s *Service) DisableServiceRequest(ctx *models.Context, uniqueID string) error {

	if err := ctx.Session.StartTransaction(); err != nil {
		return err
	}
	defer ctx.Session.EndSession(ctx.CTX)
	if err := mongo.WithSession(ctx.CTX, ctx.Session, func(sc mongo.SessionContext) error {
		debrr := s.Daos.DisableServiceRequest(ctx, uniqueID)
		if debrr != nil {
			return debrr
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

//DeleteServiceRequest : ""
func (s *Service) DeleteServiceRequest(ctx *models.Context, UniqueID string) error {
	if err := ctx.Session.StartTransaction(); err != nil {
		return err
	}
	defer ctx.Session.EndSession(ctx.CTX)
	if err := mongo.WithSession(ctx.CTX, ctx.Session, func(sc mongo.SessionContext) error {

		err := s.Daos.DeleteServiceRequest(ctx, UniqueID)
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

// InProgressServiceRequest : ""
func (s *Service) InProgressServiceRequest(ctx *models.Context, uniqueID string) error {

	if err := ctx.Session.StartTransaction(); err != nil {
		return err
	}
	defer ctx.Session.EndSession(ctx.CTX)
	if err := mongo.WithSession(ctx.CTX, ctx.Session, func(sc mongo.SessionContext) error {
		dberr := s.Daos.InProgressServiceRequest(ctx, uniqueID)
		if dberr != nil {
			return dberr
		}
		if err := sc.CommitTransaction(sc); err != nil {
			return errors.New("Not able to commit - " + err.Error())
		}
		return nil
	}); err != nil {
		if err1 := ctx.Session.AbortTransaction(ctx.CTX); err1 != nil {
			log.Println("err in abort")
			return errors.New("Transaction Aborted with error" + err1.Error())
		}
		return err
	}

	return nil
}

// PendingServiceRequest : ""
func (s *Service) PendingServiceRequest(ctx *models.Context, uniqueID string) error {

	if err := ctx.Session.StartTransaction(); err != nil {
		return err
	}
	defer ctx.Session.EndSession(ctx.CTX)
	if err := mongo.WithSession(ctx.CTX, ctx.Session, func(sc mongo.SessionContext) error {
		debrr := s.Daos.PendingServiceRequest(ctx, uniqueID)
		if debrr != nil {
			return debrr
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

//InitServiceRequest : ""
func (s *Service) InitServiceRequest(ctx *models.Context, UniqueID string) error {
	if err := ctx.Session.StartTransaction(); err != nil {
		return err
	}
	defer ctx.Session.EndSession(ctx.CTX)
	if err := mongo.WithSession(ctx.CTX, ctx.Session, func(sc mongo.SessionContext) error {

		err := s.Daos.InitServiceRequest(ctx, UniqueID)
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

//CompletedServiceRequest : ""
func (s *Service) CompletedServiceRequest(ctx *models.Context, serviceRequest *models.ServiceRequest) error {
	if err := ctx.Session.StartTransaction(); err != nil {
		return err
	}
	defer ctx.Session.EndSession(ctx.CTX)
	if err := mongo.WithSession(ctx.CTX, ctx.Session, func(sc mongo.SessionContext) error {

		err := s.Daos.CompletedServiceRequest(ctx, serviceRequest)
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

// FilterServiceRequest : ""
func (s *Service) FilterServiceRequest(ctx *models.Context, serviceRequest *models.FilterServiceRequest, pagination *models.Pagination) (ServiceRequests []models.RefServiceRequest, err error) {
	return s.Daos.FilterServiceRequest(ctx, serviceRequest, pagination)
}

func (s *Service) GetDetailServiceRequest(ctx *models.Context, UniqueID string) (*models.RefServiceRequest, error) {
	ServiceRequest, err := s.Daos.GetDetailServiceRequest(ctx, UniqueID)
	if err != nil {
		return nil, err
	}
	return ServiceRequest, nil
}

//UpdateServiceRequest : ""
func (s *Service) AssignServiceRequest(ctx *models.Context, serviceRequest *models.ServiceRequest) error {
	if err := ctx.Session.StartTransaction(); err != nil {
		return err
	}
	defer ctx.Session.EndSession(ctx.CTX)
	if err := mongo.WithSession(ctx.CTX, ctx.Session, func(sc mongo.SessionContext) error {

		err := s.Daos.AssignServiceRequest(ctx, serviceRequest)
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
