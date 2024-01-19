package models

type PropertyRequiredDocument struct {
	Name       string    `json:"name,omitempty" bson:"name,omitempty"`
	Status     string    `json:"status,omitempty" bson:"status,omitempty"`
	UniqueID   string    `json:"uniqueId,omitempty" bson:"uniqueId,omitempty"`
	For        string    `json:"for,omitempty" bson:"for,omitempty"`
	DocumentID string    `json:"documentId,omitempty" bson:"documentId,omitempty"`
	Created    CreatedV2 `json:"created,omitempty" bson:"created,omitempty"`
	UpdatedLog []Updated `json:"updatedLog,omitempty" bson:"updatedLog,omitempty"`
}

//RefPropertyRequiredDocument :""
type RefPropertyRequiredDocument struct {
	PropertyRequiredDocument `bson:",inline"`
	Ref                      struct {
		DocumentID DocumentList      `json:"documentId,omitempty" bson:"documentId,omitempty"`
		Document   PropertyDocuments `json:"document,omitempty" bson:"document,omitempty"`
	} `json:"ref,omitempty" bson:"ref,omitempty"`
}

//PropertyRequiredDocument :""
type PropertyRequiredDocumentFilter struct {
	Status                     []string `json:"status,omitempty" bson:"status,omitempty"`
	For                        []string `json:"for,omitempty" bson:"for,omitempty"`
	IsPropertyDcumentsUploaded bool     `json:"isPropertyDocumentUploaded"`
	SortBy                     string   `json:"sortBy,omitempty" bson:"sortBy,omitempty"`
	SortOrder                  int      `json:"sortOrder,omitempty" bson:"sortOrder,omitempty"`
}
