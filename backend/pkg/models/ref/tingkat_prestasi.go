package models

import (
	"time"
)

// TingkatPrestasi mewakili struktur tabel ref.tingkat_prestasi
type TingkatPrestasi struct {
	// Key: Primary Key, Type: int4 -> int32
	TingkatPrestasiID int32 `gorm:"primaryKey;column:tingkat_prestasi_id;type:int4;not null" json:"tingkat_prestasi_id"`

	// Informasi Tingkat Prestasi
	Nama string `gorm:"column:nama;type:varchar(50);not null" json:"nama"`

	// Timestamp System
	CreateDate  time.Time  `gorm:"column:create_date;type:timestamp;not null;default:CURRENT_TIMESTAMP" json:"create_date"`
	LastUpdate  time.Time  `gorm:"column:last_update;type:timestamp;not null;default:CURRENT_TIMESTAMP" json:"last_update"`
	ExpiredDate *time.Time `gorm:"column:expired_date;type:timestamp" json:"expired_date"` // Allow Null
	LastSync    time.Time  `gorm:"column:last_sync;type:timestamp;not null;default:'1901-01-01 00:00:00'" json:"last_sync"`
}

func (TingkatPrestasi) TableName() string {
	return "ref.tingkat_prestasi"
}
