package services

import (
	"fmt"
	"nicessm-api-service/models"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

// GetContentDisseminationUserAndFarmer : ""
func (s *Service) GetContentDisseminationUserAndFarmer(ctx *models.Context, contentID string) (*models.ContentDissiminateUserAndFarmer, error) {
	id, err := primitive.ObjectIDFromHex(contentID)
	fmt.Println("id===>", id)
	fmt.Println("contentId===>", contentID)
	if err != nil {
		return nil, err
	}
	ContentDataAccess, err := s.GetContentDataAccess(ctx, contentID)
	farmers, err := s.GetContentDisseminationFarmer(ctx, ContentDataAccess)
	if err != nil {
		return nil, err
	}
	users, err := s.GetContentDisseminationUser(ctx, ContentDataAccess)
	if err != nil {
		return nil, err
	}
	contentDissiminateUserAndFarmer := new(models.ContentDissiminateUserAndFarmer)
	contentDissiminateUserAndFarmer.Users = users
	contentDissiminateUserAndFarmer.Farmers = farmers
	contentDissiminateUserAndFarmer.FarmersCount = len(farmers)
	contentDissiminateUserAndFarmer.UsersCount = len(users)

	return contentDissiminateUserAndFarmer, nil
}
func (s *Service) GetContentDataAccess(ctx *models.Context, contentID string) (*models.ContentDataAccess, error) {

	refContent, err := s.Daos.GetSingleContent(ctx, contentID)
	if err != nil {
		return nil, err
	}

	contentDataAccess := new(models.ContentDataAccess)
	if !refContent.Organisation.IsZero() {
		contentDataAccess.Organisation = append(contentDataAccess.Organisation, refContent.Organisation)
	}
	if !refContent.Project.IsZero() {
		contentDataAccess.Project = append(contentDataAccess.Project, refContent.Project)
	}
	if !refContent.IndexingData.State.IsZero() {
		contentDataAccess.State = append(contentDataAccess.State, refContent.IndexingData.State)
	}
	if !refContent.IndexingData.District.IsZero() {
		contentDataAccess.District = append(contentDataAccess.District, refContent.IndexingData.District)
	}
	if !refContent.IndexingData.Block.IsZero() {
		contentDataAccess.Block = append(contentDataAccess.Block, refContent.IndexingData.Block)
	}
	if !refContent.IndexingData.Gram_panchayat.IsZero() {
		contentDataAccess.GramPanchayat = append(contentDataAccess.GramPanchayat, refContent.IndexingData.Gram_panchayat)
	}
	if !refContent.IndexingData.Village.IsZero() {
		contentDataAccess.Village = append(contentDataAccess.Village, refContent.IndexingData.Village)
	}
	return contentDataAccess, nil
}
func (s *Service) GetContentDisseminationUserAndFarmerCount(ctx *models.Context, content *models.ContentDataAccess) (*models.ContentDissiminateUserAndFarmer, error) {

	farmers, err := s.GetContentDisseminationFarmer(ctx, content)
	if err != nil {
		return nil, err
	}
	users, err := s.GetContentDisseminationUser(ctx, content)
	if err != nil {
		return nil, err
	}
	contentDissiminateUserAndFarmer := new(models.ContentDissiminateUserAndFarmer)
	contentDissiminateUserAndFarmer.Users = nil
	contentDissiminateUserAndFarmer.Farmers = nil
	contentDissiminateUserAndFarmer.FarmersCount = len(farmers)
	contentDissiminateUserAndFarmer.UsersCount = len(users)

	return contentDissiminateUserAndFarmer, nil
}
