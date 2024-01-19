package constants

// otp related constants
const (
	CONSUMERLOGIN       = "consumerlogin"
	VALIDATEOTP         = "NO"
	INVALIDOTP          = "Invalid OTP"
	TOKENOTPLENGTH      = 4
	PHONEOTPLENGTH      = 4
	OTPEXPIRY           = 900
	USERLOGIN           = "userlogin"
	FARMERLOGIN         = "farmerlogin"
	FARMERREGISTERATIOM = "farmerregisteration"
	USERREGISTERATION   = "userregisteration"
)

//Otp Scenario
const (
	OTPSCENARIOPASSWORD = "forgotpassword"
	OTPSCENARIOTOKEN    = "token"
)

//SMS Config
const (
	SMSURL      = "SMS_URL"
	SMSUSERNAME = "SMS_USERNAME"
	SMSPASSWORD = "SMS_PASSWORD"
	SMSSENDERID = "SMS_SENDER_ID"
	SMSCHANNEL  = "SMS_CHANNEL"
	SMSDCS      = "SMS_DCS"
	SMSFLASH    = "SMS_FLASH"
	SMSROUTE    = "SMS_ROUTE"
)

//Email Config
const (
	FROMEMAIL         = "EMAIL_FROM"
	FROMEMAILPASSWORD = "EMAIL_FROM_PASSWORD"
	SMTPHOST          = "EMAIL_SMTPHOST"
	SMTPPORT          = "EMAIL_SMTPPORT"
)

// otp related error constants
const (
	INTERNALSERVERERROR = "internal server error"
)
const (
	WHATSAPP_URL = "WHATSAPP_URL"

	WHATSAPP_BUSINESSID = "WHATSAPP_BUSINESSID"
)
const (
	WEATHERDATA_URL = "WEATHERDATA_URL"

	WEATHERDATA_ID = "WEATHERDATA_ID"
)

const (
	DBURL          = "mongodb_url"
	DOCLOC         = "FILE_URL"
	COMPENDIUMFILE = "COMPENDIUM_FILE_URL"
	UILOC          = "UI_URL"
	DOCLOCD        = "FILE_URL_D"
	TEMPLATELOC    = "TEMPLATE_URL"
)

const (
	SMSTEMPLATE = `Hi {#VAR#},
	You have received a notification from NICESSM regarding {#VAR#}{#VAR#}{#VAR#}{#VAR#}{#VAR#}.
	Notification - {#VAR#}{#VAR#}{#VAR#}{#VAR#}{#VAR#}.
	Regards,
	From NICESSM`
)

// otp related constants
const (
	DEALERREGISTRATION   = "dealerregistration"
	DEALERLOGIN          = "dealerlogin"
	DEALERVALIDATEOTP    = "NO"
	DEALERINVALIDOTP     = "Invalid OTP"
	DEALERTOKENOTPLENGTH = 4
	DEALERPHONEOTPLENGTH = 4
	DEALEROTPEXPIRY      = 900
)
