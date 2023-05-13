package services

import (
	"errors"
	"fmt"
	"hehe-marketplace/api/database"
	"hehe-marketplace/api/helpers"
	"hehe-marketplace/api/models"
	"os"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type JwtClaims struct {
	jwt.RegisteredClaims
	UserID  uint  `json:"user_id"`
	IsAdmin *bool `json:"is_admin"`
}

type Credentials struct {
	Phone    string `validate:"required" json:"phone"`
	Password string `validate:"required" json:"password,omitempty"`
}

// Register
func Register(c *fiber.Ctx) error {
	db, errDB := database.ConnectDB()
	if errDB != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": errDB.Error(),
		})
	}

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
		})
	}

	if result := db.Save(&newUser); result.Error != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": result.Error.Error(),
		})
	}

	store := &models.Store{
		UserID:    newUser.ID,
		StoreName: fmt.Sprintf("%s_store", helpers.GenerateSlug(newUser.Name)),
		StoreURL:  fmt.Sprintf("https://example.com/store/%s", helpers.GenerateSlug(newUser.Name)+helpers.GenerateString(12)),
	}
	if result := db.Save(&store); result.Error != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": result.Error.Error(),
		})
	}

	return c.SendStatus(fiber.StatusOK)
}

// Login
func Login(c *fiber.Ctx) error {
	db, errDB := database.ConnectDB()
	if errDB != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": errDB.Error(),
		})
	}

	var cred *Credentials
	var user *models.User
	if err := c.BodyParser(&cred); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": err.Error(),
		})
	}
	if err := helpers.ValidateStruct(cred); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(err)
	}

	if result := db.Select([]string{"id", "password", "is_admin"}).
		Where("phone = ?", cred.Phone).
		First(&user); result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"message": "invalid phone number",
			})
		}
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": result.Error.Error(),
		})
	}

	if err := bcrypt.CompareHashAndPassword(
		[]byte(user.Password),
		[]byte(cred.Password),
	); err != nil {
		if errors.Is(err, bcrypt.ErrMismatchedHashAndPassword) {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"message": "wrong password",
			})
		}
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": err.Error(),
		})
	}

	claims := &JwtClaims{
		jwt.RegisteredClaims{
			Issuer:   os.Getenv("BASE_URL"),
			IssuedAt: jwt.NewNumericDate(time.Now()),
			ExpiresAt: jwt.NewNumericDate(
				time.Now().Add(24 * time.Hour),
			),
		},
		user.ID,
		user.IsAdmin,
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	signed, errToken := token.SignedString([]byte(os.Getenv("JWT_SECRET")))
	if errToken != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": errToken.Error(),
		})
	}

	return c.JSON(fiber.Map{
		"token": signed,
	})
}
