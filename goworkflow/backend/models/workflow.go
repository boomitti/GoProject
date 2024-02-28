package models

import (
	"time"
)

type WFProcess struct {
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
