package models

type Notification struct {
	UniqueID string     `json:"uniqueId" bson:"uniqueId,omitempty"`
	Status   string     `json:"status" bson:"status,omitempty"`
	Created  *CreatedV2 `json:"created" bson:"created,omitempty"`
	Content  string     `json:"content" bson:"content,omitempty"`
	Title    string     `json:"title" bson:"title,omitempty"`
	Desc     string     `json:"desc" bson:"desc,omitempty"`
	Img      string     `json:"img" bson:"img,omitempty"`
}

type NotificationFilter struct {
	Status     []string `json:"status" bson:"status,omitempty"`
	UniqueID   []string `json:"uniqueId" bson:"uniqueId,omitempty"`
	SearchText struct {
		UniqueID string `json:"uniqueId" bson:"uniqueId,omitempty"`
		Content  string `json:"content" bson:"content,omitempty"`
	} `json:"searchText" bson:"searchText"`
}
type RefNotification struct {
	Notification `bson:",inline"`
}
