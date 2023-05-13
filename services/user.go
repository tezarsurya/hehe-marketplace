package services

import (
	"errors"
	"hehe-marketplace/api/database"
	"hehe-marketplace/api/models"

	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

func GetMyProfile(c *fiber.Ctx) error {
	db, errDB := database.ConnectDB()
	if errDB != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": errDB.Error(),
		})
	}

	var user *models.User
	userID := c.GetRespHeader("user_id")
	result := db.Select([]string{
		"id",
		"name",
		"phone",
		"birthday",
		"gender",
		"about",
		"profession",
		"email",
		"province_id",
		"city_id",
	}).
		Preload("Store").
		Where("id = ?", &userID).
		First(&user)

	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
				"message": result.Error.Error(),
			})
		}
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": result.Error.Error(),
		})
	}

	return c.JSON(user)
}
