package middlewares

import (
	"hehe-marketplace/api/services"
	"os"
	"regexp"
	"strconv"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
)

func AuthRequired(c *fiber.Ctx) error {
	reg := regexp.MustCompile(`Bearer [a-zA-z0-9]+[.][a-zA-z0-9]+[.][a-zA-z0-9-]+`)
	token := c.Get("Authorization")

	if !reg.MatchString(token) {
		return c.SendStatus(fiber.StatusUnauthorized)
	}

	token = token[7:]
	var claims services.JwtClaims

	parsed, err := jwt.ParseWithClaims(token, &claims, func(t *jwt.Token) (interface{}, error) {
		return []byte(os.Getenv("JWT_SECRET")), nil
	}, jwt.WithValidMethods(jwt.GetAlgorithms()))

	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": err.Error(),
		})
	}
	if !parsed.Valid {
		return c.SendStatus(fiber.StatusUnauthorized)
	}

	c.Set("user_id", strconv.FormatUint(uint64(claims.UserID), 10))
	var isAdmin string
	if bool(*claims.IsAdmin) {
		isAdmin = "1"
	} else {
		isAdmin = "0"
	}
	c.Set("is_admin", isAdmin)

	return c.Next()
}

func AdminOnly(c *fiber.Ctx) error {
	isAdmin := c.GetRespHeader("is_admin")
	if isAdmin == "0" {
		return c.SendStatus(fiber.StatusUnauthorized)
	}

	return c.Next()
}
