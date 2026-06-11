package models

import "time"

type TahunAjaran struct {
	TahunAjaranID  int16      `gorm:"primaryKey;column:tahun_ajaran_id;type:numeric(4,0);not null" json:"tahun_ajaran_id"`
	Nama           string     `gorm:"column:nama;type:varchar(10);not null" json:"nama"`
	PeriodeAktif   int16      `gorm:"column:periode_aktif;type:numeric(1,0);not null" json:"periode_aktif"`
	TanggalMulai   time.Time  `gorm:"column:tanggal_mulai;type:date;not null" json:"tanggal_mulai"`
	TanggalSelesai time.Time  `gorm:"column:tanggal_selesai;type:date;not null" json:"tanggal_selesai"`
	CreateDate     time.Time  `gorm:"column:create_date;type:timestamp;not null;default:CURRENT_TIMESTAMP" json:"create_date"`
	LastUpdate     time.Time  `gorm:"column:last_update;type:timestamp;not null;default:CURRENT_TIMESTAMP" json:"last_update"`
	ExpiredDate    *time.Time `gorm:"column:expired_date;type:timestamp" json:"expired_date"`
	LastSync       time.Time  `gorm:"column:last_sync;type:timestamp;not null;default:'1901-01-01 00:00:00'" json:"last_sync"`

	// ========================================================
	// RELASI (HAS MANY)
	// ========================================================
	// Satu Tahun Ajaran memiliki banyak Semester.
	// Buka komentar di bawah jika Anda ingin GORM bisa memuat daftar semester secara otomatis.
	// Semesters []Semester `gorm:"foreignKey:TahunAjaranID;references:TahunAjaranID" json:"semesters,omitempty"`
}

func (TahunAjaran) TableName() string {
	return "ref.tahun_ajaran"
}
