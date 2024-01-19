package services

import (
	"errors"

	"hrms-services/constants"
	"hrms-services/models"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
)

//SaveEmployeeDocuments :""
func (s *Service) SaveEmployeeDocuments(ctx *models.Context, employeeDocuments *models.EmployeeDocuments) error {
	log.Println("transaction start")
	//Start Transaction
	if err := ctx.Session.StartTransaction(); err != nil {
		return err
	}
	defer ctx.Session.EndSession(ctx.CTX)
	employeeDocuments.UniqueID = s.Daos.GetUniqueID(ctx, constants.COLLECTIONEMPLOYEEDOCUMENTS)
	employeeDocuments.Status = constants.EMPLOYEEDOCUMENTSSTATUSACTIVE
	t := time.Now()
	created := models.Created{}
	created.On = &t
	created.By = constants.SYSTEM
	log.Println("b4 EmployeeDocuments.created")
	employeeDocuments.Created = created
	log.Println("b4 EmployeeDocuments.created")
	if err := mongo.WithSession(ctx.CTX, ctx.Session, func(sc mongo.SessionContext) error {
		dberr := s.Daos.SaveEmployeeDocuments(ctx, employeeDocuments)
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

//UpdateEmployeeDocuments : ""
func (s *Service) UpdateEmployeeDocuments(ctx *models.Context, employeeDocuments *models.EmployeeDocuments) error {
	if err := ctx.Session.StartTransaction(); err != nil {
		return err
	}
	defer ctx.Session.EndSession(ctx.CTX)
	if err := mongo.WithSession(ctx.CTX, ctx.Session, func(sc mongo.SessionContext) error {

		err := s.Daos.UpdateEmployeeDocuments(ctx, employeeDocuments)
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

//EnableEmployeeDocuments : ""
func (s *Service) EnableEmployeeDocuments(ctx *models.Context, UniqueID string) error {
	if err := ctx.Session.StartTransaction(); err != nil {
		return err
	}
	defer ctx.Session.EndSession(ctx.CTX)
	if err := mongo.WithSession(ctx.CTX, ctx.Session, func(sc mongo.SessionContext) error {

		err := s.Daos.EnableEmployeeDocuments(ctx, UniqueID)
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

//DisableEmployeeDocuments : ""
func (s *Service) DisableEmployeeDocuments(ctx *models.Context, UniqueID string) error {
	if err := ctx.Session.StartTransaction(); err != nil {
		return err
	}
	defer ctx.Session.EndSession(ctx.CTX)
	if err := mongo.WithSession(ctx.CTX, ctx.Session, func(sc mongo.SessionContext) error {

		err := s.Daos.DisableEmployeeDocuments(ctx, UniqueID)
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

//DeleteEmployeeDocuments : ""
func (s *Service) DeleteEmployeeDocuments(ctx *models.Context, UniqueID string) error {
	if err := ctx.Session.StartTransaction(); err != nil {
		return err
	}
	defer ctx.Session.EndSession(ctx.CTX)
	if err := mongo.WithSession(ctx.CTX, ctx.Session, func(sc mongo.SessionContext) error {

		err := s.Daos.DeleteEmployeeDocuments(ctx, UniqueID)
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

//GetSingleEmployeeDocuments :""
func (s *Service) GetSingleEmployeeDocuments(ctx *models.Context, UniqueID string) (*models.RefEmployeeDocuments, error) {
	employeeDocuments, err := s.Daos.GetSingleEmployeeDocuments(ctx, UniqueID)
	if err != nil {
		return nil, err
	}
	return employeeDocuments, nil
}

//FilterEmployeeDocuments :""
func (s *Service) FilterEmployeeDocuments(ctx *models.Context, filter *models.FilterEmployeeDocuments, pagination *models.Pagination) ([]models.RefEmployeeDocuments, error) {
	defer ctx.Session.EndSession(ctx.CTX)
	return s.Daos.FilterEmployeeDocuments(ctx, filter, pagination)

}
func (s *Service) EmployeeDocumentsList(ctx *models.Context, filter *models.FilterEmployeeDocumentslist) (*models.EmployeeDocumentsList, error) {
	defer ctx.Session.EndSession(ctx.CTX)
	return s.Daos.EmployeeDocumentsList(ctx, filter)

}
func (s *Service) UpdateEmployeeDocumentsWithUpsert(ctx *models.Context, employeeDocuments *models.EmployeeDocuments) error {
	if err := ctx.Session.StartTransaction(); err != nil {
		return err
	}
	defer ctx.Session.EndSession(ctx.CTX)
	if employeeDocuments.UniqueID == "" {
		employeeDocuments.UniqueID = s.Daos.GetUniqueID(ctx, constants.COLLECTIONEMPLOYEEDOCUMENTS)
	}
	employeeDocuments.Status = constants.EMPLOYEEDOCUMENTSSTATUSACTIVE
	t := time.Now()
	created := models.Created{}
	created.On = &t
	created.By = constants.SYSTEM
	log.Println("b4 EmployeeDocuments.created")
	employeeDocuments.Created = created
	if err := mongo.WithSession(ctx.CTX, ctx.Session, func(sc mongo.SessionContext) error {

		err := s.Daos.UpdateEmployeeDocumentsWithUpsert(ctx, employeeDocuments)
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
func (s *Service) RemoveEmployeeDocuments(ctx *models.Context, employeeDocuments *models.EmployeeDocuments) error {
	if err := ctx.Session.StartTransaction(); err != nil {
		return err
	}
	defer ctx.Session.EndSession(ctx.CTX)
	if employeeDocuments.UniqueID == "" {
		employeeDocuments.UniqueID = s.Daos.GetUniqueID(ctx, constants.COLLECTIONEMPLOYEEDOCUMENTS)
	}
	employeeDocuments.Status = constants.EMPLOYEEDOCUMENTSSTATUSACTIVE
	t := time.Now()
	created := models.Created{}
	created.On = &t
	created.By = constants.SYSTEM
	log.Println("b4 EmployeeDocuments.created")
	employeeDocuments.Created = created
	if err := mongo.WithSession(ctx.CTX, ctx.Session, func(sc mongo.SessionContext) error {

		err := s.Daos.RemoveEmployeeDocuments(ctx, employeeDocuments)
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
