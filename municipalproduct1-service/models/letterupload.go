package models

import "time"

//LetterUpload : ""
type LetterUpload struct {
	UniqueID string     `json:"uniqueId" bson:"uniqueId,omitempty"`
	Name     string     `json:"name" bson:"name,omitempty"`
	Desc     string     `json:"desc" bson:"desc,omitempty"`
	Subject  string     `json:"subject" bson:"subject,omitempty"`
	Status   string     `json:"status" bson:"status,omitempty"`
	Date     *time.Time `json:"date" bson:"date,omitempty"`
	NO       string     `json:"no" bson:"no,omitempty"`
	Created  *CreatedV2 `json:"created,omitempty"  bson:"created,omitempty"`
	Updated  []Updated  `json:"updated,omitempty"  bson:"updated,omitempty"`
	From     string     `json:"from,omitempty"  bson:"from,omitempty"`
	URL      string     `json:"url,omitempty"  bson:"url,omitempty"`
}

//RefLetterUpload : ""
type RefLetterUpload struct {
	LetterUpload `bson:",inline"`

	Ref struct {
		Creator     User     `json:"creator,omitempty" bson:"creator,omitempty"`
		CreatorType UserType `json:"creatorType,omitempty" bson:"creatorType,omitempty"`
	} `json:"ref,omitempty" bson:"ref,omitempty"`
}

//LetterUploadFilter : ""
type LetterUploadFilter struct {
	Status    []string  `json:"status,omitempty" bson:"status,omitempty"`
	DateRange DateRange `json:"dateRange,omitempty"  bson:"dateRange,omitempty"`
	Regex     struct {
		Name string `json:"name" bson:"name"`
		No   string `json:"no" bson:"no"`
		From string `json:"from" bson:"from"`
	} `json:"regex" bson:"regex"`
}
