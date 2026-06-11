// File: backend/pkg/database/reference_seeder.go
package database

import (
	"log"
	"os"
)

func LoadNavicatDump() {
	// 1. Buat skema 'ref' terlebih dahulu
	if err := DB.Exec("CREATE SCHEMA IF NOT EXISTS ref;").Error; err != nil {
		log.Fatalf("failed to create schema ref: %v", err)
	}

	// 2. Cek apakah tabel ref.agama sudah ada isinya
	// Jika sudah ada isinya, hentikan eksekusi agar tidak mengulang import data besar
	var count int64
	// Kita abaikan error di sini, karena kalau tabel belum ada count akan 0
	DB.Table("ref.agama").Count(&count)
	if count > 0 {
		log.Println("ref available, skip process...")
		return
	}

	// 3. Baca file dump dari path yang Anda tentukan
	// Asumsi Anda menjalankan aplikasi (go run) dari root folder "backend"
	filePath := "pkg/database/nav_ref.sql"
	sqlBytes, err := os.ReadFile(filePath)
	if err != nil {
		log.Fatalf("Failed to read SQL file in %s: %v", filePath, err)
	}

	// 4. Eksekusi seluruh isi file
	log.Println("Executing ref, this process may take a few moments")
	if err := DB.Exec(string(sqlBytes)).Error; err != nil {
		log.Fatalf("Failed to execute SQL dump: %v", err)
	}

	log.Println("Success! All reference tables have been entered into the database.")
}
