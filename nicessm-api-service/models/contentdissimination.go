package models

type ContentDissiminateUserAndFarmer struct {
	Users        []DissiminateUser   `json:"users" bson:"users,omitempty"`
	Farmers      []DissiminateFarmer `json:"farmers" bson:"farmers,omitempty"`
	UsersCount   int                 `json:"usersCount" bson:"usersCount,omitempty"`
	FarmersCount int                 `json:"farmersCount" bson:"farmersCount,omitempty"`
}
