package models

//Asset : ""
type ExcelUploadError struct {
	ID      string `json:"id" form:"id," bson:"_id,omitempty"`
	Error   string `json:"error"  bson:"error,omitempty"`
	Message string `json:"message,omitempty"  bson:"message,omitempty"`
}
