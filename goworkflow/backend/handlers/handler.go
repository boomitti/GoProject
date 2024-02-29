package handlers

import (
	"goworkflow/backend/models"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Handler struct {
	DB *gorm.DB
}

type ProcessHandler struct {
	Handler
}

func NewProcessHandler(db *gorm.DB) *ProcessHandler {
	return &ProcessHandler{Handler: Handler{DB: db}}
}

func (handler *ProcessHandler) GetProcesses(c *fiber.Ctx) error {

	// Implement logic to fetch processes from the database
	var workflows []models.WFProcess

	// Use Find to retrieve all records from the table
	err := handler.DB.Order("created_date").Find(&workflows).Error
	if err != nil {
		// Handle the error and return an appropriate response
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error":   true,
			"message": "Failed to fetch processes from the database",
		})
	}

	// Return the fetched processes as JSON response
	return c.JSON(workflows)
}

func (handler *ProcessHandler) CreateProcess(c *fiber.Ctx) error {
	// Extract data from the request, assuming it's in JSON format
	var requestData models.WFProcess

	// Unmarshal the JSON request body into the WFProcess struct
	if err := c.BodyParser(&requestData); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error":   true,
			"message": "Failed to parse request body",
		})
	}

	requestData.WFLinkFormHeaderProcessID = uuid.NewString()

	// Set the CreatedDate field
	requestData.CreatedDate = time.Now()

	// Create three new processes in the database
	for i := 1; i <= 3; i++ {
		// Set WFProcessSeq
		requestData.WFProcessSeq = i
		requestData.WFProcessID = uuid.New()

		switch requestData.WFProcessSeq {
		case 1:
			requestData.WFStatus = "Draft"
		case 2, 3:
			requestData.WFStatus = "Not start"
		}

		// Create a new process in the database
		err := handler.DB.Create(&requestData).Error
		if err != nil {
			// Handle the error and return an appropriate response
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"error":   true,
				"message": "Failed to create process in the database",
			})
		}
	}

	// Return a success response
	return c.JSON(fiber.Map{
		"message": "Process created successfully",
	})
}

func (handler *ProcessHandler) GetProcessByID(c *fiber.Ctx) error {
	// Implement logic to fetch a process by ID from the database
	// Example: processID := c.Params("id")
	// var process models.WFProcess
	// err := handler.DB.Model(&models.WFProcess{}).Where("id = ?", processID).First(&process).Error
	// Handle the error and return the response
	return c.JSON(fiber.Map{"message": "GetProcessByID"})
}

func (handler *ProcessHandler) UpdateProcess(c *fiber.Ctx) error {
	// Extract data from the request, assuming it's in JSON format
	var updateData models.WFProcess

	// Unmarshal the JSON request body into the WFProcessUpdate struct
	if err := c.BodyParser(&updateData); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error":   true,
			"message": "Failed to parse request body",
		})
	}

	// Find the record to update based on WFLinkFormHeaderProcessID and WFProcessSeq
	var existingRecord models.WFProcess
	// Translate the modified GORM query to update WFStatus
	err := handler.DB.Where("wf_link_form_header_process_id = ? AND wf_process_seq = ?", updateData.WFLinkFormHeaderProcessID, updateData.WFProcessSeq).First(&existingRecord).Error
	if err != nil {
		// Handle the error (record not found, etc.) and return an appropriate response
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"error":   true,
			"message": "Record not found",
		})
	}

	// Step 1: Copy the record to pdms_w_wf_process_hist_vers
	histRecord := models.WFProcessHistVers{
		// Copy relevant fields from existingRecord to histRecord
		WFProcessID:                 existingRecord.WFProcessID.String(),
		WFLinkFormHeaderProcessID:   existingRecord.WFLinkFormHeaderProcessID,
		WFLinkFormVers:              existingRecord.WFLinkFormVers,
		WFProcessSeq:                existingRecord.WFProcessSeq,
		WFStatus:                    existingRecord.WFStatus,
		WFCurrentAssignStep:         existingRecord.WFCurrentAssignStep,
		WFCurrentAssignStepAssignee: existingRecord.WFCurrentAssignStepAssignee,
		WFCurrentAssignType:         existingRecord.WFCurrentAssignType,
		WFCurrentAssignID:           existingRecord.WFCurrentAssignID,
		WFCurrentAssignTempID:       existingRecord.WFCurrentAssignTempID,
		WFCurrentAssignName:         existingRecord.WFCurrentAssignName,
		WFCurrentAssignAccepted:     existingRecord.WFCurrentAssignAccepted,
		WFCurrentAssignPositionID:   existingRecord.WFCurrentAssignPositionID,
		WFCurrentAssignPositionName: existingRecord.WFCurrentAssignPositionName,
		WFCurrentAssignOrgID:        existingRecord.WFCurrentAssignOrgID,
		WFCurrentAssignOrgName:      existingRecord.WFCurrentAssignOrgName,
		WFCurrentAssignComment:      existingRecord.WFCurrentAssignComment,
		WFCurrentLayout:             existingRecord.WFCurrentLayout,
		WFCurrentAction:             existingRecord.WFCurrentAction,
		CreatedDate:                 time.Now(),
	}

	// Generate a unique identifier for the histRecord
	histRecord.WFProcessID = uuid.NewString()

	// Create a record in pdms_w_wf_process_hist_vers
	err = handler.DB.Create(&histRecord).Error
	if err != nil {
		// Handle the error and return an appropriate response
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error":   true,
			"message": "Failed to copy record to pdms_w_wf_process_hist_vers",
		})
	}

	// Step 2: Copy the record to pdms_w_wf_process_action_log with wf_action set to "Approve"
	actionLogRecord := models.WFProcessActionLog{
		// Copy relevant fields from existingRecord to actionLogRecord
		WFProcessID:                 existingRecord.WFProcessID.String(),
		WFLinkFormHeaderProcessID:   existingRecord.WFLinkFormHeaderProcessID,
		WFLinkFormVers:              existingRecord.WFLinkFormVers,
		WFProcessSeq:                existingRecord.WFProcessSeq,
		WFCurrentAssignStep:         existingRecord.WFCurrentAssignStep,
		WFCurrentAssignStepAssignee: existingRecord.WFCurrentAssignStepAssignee,
		WFCurrentAssignType:         existingRecord.WFCurrentAssignType,
		WFCurrentAssignID:           existingRecord.WFCurrentAssignID,
		WFCurrentAssignTempID:       existingRecord.WFCurrentAssignTempID,
		WFCurrentAssignName:         existingRecord.WFCurrentAssignName,
		WFCurrentAssignAccepted:     existingRecord.WFCurrentAssignAccepted,
		WFCurrentAssignPositionID:   existingRecord.WFCurrentAssignPositionID,
		WFCurrentAssignPositionName: existingRecord.WFCurrentAssignPositionName,
		WFCurrentAssignOrgID:        existingRecord.WFCurrentAssignOrgID,
		WFCurrentAssignOrgName:      existingRecord.WFCurrentAssignOrgName,
		WFCurrentAssignComment:      existingRecord.WFCurrentAssignComment,

		// Set the wf_action field
		WFAction: "Approve",

		// Set the CreatedDate field
		CreatedDate: time.Now(),
	}

	// Generate a unique identifier for the actionLogRecord
	actionLogRecord.WFProcessID = uuid.NewString()

	// Create a record in pdms_w_wf_process_action_log
	err = handler.DB.Create(&actionLogRecord).Error
	if err != nil {
		// Handle the error and return an appropriate response
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error":   true,
			"message": "Failed to copy record to pdms_w_wf_process_action_log",
		})
	}

	// Step 3: Update WFStatus of the existing record to "Done"
	// err = handler.DB.Model(&existingRecord).Update("wf_status", "Done").Error
	// if err != nil {
	// 	// Handle the error and return an appropriate response
	// 	return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
	// 		"error":   true,
	// 		"message": "Failed to update WFStatus to 'Done'",
	// 	})
	// }
	// Translate the modified GORM query to update WFStatus
	err = handler.DB.
		Model(&existingRecord).
		Where("wf_link_form_header_process_id = ? AND wf_process_seq = ?", updateData.WFLinkFormHeaderProcessID, updateData.WFProcessSeq).
		Update("wf_status", "Done").
		Error
	if err != nil {
		// Handle the error and return an appropriate response
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error":   true,
			"message": "Failed to update WFStatus to 'Done'",
		})
	}

	// Step 4: Update WFStatus of the record with the next WFProcessSeq to "In progress"
	nextSeq := existingRecord.WFProcessSeq + 1

	// Check if the record with the next WFProcessSeq exists before updating
	var nextRecordCount int64
	err = handler.DB.Model(&models.WFProcess{}).Where("wf_link_form_header_process_id = ? AND wf_process_seq = ?", updateData.WFLinkFormHeaderProcessID, nextSeq).Count(&nextRecordCount).Error
	if err != nil {
		// Handle the error and return an appropriate response
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error":   true,
			"message": "Failed to check if the record with the next WFProcessSeq exists",
		})
	}

	if nextRecordCount > 0 {
		// Update WFStatus to "In progress" only if the next record exists
		err = handler.DB.Model(&models.WFProcess{}).Where("wf_link_form_header_process_id = ? AND wf_process_seq = ?", updateData.WFLinkFormHeaderProcessID, nextSeq).Update("wf_status", "In progress").Error
		if err != nil {
			// Handle the error and return an appropriate response
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"error":   true,
				"message": "Failed to update WFStatus to 'In progress' for the next sequence",
			})
		}
	}

	// Return a success response
	return c.JSON(fiber.Map{
		"message": "Process updated successfully",
	})
}

func (handler *ProcessHandler) DeleteProcess(c *fiber.Ctx) error {
	// Implement logic to delete a process from the database
	// Example: processID := c.Params("id")
	// err := handler.DB.Model(&models.WFProcess{}).Where("id = ?", processID).Delete(&models.WFProcess{}).Error
	// Handle the error and return the response
	return c.JSON(fiber.Map{"message": "DeleteProcess"})
}
