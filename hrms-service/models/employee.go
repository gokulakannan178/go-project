package models

import "time"

type Employee struct {
	User `bson:",inline"`
	//UniqueID            string     `json:"uniqueId,omitempty" bson:"uniqueId,omitempty"`
	ReOnboardingDate    *time.Time             `json:"reonboardingDate,omitempty" bson:"reonboardingDate,omitempty"`
	OnboardingpolicyID  string                 `json:"onboardingpolicyId,omitempty" bson:"onboardingpolicyId,omitempty"`
	ProbationaryID      string                 `json:"probationaryId" bson:"probationaryId,omitempty"`
	ProbationaryendDate time.Time              `json:"probationaryendDate" bson:"probationaryendDate,omitempty"`
	RejectDate          *time.Time             `json:"rejectDate,omitempty" bson:"rejectDate,omitempty"`
	ConfirmationDate    *time.Time             `json:"confirmationDate,omitempty" bson:"confirmationDate,omitempty"`
	BenchDate           *time.Time             `json:"benchDate,omitempty" bson:"benchDate,omitempty"`
	NoticeID            string                 `json:"noticeId,omitempty" bson:"noticeId,omitempty"`
	NoticeendDate       time.Time              `json:"noticeEndDate,omitempty" bson:"noticeEndDate,omitempty"`
	OffboardingpolicyID string                 `json:"offboardingpolicyId,omitempty" bson:"offboardingpolicyId,omitempty"`
	OffBoardDate        *time.Time             `json:"offBoardDate,omitempty" bson:"offBoardDate,omitempty"`
	RelieveDate         *time.Time             `json:"relieveDate,omitempty" bson:"relieveDate,omitempty"`
	Office              string                 `json:"office,omitempty" bson:"office,omitempty"`
	OrganisationID      string                 `json:"organisationId,omitempty" bson:"organisationId,omitempty"`
	DocumentPolicyID    string                 `json:"documentPolicyID,omitempty" bson:"documentPolicyID,omitempty"`
	DepartmentID        string                 `json:"departmentId,omitempty" bson:"departmentId,omitempty"`
	BranchID            string                 `json:"branchId,omitempty" bson:"branchId,omitempty"`
	DesignationID       string                 `json:"designationId,omitempty" bson:"designationId,omitempty"`
	Remark              string                 `json:"remark,omitempty" bson:"remark,omitempty"`
	Created             *Created               `json:"createdOn" bson:"createdOn,omitempty"`
	WorkScheduleID      string                 `json:"workscheduleId,omitempty" bson:"workscheduleId,omitempty"`
	Updated             Updated                `json:"updated" form:"id," bson:"updated,omitempty"`
	UpdateLog           []Updated              `json:"updatedLog" form:"id," bson:"updatedLog,omitempty"`
	PolicyRuleID        string                 `json:"policyruleId,omitempty" bson:"policyruleId,omitempty"`
	LeavePolicyID       string                 `json:"leavePolicyID,omitempty" bson:"leavePolicyID,omitempty"`
	PayrollPolicyId     string                 `json:"payrollPolicyId,omitempty" bson:"payrollPolicyId,omitempty"`
	ProfileImg          string                 `json:"profileImg" bson:"profileImg,omitempty"`
	YearOfJoining       float64                `json:"yearOfJoining" bson:"yearOfJoining,omitempty"`
	Title               string                 `json:"title" bson:"-"`
	Image               string                 `json:"image" bson:"-"`
	EmergencyContact    UpdateEmergencyContact `json:"emergencyContact,omitempty" bson:"emergencyContact,omitempty"`
	PersonalInformation PersonalInformation    `json:"personalInformation,omitempty" bson:"personalInformation,omitempty"`
}

type EmployeeMoveToReject struct {
	EmployeeID string     `json:"employeeId,omitempty" bson:"employeeId,omitempty"`
	Remark     string     `json:"remark,omitempty" bson:"remark,omitempty"`
	By         string     `json:"by" bson:"by,omitempty"`
	ByID       string     `json:"byId" bson:"byId,omitempty"`
	ByType     string     `json:"byType" bson:"byType,omitempty"`
	Status     string     `json:"status,omitempty" bson:"status,omitempty"`
	RejectDate *time.Time `json:"rejectDate,omitempty" bson:"rejectDate,omitempty"`
}

type EmployeeMoveToOnboarding struct {
	EmployeeID string `json:"employeeId,omitempty" bson:"employeeId,omitempty"`
	Remark     string `json:"remark,omitempty" bson:"remark,omitempty"`
	By         string `json:"by" bson:"by,omitempty"`
	ByID       string `json:"byId" bson:"byId,omitempty"`
	ByType     string `json:"byType" bson:"byType,omitempty"`
	Status     string `json:"status,omitempty" bson:"status,omitempty"`
}

type EmployeeMoveToProbationary struct {
	EmployeeID       string `json:"employeeId,omitempty" bson:"employeeId,omitempty"`
	BranchID         string `json:"branchId,omitempty" bson:"branchId,omitempty"`
	DepartmentID     string `json:"departmentId,omitempty" bson:"departmentId,omitempty"`
	DesignationID    string `json:"designationId,omitempty" bson:"designationId,omitempty"`
	ProbationaryID   string `json:"probationaryId,omitempty" bson:"probationaryId,omitempty"`
	WorkScheduleID   string `json:"workscheduleId,omitempty" bson:"workscheduleId,omitempty"`
	DocumentPolicyID string `json:"documentPolicyID,omitempty" bson:"documentPolicyID,omitempty"`
	LeavePolicyID    string `json:"leavePolicyID,omitempty" bson:"leavePolicyID,omitempty"`
	OfficialEmail    string `json:"officialEmail" bson:"officialEmail,omitempty"`
	PayrollPolicyId  string `json:"payrollPolicyId,omitempty" bson:"payrollPolicyId,omitempty"`
	LineManager      string `json:"lineManager" bson:"lineManager,omitempty"`
	NoticeID         string `json:"noticeId,omitempty" bson:"noticeId,omitempty"`
	Remark           string `json:"remark,omitempty" bson:"remark,omitempty"`
	By               string `json:"by" bson:"by,omitempty"`
	ByID             string `json:"byId" bson:"byId,omitempty"`
	ByType           string `json:"byType" bson:"byType,omitempty"`
}

type EmployeeMoveToActive struct {
	EmployeeID string `json:"employeeId,omitempty" bson:"employeeId,omitempty"`
	Remark     string `json:"remark,omitempty" bson:"remark,omitempty"`
	By         string `json:"by" bson:"by,omitempty"`
	ByID       string `json:"byId" bson:"byId,omitempty"`
	ByType     string `json:"byType" bson:"byType,omitempty"`
}

type EmployeeMoveToBench struct {
	EmployeeID string `json:"employeeId,omitempty" bson:"employeeId,omitempty"`
	Remark     string `json:"remark,omitempty" bson:"remark,omitempty"`
	By         string `json:"by" bson:"by,omitempty"`
	ByID       string `json:"byId" bson:"byId,omitempty"`
	ByType     string `json:"byType" bson:"byType,omitempty"`
}

type EmployeeMoveToNotice struct {
	EmployeeID string `json:"employeeId,omitempty" bson:"employeeId,omitempty"`
	NoticeID   string `json:"noticeId,omitempty" bson:"noticeId,omitempty"`
	Remark     string `json:"remark,omitempty" bson:"remark,omitempty"`
	By         string `json:"by" bson:"by,omitempty"`
	ByID       string `json:"byId" bson:"byId,omitempty"`
	ByType     string `json:"byType" bson:"byType,omitempty"`
}

type EmployeeMoveToOffboard struct {
	EmployeeID          string `json:"employeeId,omitempty" bson:"employeeId,omitempty"`
	OffboardingpolicyId string `json:"offboardingpolicyId,omitempty" bson:"offboardingpolicyId,omitempty"`
	Remark              string `json:"remark,omitempty" bson:"remark,omitempty"`
	By                  string `json:"by" bson:"by,omitempty"`
	ByID                string `json:"byId" bson:"byId,omitempty"`
	ByType              string `json:"byType" bson:"byType,omitempty"`
}

type EmployeeMoveToRelieve struct {
	EmployeeID string `json:"employeeId,omitempty" bson:"employeeId,omitempty"`
	Remark     string `json:"remark,omitempty" bson:"remark,omitempty"`
	By         string `json:"by" bson:"by,omitempty"`
	ByID       string `json:"byId" bson:"byId,omitempty"`
	ByType     string `json:"byType" bson:"byType,omitempty"`
}

type RefEmployee struct {
	Employee `bson:",inline"`
	Ref      struct {
		OrganisationID   Organisation            `json:"organisationId,omitempty" bson:"organisationId,omitempty"`
		Grade            Grade                   `json:"grade,omitempty" bson:"grade,omitempty"`
		DepartmentID     Department              `json:"departmentId,omitempty" bson:"departmentId,omitempty"`
		BranchID         Branch                  `json:"branchId,omitempty" bson:"branchId,omitempty"`
		DesignationID    Designation             `json:"designationId,omitempty" bson:"designationId,omitempty"`
		Bank             BankInformation         `json:"bank" bson:"bank,omitempty"`
		LineManager      User                    `json:"lineManager" bson:"lineManager,omitempty"`
		WorkSchedule     WorkSchedule            `json:"workschedule,omitempty" bson:"workschedule,omitempty"`
		NoticeID         NoticePolicy            `json:"notice,omitempty" bson:"notice,omitempty"`
		ProbationaryID   Probationary            `json:"probationary,omitempty" bson:"probationary,omitempty"`
		DocumentPolicyID DocumentPolicy          `json:"documentPolicy,omitempty" bson:"documentPolicy,omitempty"`
		PayrollPolicyId  PayrollPolicy           `json:"payrollPolicyId,omitempty" bson:"payrollPolicyId,omitempty"`
		LeavePolicyID    LeavePolicy             `json:"leavePolicy,omitempty" bson:"leavePolicy,omitempty"`
		UserID           User                    `json:"userId" bson:"userId,omitempty"`
		Education        []EmployeeEducation     `json:"education" bson:"education,omitempty"`
		Experience       []EmployeeExperience    `json:"experience" bson:"experience,omitempty"`
		FamilyMembers    []EmployeeFamilyMembers `json:"familyMembers" bson:"familyMembers,omitempty"`
	} `json:"ref,omitempty" bson:"ref,omitempty"`
}
type AllEmployees struct {
	Employee []Employee `json:"employee" bson:"employee,omitempty"`
}
type LineManagerEmployee struct {
	Employee `bson:",inline"`
	Child    int64 `json:"child" bson:"child,omitempty"`
}
type FilterEmployee struct {
	Status         []string `json:"status,omitempty" bson:"status,omitempty"`
	OmitStatus     []string `json:"omitstatus,omitempty" bson:"omitstatus,omitempty"`
	OrganisationID []string `json:"organisation,omitempty" bson:"organisation,omitempty"`
	DepartmentID   []string `json:"department,omitempty" bson:"department,omitempty"`
	Manager        string   `json:"manager" bson:"manager,omitempty"`
	BranchID       []string `json:"branch,omitempty" bson:"branch,omitempty"`
	DesignationID  []string `json:"designation,omitempty" bson:"designation,omitempty"`
	UniqueID       []string `json:"uniqueId" bson:"uniqueId,omitempty"`
	Grade          []string `json:"grade" bson:"grade,omitempty"`
	Regex          struct {
		Name string `json:"name,omitempty" bson:"name,omitempty"`
	} `json:"regex" bson:"regex"`
	SortBy     string            `json:"sortBy"`
	SortOrder  int               `json:"sortOrder"`
	DataAccess DataAccessRequest `json:"dataAccess" bson:"dataAccess,omitempty"`
}

//Employee Login
// type EmployeeLogin struct {
// 	UserName string   `json:"userName"`
// 	PassWord string   `json:"password"`
// 	Location Location `json:"location,omitempty" bson:"location,omitempty"`
// }
type UpdateBioData struct {
	Mobile        string     `json:"mobile" bson:"mobile,omitempty"`
	Email         string     `json:"email" bson:"email,omitempty"`
	DOB           *time.Time `json:"dob" bson:"dob,omitempty"`
	Address       Address    `json:"address" bson:"address,omitempty"`
	Name          string     `json:"name" bson:"name,omitempty"`
	Gender        string     `json:"gender" bson:"gender,omitempty"`
	LineManager   string     `json:"lineManager" bson:"lineManager,omitempty"`
	ProfileImg    string     `json:"profileImg" bson:"profileImg,omitempty"`
	DepartmentID  string     `json:"departmentId,omitempty" bson:"departmentId,omitempty"`
	DesignationID string     `json:"designationId,omitempty" bson:"designationId,omitempty"`
}
type UpdateEmergencyContact struct {
	Primary struct {
		FullName     string   `json:"fullName" bson:"fullName,omitempty"`
		Relationship string   `json:"relationship,omitempty" bson:"relationship,omitempty"`
		PhoneNumber  []string `json:"phoneNumber" bson:"phoneNumber,omitempty"`
	} `json:"primary" bson:"primary,omitempty"`
	Secondary struct {
		FullName     string   `json:"fullName" bson:"fullName,omitempty"`
		Relationship string   `json:"relationship,omitempty" bson:"relationship,omitempty"`
		PhoneNumber  []string `json:"phoneNumber," bson:"phoneNumber,omitempty"`
	} `json:"secondary" bson:"secondary,omitempty"`
}
type UpdateBankInformation struct {
	BankName    string `json:"bankName" bson:"bankName,omitempty"`
	BankAccount string `json:"bankAccount" bson:"bankAccount,omitempty"`
	IfscCode    string `json:"ifscCode" bson:"ifscCode,omitempty"`
	PanNo       string `json:"panNo" bson:"panNo,omitempty"`
}
type PersonalInformation struct {
	PassportNo         string     `json:"passportNo" bson:"passportNo,omitempty"`
	PassportExpDate    *time.Time `json:"passportExpDate" bson:"passportExpDate,omitempty"`
	Telephone          string     `json:"telephone" bson:"telephone,omitempty"`
	Nationality        string     `json:"nationality" bson:"nationality	,omitempty"`
	Religion           string     `json:"religion" bson:"religion	,omitempty"`
	MaritalStatus      string     `json:"maritalStatus" bson:"maritalStatus,omitempty"`
	EmploymentOfSpouse string     `json:"employmentOfSpouse" bson:"employmentOfSpouse,omitempty"`
	NoOfChildrens      string     `json:"noOfChildrens" bson:"noOfChildrens,omitempty"`
}
type EmployeeUploadError struct {
	Name         string `json:"name" bson:"name,omitempty"`
	UserName     string `json:"userName" bson:"userName,omitempty"`
	MobileNumber string `json:"mobileNumber" bson:"mobileNumber,omitempty"`
	SNo          string `json:"sNo" bson:"sNo,omitempty"`
	Error        string `json:"error" bson:"error,omitempty"`
}
type DashboardEmployeeCount struct {
	Count int64 `json:"count" bson:"count,omitempty"`
}
type WeekCalEmployee struct {
	Employee `bson:",inline"`
	Ref      struct {
		DesignationID Designation `json:"designationId,omitempty" bson:"designationId,omitempty"`
	} `json:"ref,omitempty" bson:"ref,omitempty"`
}
