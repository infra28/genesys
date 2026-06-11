package middleware

import (
	"github.com/gofiber/fiber/v3"
)

func IsAdmin(c fiber.Ctx) error {
	// Ambil role dari Locals (pastikan saat login/JWT check, nama role disimpan di sini)
	role := c.Locals("role")

	if role != "admin" {
		return c.Status(403).JSON(fiber.Map{
			"error": "Forbidden: you don't have permission",
		})
	}

	return c.Next()
}
