package constants

//Constants for collection names
const (
	COLLCOUNTER  = "collectioncounter"
	COLLREGISTER = "collectionregister"
)

//Geolocation Collections
const (
	COLLECTIONLANGUAGE = "languages"
	COLLECTIONSTATE    = "states"
	COLLECTIONDISTRICT = "districts"
	COLLECTIONVILLAGE  = "villages"
	COLLECTIONZONE     = "zones"
	COLLECTIONWARD     = "wards"
)
const (
	COLLECTIONSTOREDCALC         = "storedcalcs"
	COLLECTIONSTOREDCALCDEMANDFY = "storedcalculationdemandfy"
	COLLECTIONSTOREDCALCDEMAND   = "storedcalculationdemand"
)

//Property Master Collections
const (
	COLLECTIONCONSTRUCTIONTYPE                       = "constructiontypes"
	COLLECTIONFLOORTYPE                              = "floortypes"
	COLLECTIONHONORIFFIC                             = "honoriffics"
	COLLECTIONMUNICIPALTYPES                         = "municipaltypes"
	COLLECTIONOCCUMANCYTYPE                          = "occupancytypes"
	COLLECTIONRELATION                               = "relations"
	COLLECTIONROADTYPE                               = "roadtypes"
	COLLECTIONUSAGETYPE                              = "usagetypes"
	COLLECTIONPROPERTYTYPE                           = "propertytypes"
	COLLECTIONAVRRANGE                               = "avrranges"
	COLLECTIONPROPERTYOTHERTAX                       = "propertyothertax"
	COLLECTIONFINANCIALYEAR                          = "financialyears"
	COLLECTIONOWNERSHIP                              = "ownerships"
	COLLECTIONVACANTLANDRATE                         = "vacantlandrates"
	COLLECTIONFLOORRATABLEAREA                       = "floorratableareas"
	COLLECTIONAVR                                    = "avr"
	COLLECTIONNONRESIDENTIALUSAGEFACTOR              = "nonresidentialusagefactors"
	COLLECTIONPROPERTYTAX                            = "propertytax"
	COLLECTIONREBATE                                 = "rebate"
	COLLECTIONPENALTY                                = "penalty"
	COLLECTIONRESIDENTIALTYPE                        = "residentialtype"
	COLLECTIONPROPERTYPAYMENT                        = "propertypayments"
	COLLECTIONPROPERTYPARTPAYMENT                    = "propertypartpayments"
	COLLECTIONPROPERTYPARTPAYMENTSFY                 = "propertypartpaymentsfy"
	COLLECTIONPROPERTYPARTPAYMENTSBASIC              = "propertypartpaymentsbasic"
	COLLECTIONPROPERTYPAYMENTBASIC                   = "propertypaymentbasics"
	COLLECTIONPROPERTYPAYMENTFY                      = "propertypaymentfys"
	COLLECTIONMONTH                                  = "months"
	COLLECTIONUSERLOCATION                           = "userslocation"
	COLLECTIONOTHERCHARGES                           = "othercharges"
	COLLECTIONTCLEDGER                               = "tcledger"
	COLLECTIONBANKDEPOSIT                            = "bankdeposits"
	COLLECTIONBASICPROPERTYUPDATELOG                 = "basicpropertyupdatelog"
	COLLECTIONBANKDEPOSITCOLLECTIONSUBMISSIONREQUEST = "bankdepositcollectionsubmissionrequest"
	COLLECTIONPENALCHARGES                           = "penalcharges"
	COLLECTIONPENALCHARGEFYRANGE                     = "penalchargefyrange"
)

//Property Collections
const (
	COLLECTIONPROPERTY                 = "properties"
	COLLECTIONPROPERTYFLOOR            = "floors"
	COLLECTIONPROPERTYOWNER            = "owners"
	COLLECTIONPROPERTYDEMANDLOG        = "propertydemandlogs"
	COLLECTIONPROPERTYFYDEMANDLOG      = "propertyfydemandlogs"
	COLLECTIONOVERALLPROPERTYDEMAND    = "overallpropertydemand"
	COLLECTIONOSTOREDPROPERTYDEMANDFYS = "storedpropertydemandfys"
)

//User Collections
const (
	COLLECTIONUSER         = "users"
	COLLECTIONORGANISATION = "organisations"
	COLLECTIONDUSERTYPE    = "usertypes"
)

//Config Collection
const (
	COLLECTIONPROPERTYCONFIGURATION = "propertyconfiguration"
)

//ACL Collections
const (
	COLLECTIONMODULE             = "aclmastermodules"
	COLLECTIONMENU               = "aclmastermenus"
	COLLECTIONTAB                = "aclmastertabs"
	COLLECTIONFEATURE            = "aclmasterfeatures"
	COLLECTIONACLUSERTYPEMODULE  = "aclmasterusetypemodules"
	COLLECTIONACLUSERTYPEMENU    = "aclmasterusetypemenus"
	COLLECTIONACLUSERTYPETAB     = "aclmasterusetypetabs"
	COLLECTIONACLUSERTYPEFEATURE = "aclmasterusetypefeatures"
)

// SurveyAndTax
const (
	COLLECTIONSURVEYANDTAX = "surveyandtax"
)

// ProductConfiguration
const (
	COLLECTIONPRODUCTCONFIGURATION = "productconfiguration"
)

// contactus
const (
	COLLECTIONCONTACTUS = "contactus"

	CONTACTUSSTATUSACTIVE   = "Active"
	CONTACTUSSTATUSDISABLED = "Disabled"
	CONTACTUSSTATUSDELETED  = "Deleted"
)

//legacy
const (
	COLLECTIONLEGACY     = "legacy"
	COLLECTIONLEGACYYEAR = "legacyyears"
)

// Property Mobile Tower
const (
	COLLECTIONPROPERTYMOBILETOWER = "propertymobiletowers"
)

// Mobile Tower
const (
	COLLECTIONMOBILETOWER          = "mobiletowers"
	COLLECTIONMOBILETOWERYEAR      = "mobiletoweryears"
	COLLECTIONMOBILETOWERUPDATELOG = "mobiletowersupdatelog"
)

// Mobile Tower Tax
const (
	COLLECTIONMOBILETOWERTAX             = "mobiletowertaxs"
	COLLECTIONMOBILETOWERREGISTRATIONTAX = "mobiletowerregistrationtaxs"
)

//MobileTowerRegistrationRateMaster
const (
	COLLECTIONMOBILETOWERREGISTRATIONRATEMASTER = "mobiletowerregistrationratemaster"
)

//Mobile Tower Payments
const (
	COLLECTIONMOBILETOWERPAYMENTS      = "mobiletowerpayments"
	COLLECTIONMOBILETOWERPAYMENTSFY    = "mobiletowerpaymentsfy"
	COLLECTIONMOBILETOWERPAYMENTSBASIC = "mobiletowerpaymentsbasic"
)

//Shop Rent Payments
const (
	COLLECTIONSHOPRENTPAYMENTS        = "shoprentpayments"
	COLLECTIONSHOPRENTPAYMENTSRECEIPT = "shoprentpaymentreceipts"
	COLLECTIONSHOPRENTPAYMENTSFY      = "shoprentpaymentsfy"
	COLLECTIONSHOPRENTPAYMENTSBASIC   = "shoprentpaymentsbasic"
	COLLECTIONSHOPRENTPAYMENTSMONTH   = "shoprentpaymentsmonth"
)

// Property Penalty
const (
	COLLECTIONPROPERTYPENALTY = "propertypenalties"
)

// Citizen Gravians
const (
	COLLECTIONCITIZENGRAVIANS = "citizengravians"
)

// Letter Generate
const (
	COLLECTIONLETTERGENERATE = "lettergenerate"
)

//Letter UploaD
const (
	COLLECTIONLETTERUPLOAD = "letterupload"
)

//Document Collections
const (
	COLLECTIONDOCUMENT                 = "documents"
	COLLECTIONDOCUMENTLIST             = "documentlist"
	COLLECTIONPROPERTYDOCUMENT         = "propertydocuments"
	COLLECTIONPROPERTYREQUIREDDOCUMENT = "requireddocuments"
)

// Basic Property Update Collections
const (
	BASICPROPERTYUPDATESTATUSACTIVE   = "Active"
	BASICPROPERTYUPDATESTATUSDISABLED = "Disabled"
	BASICPROPERTYUPDATESTATUSDELETED  = "Deleted"
)

//Payment Gateway
const (
	COLLECTIONPAYMENTGATEWAY     = "paymentgateway"
	PAYMENTGATEWAYSTATUSACTIVE   = "Active"
	PAYMENTGATEWAYSTATUSDISABLED = "Disabled"
	PAYMENTGATEWAYSTATUSDELETED  = "Deleted"
)

//GSTRateMaster
const (
	COLLECTIONGSTRATEMASTER     = "gstratemaster"
	GSTRATEMASTERSTATUSACTIVE   = "Active"
	GSTRATEMASTERSTATUSDISABLED = "Disabled"
	GSTRATEMASTERSTATUSDELETED  = "Deleted"
)

//LeaseRentShopCategory
const (
	COLLECTIONLEASERENTSHOPCATEGORY     = "leaserentshopcategory"
	LEASERENTSHOPCATEGORYSTATUSACTIVE   = "Active"
	LEASERENTSHOPCATEGORYSTATUSDISABLED = "Disabled"
	LEASERENTSHOPCATEGORYSTATUSDELETED  = "Deleted"
)

//LeaseRentShopSubCategory
const (
	COLLECTIONLEASERENTSHOPSUBCATEGORY     = "leaserentshopsubcategory"
	LEASERENTSHOPSUBCATEGORYSTATUSACTIVE   = "Active"
	LEASERENTSHOPSUBCATEGORYSTATUSDISABLED = "Disabled"
	LEASERENTSHOPSUBCATEGORYSTATUSDELETED  = "Deleted"
)

//LeaseRentRateMaster
const (
	COLLECTIONLEASERENTRATEMASTER     = "leaserentratemaster"
	LEASERENTRATEMASTERSTATUSACTIVE   = "Active"
	LEASERENTRATEMASTERSTATUSDISABLED = "Disabled"
	LEASERENTRATEMASTERSTATUSDELETED  = "Deleted"
)

//LeaseRent
const (
	COLLECTIONLEASERENT     = "leaserent"
	LEASERENTSTATUSACTIVE   = "Active"
	LEASERENTSTATUSDISABLED = "Disabled"
	LEASERENTSTATUSDELETED  = "Deleted"
)

//ShopRentShopCategory
const (
	COLLECTIONSHOPRENTSHOPCATEGORY     = "shoprentshopcategory"
	SHOPRENTSHOPCATEGORYSTATUSACTIVE   = "Active"
	SHOPRENTSHOPCATEGORYSTATUSDISABLED = "Disabled"
	SHOPRENTSHOPCATEGORYSTATUSDELETED  = "Deleted"
)

//ShopRentShopSubCategory
const (
	COLLECTIONSHOPRENTSHOPSUBCATEGORY     = "shoprentshopsubcategory"
	SHOPRENTSHOPSUBCATEGORYSTATUSACTIVE   = "Active"
	SHOPRENTSHOPSUBCATEGORYSTATUSDISABLED = "Disabled"
	SHOPRENTSHOPSUBCATEGORYSTATUSDELETED  = "Deleted"
)

//ShopRentRateMaster
const (
	COLLECTIONSHOPRENTRATEMASTER     = "shoprentratemaster"
	SHOPRENTRATEMASTERSTATUSACTIVE   = "Active"
	SHOPRENTRATEMASTERSTATUSDISABLED = "Disabled"
	SHOPRENTRATEMASTERSTATUSDELETED  = "Deleted"
)

// Shoprentindividualratemaster
const (
	COLLECTIONSHOPRENTINDIVIDUALRATEMASTER = "shoprentindividualratemaster"
)

//ShopRent
const (
	COLLECTIONSHOPRENT     = "shoprent"
	SHOPRENTSTATUSACTIVE   = "Active"
	SHOPRENTSTATUSDISABLED = "Disabled"
	SHOPRENTSTATUSDELETED  = "Deleted"
	SHOPRENTSTATUSREJECTED = "Rejected"
	SHOPRENTSTATUSPENDING  = "Pending"
)

//PropertyVisitLogRemarkType
const (
	COLLECTIONPROPERTYVISITLOGREMARKTYPE     = "propertyvisitlogremarktype"
	PROPERTYVISITLOGREMARKTYPESTATUSACTIVE   = "Active"
	PROPERTYVISITLOGREMARKTYPESTATUSDISABLED = "Disabled"
	PROPERTYVISITLOGREMARKTYPESTATUSDELETED  = "Deleted"
)

//PropertyVisitLog
const (
	COLLECTIONPROPERTYVISITLOG     = "propertyvisitlog"
	PROPERTYVISITLOGSTATUSACTIVE   = "Active"
	PROPERTYVISITLOGSTATUSDISABLED = "Disabled"
	PROPERTYVISITLOGSTATUSDELETED  = "Deleted"
)

// User Attendance
const (
	COLLECTIONUSERATTENDANCE = "userattendance"
)

// Dashboard Mobile Tower
const (
	COLLECTIONOVERALLDASHBOARD = "dashboardoverall"
)

// Dashboard Mobile Tower
const (
	COLLECTIONDASHBOARDMOBILETOWER = "dashboardmobiletower"
)

// Dashboard Shop Rent
const (
	COLLECTIONDASHBOARDSHOPRENT = "dashboardshoprent"
)

// Dashboard Property
const (
	COLLECTIONDASHBOARDPROPERTY = "dashboardproperty"
)

// Dashboard Lease
const (
	COLLECTIONDASHBOARDLEASE = "dashboardlease"
)

// Dashboard Trade LIcense
const (
	COLLECTIONDASHBOARDTRADELICENSE      = "dashboardtradelicense"
	COLLECTIONBASICTRADELICENSEUPDATELOG = "basictradelicenseupdatelog"
)

// Dashboard User Charge
const (
	COLLECTIONDASHBOARDUSERCHARGE = "dashboardusercharge"
)

// Dashboard Water Bill
const (
	COLLECTIONDASHBOARDWATERBILL = "dashboardwaterbill"
)

// Dashboard Mobile Tower Day Wise
const (
	COLLECTIONDASHBOARDMOBILETOWERDAYWISE = "dashboardmobiletowerdaywise"
)

// Dashboard Property Day Wise
const (
	COLLECTIONDASHBOARDPROPERTYDAYWISE = "dashboardpropertydaywise"
)

// Dashboard Lease Day Wise
const (
	COLLECTIONDASHBOARDLEASEDAYWISE = "dashboardleasedaywise"
)

// Dashboard Trade License Day Wise
const (
	COLLECTIONDASHBOARDTRADELICENSEDAYWISE = "dashboardtradelicensedaywise"
)

// Dashboard User Charge Day Wise
const (
	COLLECTIONDASHBOARDUSERCHARGEDAYWISE = "dashboarduserchargedaywise"
)

// Dashboard Water Bill Day Wise
const (
	COLLECTIONDASHBOARDWATERBILLDAYWISE = "dashboardwaterbilldaywise"
)

// Dashboard Shop Rent Day Wise
const (
	COLLECTIONDASHBOARDSHOPRENTDAYWISE = "dashboardshoprentdaywise"
)

// Monthly Target
const (
	COLLECTIONMONTHLYTARGET = "monthlytarget"
)

// User Ward Access
const (
	COLLECTIONUSERWARDACCESS = "userwardaccess"
)

// User Zone Access
const (
	COLLECTIONUSERZONEACCESS = "userzoneaccess"
)

// Trade License
const (
	COLLECTIONTRADELICENSE = "tradelicense"
)

// Trade License Rate Master
const (
	COLLECTIONTRADELICENSERATEMASTER = "tradelicenseratemaster"
)

// Trade License Category Type
const (
	COLLECTIONTRADELICENSECATEGORYTYPE = "tradelicensecategorytype"
)

// Trade License Business Type
const (
	COLLECTIONTRADELICENSEBUSINESSTYPE = "tradelicensebusinesstype"
)

// Trade License Payments
const (
	COLLECTIONTRADELICENSEPAYMENTS        = "tradelicensepayments"
	COLLECTIONTRADELICENSEPAYMENTSRECEIPT = "tradelicensepaymentreceipt"
	COLLECTIONTRADELICENSEPAYMENTSFY      = "tradelicensepaymentsfy"
	COLLECTIONTRADELICENSEPAYMENTSBASIC   = "tradelicensepaymentsbasic"
)

// Pm target Collection
const (
	COLLECTIONPMTARGET = "pmtarget"
)

// Pm Achievement Collection
const (
	COLLECTIONPMACHIEVEMENT = "pmachievement"
)

// Property Wallet Collection
const (
	COLLECTIONPROPERTYWALLET = "propertywallet"
)

// Property Wallet Log Collection
const (
	COLLECTIONPROPERTYWALLETLOG = "propertywalletlog"
)

// Dashboard Daily Log
const (
	COLLECTIONDAILYLOG = "dailylogs"
)

// MobileTower Receipt No
const (
	COLLECTIONMOBILETOWERPAYMENTSRECEIPT = "mobiletowerpaymentreceipt"
)

// Reassessment Request Collection
const (
	COLLECTIONREASSESSMENTREQUEST = "reassessmentrequest"
)

// ShopRent Reassessment Request
const (
	COLLECTIONSHOPRENTREASSESSMENTREQUEST = "shoprentreassessmentrequest"
)

// Trade License Reassessment Request
const (
	COLLECTIONTRADELICENSEREASSESSMENTREQUEST = "tradelicensereassessmentrequest"
)

// MobileTower Reassessment Request
const (
	COLLECTIONMOBILETOWERREASSESSMENTREQUEST = "mobiletowerreassessmentrequest"
)

// JObs Collection
const (
	COLLECTIONJOB = "jobs"
)

// Job Logs Collection
const (
	COLLECTIONJOBLOG = "joblogs"
)

// Major Update Collection
const (
	COLLECTIONMAJORUPDATE = "majorupdates"
)

// Cron Logs Collection
const (
	COLLECTIONCRONLOG = "cronlogs"
)
const (
	COLLECTIONCITIZENGRAVIANSLOG = "citizengravianslog"
)

// ticket Collection
const (
	COLLECTIONTICKET        = "tickets"
	COLLECTIONTICKETCOMMENT = "ticketcomments"
	COLLECTIONTICKETUSER    = "ticketusers"
)

//PayTmQrCode :"Used to create Dynamic Qr Code"
const (
	COLLECTIONPAYMENTQRCODE = "paymentqrcode"
)

// User Location Tracker
const (
	COLLECTIONUSERLOCATIONTRACKER = "userlocationtracker"
)

// User Location Log
const (
	COLLECTIONUSERLOCATIONLOG = "userlocationlog"
)

//
const (
	COLLECTIONTRADELICENSEREBATE = "tradelicenserebate"
)

// Trade License Payments Part2
const (
	COLLECTIONTRADELICENSEPART2                = "tradelicensepart2"
	COLLECTIONTRADELICENSEPAYMENTSPART2        = "tradelicensepaymentspart2"
	COLLECTIONTRADELICENSEPARTPAYMENTPART2     = "tradelicensepartpaymentspart2"
	COLLECTIONTRADELICENSEPAYMENTSRECEIPTPART2 = "tradelicensepaymentreceiptpart2"
	COLLECTIONTRADELICENSEPAYMENTSFYPART2      = "tradelicensepaymentsfypart2"
	COLLECTIONTRADELICENSEPAYMENTSBASICPART2   = "tradelicensepaymentsbasicpart2"
)

// Property Payment Mode Change
const (
	COLLECTIONPROPERTYPAYMENTMODECHANGE     = "propertypaymentmodechange"
	COLLECTIONUSERCHARGEPAYMENTMODECHANGE   = "userchargepaymentmodechange"
	COLLECTIONTRADELICENSEPAYMENTMODECHANGE = "tradelicensepaymentmodechange"
	COLLECTIONSHOPRENTPAYMENTMODECHANGE     = "shoprentpaymentmodechange"
)

// PropertypayeenameChange
const (
	COLLECTIONPROPERTYPAYEENAMEHANGE     = "propertypayeenamechange"
	COLLECTIONUSERCHARGEPAYEENAMEHANGE   = "userchargepayeenamechange"
	COLLECTIONTRADELICENSEPAYEENAMEHANGE = "tradelicensepayeenamechange"
	COLLECTIONSHOPRENTPAYEENAMEHANGE     = "shoprentpayeenamechange"
)

// Property Estimated Demand
const (
	COLLECTIONESTIMATEDPROPERTYDEMAND = "estimatedpropertydemand"
	COLLECTIONESTIMATEDPROPERTYFLOOR  = "estimatedpropertyfloor"
)

// Property Mutation Request Collection
const (
	COLLECTIONPROPERTYMUTATIONREQUEST = "propertymutationrequest"
)

// Property Delete Request Collection
const (
	COLLECTIONPROPERTYDELETEREQUEST     = "propertydeleterequest"
	COLLECTIONTRADELICENSEDELETEREQUEST = "tradelicensedeleterequest"
	COLLECTIONSHOPRENTDELETEREQUEST     = "shoprentdeleterequest"
	COLLECTIONMOBILETOWERDELETEREQUEST  = "mobiletowerdeleterequest"
)

// Property Mutation
const (
	COLLECTIONMUTATEDPROPERTY = "mutatedproperties"
)

// Boring Charges
const (
	COLLECTIONBORINGCHARGES = "boringcharges"
)

//legacy
const (
	COLLECTIONLEGACYV2     = "legacyv2"
	COLLECTIONLEGACYYEARV2 = "legacyyearsv2"
)

const (
	COLLECTIONSOLIDWASTEUSERCHARGE            = "solidwasteusercharge"
	COLLECTIONSOLIDWASTEUSERCHARGERATE        = "solidwasteuserchargerate"
	COLLECTIONSOLIDWASTEUSERCHARGECATEGORY    = "solidwasteuserchargecategory"
	COLLECTIONSOLIDWASTEUSERCHARGESUBCATEGORY = "solidwasteuserchargesubcategory"
)

// Solid Waste Reassessment Request
const (
	COLLECTIONSOLIDWASTEREASSESSMENTREQUEST = "solidwastereassessmentrequest"
)

// Solid Waste User Charge Payments
const (
	COLLECTIONSOLIDWASTEUSERCHARGEPAYMENTS        = "solidwasteuserchargepayments"
	COLLECTIONSOLIDWASTEUSERCHARGEPAYMENTSRECEIPT = "solidwasteuserchargepaymentsreceipts"
	COLLECTIONSOLIDWASTEUSERCHARGEPAYMENTSFY      = "solidwasteuserchargepaymentsfy"
	COLLECTIONSOLIDWASTEUSERCHARGEPAYMENTSBASIC   = "solidwasteuserchargepaymentsbasic"
	COLLECTIONSOLIDWASTEUSERCHARGEPAYMENTSMONTH   = "solidwasteuserchargepaymentsmonth"
)

const (
	COLLECTIONPROPERTYOTHERDEMAND = "propertyotherdemand"
)

// Property Other Demand Payments
const (
	COLLECTIONPROPERTYOTHERDEMANDPAYMENT           = "propertyotherdemandpayments"
	COLLECTIONPROPERTYOTHERDEMANDPARTPAYMENT       = "propertyotherdemandpartpayments"
	COLLECTIONPROPERTYOTHERDEMANDPAYMENTSRECEIPT   = "propertyotherdemandpaymentsreceipts"
	COLLECTIONPROPERTYOTHERDEMANDPAYMENTSFY        = "propertyotherdemandpaymentspaymentsfy"
	COLLECTIONPROPERTYOTHERDEMANDPAYMENTSBASIC     = "propertyotherdemandpaymentsbasic"
	COLLECTIONPROPERTYOTHERDEMANDPARTPAYMENTSBASIC = "propertyotherdemandpartpaymentsbasic"
)

//watertax
const (
	COLLECTIONWATERTAXARV            = "watertaxarv"
	COLLECTIONWATERTAXCONNECTIONTYPE = "watertaxconnectiontype"
)

//birthcertificate
const (
	COLLECTIONBIRTHCERTIFICATE = "birthcertificate"
	COLLECTIONHOSPITAL         = "hospital"
)

const (
	COLLECTIONPROPERTYFIXEDARV    = "propertyfixedarv"
	COLLECTIONPROPERTYFIXEDARVLOG = "propertyfixedarvlog"
	COLLECTIONPROPERTYFIXEDDEMAND = "propertyfixeddemand"
)

const (
	COLLECTIONCOMPOSITETAXRATEMASTER = "compositetaxratemaster"
)

// Parking Penalties
const (
	COLLECTIONPARKINGPENALTIES    = "parkingpenalties"
	COLLECTIONPARKINGPENALTIESLOG = "parkingpenaltieslog"
)

// HDFCPaymentGateway
const (
	COLLECTIONHDFCPAYMENTGATEWAY = "hdfcpaymentgateway"
)

const (
	COLLECTIONUSERCHARGECATEGORY        = "userchargecategory"
	COLLECTIONUSERCHARGEUPDATELOG       = "userchargeupdatelog"
	COLLECTIONUSERCHARGELOG             = "userchargelog"
	COLLECTIONPROPERTYUSERCHARGE        = "propertyusercharge"
	COLLECTIONUSERCHARGERATEMASTER      = "userchargeratemaster"
	COLLECTIONUSERCHARGEPAYMENTS        = "userchargepayments"
	COLLECTIONUSERCHARGEPAYMENTSRECEIPT = "userchargepaymentreceipts"
	COLLECTIONUSERCHARGEPAYMENTSFY      = "userchargepaymentfy"
	COLLECTIONUSERCHARGEPAYMENTBASICS   = "userchargepaymentbasics"
)
