package models

import "go.mongodb.org/mongo-driver/bson/primitive"

//Authentication : "using for auth"
type Authentication struct {
	UserID   primitive.ObjectID `json:"userID,omitempty"`
	UserName string             `json:"username,omitempty"`
	Status   string             `json:"status,omitempty"`
	Role     string             `json:"role,omitempty"`
}

//Token : "struct"
type Token struct {
	Token string `json:"token"`
}

//Login : "login"
type Login struct {
	UserName    string   `json:"userName"`
	PassWord    string   `json:"password"`
	Environment string   `json:"environment"`
	Location    Location `json:"location,omitempty" bson:"location,omitempty"`
}

//OTPLogin : ""
type OTPLogin struct {
	Mobile   string `json:"mobile"`
	OTP      string `json:"otp"`
	Scenario string `json:"scenario"`
}

// MpinLogin : ""
type MpinLogin struct {
	UserName string `json:"userName"`
	Mpin     string `json:"mpin"`
}
