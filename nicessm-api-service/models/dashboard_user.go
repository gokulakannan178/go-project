package models

type DashboardUserCountFilter struct {
	UserFilter `bson:",inline"`
}
type DashboardUserCountReport struct {
	ContentCreator      int `json:"contentcreator" bson:"contentcreator,omitempty"`
	SelfRegistration    int `json:"selfRegistration" bson:"selfRegistration,omitempty"`
	CallCenterAgent     int `json:"callCenterAgent" bson:"callCenterAgent,omitempty"`
	ContentManager      int `json:"contentManager" bson:"contentManager,omitempty"`
	ContentProvider     int `json:"contentProvider" bson:"contentProvider,omitempty"`
	ContentDisseminator int `json:"contentDisseminator" bson:"contentDisseminator,omitempty"`
	FieldAgent          int `json:"fieldAgent" bson:"fieldAgent,omitempty"`
	LanguageTranslator  int `json:"languageTranslator" bson:"languageTranslator,omitempty"`
	LanguageApprover    int `json:"languageApprover" bson:"languageApprover,omitempty"`
	Management          int `json:"management" bson:"management,omitempty"`
	Moderator           int `json:"moderator" bson:"moderator,omitempty"`
	SubjectMatterExpert int `json:"subjectMatterExpert" bson:"subjectMatterExpert,omitempty"`
	SystemAdmin         int `json:"systemAdmin" bson:"systemAdmin,omitempty"`
	VistorViewer        int `json:"vistorViewer" bson:"vistorViewer,omitempty"`
	Trainer             int `json:"trainer" bson:"trainer,omitempty"`
	FieldAgentLead      int `json:"fieldAgentLead" bson:"fieldAgentLead,omitempty"`
	GuestUser           int `json:"guestUser" bson:"guestUser,omitempty"`
	SuperAdmin          int `json:"superAdmin" bson:"superAdmin,omitempty"`
	DistrictAdmin       int `json:"districtAdmin" bson:"districtAdmin,omitempty"`
	AllUser             int `json:"allUser" bson:"allUser,omitempty"`
}
type DayWiseUserDemandChartReport struct {
	//ID   bool `json:"id" bson:"_id,omitempty"`
	Days []struct {
		ID   int `json:"id" bson:"_id,omitempty"`
		Data struct {
			Active   float64 `json:"active" bson:"Active,omitempty"`
			Disabled float64 `json:"disabled" bson:"Disabled,omitempty"`
		} `json:"data" bson:"data"`
	} `json:"days,omitempty" bson:"days"`
}
