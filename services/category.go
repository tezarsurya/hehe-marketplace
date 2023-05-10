package services

import (
	"errors"
	"hehe-marketplace/api/database"
	"hehe-marketplace/api/helpers"
	"hehe-marketplace/api/models"

	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

func GetAllCategories(c *fiber.Ctx) error {
	db, errDB := database.ConnectDB()
	if errDB != nil {
		c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": errDB.Error(),
		})
	}

	var categories []*models.Category
	if result := db.Find(&categories); result.Error != nil {
		c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": result.Error.Error(),
		})
	}
	return c.JSON(categories)
}

func PostCategory(c *fiber.Ctx) error {
	db, errDB := database.ConnectDB()
	if errDB != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": errDB.Error(),
		})
	}

	var newCategory *models.Category
	if err := c.BodyParser(&newCategory); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": errDB.Error(),
		})
	}
	if err := helpers.ValidateStruct(newCategory); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(err)
	}

	result := db.Create(&newCategory)
	if result.Error != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": result.Error.Error(),
		})
	}
	return c.JSON(fiber.Map{
		"rows_affected": result.RowsAffected,
	})
}

func GetCategoryByID(c *fiber.Ctx) error {
	db, errDB := database.ConnectDB()
	if errDB != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": errDB.Error(),
		})
	}

	var category *models.Category
	if result := db.Where("id = ?", c.Params("id")).First(&category); result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
				"message": result.Error.Error(),
			})
		}
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": result.Error.Error(),
		})
	}
	return c.JSON(category)
}

func UpdateCategoryByID(c *fiber.Ctx) error {
	db, errDB := database.ConnectDB()
	if errDB != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": errDB.Error(),
		})
	}

	var update *models.Category
	var category *models.Category
	if err := c.BodyParser(&update); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": err.Error(),
		})
	}
	if err := helpers.ValidateStruct(update); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(err)
	}

	if result := db.Where("id = ?", c.Params("id")).First(&category); result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
				"message": result.Error.Error(),
			})
		}
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": result.Error.Error(),
		})
	}

	update.ID = category.ID
	if result := db.Save(&update); result.Error != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": result.Error.Error(),
		})
	}
	return c.JSON(update)
}

func DeleteCategoryByID(c *fiber.Ctx) error {
	db, errDB := database.ConnectDB()
	if errDB != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": errDB.Error(),
		})
	}

	var delete *models.Category
	result := db.Where("id = ?", c.Params("id")).Delete(&delete)
	if result.Error != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": result.Error.Error(),
		})
	}
	if result.RowsAffected < 1 {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"message": "record not found",
		})
	}
	return c.JSON(fiber.Map{
		"rows_affected": result.RowsAffected,
	})
}
