package models

type RegistrationDealer struct {
	Dealer `bson:",inline"`
	OTP    string `json:"otp"`
}
