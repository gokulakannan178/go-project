package models

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type ProductConfig struct {
	ID           primitive.ObjectID `json:"id" form:"id," bson:"_id,omitempty"`
	Name         string             `json:"name" bson:"name,omitempty"`
	Logo         string             `json:"logo" bson:"logo,omitempty"`
	WaterMark    string             `json:"waterMark" bson:"waterMark,omitempty"`
	LogoWithName string             `json:"logoWithName" bson:"logoWithName,omitempty"`
	Email        struct {
		ContactUs string `json:"contactUs" bson:"contactUs,omitempty"`
		SendEmail string `json:"sendEmail" bson:"sendEmail,omitempty"`
	} `json:"email" bson:"email,omitempty"`
	Mobile                   string   `json:"mobile" bson:"mobile,omitempty"`
	Phone                    string   `json:"phone" bson:"phone,omitempty"`
	Address                  string   `json:"address" bson:"address,omitempty"`
	Copyrights               string   `json:"copyrights" bson:"copyrights,omitempty"`
	PoweredBy                string   `json:"poweredBy" bson:"poweredBy,omitempty"`
	Rights                   string   `json:"rights" bson:"rights,omitempty"`
	Status                   string   `json:"status" bson:"status,omitempty"`
	UniqueID                 string   `json:"uniqueId,omitempty" bson:"uniqueId,omitempty"`
	Created                  *Created `json:"created" bson:"created,omitempty"`
	IsDefault                bool     `json:"isdefault" bson:"isdefault,omitempty"`
	IsSingleProject          bool     `json:"issingleProject" bson:"issingleProject,omitempty"`
	IsSms                    bool     `json:"issms" bson:"issms,omitempty"`
	APIURL                   string   `json:"-" bson:"apiUrl,omitempty"`
	UIURL                    string   `json:"-" bson:"uiUrl,omitempty"`
	URL                      string   `json:"-" bson:"url,omitempty"`
	IsUser                   string   `json:"isUser" bson:"isUser,omitempty"`
	IsDealer                 string   `json:"isDealer" bson:"isDealer,omitempty"`
	IsFarmer                 string   `json:"isFarmer" bson:"isFarmer,omitempty"`
	ValidateUserregistration bool     `json:"validateUserRegistration" bson:"validateUserRegistration,omitempty"`
	Apptoken                 bool     `json:"apptoken" bson:"apptoken,omitempty"`
	WeatherAlert             bool     `json:"weatherAlert" bson:"weatherAlert,omitempty"`
	FarmerUpload             string   `json:"farmerUpload" bson:"farmerUpload,omitempty"`
	FarmerCasteUpload        string   `json:"farmerCasteUpload" bson:"farmerCasteUpload,omitempty"`
	FarmerSoilUpload         string   `json:"farmerSoilUpload" bson:"farmerSoilUpload,omitempty"`
	FarmerLandUpload         string   `json:"farmerLandUpload" bson:"farmerLandUpload,omitempty"`
	DetailedContentCreation  string   `json:"detailedContentCreation" bson:"detailedContentCreation,omitempty"`
	FarmerCropUpload         string   `json:"farmerCropUpload" bson:"farmerCropUpload,omitempty"`
	State                    struct {
		IsSingle  string             `json:"isSingle" bson:"isSingle,omitempty"`
		StateID   primitive.ObjectID `json:"stateId" bson:"stateId,omitempty"`
		StateName string             `json:"stateName" bson:"stateName,omitempty"`
	} `json:"state" bson:"state,omitempty"`
	District struct {
		IsSingle     string             `json:"isSingle" bson:"isSingle,omitempty"`
		DistrictID   primitive.ObjectID `json:"districtId" bson:"districtId,omitempty"`
		DistrictName string             `json:"districtName" bson:"districtName,omitempty"`
	} `json:"district" bson:"district,omitempty"`
	Orgnisation struct {
		IsSingle        string             `json:"isSingle" bson:"isSingle,omitempty"`
		OrgnisationID   primitive.ObjectID `json:"orgnisationId" bson:"orgnisationId,omitempty"`
		OrgnisationName string             `json:"orgnisationName" bson:"orgnisationName,omitempty"`
	} `json:"orgnisation" bson:"orgnisation,omitempty"`
	Project struct {
		IsSingle    string             `json:"isSingle" bson:"isSingle,omitempty"`
		ProjectID   primitive.ObjectID `json:"projectId" bson:"projectId,omitempty"`
		ProjectName string             `json:"projectName" bson:"projectName,omitempty"`
	} `json:"project" bson:"project,omitempty"`
	KnowledgeDomain struct {
		IsSingle            string             `json:"isSingle" bson:"isSingle,omitempty"`
		KnowledgeDomainID   primitive.ObjectID `json:"knowledgedomainId" bson:"knowledgedomainId,omitempty"`
		KnowledgeDomainName string             `json:"knowledgedomainName" bson:"knowledgedomainName,omitempty"`
	} `json:"knowledgedomain" bson:"knowledgedomain,omitempty"`
	IsLanguageTranslation string `json:"isLanguageTranslation" bson:"isLanguageTranslation,omitempty"`
	ExpiryMonth           int    `json:"expiryMonth" bson:"expiryMonth,omitempty"`
}
type ProductConfigFilter struct {
	IsDefault []bool   `json:"isdefault,omitempty"`
	Status    []string `json:"status" form:"status" bson:"status,omitempty"`
	SortBy    string   `json:"sortBy"`
	SortOrder int      `json:"sortOrder"`
	Searchbox struct {
		Name string `json:"name" bson:"name"`
	} `json:"searchbox" bson:"searchbox"`
}
type RefProductConfig struct {
	ProductConfig `bson:",inline"`
	Ref           struct {
	} `json:"ref,omitempty" bson:"ref,omitempty"`
}
