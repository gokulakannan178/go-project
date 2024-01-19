package service

import (
	"errors"
	"lgf-ccc-service/constants"
	"lgf-ccc-service/models"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
)

// SaveProperties : ""
func (s *Service) SaveProperties(ctx *models.Context, properties *models.Properties) error {
	log.Println("transaction start")
	// Start Transaction
	if err := ctx.Session.StartTransaction(); err != nil {
		return err
	}
	defer ctx.Session.EndSession(ctx.CTX)
	properties.UniqueID = s.Daos.GetUniqueID(ctx, constants.COLLECTIONPROPERTIES)
	properties.Status = constants.PROPERTIESSTATUSACTIVE
	t := time.Now()
	properties.RegisterDate = &t
	created := models.CreatedV2{}
	created.On = &t
	created.By = constants.SYSTEM
	log.Println("b4 Properties.created")
	properties.Created = &created
	log.Println("b4 Properties.created")
	if err := mongo.WithSession(ctx.CTX, ctx.Session, func(sc mongo.SessionContext) error {
		// tempUser, dberr := s.Daos.GetSinglePropertyWithCondition(ctx, "mobile", properties.Mobile)
		// if dberr != nil {
		// 	return errors.New("Db Error" + dberr.Error())
		// }
		// if tempUser != nil {
		// 	return errors.New("mobile no already in use")
		// }
		dberr := s.Daos.SaveProperties(ctx, properties)
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

// GetSingleProperties : ""
func (s *Service) GetSingleProperties(ctx *models.Context, UniqueID string) (*models.RefProperties, error) {
	properties, err := s.Daos.GetSingleProperties(ctx, UniqueID)
	if err != nil {
		return nil, err
	}
	return properties, nil
}
func (s *Service) GetSinglePropertiesWithHoldingNumber(ctx *models.Context, holdingNumber string) (*models.RefProperties, error) {
	properties, err := s.Daos.GetSinglePropertiesWithHoldingNumber(ctx, holdingNumber)
	if err != nil {
		return nil, err
	}
	return properties, nil
}

//UpdateProperties : ""
func (s *Service) UpdateProperties(ctx *models.Context, properties *models.Properties) error {
	if err := ctx.Session.StartTransaction(); err != nil {
		return err
	}
	defer ctx.Session.EndSession(ctx.CTX)
	if err := mongo.WithSession(ctx.CTX, ctx.Session, func(sc mongo.SessionContext) error {

		// property, err := s.Daos.CheckResgisterMobileNumber(ctx, properties.Mobile, properties.UniqueID)
		// if err != nil {
		// 	return err
		// }
		// if property != nil {
		// 	return errors.New("Mobile No Already In use" + err.Error())
		// }
		err := s.Daos.UpdateProperties(ctx, properties)
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

// EnableProperties : ""
func (s *Service) EnableProperties(ctx *models.Context, uniqueID string) error {

	if err := ctx.Session.StartTransaction(); err != nil {
		return err
	}
	defer ctx.Session.EndSession(ctx.CTX)
	if err := mongo.WithSession(ctx.CTX, ctx.Session, func(sc mongo.SessionContext) error {
		dberr := s.Daos.EnableProperties(ctx, uniqueID)
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

// DisableProperties : ""
func (s *Service) DisableProperties(ctx *models.Context, uniqueID string) error {

	if err := ctx.Session.StartTransaction(); err != nil {
		return err
	}
	defer ctx.Session.EndSession(ctx.CTX)
	if err := mongo.WithSession(ctx.CTX, ctx.Session, func(sc mongo.SessionContext) error {
		debrr := s.Daos.DisableProperties(ctx, uniqueID)
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

//DeleteProperties : ""
func (s *Service) DeleteProperties(ctx *models.Context, UniqueID string) error {
	if err := ctx.Session.StartTransaction(); err != nil {
		return err
	}
	defer ctx.Session.EndSession(ctx.CTX)
	if err := mongo.WithSession(ctx.CTX, ctx.Session, func(sc mongo.SessionContext) error {

		err := s.Daos.DeleteProperties(ctx, UniqueID)
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

// // InProgressProperties : ""
// func (s *Service) InProgressProperties(ctx *models.Context, uniqueID string) error {

// 	if err := ctx.Session.StartTransaction(); err != nil {
// 		return err
// 	}
// 	defer ctx.Session.EndSession(ctx.CTX)
// 	if err := mongo.WithSession(ctx.CTX, ctx.Session, func(sc mongo.SessionContext) error {
// 		dberr := s.Daos.InProgressProperties(ctx, uniqueID)
// 		if dberr != nil {
// 			return dberr
// 		}
// 		if err := sc.CommitTransaction(sc); err != nil {
// 			return errors.New("Not able to commit - " + err.Error())
// 		}
// 		return nil
// 	}); err != nil {
// 		if err1 := ctx.Session.AbortTransaction(ctx.CTX); err1 != nil {
// 			log.Println("err in abort")
// 			return errors.New("Transaction Aborted with error" + err1.Error())
// 		}
// 		return err
// 	}

// 	return nil
// }

// // PendingProperties : ""
// func (s *Service) PendingProperties(ctx *models.Context, uniqueID string) error {

// 	if err := ctx.Session.StartTransaction(); err != nil {
// 		return err
// 	}
// 	defer ctx.Session.EndSession(ctx.CTX)
// 	if err := mongo.WithSession(ctx.CTX, ctx.Session, func(sc mongo.SessionContext) error {
// 		debrr := s.Daos.PendingProperties(ctx, uniqueID)
// 		if debrr != nil {
// 			return debrr
// 		}
// 		if err := sc.CommitTransaction(sc); err != nil {
// 			return errors.New("Not able to commit - " + err.Error())
// 		}
// 		return nil
// 	}); err != nil {
// 		if err1 := ctx.Session.AbortTransaction(ctx.CTX); err1 != nil {
// 			log.Println("err in abort")
// 			return errors.New("Transaction Abort with error" + err1.Error())
// 		}
// 		return err
// 	}
// 	return nil
// }

// //InitProperties : ""
// func (s *Service) InitProperties(ctx *models.Context, UniqueID string) error {
// 	if err := ctx.Session.StartTransaction(); err != nil {
// 		return err
// 	}
// 	defer ctx.Session.EndSession(ctx.CTX)
// 	if err := mongo.WithSession(ctx.CTX, ctx.Session, func(sc mongo.SessionContext) error {

// 		err := s.Daos.InitProperties(ctx, UniqueID)
// 		if err != nil {
// 			if err = ctx.Session.AbortTransaction(sc); err != nil {
// 				return errors.New("Transaction Aborted with error" + err.Error())
// 			}
// 			return errors.New("Transaction Aborted - " + err.Error())
// 		}
// 		return nil

// 	}); err != nil {
// 		return err
// 	}
// 	return nil
// }

// //CompletedProperties : ""
// func (s *Service) CompletedProperties(ctx *models.Context, Properties *models.Properties) error {
// 	if err := ctx.Session.StartTransaction(); err != nil {
// 		return err
// 	}
// 	defer ctx.Session.EndSession(ctx.CTX)
// 	if err := mongo.WithSession(ctx.CTX, ctx.Session, func(sc mongo.SessionContext) error {
// 		t := time.Now()
// 		Properties.CompletionDate = &t
// 		Properties.Status = constants.PropertiesSTATUSCOMPLETED

// 		err := s.Daos.CompletedProperties(ctx, Properties)
// 		if err != nil {
// 			if err = ctx.Session.AbortTransaction(sc); err != nil {
// 				return errors.New("Transaction Aborted with error" + err.Error())
// 			}
// 			return errors.New("Transaction Aborted - " + err.Error())
// 		}
// 		return nil

// 	}); err != nil {
// 		return err
// 	}
// 	return nil
// }

// FilterProperties : ""
func (s *Service) FilterProperties(ctx *models.Context, properties *models.FilterProperties, pagination *models.Pagination) (Propertiess []models.RefProperties, err error) {
	return s.Daos.FilterProperties(ctx, properties, pagination)
}

// func (s *Service) GetDetailProperties(ctx *models.Context, UniqueID string) (*models.RefProperties, error) {
// 	Properties, err := s.Daos.GetDetailProperties(ctx, UniqueID)
// 	if err != nil {
// 		return nil, err
// 	}
// 	return Properties, nil
// }

// //UpdateProperties : ""
// func (s *Service) AssignProperties(ctx *models.Context, Properties *models.Properties) error {
// 	if err := ctx.Session.StartTransaction(); err != nil {
// 		return err
// 	}
// 	defer ctx.Session.EndSession(ctx.CTX)
// 	if err := mongo.WithSession(ctx.CTX, ctx.Session, func(sc mongo.SessionContext) error {

// 		err := s.Daos.AssignProperties(ctx, Properties)
// 		if err != nil {
// 			if err = ctx.Session.AbortTransaction(sc); err != nil {
// 				return errors.New("Transaction Aborted with error" + err.Error())
// 			}
// 			return errors.New("Transaction Aborted - " + err.Error())
// 		}
// 		return nil

// 	}); err != nil {
// 		return err
// 	}
// 	return nil
// }
