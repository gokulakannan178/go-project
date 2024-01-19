package services

import (
	"errors"
	"hrms-services/constants"
	"hrms-services/models"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
)

//SaveUser :""
func (s *Service) SaveProject(ctx *models.Context, project *models.Project) error {
	log.Println("transaction start")
	//Start Transaction
	if err := ctx.Session.StartTransaction(); err != nil {
		return err
	}
	defer ctx.Session.EndSession(ctx.CTX)

	project.UniqueID = s.Daos.GetUniqueID(ctx, constants.COLLECTIONPROJECT)
	project.Status = constants.PROJECTSTATUSACTIVE
	//user.Password = "#nature32" //Default Password
	t := time.Now()
	created := models.Created{}
	created.On = &t
	created.By = constants.SYSTEM
	log.Println("b4 user.created")
	project.Created = created
	log.Println("b4 user.created")
	if err := mongo.WithSession(ctx.CTX, ctx.Session, func(sc mongo.SessionContext) error {
		dberr := s.Daos.SaveProject(ctx, project)
		if dberr != nil {

			log.Println("err in abort out")
			return errors.New("Transaction Aborted - " + dberr.Error())
		}
		if len(project.TeamMember) > 0 {
			var members []models.ProjectMember
			for _, v := range project.TeamMember {
				var member models.ProjectMember
				member.Status = constants.PROJECTSTATUSACTIVE
				member.UniqueID = s.Daos.GetUniqueID(ctx, constants.COLLECTIONPROJECTMEMBER)
				member.ProjectID = project.UniqueID
				member.UserName = v
				members = append(members, member)
			}
			dberr := s.Daos.SaveProjectMembers(ctx, members)
			if dberr != nil {

				log.Println("err in abort out")
				return errors.New("Transaction Aborted - " + dberr.Error())
			}

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

//SaveTeamMember :""
func (s *Service) SaveProjectTeamMember(ctx *models.Context, project *models.ProjectMember) error {
	log.Println("transaction start")
	//Start Transaction
	if err := ctx.Session.StartTransaction(); err != nil {
		return err
	}
	defer ctx.Session.EndSession(ctx.CTX)

	project.UniqueID = s.Daos.GetUniqueID(ctx, constants.COLLECTIONPROJECTMEMBER)
	project.Status = constants.PROJECTTEAMMEMBERSTATUSACTIVE
	//user.Password = "#nature32" //Default Password
	t := time.Now()
	created := models.Created{}
	created.On = &t
	created.By = constants.SYSTEM
	log.Println("b4 user.created")
	project.Created = created
	log.Println("b4 user.created")
	if err := mongo.WithSession(ctx.CTX, ctx.Session, func(sc mongo.SessionContext) error {
		dberr := s.Daos.SaveProjectTeamMember(ctx, project)
		if dberr != nil {

			log.Println("err in abort out")
			return errors.New("Transaction Aborted - " + dberr.Error())
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

//DisableProjectTeamMember : ""
func (s *Service) DisableProjectTeamMember(ctx *models.Context, UniqueID string) error {
	if err := ctx.Session.StartTransaction(); err != nil {
		return err
	}
	defer ctx.Session.EndSession(ctx.CTX)
	if err := mongo.WithSession(ctx.CTX, ctx.Session, func(sc mongo.SessionContext) error {

		err := s.Daos.DisableProjectTeamMember(ctx, UniqueID)
		if err != nil {

			return errors.New("Transaction Aborted - " + err.Error())
		}
		return nil

	}); err != nil {
		if err = ctx.Session.AbortTransaction(ctx.CTX); err != nil {
			return errors.New("Transaction Aborted with error" + err.Error())
		}
		return err
	}
	return nil
}

//GetSingleProject :""
func (s *Service) GetSingleProject(ctx *models.Context, UniqueID string) (*models.Project, error) {
	user, err := s.Daos.GetSingleProject(ctx, UniqueID)
	if err != nil {
		return nil, err
	}
	return user, nil
}

//UpdateProject : ""
func (s *Service) UpdateProject(ctx *models.Context, user *models.Project) error {
	if err := ctx.Session.StartTransaction(); err != nil {
		return err
	}
	defer ctx.Session.EndSession(ctx.CTX)
	if err := mongo.WithSession(ctx.CTX, ctx.Session, func(sc mongo.SessionContext) error {

		err := s.Daos.UpdateProject(ctx, user)
		if err != nil {
			log.Println("err in abort out")
			return errors.New("Transaction Aborted - " + err.Error())
		}
		return nil

	}); err != nil {
		if err = ctx.Session.AbortTransaction(ctx.CTX); err != nil {
			return errors.New("Transaction Aborted with error" + err.Error())

			//return errors.New("Transaction Aborted - " + err.Error())
		}
		return err
	}
	return nil
}

//EnableProject : ""
func (s *Service) EnableProject(ctx *models.Context, UniqueID string) error {
	if err := ctx.Session.StartTransaction(); err != nil {
		return err
	}
	defer ctx.Session.EndSession(ctx.CTX)
	if err := mongo.WithSession(ctx.CTX, ctx.Session, func(sc mongo.SessionContext) error {

		err := s.Daos.EnableProject(ctx, UniqueID)
		if err != nil {

			return errors.New("Transaction Aborted - " + err.Error())
		}
		return nil

	}); err != nil {
		if err = ctx.Session.AbortTransaction(ctx.CTX); err != nil {
			return errors.New("Transaction Aborted with error" + err.Error())
		}

		return err
	}
	return nil
}

//DisableProject : ""
func (s *Service) DisableProject(ctx *models.Context, UniqueID string) error {
	if err := ctx.Session.StartTransaction(); err != nil {
		return err
	}
	defer ctx.Session.EndSession(ctx.CTX)
	if err := mongo.WithSession(ctx.CTX, ctx.Session, func(sc mongo.SessionContext) error {

		err := s.Daos.DisableProject(ctx, UniqueID)
		if err != nil {

			return errors.New("Transaction Aborted - " + err.Error())
		}
		return nil

	}); err != nil {
		if err = ctx.Session.AbortTransaction(ctx.CTX); err != nil {
			return errors.New("Transaction Aborted with error" + err.Error())
		}
		return err
	}
	return nil
}

//DeleteProject : ""
func (s *Service) DeleteProject(ctx *models.Context, UniqueID string) error {
	if err := ctx.Session.StartTransaction(); err != nil {
		return err
	}
	defer ctx.Session.EndSession(ctx.CTX)
	if err := mongo.WithSession(ctx.CTX, ctx.Session, func(sc mongo.SessionContext) error {

		err := s.Daos.DeleteProject(ctx, UniqueID)
		if err != nil {

			return errors.New("Transaction Aborted - " + err.Error())
		}
		return nil

	}); err != nil {
		if err = ctx.Session.AbortTransaction(ctx.CTX); err != nil {
			return errors.New("Transaction Aborted with error" + err.Error())
		}
		return err
	}
	return nil
}

//ProjectFilter:""
func (s *Service) ProjectFilter(ctx *models.Context, projectfilter *models.ProjectFilter, pagination *models.Pagination) (user []models.Project, err error) {
	return s.Daos.ProjectFilter(ctx, projectfilter, pagination)

}
