package main

import (
	"fmt"
	"goworkflow/backend/database"
	"goworkflow/backend/handlers"
	"goworkflow/backend/models"
	"goworkflow/backend/routes"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
)

func main() {
	app := fiber.New()

	// Initialize the database connection
	db, err := database.ConnectDB()
	if err != nil {
		fmt.Printf("Failed to connect to the database: %v\n", err)
		return
	}

	// Auto-migrate the GORM models
	err = db.AutoMigrate(&models.WFProcess{}, &models.WFProcessHistVers{}, &models.WFProcessActionLog{})
	if err != nil {
		fmt.Printf("Failed to auto-migrate database models: %v\n", err)
		return
	}

	// Create a new instance of ProcessHandler
	processHandler := handlers.NewProcessHandler(db)

	// Setup routes and handlers
	routes.SetupRoutes(app, processHandler)

	// Run the Fiber application
	err = app.Listen(":3000")
	if err != nil {
		fmt.Printf("Error starting Fiber application: %v\n", err)
	}

	app.Use(logger.New())
}
