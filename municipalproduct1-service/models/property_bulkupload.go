package models

type PropertyUploadError struct {
	MobileNumber string `json:"mobileNumber" bson:"mobileNumber,omitempty"`
	Error        string `json:"error" bson:"error,omitempty"`
}
