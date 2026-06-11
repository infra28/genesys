package seeders

import (
	models "backend/pkg/models/ref" // sesuaikan dengan path modul Anda
	"log"

	"gorm.io/gorm"
)

// RunSeeders akan dipanggil oleh main.go
func RunSeeders(db *gorm.DB) {
	seedList := []struct {
		fileName string
		model    any
	}{
		{"agama.json", models.Agama{}},
		{"jenjang_pendidikan.json", models.JenjangPendidikan{}},
		{"mst_wilayah.json", models.MstWilayah{}},
	}

	for _, s := range seedList {
		err := SeedFromJSON(db, s.fileName, s.model)
		if err != nil {
			log.Fatalf("Error seeding %s: %v", s.fileName, err)
		}
	}
}
