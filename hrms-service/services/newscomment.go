package services

import (
	"errors"
	"hrms-services/constants"
	"hrms-services/models"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
)

// SaveNewsComment : ""
func (s *Service) SaveNewsComment(ctx *models.Context, newscomment *models.NewsComment) error {
	log.Println("transaction start")
	// Start Transaction
	if err := ctx.Session.StartTransaction(); err != nil {
		return err
	}
	defer ctx.Session.EndSession(ctx.CTX)
	newscomment.UniqueID = s.Daos.GetUniqueID(ctx, constants.COLLECTIONNEWSCOMMENT)
	newscomment.Status = constants.NEWSCOMMENTSTATUSACTIVE
	t := time.Now()
	newscomment.Date = &t
	created := models.Created{}
	created.On = &t
	created.By = constants.SYSTEM
	log.Println("b4 News Comment.created")
	newscomment.Created = &created
	log.Println("b4 News Comment.created")
	if err := mongo.WithSession(ctx.CTX, ctx.Session, func(sc mongo.SessionContext) error {
		dberr := s.Daos.SaveNewsComment(ctx, newscomment)
		if dberr != nil {
			return dberr
		}

		news, err := s.Daos.GetSingleNewsWithoutRef(ctx, newscomment.NewsId)
		if err != nil {
			return err
		}
		news.NewsCommentCount = news.NewsCommentCount + 1
		dberr = s.Daos.UpdateNews(ctx, news)
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

// GetSingleNewsComment : ""
func (s *Service) GetSingleNewsComment(ctx *models.Context, UniqueID string) (*models.RefNewsComment, error) {
	news, err := s.Daos.GetSingleNewsComment(ctx, UniqueID)
	if err != nil {
		return nil, err
	}
	return news, nil
}

//UpdateNewsComment : ""
func (s *Service) UpdateNewsComment(ctx *models.Context, newscomment *models.NewsComment) error {
	if err := ctx.Session.StartTransaction(); err != nil {
		return err
	}
	defer ctx.Session.EndSession(ctx.CTX)
	if err := mongo.WithSession(ctx.CTX, ctx.Session, func(sc mongo.SessionContext) error {

		err := s.Daos.UpdateNewsComment(ctx, newscomment)
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

// EnableNewsComment : ""
func (s *Service) EnableNewsComment(ctx *models.Context, uniqueID string) error {

	if err := ctx.Session.StartTransaction(); err != nil {
		return err
	}
	defer ctx.Session.EndSession(ctx.CTX)
	if err := mongo.WithSession(ctx.CTX, ctx.Session, func(sc mongo.SessionContext) error {
		dberr := s.Daos.EnableNewsComment(ctx, uniqueID)
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

// DisableNewsComment : ""
func (s *Service) DisableNewsComment(ctx *models.Context, uniqueID string) error {

	if err := ctx.Session.StartTransaction(); err != nil {
		return err
	}
	defer ctx.Session.EndSession(ctx.CTX)
	if err := mongo.WithSession(ctx.CTX, ctx.Session, func(sc mongo.SessionContext) error {
		debrr := s.Daos.DisableNewsComment(ctx, uniqueID)
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

//DeleteNewsComment : ""
func (s *Service) DeleteNewsComment(ctx *models.Context, UniqueID string) error {
	if err := ctx.Session.StartTransaction(); err != nil {
		return err
	}
	defer ctx.Session.EndSession(ctx.CTX)
	if err := mongo.WithSession(ctx.CTX, ctx.Session, func(sc mongo.SessionContext) error {

		err := s.Daos.DeleteNewsComment(ctx, UniqueID)
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

// FilterNewsComment : ""
func (s *Service) FilterNewsComment(ctx *models.Context, newscomment *models.FilterNewsComment, pagination *models.Pagination) (newscomments []models.RefNewsComment, err error) {
	return s.Daos.FilterNewsComment(ctx, newscomment, pagination)
}
