package models

import "time"

type ProductConfiguration struct {
	Name         string `json:"name" bson:"name,omitempty"`
	Lang         string `json:"lang" bson:"lang,omitempty"`
	Logo         string `json:"logo" bson:"logo,omitempty"`
	WaterMark    string `json:"waterMark" bson:"waterMark,omitempty"`
	LogoWithName string `json:"logoWithName" bson:"logoWithName,omitempty"`
	Email        struct {
		ContactUs string `json:"contactUs" bson:"contactUs,omitempty"`
		SendEmail string `json:"sendEmail" bson:"sendEmail,omitempty"`
	} `json:"email" bson:"email,omitempty"`
	Mobile                        string     `json:"mobile" bson:"mobile,omitempty"`
	Phone                         string     `json:"phone" bson:"phone,omitempty"`
	Address                       string     `json:"address" bson:"address,omitempty"`
	Copyrights                    string     `json:"copyrights" bson:"copyrights,omitempty"`
	PoweredBy                     string     `json:"poweredBy" bson:"poweredBy,omitempty"`
	Rights                        string     `json:"rights" bson:"rights,omitempty"`
	DOAStartYear                  string     `json:"doaStartYear" bson:"doaStartYear,omitempty"`
	ISLegacy                      string     `json:"isLegacy" bson:"isLegacy,omitempty"`
	IsLegacyV1                    string     `json:"isLegacyV1" bson:"isLegacyV1,omitempty"`
	IsLegacyV2                    string     `json:"isLegacyV2" bson:"isLegacyV2,omitempty"`
	IsCGBhilai                    string     `json:"isCgBhilai" bson:"isCgBhilai,omitempty"`
	ISMobileTower                 string     `json:"isMobileTower" bson:"isMobileTower,omitempty"`
	IsDefault                     bool       `json:"isDefault" bson:"isDefault,omitempty"`
	IsCaptcha                     string     `json:"isCaptcha" bson:"isCaptcha,omitempty"`
	FixedARVTaxCalc               string     `json:"fixedARVTaxCalc" bson:"fixedARVTaxCalc,omitempty"`
	PropertyHoldPayment           string     `json:"propertyHoldPayment" bson:"propertyHoldPayment,omitempty"`
	LocationID                    string     `json:"locationId" bson:"locationId,omitempty"`
	FYSelection                   string     `json:"fySelection" bson:"fySelection,omitempty"`
	Status                        string     `json:"status" bson:"status,omitempty"`
	UniqueID                      string     `json:"uniqueId,omitempty" bson:"uniqueId,omitempty"`
	TLPart2                       string     `json:"tlPart2" bson:"tlPart2"`
	Created                       Created    `json:"created" bson:"created,omitempty"`
	Location                      Location   `json:"location" bson:"location,omitempty"`
	ReportMgmt                    string     `json:"reportMgmt" bson:"reportMgmt,omitempty"`
	WaterBillMgmt                 string     `json:"waterBillMgmt" bson:"waterBillMgmt,omitempty"`
	TradeLisenceMgmt              string     `json:"tradeLisenceMgmt" bson:"tradeLisenceMgmt,omitempty"`
	TLRateMaster                  string     `json:"tlRateMaster" bson:"tlRateMaster,omitempty"`
	ShopRentMgmt                  string     `json:"shopRentMgmt" bson:"shopRentMgmt,omitempty"`
	AccountMgmt                   string     `json:"accountMgmt" bson:"accountMgmt,omitempty"`
	OperationsMgmt                string     `json:"operationsMgmt" bson:"operationsMgmt,omitempty"`
	APIURL                        string     `json:"-" bson:"apiUrl,omitempty"`
	UIURL                         string     `json:"-" bson:"uiUrl,omitempty"`
	PDFLogo                       string     `json:"pdfLogo" bson:"pdfLogo,omitempty"`
	PDFLogoWaterMark              string     `json:"pdfLogoWaterMark" bson:"pdfLogoWaterMark,omitempty"`
	PlsContactAdministrator       []string   `json:"plsContactAdministrator" bson:"plsContactAdministrator,omitempty"`
	ShowEditFloorDateFrom         string     `json:"showEditFloorDateFrom" bson:"showEditFloorDateFrom,omitempty"`
	ShowEditFloorDateTo           string     `json:"showEditFloorDateTo" bson:"showEditFloorDateTo,omitempty"`
	ShowHDFCButton                string     `json:"showHdfcButton" bson:"showHdfcButton,omitempty"`
	ShowPaytmButton               string     `json:"showPaytmButton" bson:"showPaytmButton,omitempty"`
	BaseURL                       string     `json:"baseUrl" bson:"baseUrl,omitempty"`
	OldDemandReceipt              string     `json:"oldDemandReceipt" bson:"oldDemandReceipt,omitempty"`
	ShowDemandListInConsumerLogin string     `json:"showDemandListInConsumerLogin" bson:"showDemandListInConsumerLogin,omitempty"`
	PayWhilePropertyStatusPending string     `json:"payWhilePropertyStatusPending" bson:"payWhilePropertyStatusPending,omitempty"`
	AppVersion                    float64    `json:"appVersion" bson:"appVersion,omitempty"`
	RemoveBuildUpAreaRestriction  string     `json:"removeBuildUpAreaRestriction" bson:"removeBuildUpAreaRestriction,omitempty"`
	ShowPropertyPayeeName         string     `json:"showPropertyPayeeName" bson:"showPropertyPayeeName,omitempty"`
	CompleteChequePayment         string     `json:"completeChequePayment" bson:"completeChequePayment,omitempty"`
	UserChargeDOA                 *time.Time `json:"userChargeDoa" bson:"userChargeDoa,omitempty"`
	UserChargeNewAssessmentV1     string     `json:"userChargeNewAssessmentV1" bson:"userChargeNewAssessmentV1,omitempty"`
	UserChargeNewAssessmentV2     string     `json:"userChargeNewAssessmentV2" bson:"userChargeNewAssessmentV2,omitempty"`
	PropertyManagement            string     `json:"propertyMgmt" bson:"propertyMgmt,omitempty"`
	UserChargeManagement          string     `json:"userChargeMgmt" bson:"userChargeMgmt,omitempty"`
	Municipality                  struct {
		Name string `json:"name" bson:"name,omitempty"`
		ID   string `json:"id" bson:"id,omitempty"`
	} `json:"municipality" bson:"municipality,omitempty"`
	State struct {
		Name string `json:"name" bson:"name,omitempty"`
		ID   string `json:"id" bson:"id,omitempty"`
	} `json:"state" bson:"state,omitempty"`
	District struct {
		Name string `json:"name" bson:"name,omitempty"`
		ID   string `json:"id" bson:"id,omitempty"`
	} `json:"district" bson:"district,omitempty"`
	Village struct {
		Name string `json:"name" bson:"name,omitempty"`
		ID   string `json:"id" bson:"id,omitempty"`
	} `json:"village" bson:"village,omitempty"`
	YOA struct {
		Name string `json:"name" bson:"name,omitempty"`
		ID   string `json:"id" bson:"id,omitempty"`
	} `json:"yoa" bson:"yoa,omitempty"`
}

type Logo struct {
	LogoB64  string `json:"logoB64" bson:"logoB64,omitempty"`
	UniqueID string `json:"uniqueId,omitempty" bson:"uniqueId,omitempty"`
}

type WatermarkLogo struct {
	WatermarkLogoB64 string `json:"watermarkLogoB64" bson:"watermarkLogoB64,omitempty"`
	UniqueID         string `json:"uniqueId,omitempty" bson:"uniqueId,omitempty"`
}

type ProductConfigurationFilter struct {
	SortBy    string `json:"sortBy,omitempty"`
	SortOrder int    `json:"sortOrder,omitempty"`
}

type RefProductConfiguration struct {
	ProductConfiguration `bson:",inline"`
	Ref                  struct {
	} `json:"ref,omitempty" bson:"ref,omitempty"`
}
