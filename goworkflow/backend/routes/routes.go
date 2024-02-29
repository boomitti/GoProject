package routes

import (
	"goworkflow/backend/handlers"

	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

type Handler struct {
	DB *gorm.DB
}

func SetupRoutes(app *fiber.App, handler *handlers.ProcessHandler) {
	api := app.Group("/api")

	api.Get("/process", handler.GetProcesses)
	api.Post("/process", handler.CreateProcess)
	api.Get("/process/:id", handler.GetProcessByID)
	api.Put("/approve", handler.UpdateProcess)
	api.Delete("/process/:id", handler.DeleteProcess)
}
