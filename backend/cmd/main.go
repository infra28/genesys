package main

import (
	"backend/pkg/database"
	//"backend/pkg/database/seeders"
	"backend/pkg/models"
	"backend/pkg/routes"
	"fmt"
	"log"

	"github.com/gofiber/fiber/v3"
	"github.com/gofiber/fiber/v3/middleware/cors" // Tambahkan import cors
	"github.com/spf13/viper"
)

func main() {
	// 1. Load Konfigurasi via Viper
	viper.SetConfigFile(".env")
	if err := viper.ReadInConfig(); err != nil {
		log.Fatalf("Error memuat file .env: %s", err)
	}

	// 2. Inisialisasi Database
	database.Connect()

	if !fiber.IsChild() {
		fmt.Println("Running database migration...")
		database.LoadNavicatDump()
		err := database.DB.AutoMigrate(&models.Role{}, &models.User{})
		if err != nil {
			log.Fatalf("Failed to migrate database: %s", err)
		}
		database.SeedRoles()
		database.SeedAdmin()
		//seeders.RunSeeders(database.DB)
	}

	// 3. Inisialisasi Fiber v3
	app := fiber.New(fiber.Config{
		AppName: "Genesys Version 1.0",
	})

	// PENTING: Tambahkan Middleware CORS
	app.Use(cors.New(cors.Config{
		AllowOrigins: []string{"http://localhost:5173"}, // Sesuaikan dengan URL Vite frontend Anda
		AllowHeaders: []string{"Origin", "Content-Type", "Accept", "Authorization"},
		AllowMethods: []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
	}))

	routes.SetupRoutes(app)

	// 5. Jalankan Server
	port := viper.GetString("PORT")
	if port == "" {
		port = "3000"
	}

	if !fiber.IsChild() {
		fmt.Printf("Server berjalan di port %s \n", port)
	}

	log.Fatal(app.Listen(fmt.Sprintf(":%s", port), fiber.ListenConfig{
		EnablePrefork: false, // Perhatikan: di Printf Anda bilang aktif, tapi di config ini false
	}))
}
