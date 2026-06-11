// Buat file baru: backend/pkg/handlers/user.go
package handlers

import (
	"backend/pkg/database"
	"backend/pkg/models"

	"github.com/gofiber/fiber/v3"
)

func GetMe(c fiber.Ctx) error {
	// Ambil user_id dari context yang diset oleh middleware JWT
	userID := c.Locals("user_id")
	if userID == nil {
		return c.Status(401).JSON(fiber.Map{"error": "Unauthorized"})
	}

	var user models.User
	// Query ke database dan ambil relasi Role-nya juga
	result := database.DB.Preload("Role").Where("id = ?", userID).First(&user)

	if result.Error != nil {
		return c.Status(404).JSON(fiber.Map{"error": "User not found"})
	}

	// Kembalikan data yang dibutuhkan frontend
	return c.JSON(fiber.Map{
		"status": "success",
		"data": fiber.Map{
			"id":         user.ID,
			"name":       user.Name,
			"email":      user.Email,
			"role":       user.Role.Name,
			"created_at": user.CreatedAt,
		},
	})
}
