package models

import "go.mongodb.org/mongo-driver/bson/primitive"

//UserAcl : ""
type UserAcl struct {
	ID                    primitive.ObjectID    `json:"id"  bson:"_id,omitempty"`
	UniqueID              string                `json:"uniqueId" bson:"uniqueId,omitempty"`
	UserName              primitive.ObjectID    `json:"userName" bson:"userName,omitempty"`
	User                  Access                `json:"user" bson:"user,omitempty"`
	Location              Access                `json:"location" bson:"location,omitempty"`
	KnowledgeDomain       Access                `json:"knowledgeDomain" bson:"knowledgeDomain,omitempty"`
	Crop                  Access                `json:"crop" bson:"crop,omitempty"`
	Livestock             Access                `json:"livestock" bson:"livestock,omitempty"`
	StateSeason           Access                `json:"stateSeason" bson:"stateSeason,omitempty"`
	StateCalendar         Access                `json:"stateCalendar" bson:"stateCalendar,omitempty"`
	Farmer                Access                `json:"farmer" bson:"farmer,omitempty"`
	ContingencyPlan       Access                `json:"contingencyPlan" bson:"contingencyPlan,omitempty"`
	Market                Access                `json:"market" bson:"market,omitempty"`
	Language              Access                `json:"language" bson:"language,omitempty"`
	SoilType              Access                `json:"soilType" bson:"soilType,omitempty"`
	Asset                 Access                `json:"asset" bson:"asset,omitempty"`
	Insect                Access                `json:"insect" bson:"insect,omitempty"`
	Disease               Access                `json:"disease" bson:"disease,omitempty"`
	BannedItem            Access                `json:"bannedItem" bson:"bannedItem,omitempty"`
	NARPZones             Access                `json:"narpZones" bson:"narpZones,omitempty"`
	Vaccines              Access                `json:"vaccines" bson:"vaccines,omitempty"`
	LivestockVaccines     Access                `json:"livestockVaccines" bson:"livestockVaccines,omitempty"`
	DistrictWeather       Access                `json:"districtWeather" bson:"districtWeather,omitempty"`
	Organisation          Access                `json:"organisation" bson:"organisation,omitempty"`
	Project               Access                `json:"project" bson:"project,omitempty"`
	AidLocation           Access                `json:"aidLocation" bson:"aidLocation,omitempty"`
	ContentAndQueryAccess ContentAndQueryAccess `json:"contentAndQueryAccess" bson:"contentAndQueryAccess,omitempty"`
	SpecialFeatures       SpecialFeatures       `json:"specialFeatures" bson:"specialFeatures,omitempty"`
	Status                string                `json:"status" bson:"status,omitempty"`
	Created               *Created              `json:"created,omitempty"  bson:"created,omitempty"`
	Updated               []Updated             `json:"updated,omitempty"  bson:"updated,omitempty"`
}
type Access struct {
	Read  string `json:"read" bson:"read,omitempty"`
	Write string `json:"write" bson:"write,omitempty"`
}
type ContentAndQueryAccess struct {
	CreateEdit      string `json:"createEdit" bson:"createEdit,omitempty"`
	BypassContent   string `json:"bypassContent" bson:"bypassContent,omitempty"`
	Manage          string `json:"manage" bson:"manage,omitempty"`
	TranslateReview string `json:"translateReview" bson:"translateReview,omitempty"`
	Upload          string `json:"upload" bson:"upload,omitempty"`
	QueryEdit       string `json:"queryEdit" bson:"queryEdit,omitempty"`
	Delete          string `json:"delete" bson:"delete,omitempty"`
	Review          string `json:"review" bson:"review,omitempty"`
	Translate       string `json:"translate" bson:"translate,omitempty"`
	Disseminate     string `json:"disseminate" bson:"disseminate,omitempty"`
	Search          string `json:"search" bson:"search,omitempty"`
	Reassign        string `json:"reassign" bson:"reassign,omitempty"`
}
type SpecialFeatures struct {
	WeatherData      string `json:"weatherData" bson:"weatherData,omitempty"`
	PickResolveQuery string `json:"pickResolveQuery" bson:"pickResolveQuery,omitempty"`
}

//RefUserAcl : ""
type RefUserAcl struct {
	UserAcl `bson:",inline"`
	Ref     struct {
	} `json:"ref,omitempty" bson:"ref,omitempty"`
}

//UserAclFilter : ""
type UserAclFilter struct {
	Status    []string             `json:"status,omitempty" bson:"status,omitempty"`
	UserName  []primitive.ObjectID `json:"userName" bson:"userName,omitempty"`
	SortBy    string               `json:"sortBy,omitempty" bson:"sortBy,omitempty"`
	SortOrder int                  `json:"sortOrder,omitempty" bson:"sortOrder,omitempty"`
}
