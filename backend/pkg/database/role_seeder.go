package database

import (
	"backend/pkg/models"
	"fmt"
)

func SeedRoles() {
	roles := []models.Role{
		{ID: 1, Name: "admin", Description: "Administrator"},
		{ID: 2, Name: "siswa", Description: "Siswa"},
		{ID: 3, Name: "guru", Description: "Guru"},
		{ID: 4, Name: "tata_usaha", Description: "Tata Usaha"},
		{ID: 5, Name: "kepsek", Description: "Kepala sekolah"},
		{ID: 6, Name: "walas", Description: "Wali kelas"},
		{ID: 7, Name: "bk", Description: "Bimbingan konseling"},
		{ID: 8, Name: "kesiswaan", Description: "Kesiswaan"},
		{ID: 9, Name: "kurikulum", Description: "Kurikulum"},
		{ID: 10, Name: "perpustakaan", Description: "Perpustakaan"},
		{ID: 11, Name: "toolman", Description: "Toolman"},
		{ID: 12, Name: "sarpras", Description: "Sarana dan prasarana"},
		{ID: 13, Name: "kakomli", Description: "Kakomli"},
	}

	for _, r := range roles {
		// Menggunakan FirstOrCreate agar tidak terjadi error duplicate key saat restart
		if err := DB.FirstOrCreate(&models.Role{}, r).Error; err != nil {
			fmt.Printf("Gagal seed role %s: %v\n", r.Name, err)
		}
	}
	fmt.Println("Seeder: Roles initialized")
}
