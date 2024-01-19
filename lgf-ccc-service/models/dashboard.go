package models

//Dashboard : ""
type Dashboard struct {
	Total    int64 `json:"total" bson:"total,omitempty"`
	Active   int64 `json:"active" bson:"active,omitempty"`
	Disabled int64 `json:"disabled" bson:"disabled,omitempty"`
}
type UserTypeCount struct {
	Admin            int64 `json:"admin" bson:"admin,omitempty"`
	GarbageCollector int64 `json:"garbageCollector" bson:"garbageCollector,omitempty"`
	ProjectManager   int64 `json:"projectManager" bson:"projectManager,omitempty"`
	DumbSiteUser     int64 `json:"dumbSiteUser" bson:"dumbSiteUser,omitempty"`
	Citizen          int64 `json:"citizen" bson:"citizen,omitempty"`
	Driver           int64 `json:"driver" bson:"driver,omitempty"`
}

type PropertyCount struct {
	TotalProperty int64 `json:"totalProperty" bson:"totalProperty,omitempty"`
	Today         int64 `json:"today" bson:"today,omitempty"`
	ThisWeek      int64 `json:"thisWeek" bson:"thisWeek,omitempty"`
	ThisMonth     int64 `json:"thisMonth" bson:"thisMonth,omitempty"`
	ThisYear      int64 `json:"thisYear" bson:"thisYear,omitempty"`
}

type HousevisitedCount struct {
	Total        int64 `json:"total" bson:"total,omitempty"`
	Collected    int64 `json:"collected" bson:"collected,omitempty"`
	NotCollected int64 `json:"notCollected" bson:"notCollected,omitempty"`
}
type DumbSiteCount struct {
	Total        int64 `json:"total" bson:"total,omitempty"`
	TotalVehicle int64 `json:"totalVehicle" bson:"totalVehicle,omitempty"`
	Quantity     int64 `json:"quantity" bson:"quantity,omitempty"`
}
