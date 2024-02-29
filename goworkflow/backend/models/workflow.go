package models

import (
	"time"

	"github.com/google/uuid"
)

type WFProcess struct {
	ID                          uint      `gorm:"primaryKey"`
	WFProcessID                 uuid.UUID `gorm:"type:uuid;default:gen_random_uuid()"` // Assuming you're using a UUID for wf_process_id
	WFLinkFormHeaderProcessID   string    `gorm:"not null"`
	WFLinkFormVers              int       `gorm:"not null"`
	WFProcessSeq                int       `gorm:"not null"`
	WFStatus                    string    `gorm:"not null"`
	WFCurrentAssignStep         string
	WFCurrentAssignStepAssignee string
	WFCurrentAssignType         string
	WFCurrentAssignID           string
	WFCurrentAssignTempID       string
	WFCurrentAssignName         string
	WFCurrentAssignAccepted     string
	WFCurrentAssignPositionID   string
	WFCurrentAssignPositionName string
	WFCurrentAssignOrgID        string
	WFCurrentAssignOrgName      string
	WFCurrentAssignComment      string `gorm:"type:text"`
	WFCurrentLayout             string `gorm:"type:text"`
	WFCurrentAction             string `gorm:"type:text"`
	Remark                      string
	CreatedDate                 time.Time `gorm:"not null"`
	CreatedBy                   string    `gorm:"not null"`
	UpdatedDate                 *time.Time
	UpdatedBy                   string
}

// Explicitly specify the table name
func (WFProcess) TableName() string {
	return "pdms_w_wf_process"
}

type WFProcessHistVers struct {
	ID                          uint   `gorm:"primaryKey"`
	WFProcessID                 string `gorm:"not null"`
	WFLinkFormHeaderProcessID   string `gorm:"not null"`
	WFLinkFormVers              int    `gorm:"not null"`
	WFProcessSeq                int    `gorm:"not null"`
	WFStatus                    string `gorm:"not null"`
	WFCurrentAssignStep         string
	WFCurrentAssignStepAssignee string
	WFCurrentAssignType         string
	WFCurrentAssignID           string
	WFCurrentAssignTempID       string
	WFCurrentAssignName         string
	WFCurrentAssignAccepted     string
	WFCurrentAssignPositionID   string
	WFCurrentAssignPositionName string
	WFCurrentAssignOrgID        string
	WFCurrentAssignOrgName      string
	WFCurrentAssignComment      string `gorm:"type:text"`
	WFCurrentLayout             string `gorm:"type:text"`
	WFCurrentAction             string `gorm:"type:text"`
	Remark                      string
	CreatedDate                 time.Time `gorm:"not null"`
	CreatedBy                   string    `gorm:"not null"`
	UpdatedDate                 *time.Time
	UpdatedBy                   string
}

// Explicitly specify the table name
func (WFProcessHistVers) TableName() string {
	return "pdms_w_wf_process_hist_vers"
}

type WFProcessActionLog struct {
	ID                          int       `gorm:"column:id;primaryKey"`
	WFProcessID                 string    `gorm:"column:wf_process_id"`
	WFLinkFormHeaderProcessID   string    `gorm:"column:wf_link_form_header_process_id"`
	WFLinkFormVers              int       `gorm:"column:wf_link_form_vers"`
	WFProcessSeq                int       `gorm:"column:wf_process_seq"`
	WFAction                    string    `gorm:"column:wf_action"`
	WFCurrentAssignStep         string    `gorm:"column:wf_current_assign_step"`
	WFCurrentAssignStepAssignee string    `gorm:"column:wf_current_assign_step_assignee"`
	WFCurrentAssignType         string    `gorm:"column:wf_current_assign_type"`
	WFCurrentAssignID           string    `gorm:"column:wf_current_assign_id"`
	WFCurrentAssignTempID       string    `gorm:"column:wf_current_assign_temp_id"`
	WFCurrentAssignName         string    `gorm:"column:wf_current_assign_name"`
	WFCurrentAssignAccepted     string    `gorm:"column:wf_current_assign_accepted"`
	WFCurrentAssignPositionID   string    `gorm:"column:wf_current_assign_position_id"`
	WFCurrentAssignPositionName string    `gorm:"column:wf_current_assign_position_name"`
	WFCurrentAssignOrgID        string    `gorm:"column:wf_current_assign_org_id"`
	WFCurrentAssignOrgName      string    `gorm:"column:wf_current_assign_org_name"`
	WFCurrentAssignComment      string    `gorm:"column:wf_current_assign_comment"`
	Remark                      string    `gorm:"column:remark"`
	CreatedDate                 time.Time `gorm:"column:created_date"`
	CreatedBy                   string    `gorm:"column:created_by"`
}

// TableName sets the table name for the WFProcessActionLog model
func (WFProcessActionLog) TableName() string {
	return "pdms_w_wf_process_action_log"
}
