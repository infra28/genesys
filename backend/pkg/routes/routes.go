// File: backend/pkg/routes/routes.go
package routes

import (
	"backend/pkg/handlers"
	"backend/pkg/middleware"

	"github.com/gofiber/fiber/v3"
)

// SetupRoutes mengatur semua alur rute API untuk aplikasi Genesys
func SetupRoutes(app *fiber.App) {
	// Struktur Routing v1
	api := app.Group("/api")
	v1 := api.Group("/v1")

	// Endpoint Health Check untuk frontend
	v1.Get("/health", func(c fiber.Ctx) error {
		return c.JSON(fiber.Map{
			"status": "ok",
		})
	})

	// Auth Group (Public)
	auth := v1.Group("/auth")
	auth.Post("/register", handlers.Register)
	auth.Post("/login", handlers.Login)
	auth.Post("/logout", handlers.Logout)
	// Nanti Anda bisa tambahkan:
	// auth.Post("/forgot-password", handlers.ForgotPassword)

	// User Group (Protected)
	users := v1.Group("/users")
	users.Use(middleware.JWTProtected) // Middleware akan berlaku untuk rute di bawahnya
	users.Get("/me", handlers.GetMe)
}
