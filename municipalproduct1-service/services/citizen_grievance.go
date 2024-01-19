package services

import (
	"errors"
	"log"
	"municipalproduct1-service/constants"
	"municipalproduct1-service/models"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
)

// SaveCitizenGrievance : ""
func (s *Service) SaveCitizenGrievance(ctx *models.Context, citizens *models.CitizenGrievance) error {
	log.Println("transaction start")
	// Start Transaction
	if err := ctx.Session.StartTransaction(); err != nil {
		return err
	}
	defer ctx.Session.EndSession(ctx.CTX)
	citizens.UniqueID = s.Daos.GetUniqueID(ctx, constants.COLLECTIONCITIZENGRAVIANS)
	citizens.Status = constants.CITIZENGRAVIANSSTATUSINIT
	t := time.Now()
	citizens.Requestor.On = &t
	citizens.Created = new(models.CreatedV2)
	citizens.Created.On = &t
	citizens.Created.By = constants.SYSTEM
	if err := mongo.WithSession(ctx.CTX, ctx.Session, func(sc mongo.SessionContext) error {

		dberr := s.Daos.SaveCitizenGrievance(ctx, citizens)
		if dberr != nil {
			return dberr
		}
		// refUser, err := s.Daos.GetSingleUser(ctx, citizens.ByID)
		// if err != nil {
		// 	return errors.New("error in getting the user - " + err.Error())
		// }
		citizenLog := new(models.CitizenGraviansLog)
		citizenLog.On = &t
		citizenLog.UniqueID = s.Daos.GetUniqueID(ctx, constants.COLLECTIONCITIZENGRAVIANSLOG)
		citizenLog.CitizenGraviansID = citizens.UniqueID
		citizenLog.PreviousStatus = "NA"
		citizenLog.NewStatus = constants.CITIZENGRAVIANSSTATUSINIT
		citizenLog.Desc = "Citizen Gravians Updated for UniqueId" + citizens.UniqueID
		// citizenLog.By = refUser.Name
		citizenLog.ByID = citizens.ByID
		citizenLog.ByType = citizens.ByType
		citizenLog.Type = citizens.Type

		err := s.Daos.SaveCitizenGraviansLog(ctx, citizenLog)
		if err != nil {
			return err
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

//GetSingleCitizenGrievance :""
func (s *Service) GetSingleCitizenGrievance(ctx *models.Context, UniqueID string) (*models.RefCitizenGrievance, error) {
	tower, err := s.Daos.GetSingleCitizenGrievance(ctx, UniqueID)
	if err != nil {
		return nil, err
	}
	return tower, nil
}

// UpdateCitizenGrievance : ""
func (s *Service) UpdateCitizenGrievance(ctx *models.Context, citizens *models.CitizenGrievance) error {
	if err := ctx.Session.StartTransaction(); err != nil {
		return err
	}
	defer ctx.Session.EndSession(ctx.CTX)
	if err := mongo.WithSession(ctx.CTX, ctx.Session, func(sc mongo.SessionContext) error {
		citizengravians, err := s.GetSingleCitizenGrievance(ctx, citizens.UniqueID)
		if err != nil {
			return err
		}
		if citizengravians == nil {
			return errors.New("error in getting the citizengravians - " + err.Error())
		}

		err = s.Daos.UpdateCitizenGrievance(ctx, citizens)
		if err != nil {
			return err
		}
		refUser, err := s.Daos.GetSingleUser(ctx, citizens.ByID)
		if err != nil {
			return errors.New("error in getting the user - " + err.Error())
		}
		citizenLog := new(models.CitizenGraviansLog)
		citizenLog.UniqueID = s.Daos.GetUniqueID(ctx, constants.COLLECTIONCITIZENGRAVIANSLOG)
		citizenLog.CitizenGraviansID = citizens.UniqueID

		t := time.Now()
		citizenLog.On = &t
		citizenLog.PreviousStatus = citizengravians.Status
		citizenLog.NewStatus = citizens.Status
		citizenLog.Desc = "Citizen Gravians Updated for UniqueId" + citizenLog.UniqueID
		citizenLog.By = refUser.Name
		citizenLog.ByID = citizens.ByID
		citizenLog.ByType = citizens.ByType
		err = s.Daos.SaveCitizenGraviansLog(ctx, citizenLog)
		if err != nil {
			return err
		}

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

// EnableCitizenGrievance : ""
func (s *Service) EnableCitizenGrievance(ctx *models.Context, citizen *models.CitizenGrievance) error {
	if err := ctx.Session.StartTransaction(); err != nil {
		return err
	}
	defer ctx.Session.EndSession(ctx.CTX)
	if err := mongo.WithSession(ctx.CTX, ctx.Session, func(sc mongo.SessionContext) error {
		citizengravians, err := s.GetSingleCitizenGrievance(ctx, citizen.UniqueID)
		if err != nil {
			return err
		}
		if citizengravians == nil {
			return errors.New("error in getting the citizengravians - " + err.Error())
		}

		err = s.Daos.EnableCitizenGrievance(ctx, citizen)
		if err != nil {
			return err
		}
		refUser, err := s.Daos.GetSingleUser(ctx, citizen.ByID)
		if err != nil {
			return errors.New("error in getting the user - " + err.Error())
		}
		citizenLog := new(models.CitizenGraviansLog)
		citizenLog.UniqueID = s.Daos.GetUniqueID(ctx, constants.COLLECTIONCITIZENGRAVIANSLOG)
		citizenLog.CitizenGraviansID = citizen.UniqueID
		t := time.Now()
		citizenLog.On = &t
		citizenLog.PreviousStatus = citizengravians.Status
		citizenLog.NewStatus = constants.CITIZENGRAVIANSSTATUSACTIVE
		citizenLog.Desc = "Citizen Gravians Updated for UniqueId" + citizenLog.UniqueID
		citizenLog.By = refUser.Name
		citizenLog.ByID = citizen.ByID
		citizenLog.ByType = citizen.ByType

		err = s.Daos.SaveCitizenGraviansLog(ctx, citizenLog)
		if err != nil {
			return err
		}

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

//DisableCitizenGrievance : ""
func (s *Service) DisableCitizenGrievance(ctx *models.Context, citizen *models.RejectedCitizengravians) error {
	if err := ctx.Session.StartTransaction(); err != nil {
		return err
	}
	defer ctx.Session.EndSession(ctx.CTX)
	if err := mongo.WithSession(ctx.CTX, ctx.Session, func(sc mongo.SessionContext) error {
		citizengravians, err := s.GetSingleCitizenGrievance(ctx, citizen.UniqueID)
		if err != nil {
			return err
		}
		if citizengravians == nil {
			return errors.New("error in getting the citizengravians - " + err.Error())
		}

		err = s.Daos.DisableCitizenGrievance(ctx, citizen.UniqueID)
		if err != nil {
			return err
		}
		refUser, err := s.Daos.GetSingleUser(ctx, citizen.ByID)
		if err != nil {
			return errors.New("error in getting the user - " + err.Error())
		}
		citizenLog := new(models.CitizenGraviansLog)
		citizenLog.UniqueID = s.Daos.GetUniqueID(ctx, constants.COLLECTIONCITIZENGRAVIANSLOG)
		citizenLog.CitizenGraviansID = citizen.UniqueID
		t := time.Now()
		citizenLog.On = &t
		citizenLog.PreviousStatus = citizengravians.Status
		citizenLog.NewStatus = constants.CITIZENGRAVIANSSTATUSDISABLED
		citizenLog.Desc = "Citizen Gravians Updated for UniqueId" + citizenLog.UniqueID
		citizenLog.By = refUser.Name
		citizenLog.ByID = citizen.ByID
		citizenLog.ByType = citizen.ByType

		err = s.Daos.SaveCitizenGraviansLog(ctx, citizenLog)
		if err != nil {
			return err
		}

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

// DeleteCitizenGrievance : ""
func (s *Service) DeleteCitizenGrievance(ctx *models.Context, citizen *models.RejectedCitizengravians) error {
	if err := ctx.Session.StartTransaction(); err != nil {
		return err
	}
	defer ctx.Session.EndSession(ctx.CTX)
	if err := mongo.WithSession(ctx.CTX, ctx.Session, func(sc mongo.SessionContext) error {
		citizengravians, err := s.GetSingleCitizenGrievance(ctx, citizen.UniqueID)
		if err != nil {
			return err
		}
		if citizengravians == nil {
			return errors.New("error in getting the citizengravians - " + err.Error())
		}

		err = s.Daos.DeleteCitizenGrievance(ctx, citizen.UniqueID)
		if err != nil {
			return err
		}
		refUser, err := s.Daos.GetSingleUser(ctx, citizen.ByID)
		if err != nil {
			return errors.New("error in getting the user - " + err.Error())
		}
		citizenLog := new(models.CitizenGraviansLog)
		citizenLog.UniqueID = s.Daos.GetUniqueID(ctx, constants.COLLECTIONCITIZENGRAVIANSLOG)
		citizenLog.CitizenGraviansID = citizen.UniqueID
		t := time.Now()
		citizenLog.On = &t
		citizenLog.PreviousStatus = citizengravians.Status
		citizenLog.NewStatus = constants.CITIZENGRAVIANSSTATUSDELETED
		citizenLog.Desc = "Citizen Gravians Updated for UniqueId" + citizenLog.UniqueID
		citizenLog.By = refUser.Name
		citizenLog.ByID = citizen.ByID
		citizenLog.ByType = citizen.ByType
		err = s.Daos.SaveCitizenGraviansLog(ctx, citizenLog)
		if err != nil {
			return err
		}

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

// CompletedCitizenGrievance: ""
func (s *Service) CompletedCitizenGrievance(ctx *models.Context, citizen *models.CitizenGrievance) error {
	if err := ctx.Session.StartTransaction(); err != nil {
		return err
	}
	defer ctx.Session.EndSession(ctx.CTX)
	if err := mongo.WithSession(ctx.CTX, ctx.Session, func(sc mongo.SessionContext) error {
		citizengravians, err := s.GetSingleCitizenGrievance(ctx, citizen.UniqueID)
		if err != nil {
			return err
		}
		if citizengravians == nil {
			return errors.New("error in getting the citizengravians - " + err.Error())
		}
		err = s.Daos.CompletedCitizenGrievance(ctx, citizen)
		if err != nil {
			return err
		}
		refUser, err := s.Daos.GetSingleUser(ctx, citizen.ByID)
		if err != nil {
			return errors.New("error in getting the user - " + err.Error())
		}
		citizenLog := new(models.CitizenGraviansLog)
		citizenLog.UniqueID = s.Daos.GetUniqueID(ctx, constants.COLLECTIONCITIZENGRAVIANSLOG)
		citizenLog.CitizenGraviansID = citizen.UniqueID
		t := time.Now()
		citizenLog.On = &t
		citizenLog.PreviousStatus = citizengravians.Status
		citizenLog.NewStatus = constants.CITIZENGRAVIANSSTATUSCOMPLETED
		citizenLog.Desc = "Citizen Gravians Updated for UniqueId" + citizenLog.UniqueID
		citizenLog.By = refUser.Name
		citizenLog.ByID = citizen.ByID
		citizenLog.ByType = citizen.ByType
		err = s.Daos.SaveCitizenGraviansLog(ctx, citizenLog)
		if err != nil {
			return err
		}
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

// RejectedCitizenGrievance : ""
func (s *Service) RejectedCitizenGrievance(ctx *models.Context, citizen *models.RejectedCitizengravians) error {
	if err := ctx.Session.StartTransaction(); err != nil {
		return err
	}
	defer ctx.Session.EndSession(ctx.CTX)
	if err := mongo.WithSession(ctx.CTX, ctx.Session, func(sc mongo.SessionContext) error {
		citizengravians, err := s.GetSingleCitizenGrievance(ctx, citizen.UniqueID)
		if err != nil {
			return err
		}
		if citizengravians == nil {
			return errors.New("error in getting the citizengravians - " + err.Error())
		}

		err = s.Daos.RejectedCitizenGrievance(ctx, citizen.UniqueID)
		if err != nil {
			return err
		}
		refUser, err := s.Daos.GetSingleUser(ctx, citizen.ByID)
		if err != nil {
			return errors.New("error in getting the user - " + err.Error())
		}
		citizenLog := new(models.CitizenGraviansLog)
		citizenLog.UniqueID = s.Daos.GetUniqueID(ctx, constants.COLLECTIONCITIZENGRAVIANSLOG)
		citizenLog.CitizenGraviansID = citizen.UniqueID
		t := time.Now()
		citizenLog.On = &t
		citizenLog.PreviousStatus = citizengravians.Status
		citizenLog.NewStatus = constants.CITIZENGRAVIANSSTATUSREJECTED
		citizenLog.Desc = "Citizen Gravians Updated for UniqueId" + citizen.UniqueID
		citizenLog.By = refUser.Name
		citizenLog.ByID = citizen.ByID
		citizenLog.ByType = citizen.ByType

		err = s.Daos.SaveCitizenGraviansLog(ctx, citizenLog)
		if err != nil {
			return err
		}

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

// FilterCitizenGrievance : ""
func (s *Service) FilterCitizenGrievance(ctx *models.Context, filter *models.CitizenGrievanceFilter, pagination *models.Pagination) ([]models.RefCitizenGrievance, error) {
	return s.Daos.FilterCitizenGrievance(ctx, filter, pagination)

}

// UpdateCitizenGrievance : ""
func (s *Service) UpdateCitizenGrievanceSolution(ctx *models.Context, solution *models.CitizenGrievanceSolution) error {
	if err := ctx.Session.StartTransaction(); err != nil {
		return err
	}
	defer ctx.Session.EndSession(ctx.CTX)
	if err := mongo.WithSession(ctx.CTX, ctx.Session, func(sc mongo.SessionContext) error {
		citizengravians, err := s.GetSingleCitizenGrievance(ctx, solution.CitizenGrievanceID)
		if err != nil {
			return err
		}
		if citizengravians == nil {
			return errors.New("error in getting the citizengravians - " + err.Error())
		}
		refUser, err := s.Daos.GetSingleUser(ctx, solution.ByID)
		if err != nil {
			return errors.New("error in getting the user - " + err.Error())
		}
		t := time.Now()
		solution.SolutionDate = &t
		solution.Status = constants.CITIZENGRAVIANSLOGSTATUSCOMPLETED
		solution.By = refUser.Name
		err = s.Daos.UpdateCitizenGrievanceSolution(ctx, solution)
		if err != nil {
			return err
		}

		citizenLog := new(models.CitizenGraviansLog)
		citizenLog.UniqueID = s.Daos.GetUniqueID(ctx, constants.COLLECTIONCITIZENGRAVIANSLOG)
		citizenLog.CitizenGraviansID = solution.CitizenGrievanceID

		t1 := time.Now()
		citizenLog.On = &t1
		citizenLog.PreviousStatus = citizengravians.Status
		citizenLog.NewStatus = citizengravians.Status
		citizenLog.Desc = "Citizen Gravians Updated for UniqueId" + citizenLog.UniqueID
		citizenLog.By = refUser.Name
		citizenLog.ByID = solution.ByID
		citizenLog.ByType = solution.ByType
		err = s.Daos.SaveCitizenGraviansLog(ctx, citizenLog)
		if err != nil {
			return err
		}

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
