package models

import "go.mongodb.org/mongo-driver/bson/primitive"

//Authentication : "using for auth"
type FarmerAuthentication struct {
	FarmerID primitive.ObjectID `json:"farmerID,omitempty"`
	Name     string             `json:"name,omitempty"`
	Status   string             `json:"status,omitempty"`
	Role     string             `json:"role,omitempty"`
	Type     string             `json:"type,omitempty"`
}

//Token : "struct"
type FarmerToken struct {
	Token string `json:"token"`
}

//Login : "login"
type FarmerLogin struct {
	MobileNumber string   `json:"mobileNumber"`
	PassWord     string   `json:"password"`
	Location     Location `json:"location,omitempty" bson:"location,omitempty"`
}

//OTPLogin : ""
type FarmerOTPLogin struct {
	Farmer   `bson:",inline"`
	OTP      string `json:"otp"`
	Scenario string `json:"scenario"`
}
