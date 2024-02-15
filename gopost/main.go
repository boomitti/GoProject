package main

import (
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

// User structure to represent data from the Users table
type User struct {
	gorm.Model
	UserName string `json:"userName" gorm:"column:username"`
	Email    string `json:"email"`
	Password string `json:"-"`
}

var db *gorm.DB

func init() {
	var err error

	// Replace the connection string with your PostgreSQL database connection details
	dsn := "user=postgres password=MgC&*s%3yU#Es9 dbname=godemo sslmode=disable"
	db, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal(err)
	}

	// Auto Migrate the Users table
	err = db.AutoMigrate(&User{})
	if err != nil {
		log.Fatal(err)
	}
}

func getAllUsersHandler(c *fiber.Ctx) error {
	var users []User
	result := db.Order("userid").Find(&users)

	if result.Error != nil {
		log.Println(result.Error)
		return c.Status(http.StatusInternalServerError).SendString("Internal Server Error")
	}

	return c.JSON(users)
}

// Get a specific user by ID
func getUserHandler(c *fiber.Ctx) error {
	userID, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid user ID",
		})
	}

	var user User
	result := db.First(&user, userID)
	if result.Error != nil {
		return c.Status(http.StatusNotFound).JSON(fiber.Map{
			"error": "User not found",
		})
	}

	return c.JSON(user)
}

func createUserHandler(c *fiber.Ctx) error {
	var newUser User
	if err := c.BodyParser(&newUser); err != nil {
		return c.Status(http.StatusBadRequest).SendString("Invalid request body")
	}

	result := db.Create(&newUser)
	if result.Error != nil {
		log.Println(result.Error)
		return c.Status(http.StatusInternalServerError).SendString("Failed to create user")
	}

	return c.JSON(newUser)
}

func updateUserHandler(c *fiber.Ctx) error {
	userID := c.Params("id")
	var updatedUser User

	if err := c.BodyParser(&updatedUser); err != nil {
		log.Println("Error parsing request body:", err)
		return c.Status(http.StatusBadRequest).SendString("Invalid request body")
	}

	var existingUser User
	result := db.First(&existingUser, userID)
	if result.Error != nil {
		log.Println("Error finding user:", result.Error)
		return c.Status(http.StatusNotFound).SendString("User not found")
	}

	// Update fields of the existing user
	existingUser.UserName = updatedUser.UserName
	existingUser.Email = updatedUser.Email
	existingUser.Password = updatedUser.Password

	result = db.Save(&existingUser)
	if result.Error != nil {
		log.Println("Error updating user:", result.Error)
		return c.Status(http.StatusInternalServerError).SendString("Failed to update user")
	}

	return c.JSON(existingUser)
}

func deleteUserHandler(c *fiber.Ctx) error {
	userID := c.Params("id")

	result := db.Delete(&User{}, userID)
	if result.Error != nil {
		log.Println("Error deleting user:", result.Error)
		return c.Status(http.StatusInternalServerError).SendString("Failed to delete user")
	}

	return c.Status(http.StatusOK).SendString("User deleted successfully")
}

func main() {
	app := fiber.New()

	// Use CORS middleware to enable cross-origin requests
	app.Use(cors.New())

	// Define API routes
	app.Get("/users", getAllUsersHandler)
	app.Get("/users/:id", getUserHandler)
	app.Post("/users", createUserHandler)
	app.Put("/users/:id", updateUserHandler)
	app.Delete("/users/:id", deleteUserHandler)

	// Start the HTTP server
	port := 8080
	fmt.Printf("Server running on :%d...\n", port)
	log.Fatal(app.Listen(fmt.Sprintf(":%d", port)))
}
