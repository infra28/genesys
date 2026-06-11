package database

import (
	"backend/pkg/models"
	"backend/pkg/utils"
	"fmt"
)

func SeedAdmin() {
	adminEmail := "admin@example.com"

	var exists int64
	DB.Model(&models.User{}).Where("email = ?", adminEmail).Count(&exists)

	if exists == 0 {
		hashedPassword, _ := utils.HashPassword("admin123")

		admin := models.User{
			Name:     "Administrator",
			Email:    adminEmail,
			Password: hashedPassword,
			RoleID:   1, // Merujuk pada ID Role 'admin'
		}

		if err := DB.Create(&admin).Error; err != nil {
			fmt.Printf("Gagal membuat seeder admin: %v\n", err)
		} else {
			fmt.Println("Seeder: Admin user created")
		}
	} else {
		fmt.Println("Seeder: Admin already exists, skipping...")
	}
}
