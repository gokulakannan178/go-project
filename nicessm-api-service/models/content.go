package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

//Content : ""
type Content struct {
	ID              primitive.ObjectID `json:"id" form:"id," bson:"_id,omitempty"`
	Class           string             `json:"class,omitempty" bson:"class,omitempty"`
	ActiveStatus    bool               `json:"activeStatus,omitempty" bson:"activeStatus,omitempty"`
	ByPass          bool               `json:"byPass,omitempty" bson:"byPass,omitempty"`
	Author          primitive.ObjectID `json:"author" bson:"author,omitempty"`
	Controller      string             `json:"controller,omitempty" bson:"controller,omitempty"`
	Content         string             `json:"content" bson:"content,omitempty"`
	LinkType        string             `json:"linkType" bson:"linkType,omitempty"`
	ContentTitle    string             `json:"contentTitle" bson:"contentTitle,omitempty"`
	Document        string             `json:"document" bson:"document,omitempty"`
	DateCreated     *time.Time         `json:"dateCreated" form:"dateCreated" bson:"dateCreated,omitempty"`
	IgnoredIndex    []string           `json:"ignoredIndex,omitempty" bson:"ignoredIndex,omitempty"`
	IndexingData    IndexingData       `json:"indexingData,omitempty" bson:"indexingData,omitempty"`
	KnowledgeDomain primitive.ObjectID `json:"knowledgeDomain" bson:"knowledgeDomain,omitempty"`
	LocationRank    int                `json:"locationRank,omitempty" bson:"locationRank,omitempty"`
	Source          string             `json:"source,omitempty" bson:"source,omitempty"`
	Status          string             `json:"status,omitempty" bson:"status,omitempty"`
	SubDomain       primitive.ObjectID `json:"subDomain" bson:"subDomain,omitempty"`
	SubTopic        primitive.ObjectID `json:"subTopic" bson:"subTopic,omitempty"`
	TimeApplicable  struct {
		Month              int    `json:"month" bson:"month,omitempty"`
		Date               int    `json:"date" bson:"date,omitempty"`
		FromDate           int    `json:"fromDate" bson:"fromDate,omitempty"`
		ToDate             int    `json:"toDate" bson:"toDate,omitempty"`
		Type               int    `json:"type" bson:"type,omitempty"`
		TimeApplicableType string `json:"timeApplicableType" bson:"timeApplicableType,omitempty"`
	} `json:"timeApplicable" bson:"timeApplicable,omitempty"`
	Topic               primitive.ObjectID `json:"topic" bson:"topic,omitempty"`
	Type                string             `json:"type,omitempty" bson:"type,omitempty"`
	Version             int                `json:"version" form:"version" bson:"version,omitempty"`
	RecordId            string             `json:"recordId" bson:"recordId,omitempty"`
	DateReviewed        *time.Time         `json:"dateReviewed" form:"dateReviewed" bson:"dateReviewed,omitempty"`
	OriginalContent     string             `json:"originalContent,omitempty" bson:"originalContent,omitempty"`
	ReviewedBy          primitive.ObjectID `json:"reviewedBy,omitempty" bson:"reviewedBy,omitempty"`
	Organisation        primitive.ObjectID `json:"organisation" bson:"organisation,omitempty"`
	Tag                 string             `json:"tag,omitempty" bson:"tag,omitempty"`
	Tentativedate       string             `json:"tentativedate,omitempty" bson:"tentativedate,omitempty"`
	Compendium          primitive.ObjectID `json:"compendium" bson:"compendium,omitempty"`
	Project             primitive.ObjectID `json:"project" bson:"project,omitempty"`
	SmsType             string             `json:"smsType,omitempty" bson:"smsType,omitempty"`
	FarmerViewCount     int64              `json:"farmerViewCount" bson:"farmerViewCount,omitempty"`
	UsersViewCount      int64              `json:"usersViewCount" bson:"usersViewCount,omitempty"`
	GuestUsersViewCount int64              `json:"guestUsersViewCount" bson:"guestUsersViewCount,omitempty"`
	Comment             string             `json:"comment,omitempty" bson:"comment,omitempty"`
	Note                string             `json:"note,omitempty" bson:"note,omitempty"`
	Sms                 string             `json:"sms,omitempty" bson:"sms,omitempty"`
	SmsContentType      string             `json:"smsContentType,omitempty" bson:"smsContentType,omitempty"`
}
type IndexingData struct {
	Classfication  string             `json:"CLASSIFICATION" bson:"CLASSIFICATION,omitempty"`
	Cause          string             `json:"CAUSE,omitempty" bson:"CAUSE,omitempty"`
	Cause_type     string             `json:"CAUSE_TYPE,omitempty" bson:"CAUSE_TYPE,omitempty"`
	Irrigation     string             `json:"IRRIGATION,omitempty" bson:"IRRIGATION,omitempty"`
	State          primitive.ObjectID `json:"STATE"  bson:"STATE,omitempty"`
	District       primitive.ObjectID `json:"DISTRICT"  bson:"DISTRICT,omitempty"`
	Block          primitive.ObjectID `json:"BLOCK"  bson:"BLOCK,omitempty"`
	Season         primitive.ObjectID `json:"SEASON"  bson:"SEASON,omitempty"`
	Stage          primitive.ObjectID `json:"STAGE"  bson:"STAGE,omitempty"`
	Village        primitive.ObjectID `json:"VILLAGE"  bson:"VILLAGE,omitempty"`
	Gram_panchayat primitive.ObjectID `json:"GRAM_PANCHAYAT"  bson:"GRAM_PANCHAYAT,omitempty"`
	Function       primitive.ObjectID `json:"FUNCTION"  bson:"FUNCTION,omitempty"`
	Soil_type      primitive.ObjectID `json:"SOIL_TYPE"  bson:"SOIL_TYPE,omitempty"`
	Market         primitive.ObjectID `json:"MARKET"  bson:"MARKET,omitempty"`
	Causative      primitive.ObjectID `json:"CAUSATIVE"  bson:"CAUSATIVE,omitempty"`
	Variety        primitive.ObjectID `json:"VARIETY"  bson:"VARIETY,omitempty"`
	Category       primitive.ObjectID `json:"CATEGORY,omitempty" bson:"CATEGORY,omitempty"`
	Commodity      primitive.ObjectID `json:"COMMODITY,omitempty" bson:"COMMODITY,omitempty"`
	SubVariety     primitive.ObjectID `json:"SUB_VARIETY,omitempty" bson:"SUB_VARIETY,omitempty"`
	SoilData       primitive.ObjectID `json:"SOIL_DATA,omitempty" bson:"SOIL_DATA,omitempty"`
}

type ContentFilter struct {
	Classfication     []string             `json:"classfication,omitempty" bson:"classfication,omitempty"`
	ContentTitle      []string             `json:"contentTitle,omitempty" bson:"contentTitle,omitempty"`
	Status            []string             `json:"status" form:"status" bson:"status,omitempty"`
	OmitStatus        []string             `json:"omitStatus" form:"omitStatus" bson:"omitStatus,omitempty"`
	Irrigation        string               `json:"irrigation,omitempty" bson:"irrigation,omitempty"`
	KnowledgeDomain   []primitive.ObjectID `json:"knowledgeDomain" bson:"knowledgeDomain,omitempty"`
	State             []primitive.ObjectID `json:"state"  bson:"state,omitempty"`
	Soil_type         []primitive.ObjectID `json:"soil_type"  bson:"soil_type,omitempty"`
	Organisation      []primitive.ObjectID `json:"organisation" bson:"organisation,omitempty"`
	Author            []primitive.ObjectID `json:"author" bson:"author,omitempty"`
	SubDomain         []primitive.ObjectID `json:"subDomain" bson:"subDomain,omitempty"`
	SubTopic          []primitive.ObjectID `json:"subTopic" bson:"subTopic,omitempty"`
	Topic             []primitive.ObjectID `json:"Topic" bson:"Topic,omitempty"`
	ReviewedBy        []primitive.ObjectID `json:"reviewedBy,omitempty" bson:"reviewedBy,omitempty"`
	Project           []primitive.ObjectID `json:"project" bson:"project,omitempty"`
	SmsType           []string             `json:"smsType,omitempty" bson:"smsType,omitempty"`
	SmsContentType    []string             `json:"smsContentType,omitempty" bson:"smsContentType,omitempty"`
	RecordId          []string             `json:"recordId,omitempty" bson:"recordId,omitempty"`
	Source            []string             `json:"source,omitempty" bson:"source,omitempty"`
	TranslationStatus []string             `json:"translationStatus,omitempty" bson:"translationStatus,omitempty"`
	Type              []string             `json:"type,omitempty" bson:"type,omitempty"`
	District          []primitive.ObjectID `json:"district"  bson:"district,omitempty"`
	Block             []primitive.ObjectID `json:"block"  bson:"block,omitempty"`
	Season            []primitive.ObjectID `json:"season"  bson:"season,omitempty"`
	Stage             []primitive.ObjectID `json:"stage"  bson:"stage,omitempty"`
	Village           []primitive.ObjectID `json:"village"  bson:"village,omitempty"`
	Gram_panchayat    []primitive.ObjectID `json:"gram_panchayat"  bson:"gram_panchayat,omitempty"`
	Function          []primitive.ObjectID `json:"function"  bson:"function,omitempty"`
	Market            []primitive.ObjectID `json:"market"  bson:"market,omitempty"`
	Causative         []primitive.ObjectID `json:"causative"  bson:"causative,omitempty"`
	Variety           []primitive.ObjectID `json:"variety"  bson:"variety,omitempty"`
	SubVariety        []primitive.ObjectID `json:"subVariety,omitempty" bson:"subVariety,omitempty"`
	SoilData          []primitive.ObjectID `json:"soilData,omitempty" bson:"soilData,omitempty"`
	Category          []string             `json:"category,omitempty" bson:"category,omitempty"`
	Commodity         []string             `json:"commodity,omitempty" bson:"commodity,omitempty"`
	Cause             []string             `json:"cause,omitempty" bson:"cause,omitempty"`
	Cause_type        []string             `json:"cause_type,omitempty" bson:"cause_type,omitempty"`
	CreatedFrom       struct {
		StartDate *time.Time `json:"startDate"`
		EndDate   *time.Time `json:"endDate"`
	} `json:"createdFrom"`
	SortBy    string `json:"sortBy"`
	SortOrder int    `json:"sortOrder"`
	SearchBox struct {
		Content      string `json:"content,omitempty" bson:"content,omitempty"`
		ContentTitle string `json:"contentTitle" bson:"contentTitle,omitempty"`
		RecordId     string `json:"recordId,omitempty" bson:"recordId,omitempty"`
		Comment      string `json:"comment,omitempty" bson:"comment,omitempty"`
	} `json:"searchBox" bson:"searchBox"`
	DataAccess DataAccessRequest `json:"dataAccess" bson:"dataAccess,omitempty"`
}

type ContentCount struct {
	Total                   int `json:"total" bson:"total,omitempty"`
	SMSCount                int `json:"smsCount" bson:"smsCount,omitempty"`
	MonitoringSMSCount      int `json:"monitoringSmsCount" bson:"monitoringSmsCount,omitempty"`
	OrganicSMSCount         int `json:"organicSmsCount" bson:"organicSmsCount,omitempty"`
	InOrganicSMSCount       int `json:"inOrganicSmsCount" bson:"inOrganicSmsCount,omitempty"`
	VoiceSmsCount           int `json:"voiceSmsCount" bson:"voiceSmsCount,omitempty"`
	MonitoringVoiceSmsCount int `json:"monitoringVoiceSmsCount" bson:"monitoringVoiceSmsCount,omitempty"`
	OrganicVoiceSmsCount    int `json:"organicVoiceSmsCount" bson:"organicVoiceSmsCount,omitempty"`
	InOrganicVoiceSmsCount  int `json:"inOrganicVoiceSmsCount" bson:"inOrganicVoiceSmsCount,omitempty"`
	VideoUrlCount           int `json:"videoUrlCount" bson:"videoUrlCount,omitempty"`
	DocumentCount           int `json:"documentCount" bson:"documentCount,omitempty"`
	OnePageHtmlCount        int `json:"onePageHtmlCount" bson:"onePageHtmlCount,omitempty"`
}

// type RefContentCount struct {
// 	ContentCount []ContentCount `json:"contentCount,omitempty" bson:"contentCount,omitempty"`
// }

type RefContent struct {
	Content `bson:",inline"`
	Ref     struct {
		KnowledgeDomain    KnowledgeDomain         `json:"knowledgeDomain" bson:"knowledgeDomain,omitempty"`
		Organisation       Organisation            `json:"organisation" bson:"organisation,omitempty"`
		Project            Project                 `json:"project" bson:"project,omitempty"`
		SubDomain          SubDomain               `json:"subDomain" bson:"subDomain,omitempty"`
		SubTopic           SubTopic                `json:"subTopic" bson:"subTopic,omitempty"`
		Topic              Topic                   `json:"topic" bson:"topic,omitempty"`
		State              State                   `json:"state"  bson:"state,omitempty"`
		District           District                `json:"district"  bson:"district,omitempty"`
		Block              Block                   `json:"block"  bson:"block,omitempty"`
		Season             Cropseason              `json:"season"  bson:"season,omitempty"`
		Stage              CommodityStage          `json:"stage"  bson:"stage,omitempty"`
		Village            Village                 `json:"village"  bson:"village,omitempty"`
		Gram_panchayat     GramPanchayat           `json:"gram_panchayat"  bson:"gram_panchayat,omitempty"`
		Function           CommodityFunction       `json:"function"  bson:"function,omitempty"`
		Soil_type          SoilType                `json:"soil_type"  bson:"soil_type,omitempty"`
		Market             Market                  `json:"market"  bson:"market,omitempty"`
		Variety            CommodityVariety        `json:"variety"  bson:"variety,omitempty"`
		Author             User                    `json:"author" bson:"author,omitempty"`
		ReviewedBy         User                    `json:"reviewedBy,omitempty" bson:"reviewedBy,omitempty"`
		Category           CommodityCategory       `json:"category,omitempty" bson:"category,omitempty"`
		Commodity          Commodity               `json:"commodity,omitempty" bson:"commodity,omitempty"`
		CausativeInsect    Insect                  `json:"causativeInsect"  bson:"causativeInsect,omitempty"`
		CausativeDisease   Disease                 `json:"causativeDisease"  bson:"causativeDisease,omitempty"`
		SubVariety         CommoditySubVariety     `json:"subVariety,omitempty" bson:"subVariety,omitempty"`
		SoilData           FarmerSoilData          `json:"soilData,omitempty" bson:"soilData,omitempty"`
		TranslatedContents []RefContentTranslation `json:"translatedContents,omitempty" bson:"translatedContents,omitempty"`
	} `json:"ref,omitempty" bson:"ref,omitempty"`
}
type ApprovedContent struct {
	ID          primitive.ObjectID `json:"id" form:"id," bson:"_id,omitempty"`
	Status      string             `json:"status,omitempty" bson:"status,omitempty"`
	Content     string             `json:"content,omitempty" bson:"content,omitempty"`
	ReviewedBy  primitive.ObjectID `json:"reviewedBy" bson:"reviewedBy,omitempty"`
	DateCreated *time.Time         `json:"dateCreated" form:"dateCreated" bson:"dateCreated,omitempty"`
}
type RejectedContent struct {
	ID          primitive.ObjectID `json:"id" form:"id," bson:"_id,omitempty"`
	Status      string             `json:"status,omitempty" bson:"status,omitempty"`
	ReviewedBy  primitive.ObjectID `json:"reviewedBy" bson:"reviewedBy,omitempty"`
	DateCreated *time.Time         `json:"dateCreated" form:"dateCreated" bson:"dateCreated,omitempty"`
	Note        string             `json:"note,omitempty" bson:"note,omitempty"`
}
type ContentViewCount struct {
	ContentId primitive.ObjectID `json:"contentId,omitempty" bson:"contentId,omitempty"`
	UserType  string             `json:"userType,omitempty" bson:"userType,omitempty"`
	UserId    primitive.ObjectID `json:"userId" bson:"userId,omitempty"`
}
type ContentDataAccess struct {
	Organisation  []primitive.ObjectID `json:"organisation" bson:"organisation,omitempty"`
	Project       []primitive.ObjectID `json:"project" bson:"project,omitempty"`
	State         []primitive.ObjectID `json:"state" bson:"state,omitempty"`
	District      []primitive.ObjectID `json:"district" bson:"district,omitempty"`
	Block         []primitive.ObjectID `json:"block" bson:"block,omitempty"`
	GramPanchayat []primitive.ObjectID `json:"gramPanchayat" bson:"gramPanchayat,omitempty"`
	Village       []primitive.ObjectID `json:"village" bson:"village,omitempty"`
}
type DuplicateContentFilter struct {
	UserFilter `bson:",inline"`
	By         string `json:"by" bson:"by,omitempty"`
}
type DuplicateContentReport struct {
	ID struct {
		Kd       primitive.ObjectID `json:"kd" bson:"kd,omitempty"`
		Sd       primitive.ObjectID `json:"sd" bson:"sd,omitempty"`
		Topic    primitive.ObjectID `json:"topic" bson:"topic,omitempty"`
		Subtopic primitive.ObjectID `json:"subtopic" bson:"subtopic,omitempty"`
		Season   primitive.ObjectID `json:"season" bson:"season,omitempty"`
	} `json:"id" bson:"_id,omitempty"`
	Contents []struct {
		RecordId string `json:"recordId" bson:"recordId,omitempty"`
	} `json:"contents" bson:"contents,omitempty"`
	Ref struct {
		KnowledgeDomain KnowledgeDomain `json:"knowledgeDomain" bson:"knowledgeDomain,omitempty"`
		SubDomain       SubDomain       `json:"subDomain" bson:"subDomain,omitempty"`
		SubTopic        SubTopic        `json:"subTopic" bson:"subTopic,omitempty"`
		Topic           Topic           `json:"topic" bson:"topic,omitempty"`
		Season          Cropseason      `json:"season"  bson:"season,omitempty"`
	} `json:"ref,omitempty" bson:"ref,omitempty"`
}
