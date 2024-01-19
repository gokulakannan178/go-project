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

//Property Master Collections
const (
	COLLECTIONRELATION     = "relations"
	COLLECTIONHONORIFFIC   = "honoriffics"
	COLLECTIONMONTH        = "months"
	COLLECTIONUSERLOCATION = "userslocation"
)
const (
	COLLECTIONGRADE           = "grade"
	COLLECTIONBILLCLAIMCONFIG = "billclaimconfig"
	COLLECTIONBILLCLAIMLEVELS = "billclaimlevels"
)

//User Collections
const (
	COLLECTIONUSER                = "users"
	COLLECTIONORGANISATION        = "organisations"
	COLLECTIONBRANCH              = "branch"
	COLLECTIONDEPARTMENT          = "department"
	COLLECTIONDESIGNATION         = "designation"
	COLLECTIONUSERTYPE            = "usertypes"
	COLLECTIONUSERLOCATIONTRACKER = "userlocationtracker"
	COLLECTIONUSERACL             = "useracl"
)

//Config Collection
const (
	COLLECTIONPROPERTYCONFIGURATION = "propertyconfiguration"
)
const (
	COLLECTIONORGANISATIONCONFIG = "organisationconfig"
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

//Project Collections
const (
	COLLECTIONPROJECT       = "projects"
	COLLECTIONPROJECTMEMBER = "projectmembers"
)

// Task Collections
const (
	COLLECTIONTASK        = "tasks"
	COLLECTIONTASKMEMBER  = "taskmembers"
	COLLECTIONTASKMESSAGE = "taskmessages"
)

// Attendance Collection
const (
	COLLECTIONATTENDANCE    = "attendances"
	COLLECTIONATTENDANCELOG = "attendancelog"
)

//Policy Collection
const (
	COLLECTIONPROBATIONARY = "probationarypolicy"
	COLLECTIONWORKSCHEDULE = "workschedulepolicy"
	COLLECTIONNOTICEPOLICY = "noticepolicy"
)

//Leave Collection
const (
	COLLECTIONLEAVEPOLICY            = "leavepolicy"
	COLLECTIONLEAVEMASTER            = "leavemaster"
	COLLECTIONPOLICYRULE             = "policyrule"
	COLLECTIONPAYROLLPOLICY          = "payrollpolicy"
	COLLECTIONPAYROLLPOLICYEARNING   = "payrollpolicyearning"
	COLLECTIONPAYROLLPOLICYDETECTION = "payrollpolicydetection"
)

//Offboarding Collection
const (
	COLLECTIONOFFBOARDINGCHECKLIST         = "offboardingchecklist"
	COLLECTIONOFFBOARDINGCHECKLISTMASTER   = "offboardingchecklistmaster"
	COLLECTIONOFFBOARDINGPOLICY            = "offboardingpolicy"
	COLLECTIONEMPLOYEEOFFBOARDINGCHECKLIST = "employeeoffboardingchecklist"
)

//Onboarding Collection
const (
	COLLECTIONONBOARDINGCHECKLIST         = "onboardingchecklist"
	COLLECTIONONBOARDINGCHECKLISTMASTER   = "onboardingchecklistmaster"
	COLLECTIONONBOARDINGPOLICY            = "onboardingpolicy"
	COLLECTIONEMPLOYEEONBOARDINGCHECKLIST = "employeeonboardingchecklist"
)
const (
	COLLECTIONEXPENSECATEGORY     = "expensecategory"
	COLLECTIONEXPENSESUBCATEGORY  = "expensesubcategory"
	COLLECTIONEXPENSECATEGORYLIST = "expensecategorylist"
)

//News Collection
const (
	COLLECTIONNEWS        = "news"
	COLLECTIONNEWSLIKE    = "newslike"
	COLLECTIONNEWSCOMMENT = "newscomment"
)

//Document Collection
const (
	COLLECTIONDOCUMENTTYPE            = "DocumentType"
	COLLECTIONDOCUMENTSCENARIO        = "documentscenario"
	COLLECTIONDOCUMENTMUXMASTER       = "documentmuxmaster"
	COLLECTIONDOCUMENTPOLICY          = "documentpolicy"
	COLLECTIONDOCUMENTMASTER          = "documentmaster"
	COLLECTIONDOCUMENTPOLICYDOCUMENTS = "documentpolicydocuments"
)

//Assets Collection
const (
	COLLECTIONASSETPOLICY        = "assetpolicy"
	COLLECTIONASSETMASTER        = "assetmaster"
	COLLECTIONASSETPOLICYASSETS  = "assetpolicyassets"
	COLLECTIONASSET              = "asset"
	COLLECTIONASSETTYPE          = "assettype"
	COLLECTIONASSETTYPEPROPERTYS = "assettypepropertys"
	COLLECTIONASSETPROPERTYS     = "assetpropertys"
	COLLECTIONEMPLOYEEASSETS     = "employeeassets"
	COLLECTIONASSETLOG           = "assetlog"
)

//EMPLOYEE Collection
const (
	COLLECTIONEMPLOYEE                   = "employee"
	COLLECTIONEMPLOYEEJOB                = "employeejob"
	COLLECTIONEMPLOYEEPAYSLIP            = "employeepayslip"
	COLLECTIONEMPLOYEEHISTORY            = "employeehistory"
	COLLECTIONEMERGENCYCONTACT           = "emergecycontact"
	COLLECTIONEMPLOYEELOG                = "employeelog"
	COLLECTIONEMPLOYEELEAVE              = "employeeleave"
	COLLECTIONEMPLOYEELEAVELOG           = "employeeleavelog"
	COLLECTIONJOBTIMELINE                = "jobtimeline"
	COLLECTIONEMPLOYEEEXPERIENCE         = "employeeexperience"
	COLLECTIONEMPLOYEEEDUCATION          = "employeeeducation"
	COLLECTIONEMPLOYEEFAMILYMEMBERS      = "employeefamilymembers"
	COLLECTIONEMPLOYEETIMEOFF            = "employeetimeoff"
	COLLECTIONEMPLOYEESALARY             = "employeesalary"
	COLLECTIONEMPLOYEEATTENDANCECALENDAR = "employeeattendancecalendar"
)
const (
	COLLECTIONSALARYCONFIG    = "salaryconfig"
	COLLECTIONSALARYCONFIGLOG = "salaryconfiglog"
)

//BillClaim Collection
const (
	COLLECTIONBILLCLAIM    = "billclaim"
	COLLECTIONBILLCLAIMLOG = "billclaimlog"
)

// BANK Collection
const (
	COLLECTIONBANKINFORMATION         = "bank"
	COLLECTIONSALARY                  = "salary"
	COLLECTIONPAYMENT                 = "payment"
	COLLECTIONEMPLOYEEEARNINGMASTER   = "employeeearningmaster"
	COLLECTIONEMPLOYEEDEDUCTIONMASTER = "employeedeductionmaster"
	COLLECTIONEMPLOYEEEARNING         = "employeeearning"
	COLLECTIONEMPLOYEEDEDUCTION       = "employeededuction"
	COLLECTIONEMPLOYEEPAYROLL         = "employeepayroll"
)

//SMSLOG Collection
const (
	COLLECTIONSMSLOG = "smslog"
)

//EMAILLOG Collection
const (
	COLLECTIONEMAILLOG = "emaillog"
)

//HOLIDAYS Collection
const (
	COLLECTIONHOLIDAYS = "holidays"
)

//APPTOKEN Collection
const (
	COLLECTIONAPPTOKEN = "apptoken"
)
const (
	COLLECTIONDEMOUSER = "demouser"
)
const (
	COLLECTIONPAYROLL = "payroll"
)
const (
	COLLECTIONPAYROLLLOG = "payrolllog"
)
