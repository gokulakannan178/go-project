package models

type NotificationLog struct {
	UniqueID       string     `json:"uniqueId" bson:"uniqueId,omitempty"`
	NotificationID string     `json:"notificationID" bson:"notificationID,omitempty"`
	Status         string     `json:"status" bson:"status,omitempty"`
	Created        *CreatedV2 `json:"created" bson:"created,omitempty"`
	Content        string     `json:"content" bson:"content,omitempty"`
	To             []struct {
		UserID string `json:"userId" bson:"userId,omitempty"`
		Name   string `json:"name" bson:"name,omitempty"`
		Type   string `json:"type" bson:"type,omitempty"`
	} `json:"to" bson:"to,omitempty"`
}

type NotificationLogFilter struct {
	Status     []string `json:"status" bson:"status,omitempty"`
	UniqueID   []string `json:"uniqueId" bson:"uniqueId,omitempty"`
	SearchText struct {
		UniqueID string `json:"uniqueId" bson:"uniqueId,omitempty"`
		Content  string `json:"content" bson:"content,omitempty"`
	} `json:"searchText" bson:"searchText"`
}

type RefNotificationLog struct {
	NotificationLog `bson:",inline"`
	Ref             struct {
		Notification Notification `json:"notification,omitempty" bson:"notification,omitempty"`
	} `json:"ref" bson:"ref"`
}
