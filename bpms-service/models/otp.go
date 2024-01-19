package models

//OTPLogin : ""
type OTPLogin struct {
	Mobile   string `json:"mobile"`
	OTP      string `json:"otp"`
	Scenario string `json:"scenario"`
}
