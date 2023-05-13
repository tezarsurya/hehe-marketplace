package main

import (
	"hehe-marketplace/api/database"
	"hehe-marketplace/api/middlewares"
	"hehe-marketplace/api/models"
	"hehe-marketplace/api/services"
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"
)

func main() {
	// Load .env
	if err := godotenv.Load(); err != nil {
		log.Fatal("failed to load .env")
	}

	// Init fiber instance
	app := fiber.New()

	// Database connection
	db, errDB := database.ConnectDB()
	if errDB != nil {
		log.Fatal(errDB)
	}

	// Automatic Migration
	db.AutoMigrate(
		&models.User{},
		&models.Store{},
		&models.Address{},
		&models.Category{},
		&models.Product{},
		&models.ProductPicture{},
		&models.Transaction{},
		&models.TransactionDetail{},
		&models.ProductLog{},
	)

	// Routes
	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hi there!")
	})

	// API Routes
	api := app.Group("/api")

	// Auth
	api.Post("/register", services.Register)
	api.Post("/login", services.Login)

	// Category
	category := api.Group("/category", middlewares.AuthRequired, middlewares.AdminOnly)
	category.Get("/", services.GetAllCategories)
	category.Get("/:id<int>", services.GetCategoryByID)
	category.Post("/", services.PostCategory)
	category.Put("/:id", services.UpdateCategoryByID)
	category.Delete("/:id", services.DeleteCategoryByID)

	// Profile
	profile := api.Group("/profile", middlewares.AuthRequired)
	profile.Get("/", services.GetMyProfile)

	// PORT 8080
	app.Listen(":8080")
}
