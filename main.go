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

	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hi there!")
	})

	app.Post("/register", services.Register)

	// PORT 8080
	app.Listen(":8080")
}
