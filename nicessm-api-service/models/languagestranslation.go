package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type LanguageTranslations struct {
	ID                     primitive.ObjectID `json:"id" form:"id," bson:"_id,omitempty"`
	LanguageType           string             `json:"languageType" form:"languageType" bson:"languageType,omitempty"`
	LoginButton            string             `json:"loginButton" form:"loginButton" bson:"loginButton,omitempty"`
	WeatherData            string             `json:"weatherData" form:"weatherData" bson:"weatherData,omitempty"`
	ViewMore               string             `json:"viewMore" form:"viewMore" bson:"viewMore,omitempty"`
	LoginHeader            string             `json:"loginHeader" form:"loginHeader" bson:"loginHeader,omitempty"`
	UserName               string             `json:"userName" form:"userName" bson:"userName,omitempty"`
	Password               string             `json:"password" form:"password" bson:"password,omitempty"`
	LogIn                  string             `json:"logIn" bson:"logIn,omitempty"`
	ForgotPassword         string             `json:"forgotPassword" bson:"forgotPassword,omitempty"`
	UserRegistration       string             `json:"userRegistration" bson:"userRegistration,omitempty"`
	FarmerRegistration     string             `json:"farmerRegistration" bson:"farmerRegistration,omitempty"`
	College                string             `json:"college" bson:"college,omitempty"`
	NameOfCity             string             `json:"nameOfCity" bson:"nameOfCity,omitempty"`
	Header                 string             `json:"header" bson:"header,omitempty"`
	ImportantLink          string             `json:"importantLink" bson:"importantLink,omitempty"`
	ImportantLink1         string             `json:"importantLink1" bson:"importantLink1,omitempty"`
	Link1                  string             `json:"link1" bson:"link1,omitempty"`
	Link2                  string             `json:"link2" bson:"link2,omitempty"`
	Link3                  string             `json:"link3" bson:"link3,omitempty"`
	Link4                  string             `json:"link4" bson:"link4,omitempty"`
	Link5                  string             `json:"link5" bson:"link5,omitempty"`
	Link6                  string             `json:"link6" bson:"link6,omitempty"`
	Link7                  string             `json:"link7" bson:"link7,omitempty"`
	ImportantLink2         string             `json:"importantLink2" bson:"importantLink2,omitempty"`
	ImportantLink3         string             `json:"importantLink3" bson:"importantLink3,omitempty"`
	ImportantLink4         string             `json:"importantLink4" bson:"importantLink4,omitempty"`
	ImportantLink5         string             `json:"importantLink5" bson:"importantLink5,omitempty"`
	ImportantLink6         string             `json:"importantLink6" bson:"importantLink6,omitempty"`
	ImportantLink7         string             `json:"importantLink7" bson:"importantLink7,omitempty"`
	FooterInformation      string             `json:"footerInformation" bson:"footerInformation,omitempty"`
	PowerdBy               string             `json:"powerdBy" bson:"powerdBy,omitempty"`
	PrivacyPolicy          string             `json:"privacyPolicy" bson:"privacyPolicy,omitempty"`
	Credit                 string             `json:"credit" bson:"credit,omitempty"`
	Disclaimer             string             `json:"disclaimer" bson:"disclaimer,omitempty"`
	ContactDetails         string             `json:"contactDetails" bson:"contactDetails,omitempty"`
	Person1Title           string             `json:"person1Title" bson:"person1Title,omitempty"`
	Person1Name            string             `json:"person1Name" bson:"person1Name,omitempty"`
	Person1Message         string             `json:"person1Message" bson:"person1Message,omitempty"`
	Person1Pdf             string             `json:"person1Pdf" bson:"person1Pdf,omitempty"`
	Person2Title           string             `json:"person2Title" bson:"person2Title,omitempty"`
	Person2Name            string             `json:"person2Name" bson:"person2Name,omitempty"`
	Person2Message         string             `json:"person2Message" bson:"person2Message,omitempty"`
	Person2Pdf             string             `json:"person2Pdf" bson:"person2Pdf,omitempty"`
	Person3Title           string             `json:"person3Title" bson:"person3Title,omitempty"`
	Person3Name            string             `json:"person3Name" bson:"person3Name,omitempty"`
	Person3Message         string             `json:"person3Message" bson:"person3Message,omitempty"`
	Person3Pdf             string             `json:"person3Pdf" bson:"person3Pdf,omitempty"`
	Person4Title           string             `json:"person4Title" bson:"person4Title,omitempty"`
	Person4Name            string             `json:"person4Name" bson:"person4Name,omitempty"`
	Person4Message         string             `json:"person4Message" bson:"person4Message,omitempty"`
	Person4Pdf             string             `json:"person4Pdf" bson:"person4Pdf,omitempty"`
	Organisation           string             `json:"organisation" bson:"organisation,omitempty"`
	Project                string             `json:"project" bson:"project,omitempty"`
	FirstName              string             `json:"firstName" bson:"firstName,omitempty"`
	Name                   string             `json:"name" bson:"name,omitempty"`
	Lastname               string             `json:"lastname" bson:"lastname,omitempty"`
	DateOfBirth            string             `json:"dateOfBirth" bson:"dateOfBirth,omitempty"`
	Gender                 string             `json:"gender" bson:"gender,omitempty"`
	FatherName             string             `json:"fatherName" bson:"fatherName,omitempty"`
	Address                string             `json:"address" bson:"address,omitempty"`
	State                  string             `json:"state" bson:"state,omitempty"`
	District               string             `json:"district" bson:"district,omitempty"`
	Block                  string             `json:"block" bson:"block,omitempty"`
	Grampanchayat          string             `json:"grampanchayat" bson:"grampanchayat,omitempty"`
	Village                string             `json:"village" bson:"village,omitempty"`
	PinCode                string             `json:"pinCode" bson:"pinCode,omitempty"`
	Mobile                 string             `json:"mobileNumber" bson:"mobileNumber,omitempty"`
	AlternateNumber        string             `json:"alternateNumber" bson:"alternateNumber,omitempty"`
	Email                  string             `json:"email" bson:"email,omitempty"`
	Role                   string             `json:"role" bson:"role,omitempty"`
	GenerateOtp            string             `json:"generateOtp" bson:"generateOtp,omitempty"`
	NewPassword            string             `json:"newPassword" bson:"newPassword,omitempty"`
	EnterThePasswordAgain  string             `json:"enterThePasswordAgain" bson:"enterThePasswordAgain,omitempty"`
	VerifyOtp              string             `json:"verifyOtp" bson:"verifyOtp,omitempty"`
	EnterOtp               string             `json:"enterOtp" bson:"enterOtp,omitempty"`
	Healthylifehealthysoil string             `json:"healthylifehealthysoil" bson:"healthylifehealthysoil,omitempty"`
	StartDate              string             `json:"startDate" bson:"startDate,omitempty"`
	LastDate               string             `json:"lastDate" bson:"lastDate,omitempty"`
	Search                 string             `json:"search" bson:"search,omitempty"`
	Reset                  string             `json:"reset" bson:"reset,omitempty"`
	Rain                   string             `json:"rain" bson:"rain,omitempty"`
	MinimumTemperature     string             `json:"minimumTemperature" bson:"minimumTemperature,omitempty"`
	MaximumTemperature     string             `json:"maximumTemperature" bson:"maximumTemperature,omitempty"`
	Damp                   string             `json:"damp" bson:"damp,omitempty"`
	WindSpeed              string             `json:"windSpeed" bson:"windSpeed,omitempty"`
	First                  string             `json:"first" bson:"first,omitempty"`
	Last                   string             `json:"last" bson:"last,omitempty"`
	Next                   string             `json:"next" bson:"next,omitempty"`
	Previous               string             `json:"previous" bson:"previous,omitempty"`
	ForgotPasswordTitle    string             `json:"forgotPasswordTitle" bson:"forgotPasswordTitle,omitempty"`
	BackToLogin            string             `json:"backToLogin" bson:"backToLogin,omitempty"`
	NickName               string             `json:"nickName" bson:"nickName,omitempty"`
	City                   string             `json:"city" bson:"city,omitempty"`
	UserSelfRegistration   string             `json:"userSelfRegistration" bson:"userSelfRegistration,omitempty"`
	FarmerSelfRegistration string             `json:"farmerSelfRegistration" bson:"farmerSelfRegistration,omitempty"`
	UserRegistrationTitle  string             `json:"userRegistrationTitle" bson:"userRegistrationTitle,omitempty"`
	ChooseRole             string             `json:"chooseRole" bson:"chooseRole,omitempty"`
	SelectGender           string             `json:"selectGender" bson:"selectGender,omitempty"`
	Male                   string             `json:"male" bson:"male,omitempty"`
	Female                 string             `json:"female" bson:"female,omitempty"`
	Others                 string             `json:"others" bson:"others,omitempty"`
	WeathersData           string             `json:"weathersData" bson:"weathersData,omitempty"`
	OtpAlert               string             `json:"otpAlert" bson:"otpAlert,omitempty"`
	UserNameAlert          string             `json:"userNameAlert" bson:"userNameAlert,omitempty"`
	EmailAlert             string             `json:"emailAlert" bson:"emailAlert,omitempty"`
	MobileAlert            string             `json:"mobileAlert" bson:"mobileAlert,omitempty"`
	ValidateMobileAlert    string             `json:"validateMobileAlert" bson:"validateMobileAlert,omitempty"`
	SelectState            string             `json:"selectState" bson:"selectState,omitempty"`
	SelectDistrict         string             `json:"selectDistrict" bson:"selectDistrict,omitempty"`
	AddBlock               string             `json:"addBlock" bson:"addBlock,omitempty"`
	Save                   string             `json:"save" bson:"save,omitempty"`
	AddGrampanchayat       string             `json:"addGrampanchayat" bson:"addGrampanchayat,omitempty"`
	NameOfGrampanchayat    string             `json:"nameOfGrampanchayat" bson:"nameOfGrampanchayat,omitempty"`
	AddVillage             string             `json:"addVillage" bson:"addVillage,omitempty"`
	Population             string             `json:"population" bson:"population,omitempty"`
	Latitude               string             `json:"latitude" bson:"latitude,omitempty"`
	Longitude              string             `json:"longitude" bson:"longitude,omitempty"`
	CommitteeStatement     string             `json:"committeeStatement" bson:"committeeStatement,omitempty"`
	School                 string             `json:"school" bson:"school,omitempty"`
	Status                 string             `json:"status" form:"status" bson:"status,omitempty"`
	Created                *Created           `json:"created" form:"created" bson:"created,omitempty"`
}
type LanguageTranslationFilter struct {
	ActiveStatus []bool   `json:"activestatus,omitempty" form:"activestatus" bson:"activestatus,omitempty"`
	Status       []string `json:"status" form:"status" bson:"status,omitempty"`
	Version      []string `json:"version,omitempty" form:"version" bson:"version,omitempty"`
	SortBy       string   `json:"sortBy"`
	SortOrder    int      `json:"sortOrder"`
	Regex        struct {
		Name string `json:"name" bson:"name"`
	} `json:"regex" bson:"regex"`
}
type RefLanguageTranslation struct {
	LanguageTranslations `bson:",inline"`
	Ref                  struct {
	} `json:"ref,omitempty" bson:"ref,omitempty"`
}
