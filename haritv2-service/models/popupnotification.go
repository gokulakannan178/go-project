package models

type PopupNotification struct {
	UniqueID string   `json:"uniqueId" bson:"uniqueId,omitempty"`
	Status   string   `json:"status" bson:"status,omitempty"`
	Created  *Created `json:"createdOn" bson:"createdOn,omitempty"`
	Type     string   `json:"type" bson:"type,omitempty"`
	ImgUrl   string   `json:"imgurl" bson:"imgurl,omitempty"`
	IsPop    bool     `json:"ispop" bson:"ispop,omitempty"`
	Title    string   `json:"title" bson:"title,omitempty"`
}
type PopupNotificationFilter struct {
	Status     []string `json:"status" bson:"status,omitempty"`
	Type       []string `json:"type" bson:"type,omitempty"`
	IsPop      []bool   `json:"ispop" bson:"ispop,omitempty"`
	SearchText struct {
		Title string `json:"title" bson:"title,omitempty"`
	} `json:"searchtext" bson:"searchtext"`
}
type RefPopupNotification struct {
	PopupNotification `bson:",inline"`
}
type Ispoptrue struct {
	UniqueID string `json:"uniqueId" bson:"uniqueId,omitempty"`
	Type     string `json:"type" bson:"type,omitempty"`
}
