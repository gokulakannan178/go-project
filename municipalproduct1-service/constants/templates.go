package constants

const (

	// PROPERTYDEMANDSMSTOCUSTOMER = "Hi <owner name>,You have received a notification from <Bihar municipal corporation> regarding <Property tax demand>\n
	// Notification - <content>\nRegards,\nFrom BRMNCP\nFor help/queries contact <config>"
	COMMONSMSTEMPLATE        = "Hi %v,You have received a notification from %v regarding %v\n Notification - %v\nRegards,\nFrom %v\nFor help/queries contact %v"
	PROPERTYTAXDEMANDCONTENT = "A sum of %v is demand on the property with Holding no %v on the Financial year %v"
	// Bihar municipal corporation - Demand Pending for <UniqueID>

	PROPERTYDEMANDEMAILSUBJECTTOCUSTOMER = "Bihar Municipal Corporation - Demand Pending for (%v)"
	COMMONEMAILBODYTOCUSTOMER            = "Hi %v,You have received a notification from %v regarding %v\n Notification - %v. Please visit citizens counter to pay. \nRegards,\nFrom %v\nFor help/queries contact %v"
	// COMMONEMAILBODYTOCUSTOMER = "Hi %v,You have received a notification from %v regarding %v\n Notification - A sum of %v is demand on the property with Holding no %v on the Financial year %v. \nPlease visit citizens counter to pay. \nRegards,\nFrom %v\nFor help/queries contact %v"
	COMMONWHATSAPPTEMPLATE = "Hi %v,You have received a notification from %v regarding %v. Notification - %v. Download Demand PDF - %v"
)
