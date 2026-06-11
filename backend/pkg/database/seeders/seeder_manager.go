package seeders

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

// SeedFromJSON adalah fungsi generic untuk semua tabel
func SeedFromJSON[T any](db *gorm.DB, fileName string, model T) error {
	path := filepath.Join("database/data", fileName)
	file, err := os.ReadFile(path)
	if err != nil {
		return fmt.Errorf("gagal membaca file %s: %w", fileName, err)
	}

	var data []T
	if err := json.Unmarshal(file, &data); err != nil {
		return fmt.Errorf("gagal unmarshal file %s: %w", fileName, err)
	}

	// Gunakan Upsert agar bisa dijalankan berkali-kali tanpa error
	err = db.Clauses(clause.OnConflict{
		UpdateAll: true,
	}).Create(&data).Error

	if err != nil {
		return fmt.Errorf("gagal seeding database %s: %w", fileName, err)
	}

	fmt.Printf("Berhasil seeding: %s\n", fileName)
	return nil
}
