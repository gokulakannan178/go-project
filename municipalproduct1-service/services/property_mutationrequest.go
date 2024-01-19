package services

import (
	"errors"
	"fmt"
	"log"
	"math"
	"municipalproduct1-service/constants"
	"municipalproduct1-service/models"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
)

// BasicUpdatePropertyMutationRequest : ""
func (s *Service) BasicUpdatePropertyMutationRequest(ctx *models.Context, request *models.PropertyMutationRequest) error {
	if err := ctx.Session.StartTransaction(); err != nil {
		return err
	}
	defer ctx.Session.EndSession(ctx.CTX)
	if err := mongo.WithSession(ctx.CTX, ctx.Session, func(sc mongo.SessionContext) error {
		oldPropertyData, err := s.Daos.GetSingleProperty(ctx, request.PropertyID)
		if err != nil {
			return errors.New("error in geting old property" + err.Error())
		}
		if oldPropertyData == nil {
			return errors.New("property Not Found")
		}
		t := time.Now()
		request.UniqueID = s.Daos.GetUniqueID(ctx, constants.COLLECTIONPROPERTYMUTATIONREQUEST)
		request.Property = oldPropertyData
		request.Status = constants.PROPERTYMUTATIONREQUESTSTATUSINIT
		request.Requester.On = &t
		request.RequestedDate = &t
		request.Requester.By = request.UserName
		request.Requester.ByType = request.UserType
		request.Requester.Scenario = "Property Requested for Mutation"
		err = s.Daos.SavePropertyMutationRequest(ctx, request)
		if err != nil {
			return errors.New("Error in updating log" + err.Error())
		}
		if err = ctx.Session.CommitTransaction(sc); err != nil {
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

// GetSinglePropertyMutationRequest : ""
func (s *Service) GetSinglePropertyMutationRequest(ctx *models.Context, UniqueID string) (*models.RefPropertyMutationRequest, error) {
	pmr, err := s.Daos.GetSinglePropertyMutationRequest(ctx, UniqueID)
	if err != nil {
		return nil, err
	}
	return pmr, nil
}

// AcceptPropertyMutationRequestUpdate : ""
func (s *Service) AcceptPropertyMutationRequestUpdate(ctx *models.Context, accept *models.AcceptPropertyMutationRequestUpdate) error {
	if err := ctx.Session.StartTransaction(); err != nil {
		return err
	}
	defer ctx.Session.EndSession(ctx.CTX)
	if err := mongo.WithSession(ctx.CTX, ctx.Session, func(sc mongo.SessionContext) error {

		err := s.Daos.AcceptPropertyMutationRequestUpdate(ctx, accept)
		if err != nil {
			return err
		}
		var resPropertyMutation *models.RefPropertyMutationRequest
		resPropertyMutation, err = s.GetSinglePropertyMutationRequest(ctx, accept.UniqueID)
		if err != nil {
			return errors.New("error in getting the property mutation request" + err.Error())
		}
		err = s.Daos.UpdatePropertyEndDate(ctx, resPropertyMutation.PropertyID, resPropertyMutation.PropertyMutatedDate)
		// err = s.Daos.UpdatePropertyEndDate(ctx, resPropertyMutation.PropertyID)
		if err != nil {
			return errors.New("error in updating property end date" + err.Error())
		}
		if err = ctx.Session.CommitTransaction(sc); err != nil {
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

// RejectPropertyMutationRequestUpdate : ""
func (s *Service) RejectPropertyMutationRequestUpdate(ctx *models.Context, req *models.RejectPropertyMutationRequestUpdate) error {
	if err := ctx.Session.StartTransaction(); err != nil {
		return err
	}
	defer ctx.Session.EndSession(ctx.CTX)
	if err := mongo.WithSession(ctx.CTX, ctx.Session, func(sc mongo.SessionContext) error {

		err := s.Daos.RejectPropertyMutationRequestUpdate(ctx, req)
		if err != nil {
			return nil
		}
		if err = ctx.Session.CommitTransaction(sc); err != nil {
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

// FilterPropertyMutationRequest : ""
func (s *Service) FilterPropertyMutationRequest(ctx *models.Context, filter *models.PropertyMutationRequestFilter, pagination *models.Pagination) ([]models.RefPropertyMutationRequest, error) {
	return s.Daos.FilterPropertyMutationRequest(ctx, filter, pagination)
}

// SaveMutatedProperty : ""
func (s *Service) SaveMutatedProperty(ctx *models.Context, mutatedProperty *models.MutatedProperty) error {

	t := time.Now()
	mutatedProperty.UniqueID = s.Daos.GetUniqueID(ctx, constants.COLLECTIONMUTATEDPROPERTY)
	mutatedProperty.Status = constants.MUTATEDPROPERTYSTATUSACTIVE
	Created := new(models.CreatedV2)
	Created.On = &t
	Created.By = constants.SYSTEM
	mutatedProperty.Created = Created

	//
	err := s.SaveProperty(ctx, mutatedProperty.Property, "")
	if err != nil {
		return errors.New("error in saving the property - " + err.Error())
	}

	mutatedProperty.ChildID = mutatedProperty.Property.UniqueID

	res := new(models.RefRemainingOfMutatedProperty)
	res, err = s.RemainingAreaOfMutatedProperty(ctx, mutatedProperty.ParentID)
	if err != nil {
		return errors.New("error in remaining area of mutated property - " + err.Error())
	}

	mutatedProperty.RemainingAreaOfPlot = res.RemainingAreaOfPlot
	mutatedProperty.RemainingBuiltUpArea = res.RemainingBuiltUpArea
	mutatedProperty.PercentAreaOfPlotFilled = res.PercentAreaOfPlotFilled
	mutatedProperty.PercentBuiltUpAreaFilled = res.PercentBuiltUpAreaFilled
	mutatedProperty.TotalAreaOfPlot = res.TotalAreaOfPlot
	mutatedProperty.TotalBuiltUpArea = res.TotalBuiltUpArea
	err = s.Daos.SaveMutatedProperty(ctx, mutatedProperty)
	if err != nil {
		return errors.New("Error in updating log" + err.Error())
	}

	return nil

}

// GetSingleMutatedProperty : ""
func (s *Service) GetSingleMutatedProperty(ctx *models.Context, UniqueID string) (*models.RefMutatedProperty, error) {
	mutatedProperty, err := s.Daos.GetSingleMutatedProperty(ctx, UniqueID)
	if err != nil {
		return nil, err
	}
	return mutatedProperty, nil
}

// FilterMutatedProperty : ""
func (s *Service) FilterMutatedProperty(ctx *models.Context, filter *models.MutatedPropertyFilter, pagination *models.Pagination) ([]models.RefMutatedProperty, error) {
	return s.Daos.FilterMutatedProperty(ctx, filter, pagination)
}

// RemainingAreaOfMutatedProperty : ""
func (s *Service) RemainingAreaOfMutatedProperty(ctx *models.Context, PropertyID string) (*models.RefRemainingOfMutatedProperty, error) {

	resProperty := new(models.RefProperty)
	resProperty, err := s.Daos.GetSingleProperty(ctx, PropertyID)
	if err != nil {
		return nil, errors.New("error in geting Property - " + err.Error())
	}
	if resProperty == nil {
		return nil, errors.New("property is nil")
	}
	if resProperty.DOA == nil {
		return nil, errors.New("psroperty is DOA is not valid")
	}
	filter := new(models.MutatedPropertyFilter)
	filter.ParentID = append(filter.ParentID, PropertyID)

	resMutatedProperty, err := s.FilterMutatedProperty(ctx, filter, nil)
	var totalPlotArea, totalBuiltUpArea float64
	fmt.Println("resMutatedProperty =======>", resMutatedProperty)
	if len(resMutatedProperty) > 0 {

		for _, v := range resMutatedProperty {

			totalPlotArea = totalPlotArea + v.Property.AreaOfPlot

			totalBuiltUpArea = totalBuiltUpArea + v.Property.BuiltUpArea

			fmt.Println("totalPlotArea =========>", totalPlotArea)
			fmt.Println("totalBuiltUpArea =========>", totalBuiltUpArea)
		}
	}

	fmt.Println("Parent Property AreaOfPlot =======>", resProperty.Property.AreaOfPlot)
	fmt.Println("Parent Property BuildUpArea =======>", resProperty.Property.BuiltUpArea)
	fmt.Println("Total AreaOfPlot of Child Properties =======>", totalPlotArea)
	fmt.Println("Total BuildUpArea of Child Properties =======>", totalBuiltUpArea)

	res := new(models.RefRemainingOfMutatedProperty)
	res.RemainingAreaOfPlot = resProperty.Property.AreaOfPlot - totalPlotArea
	res.RemainingBuiltUpArea = resProperty.Property.BuiltUpArea - totalBuiltUpArea
	res.PercentAreaOfPlotFilled = (totalPlotArea / resProperty.Property.AreaOfPlot) * 100
	if math.IsNaN(res.PercentAreaOfPlotFilled) {
		res.PercentAreaOfPlotFilled = 0
	}
	res.PercentBuiltUpAreaFilled = (totalBuiltUpArea / resProperty.Property.BuiltUpArea) * 100
	if math.IsNaN(res.PercentBuiltUpAreaFilled) {
		res.PercentBuiltUpAreaFilled = 0
	}
	res.TotalAreaOfPlot = resProperty.Property.AreaOfPlot
	res.TotalBuiltUpArea = resProperty.Property.BuiltUpArea
	fmt.Println("Percentage of Filled Area Of Plot =======>", res.PercentAreaOfPlotFilled)
	fmt.Println("Percentage of Filled BuiltUp Area =======>", res.PercentBuiltUpAreaFilled)
	fmt.Println("Total BuildUpArea of Child Properties =======>", totalBuiltUpArea)
	s.Shared.BsonToJSONPrintTag("RemainingAreaOfMutatedProperty", res)
	return res, nil
}

// UpdatePropertyMutationRequestPropertyID : ""
func (s *Service) UpdatePropertyMutationRequestPropertyID(ctx *models.Context, uniqueIds *models.UpdatePropertyUniqueID) error {
	if err := ctx.Session.StartTransaction(); err != nil {
		return err
	}
	defer ctx.Session.EndSession(ctx.CTX)
	if err := mongo.WithSession(ctx.CTX, ctx.Session, func(sc mongo.SessionContext) error {
		for _, v := range uniqueIds.UniqueIDs {
			resProperty, err := s.GetSingleProperty(ctx, v)
			if err != nil {
				return errors.New("Not able to get property - " + err.Error())
			}

			uniqueIds.UniqueID = resProperty.OldUniqueID
			uniqueIds.OldUniqueID = resProperty.OldUniqueID
			uniqueIds.NewUniqueID = resProperty.NewUniqueID
			err = s.Daos.UpdatePropertyMutationRequestPropertyID(ctx, uniqueIds)
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
