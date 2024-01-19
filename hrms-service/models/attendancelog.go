package models

import (
	"time"
)

type AttendanceLog struct {
	UniqueID       string     `json:"uniqueId,omitempty" bson:"uniqueId,omitempty"`
	AttendanceID   string     `json:"attendanceId,omitempty" bson:"attendanceId,omitempty"`
	LoginMode      string     `json:"loginMode,omitempty" bson:"loginMode,omitempty"`
	Status         string     `json:"status,omitempty" bson:"status,omitempty"`
	PunchinTime    *time.Time `json:"punchinTime,omitempty" bson:"punchinTime,omitempty"`
	PunchoutTime   *time.Time `json:"punchoutTime,omitempty" bson:"punchoutTime,omitempty"`
	EmployeeId     string     `json:"employeeId,omitempty" bson:"employeeId,omitempty"`
	OrganisationId string     `json:"organisationId" bson:"organisationId,omitempty"`
	WorkingMins    float64    `json:"workingMins" bson:"workingMins,omitempty"`
	OverTime       float64    `json:"overtime,omitempty" bson:"overtime,omitempty"`
	BreakMins      float64    `json:"breakMins" bson:"breakMins,omitempty"`
	Date           *time.Time `json:"date,omitempty" bson:"date,omitempty"`
	Notes          string     `json:"notes" bson:"notes,omitempty"`
	Created        *Created   `json:"createdOn,omitempty" bson:"createdOn,omitempty"`
}

type RefAttendanceLog struct {
	AttendanceLog `bson:",inline"`
	Ref           struct {
	} `json:"ref,omitempty" bson:"ref,omitempty"`
}

type AttendanceEmployeeTodayStatus struct {
	//AttendanceLog        `bson:",inline"`
	RecentPunchinTime    *time.Time `json:"recentpunchinTime" bson:"recentpunchinTime,omitempty"`
	RecentPunchoutTime   *time.Time `json:"recentpunchoutTime" bson:"recentpunchoutTime,omitempty"`
	FirstPunchinTime     *time.Time `json:"firstpunchinTime" bson:"firstpunchinTime,omitempty"`
	LastpunchoutTime     *time.Time `json:"lastpunchoutTime" bson:"lastpunchoutTime,omitempty"`
	TotalWorkingHours    float64    `json:"totalworkingHours" bson:"totalworkingHours,omitempty"`
	TotalBreakHours      float64    `json:"totalbreakHours" bson:"totalbreakHours,omitempty"`
	OverTime             float64    `json:"overTime" bson:"overTime,omitempty"`
	TotalWorkingHoursStr string     `json:"totalworkingHoursStr" bson:"totalworkingHoursStr,omitempty"`
	TotalBreakHoursStr   string     `json:"totalbreakHoursStr" bson:"totalbreakHoursStr,omitempty"`
	OverTimeStr          string     `json:"overTimeStr" bson:"overTimeStr,omitempty"`
	CurrenntStatus       string     `json:"currentStatus" bson:"currentStatus,omitempty"`
	IsTimeoffCheck       bool       `json:"isTimeoffCheck" bson:"isTimeoffCheck,omitempty"`
}
type FilterAttendanceLog struct {
	Status     []string `json:"status,omitempty" bson:"status,omitempty"`
	State      []string `json:"state,omitempty" bson:"state,omitempty"`
	EmployeeId []string `json:"employeeId,omitempty" bson:"employeeId,omitempty"`
	UniqueID   []string `json:"uniqueId,omitempty" bson:"uniqueId,omitempty"`
	// DateRange struct {
	// 	From *time.Time `json:"from,omitempty"  bson:"from,omitempty"`
	// 	To   *time.Time `json:"to,omitempty"  bson:"to,omitempty"`
	// } `json:"dateRange,omitempty"  bson:"dateRange,omitempty"`
}
