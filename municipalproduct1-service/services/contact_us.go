package services

import (
	"errors"
	"log"
	"municipalproduct1-service/constants"
	"municipalproduct1-service/models"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
)

// SendContactUs : ""
func (s *Service) SendContactUs(ctx *models.Context, cu *models.ContactUs) error {
	resProduct, err := s.Daos.GetSingleDefaultProductConfiguration(ctx)
	if err != nil {
		return err
	}
	if resProduct == nil {
		return errors.New("product configuration is nil")

	}
	err = s.SendEmail("Reg:Contacted by - "+cu.Name, []string{resProduct.Email.ContactUs}, cu.ConvertEmailMsg())
	if err != nil {
		log.Println("email not sent - ", err.Error())
	}
	err = s.SendEmail("Reg:Thanks for Contacting us - "+cu.Name, []string{cu.Email}, cu.AutoResponse())
	if err != nil {
		log.Println("autoresponse email not sent - ", err.Error())
	}
	return nil
}

//SaveContact :""
func (s *Service) SaveContact(ctx *models.Context, contact *models.ContactUs) error {
	log.Println("transaction start")
	//Start Transaction
	if err := ctx.Session.StartTransaction(); err != nil {
		return err
	}
	defer ctx.Session.EndSession(ctx.CTX)

	contact.UniqueID = s.Daos.GetUniqueID(ctx, constants.COLLECTIONCONTACTUS)
	contact.Status = constants.CONTACTUSSTATUSACTIVE
	t := time.Now()
	// created := models.Created{}
	created := new(models.CreatedV2)
	created.On = &t
	created.By = constants.SYSTEM
	log.Println("b4 user.created")
	contact.Created = created

	log.Println("b4 user.created")
	err := s.SendContactUs(ctx, contact)
	if err != nil {
		log.Println("email not sent - ", err.Error())
	}
	if err := mongo.WithSession(ctx.CTX, ctx.Session, func(sc mongo.SessionContext) error {
		dberr := s.Daos.SaveContact(ctx, contact)
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

//UpdateUser : ""
func (s *Service) UpdateContact(ctx *models.Context, contact *models.ContactUs) error {
	if err := ctx.Session.StartTransaction(); err != nil {
		return err
	}
	defer ctx.Session.EndSession(ctx.CTX)
	if err := mongo.WithSession(ctx.CTX, ctx.Session, func(sc mongo.SessionContext) error {

		err := s.Daos.UpdateContact(ctx, contact)
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

//EnableUser : ""
func (s *Service) EnableContact(ctx *models.Context, UniqueID string) error {
	if err := ctx.Session.StartTransaction(); err != nil {
		return err
	}
	defer ctx.Session.EndSession(ctx.CTX)
	if err := mongo.WithSession(ctx.CTX, ctx.Session, func(sc mongo.SessionContext) error {

		err := s.Daos.EnableContact(ctx, UniqueID)
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

//DisableUser : ""
func (s *Service) DisableContact(ctx *models.Context, UniqueID string) error {
	if err := ctx.Session.StartTransaction(); err != nil {
		return err
	}
	defer ctx.Session.EndSession(ctx.CTX)
	if err := mongo.WithSession(ctx.CTX, ctx.Session, func(sc mongo.SessionContext) error {

		err := s.Daos.DisableContact(ctx, UniqueID)
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

//DeleteUser : ""
func (s *Service) DeleteContact(ctx *models.Context, UniqueID string) error {
	if err := ctx.Session.StartTransaction(); err != nil {
		return err
	}
	defer ctx.Session.EndSession(ctx.CTX)
	if err := mongo.WithSession(ctx.CTX, ctx.Session, func(sc mongo.SessionContext) error {

		err := s.Daos.DeleteContact(ctx, UniqueID)
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

//GetSingleUser :""
func (s *Service) GetSingleContact(ctx *models.Context, UniqueID string) (*models.RefContactUs, error) {
	contact, err := s.Daos.GetSingleContact(ctx, UniqueID)
	if err != nil {
		return nil, err
	}
	return contact, nil
}

//FilterUser :""
func (s *Service) FilterContactUs(ctx *models.Context, contactfilter *models.FilterContactUs, pagination *models.Pagination) ([]models.RefContactUs, error) {
	return s.Daos.FilterContactUs(ctx, contactfilter, pagination)

}
