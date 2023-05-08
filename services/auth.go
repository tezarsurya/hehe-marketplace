package services

import (
	"hehe-marketplace/api/database"
	"hehe-marketplace/api/helpers"
	"hehe-marketplace/api/models"

	"github.com/gofiber/fiber/v2"
	"golang.org/x/crypto/bcrypt"
)

var db, _ = database.ConnectDB()

// Register
func Register(c *fiber.Ctx) error {
	var newUser = models.User{}
	var checkUser = models.User{}

	if err := c.BodyParser(&newUser); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": err.Error(),
		})
	}
	if err := helpers.ValidateStruct(newUser); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(err)
	}

	hashed, errHash := bcrypt.GenerateFromPassword([]byte(newUser.Password), 10)
	if errHash != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": errHash.Error(),
		})
	}
	newUser.Password = string(hashed)

	checkDuplicate := db.Select("id").
		Where("email = ?", newUser.Email).
		Or("phone = ?", newUser.Phone).
		Find(&checkUser)
	if checkDuplicate.RowsAffected > 0 {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "email or phone already exists",
			"user":    &checkUser,
		})
	}

	result := db.Create(&newUser)
	if result.Error != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": result.Error.Error(),
		})
	}
	return c.JSON(fiber.Map{
		"rows_affected": result.RowsAffected,
	})

}

// Login
