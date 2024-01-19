package models

//Login : "login"
type Login struct {
	UserName string   `json:"userName"`
	PassWord string   `json:"password"`
	Location Location `json:"location,omitempty" bson:"location,omitempty"`
}
