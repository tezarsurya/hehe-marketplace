package main

import (
	"hehe-marketplace/api/database"
	"hehe-marketplace/api/models"
	"hehe-marketplace/api/services"
	"log"

	"github.com/gofiber/fiber/v2"
)

func main() {
	// Init fiber instance
	app := fiber.New()

	// Database connection
	db, errDB := database.ConnectDB()
	if errDB != nil {
		log.Panicln(errDB)
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

	api := app.Group("/api")
	api.Post("/register", services.Register)
	api.Get("/category", services.GetAllCategories)
	api.Get("/category/:id<int>", services.GetCategoryByID)
	api.Post("/category", services.PostCategory)
	api.Put("/category/:id", services.UpdateCategoryByID)
	api.Delete("/category/:id", services.DeleteCategoryByID)

	// PORT 8080
	app.Listen(":8080")
}
