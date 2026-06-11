package models

import (
	"time"
)

// MataPelajaran mewakili struktur tabel ref.mata_pelajaran
type MataPelajaran struct {
	// Key: Primary Key, Type: int4 -> int32
	MataPelajaranID int32 `gorm:"primaryKey;column:mata_pelajaran_id;type:int4;not null" json:"mata_pelajaran_id"`

	// Informasi Mata Pelajaran
	Nama string `gorm:"column:nama;type:varchar(80);not null" json:"nama"`

	// Flags (numeric 1,0 -> int16)
	PilihanSekolah      int16 `gorm:"column:pilihan_sekolah;type:numeric(1,0);not null" json:"pilihan_sekolah"`
	PilihanBuku         int16 `gorm:"column:pilihan_buku;type:numeric(1,0);not null" json:"pilihan_buku"`
	PilihanKepengawasan int16 `gorm:"column:pilihan_kepengawasan;type:numeric(1,0);not null" json:"pilihan_kepengawasan"`
	PilihanEvaluasi     int16 `gorm:"column:pilihan_evaluasi;type:numeric(1,0);not null;default:0" json:"pilihan_evaluasi"`

	// Foreign Key: Jurusan
	JurusanID *string  `gorm:"column:jurusan_id;type:varchar(25)" json:"jurusan_id"`
	Jurusan   *Jurusan `gorm:"foreignKey:JurusanID;references:JurusanID" json:"jurusan,omitempty"`

	// Timestamp System
	CreateDate  time.Time  `gorm:"column:create_date;type:timestamp;not null;default:CURRENT_TIMESTAMP" json:"create_date"`
	LastUpdate  time.Time  `gorm:"column:last_update;type:timestamp;not null;default:CURRENT_TIMESTAMP" json:"last_update"`
	ExpiredDate *time.Time `gorm:"column:expired_date;type:timestamp" json:"expired_date"`
	LastSync    time.Time  `gorm:"column:last_sync;type:timestamp;not null;default:'1901-01-01 00:00:00'" json:"last_sync"`
}

func (MataPelajaran) TableName() string {
	return "ref.mata_pelajaran"
}
