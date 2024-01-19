package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type BasicFarmer struct {
	ID   primitive.ObjectID `json:"id" form:"id," bson:"_id,omitempty"`
	Name string             `json:"name" bson:"name,omitempty"`
}

//Farmer : "Holds single Farmer data"
type Farmer struct {
	BasicFarmer               `bson:",inline"`
	Gender                    string              `json:"gender" bson:"gender,omitempty"`
	AlternateNumber           string              `json:"alternateNumber" bson:"alternateNumber,omitempty"`
	AadhaarNumber             string              `json:"aadhaarNumber" bson:"aadhaarNumber,omitempty"`
	Education                 string              `json:"education" bson:"education,omitempty"`
	FatherName                string              `json:"fatherName" bson:"fatherName,omitempty"`
	CurrentLiveStocks         []CurrentLiveStocks `json:"currentLiveStocks" bson:"currentLiveStocks,omitempty"`
	CurrentCrops              []CurrentCrops      `json:"currentCrops" bson:"currentCrops,omitempty"`
	CultivationPractice       string              `json:"cultivationPractice" bson:"cultivationPractice,omitempty"`
	DateOfBirth               *time.Time          `json:"dateOfBirth" bson:"dateOfBirth,omitempty"`
	Status                    string              `json:"status" bson:"status,omitempty"`
	District                  primitive.ObjectID  `json:"district"  bson:"district,omitempty"`
	GramPanchayat             primitive.ObjectID  `json:"gramPanchayat"  bson:"gramPanchayat,omitempty"`
	IsMemberInvolvedInCbo     bool                `json:"isMemberInvolvedInCbo" bson:"isMemberInvolvedInCbo,omitempty"`
	ActiveStatus              bool                `json:"activeStatus" bson:"activeStatus,omitempty"`
	IsSMS                     bool                `json:"isSMS" bson:"isSMS,omitempty"`
	FeminineMobile            bool                `json:"feminineMobile" bson:"feminineMobile,omitempty"`
	CreditAvailed             bool                `json:"creditAvailed" bson:"creditAvailed,omitempty"`
	HasKitchenGarden          bool                `json:"hasKitchenGarden" bson:"hasKitchenGarden,omitempty"`
	LikeToReceiveSMS          bool                `json:"likeToReceiveSMS" bson:"likeToReceiveSMS,omitempty"`
	LikeToReceiveVoicecall    bool                `json:"likeToReceiveVoicecall" bson:"likeToReceiveVoicecall,omitempty"`
	HasCredits                bool                `json:"hasCredits" bson:"hasCredits,omitempty"`
	MemberOfMGNREGNA          bool                `json:"memberOfMGNREGNA" bson:"memberOfMGNREGNA,omitempty"`
	IsPhysicallyDisabled      bool                `json:"isPhysicallyDisabled" bson:"isPhysicallyDisabled,omitempty"`
	IsAnyMmemberOfFamilyInCBO bool                `json:"isAnyMmemberOfFamilyInCBO" bson:"isAnyMmemberOfFamilyInCBO,omitempty"`
	Created                   Created             `json:"createdOn" bson:"createdOn,omitempty"`
	Block                     primitive.ObjectID  `json:"block"  bson:"block,omitempty"`
	Version                   int                 `json:"version"  bson:"version,omitempty"`
	IsVoiceSMS                bool                `json:"isVoiceSMS" bson:"isVoiceSMS,omitempty"`
	KitchenGarden             bool                `json:"kitchenGarden" bson:"kitchenGarden,omitempty"`
	LeasedInIrrigated         float64             `json:"leasedInIrrigated"  bson:"leasedInIrrigated,omitempty"`
	LeasedInRainfed           float64             `json:"leasedInRainfed"  bson:"leasedInRainfed,omitempty"`
	LeasedOutIrrigated        float64             `json:"leasedOutIrrigated"  bson:"leasedOutIrrigated,omitempty"`
	LeasedOutRainfed          float64             `json:"leasedOutRainfed"  bson:"leasedOutRainfed,omitempty"`
	MembershipInMgnrega       bool                `json:"membershipInMgnrega" bson:"membershipInMgnrega,omitempty"`
	MobileNumber              string              `json:"mobileNumber"  bson:"mobileNumber,omitempty"`
	LandLineNumber            string              `json:"landLineNumber"  bson:"landLineNumber,omitempty"`
	OwnedIrrigated            float64             `json:"ownedIrrigated"  bson:"ownedIrrigated,omitempty"`
	OwnedRainfed              float64             `json:"ownedRainfed"  bson:"ownedRainfed,omitempty"`
	SoilType                  primitive.ObjectID  `json:"soilType"  bson:"soilType,omitempty"`
	SpouseName                string              `json:"spouseName"  bson:"spouseName,omitempty"`
	State                     primitive.ObjectID  `json:"state"  bson:"state,omitempty"`
	//	TotalLand                    float64              `json:"totalLand"  bson:"totalLand,omitempty"`
	TotalMobiles                 float64              `json:"totalMobiles"  bson:"totalMobiles,omitempty"`
	Village                      primitive.ObjectID   `json:"village"  bson:"village,omitempty"`
	YearlyIncome                 string               `json:"yearlyIncome"  bson:"yearlyIncome,omitempty"`
	FarmerOrg                    primitive.ObjectID   `json:"farmerOrg"  bson:"farmerOrg,omitempty"`
	FarmerID                     string               `json:"farmerID"  bson:"farmerID,omitempty"`
	DoorNumber                   string               `json:"doorNo"  bson:"doorNo,omitempty"`
	LandMark                     string               `json:"landMark"  bson:"landMark,omitempty"`
	Street                       string               `json:"street"  bson:"street,omitempty"`
	PreferredMarkets             []string             `json:"preferredMarkets"  bson:"preferredMarkets,omitempty"`
	IsDisabled                   bool                 `json:"isDisabled"  bson:"isDisabled,omitempty"`
	FertilizerQuantity           float64              `json:"fertilizerQuantity"  bson:"fertilizerQuantity,omitempty"`
	RainDisasterFrom             float64              `json:"rainDisasterFrom"  bson:"rainDisasterFrom,omitempty"`
	RainDisasterTo               float64              `json:"rainDisasterTo"  bson:"rainDisasterTo,omitempty"`
	RainfallMedFrom              float64              `json:"rainfallMedFrom"  bson:"rainfallMedFrom,omitempty"`
	RainfallMedTo                float64              `json:"rainfallMedTo"  bson:"rainfallMedTo,omitempty"`
	RelativeHumidityDisasterFrom float64              `json:"relativeHumidityDisasterFrom"  bson:"relativeHumidityDisasterFrom,omitempty"`
	SeedQuantity                 float64              `json:"seedQuantity"  bson:"seedQuantity,omitempty"`
	TemperatureDisasterFrom      float64              `json:"temperatureDisasterFrom"  bson:"temperatureDisasterFrom,omitempty"`
	TemperatureDisasterTo        float64              `json:"temperatureDisasterTo"  bson:"temperatureDisasterTo,omitempty"`
	WindDirectionDisasterTo      float64              `json:"windDirectionDisasterTo"  bson:"windDirectionDisasterTo,omitempty"`
	WindDirectionMedFrom         float64              `json:"windDirectionMedFrom"  bson:"windDirectionMedFrom,omitempty"`
	RelativeHumidityDisasterTo   float64              `json:"relativeHumidityDisasterTo"  bson:"relativeHumidityDisasterTo,omitempty"`
	WindDirectionMedTo           float64              `json:"windDirectionMedTo"  bson:"windDirectionMedTo,omitempty"`
	WindSpeedDisasterFrom        float64              `json:"windSpeedDisasterFrom"  bson:"windSpeedDisasterFrom,omitempty"`
	WindSpeedDisasterTo          float64              `json:"windSpeedDisasterTo"  bson:"windSpeedDisasterTo,omitempty"`
	WindSpeedMedFrom             float64              `json:"windSpeedMedFrom"  bson:"windSpeedMedFrom,omitempty"`
	WindSpeedMedTo               float64              `json:"windSpeedMedTo"  bson:"windSpeedMedTo,omitempty"`
	UniqueId                     int64                `json:"uniqueId"  bson:"uniqueId,omitempty"`
	Password                     string               `json:"password"  bson:"password,omitempty"`
	Assert                       []primitive.ObjectID `json:"assert"  bson:"assert,omitempty"`
	Cast                         string               `json:"cast"  bson:"cast,omitempty"`
	Isleadwomenhasmobile         bool                 `json:"isleadwomenhasmobile" bson:"isleadwomenhasmobile"`
	WomenMobileNumber            string               `json:"womenMobileNumber" bson:"womanMobileNumber"`
	CreditName                   string               `json:"creditName"  bson:"creditName,omitempty"`
	DisabilityType               string               `json:"disabilityType"  bson:"disabilityType,omitempty"`
	MemberName                   string               `json:"memberName"  bson:"memberName,omitempty"`
	ProjectID                    primitive.ObjectID   `json:"-"  bson:"-"`
	Projects                     []primitive.ObjectID `json:"projects"  bson:"projects,omitempty"`
	CreatedDate                  *time.Time           `json:"createdDate"  bson:"createdDate"`
	Email                        string               `json:"email" bson:"email,omitempty"`
	CultivatedArea               float64              `json:"cultivatedArea"  bson:"cultivatedArea,omitempty"`
	VacantArea                   float64              `json:"vacantArea"  bson:"vacantArea,omitempty"`
	TotalLand                    float64              `json:"totalLand"  bson:"totalLand,omitempty"`
	CropCount                    float64              `json:"cropCount"  bson:"cropCount,omitempty"`
	LiveStockCount               float64              `json:"liveStockCount"  bson:"liveStockCount,omitempty"`
	AppRegistrationToken         string               `json:"appRegistrationToken" bson:"appRegistrationToken,omitempty"`
	RegistrationType             string               `json:"registrationType" bson:"registrationType,omitempty"`
	PinCode                      float64              `json:"pinCode"  bson:"pinCode,omitempty"`
	//	Location                     Location             `json:"location"  bson:"location,omitempty"`
	UploadVersion string   `json:"uploadVersion"  bson:"uploadVersion"`
	UserName      string   `json:"userName"  bson:"userName"`
	ProfileImg    string   `json:"profileImg" bson:"profileImg,omitempty"`
	Location      Location `json:"location" bson:"location,omitempty"`
	//ProfileImg                   string               `json:"profileImg"  bson:"profileImg"`
	CreatedBy primitive.ObjectID `json:"createdBy" bson:"createdBy"`
}

//RefFarmer : "Farmer with refrence data such as language..."
type RefFarmer struct {
	Farmer `bson:",inline"`
	Ref    struct {
		Projects      []RefProjectFarmer `json:"projects" bson:"projects,omitempty"`
		State         State              `json:"state"  bson:"state,omitempty"`
		Block         Block              `json:"block"  bson:"block,omitempty"`
		District      District           `json:"district"  bson:"district,omitempty"`
		GramPanchayat GramPanchayat      `json:"gramPanchayat"  bson:"gramPanchayat,omitempty"`
		Village       Village            `json:"village"  bson:"village,omitempty"`
		SoilType      SoilType           `json:"soilType"  bson:"soilType,omitempty"`
		FarmerOrg     Organisation       `json:"farmerOrg"  bson:"farmerOrg,omitempty"`
		Assert        []Asset            `json:"assert"  bson:"assert,omitempty"`
		// Land          struct {
		// 	CultivatedArea float64 `json:"cultivatedArea"  bson:"cultivatedArea,omitempty"`
		// 	VacantArea     float64 `json:"vacantArea"  bson:"vacantArea,omitempty"`
		// 	TotalLand      float64 `json:"totalLand"  bson:"totalLand,omitempty"`
		// } `json:"land"  bson:"land,omitempty"`
		Crops      float64
		LiveStocks float64
	} `json:"ref,omitempty" bson:"ref,omitempty"`
}

//RefFarmer : "Farmer with refrence data such as language..."
type RefBasicFarmer struct {
	BasicFarmer `bson:",inline"`
	Ref         struct {
		// Projects      []RefProjectFarmer `json:"projects" bson:"projects,omitempty"`
		// State         State              `json:"state"  bson:"state,omitempty"`
		// Block         Block              `json:"block"  bson:"block,omitempty"`
		// District      District           `json:"district"  bson:"district,omitempty"`
		// GramPanchayat GramPanchayat      `json:"gramPanchayat"  bson:"gramPanchayat,omitempty"`
		// Village       Village            `json:"village"  bson:"village,omitempty"`
		// SoilType      SoilType           `json:"soilType"  bson:"soilType,omitempty"`
		// FarmerOrg     Organisation       `json:"farmerOrg"  bson:"farmerOrg,omitempty"`
		// Assert        []Asset            `json:"assert"  bson:"assert,omitempty"`
		// // Land          struct {
		// // 	CultivatedArea float64 `json:"cultivatedArea"  bson:"cultivatedArea,omitempty"`
		// // 	VacantArea     float64 `json:"vacantArea"  bson:"vacantArea,omitempty"`
		// // 	TotalLand      float64 `json:"totalLand"  bson:"totalLand,omitempty"`
		// // } `json:"land"  bson:"land,omitempty"`
		// Crops      float64
		// LiveStocks float64
	} `json:"ref,omitempty" bson:"ref,omitempty"`
}

//FarmerFilter : "Used for constructing filter query"
type FarmerFilter struct {
	State         []primitive.ObjectID `json:"state"  bson:"state,omitempty"`
	District      []primitive.ObjectID `json:"district"  bson:"district,omitempty"`
	Block         []primitive.ObjectID `json:"block"  bson:"block,omitempty"`
	GramPanchayat []primitive.ObjectID `json:"gramPanchayat"  bson:"gramPanchayat,omitempty"`
	Village       []primitive.ObjectID `json:"village"  bson:"village,omitempty"`
	SoilType      []primitive.ObjectID `json:"soilType"  bson:"soilType,omitempty"`
	FarmerOrg     []primitive.ObjectID `json:"farmerOrg"  bson:"farmerOrg,omitempty"`
	Project       []primitive.ObjectID `json:"project"  bson:"project,omitempty"`
	NonFarmerOrg  []primitive.ObjectID `json:"nonFarmerOrg"  bson:"nonFarmerOrg,omitempty"`
	NonProject    []primitive.ObjectID `json:"nonProject"  bson:"nonProject,omitempty"`
	CreatedDate   *struct {
		From *time.Time `json:"from"`
		To   *time.Time `json:"to"`
	} `json:"createdDate"  bson:"createdDate"`
	OmitProjectFarmer struct {
		Is      bool               `json:"is"`
		Project primitive.ObjectID `json:"project"`
	} `json:"omitProjectFarmer"`
	ActiveStatus []bool   `json:"activeStatus" bson:"activeStatus,omitempty"`
	Status       []string `json:"status" bson:"status,omitempty"`
	SortBy       string   `json:"sortBy"`
	SortOrder    int      `json:"sortOrder"`
	Regex        struct {
		Name         string `json:"name" bson:"name"`
		MobileNumber string `json:"mobileNumber"  bson:"mobileNumber,omitempty"`
		SpouseName   string `json:"spouseName"  bson:"spouseName,omitempty"`
	} `json:"regex" bson:"regex"`
	RemoveLookup struct {
		Projects      bool `json:"projects" bson:"projects,omitempty"`
		State         bool `json:"state"  bson:"state,omitempty"`
		Block         bool `json:"block"  bson:"block,omitempty"`
		District      bool `json:"district"  bson:"district,omitempty"`
		GramPanchayat bool `json:"gramPanchayat"  bson:"gramPanchayat,omitempty"`
		Village       bool `json:"village"  bson:"village,omitempty"`
		SoilType      bool `json:"soilType"  bson:"soilType,omitempty"`
		FarmerOrg     bool `json:"farmerOrg"  bson:"farmerOrg,omitempty"`
		Assert        bool `json:"assert"  bson:"assert,omitempty"`
	} `json:"removeLookup" bson:"removeLookup"`
	DataAccess DataAccessRequest `json:"dataAccess" bson:"dataAccess,omitempty"`
}
type DuplicateFarmerFilter struct {
	FarmerFilter `bson:",inline"`
	By           string `json:"by" bson:"by,omitempty"`
}
type DuplicateFarmerReport struct {
	ID      string      `json:"id" bson:"_id,omitempty"`
	Farmers []RefFarmer `json:"farmers" bson:"farmers,omitempty"`
}
type CurrentLiveStocks struct {
	FarmerLiveStock primitive.ObjectID `json:"farmerLiveStock"  bson:"farmerLiveStock,omitempty"`
	Quantity        int                `json:"quantity"  bson:"quantity,omitempty"`
	Commodity       primitive.ObjectID `json:"commodity"  bson:"commodity,omitempty"`
	Category        primitive.ObjectID `json:"category"  bson:"category,omitempty"`
	Variety         primitive.ObjectID `json:"variety"  bson:"variety,omitempty"`
	Stage           primitive.ObjectID `json:"stage"  bson:"stage,omitempty"`
}
type CurrentCrops struct {
	FarmerCrops primitive.ObjectID `json:"farmerCrops"  bson:"farmerCrops,omitempty"`
	Commodity   primitive.ObjectID `json:"commodity"  bson:"commodity,omitempty"`
	Category    primitive.ObjectID `json:"category"  bson:"category,omitempty"`
	Variety     primitive.ObjectID `json:"variety"  bson:"variety,omitempty"`
	Season      primitive.ObjectID `json:"season"  bson:"season,omitempty"`
}
type FarmerUniquinessChk struct {
	Success bool   `json:"success" bson:"success,omitempty"`
	Message string `json:"message" bson:"message,omitempty"`
}
type DissiminateFarmer struct {
	ID                   primitive.ObjectID `json:"id" form:"id," bson:"_id,omitempty"`
	Name                 string             `json:"name" bson:"name,omitempty"`
	MobileNumber         string             `json:"mobileNumber"  bson:"mobileNumber,omitempty"`
	Email                string             `json:"email" bson:"email,omitempty"`
	FarmerID             string             `json:"farmerID"  bson:"farmerID,omitempty"`
	AppRegistrationToken string             `json:"appRegistrationToken" bson:"appRegistrationToken,omitempty"`
}
type FarmerUploadError struct {
	MobileNumber string `json:"mobileNumber" bson:"mobileNumber,omitempty"`
	Error        string `json:"error" bson:"error,omitempty"`
}
type AddProjectFarmer struct {
	ID           primitive.ObjectID `json:"id" form:"id," bson:"_id,omitempty"`
	Organisation string             `json:"organisation" bson:"organisation,omitempty"`
	Name         string             `json:"name" bson:"name,omitempty"`
}
type FarmerLocation struct {
	ID           primitive.ObjectID `json:"id" form:"id," bson:"_id,omitempty"`
	Name         string             `json:"name" bson:"name,omitempty"`
	Gender       string             `json:"gender" bson:"gender,omitempty"`
	MobileNumber string             `json:"mobileNumber"  bson:"mobileNumber,omitempty"`
	ProfileImg   string             `json:"profileImg" bson:"profileImg,omitempty"`
	Location     Location           `json:"location" bson:"location,omitempty"`
}
