package models

import (
	"time"
)

// SertifikasiISO mewakili struktur tabel ref.sertifikasi_iso
type SertifikasiISO struct {
	// Key: Primary Key, Type: int2 -> int16
	SertifikasiISOID int16 `gorm:"primaryKey;column:sertifikasi_iso_id;type:int2;not null" json:"sertifikasi_iso_id"`

	// Informasi Sertifikasi
	Nama string `gorm:"column:nama;type:varchar(20);not null" json:"nama"`

	// Timestamp System
	CreateDate  time.Time  `gorm:"column:create_date;type:timestamp;not null;default:CURRENT_TIMESTAMP" json:"create_date"`
	LastUpdate  time.Time  `gorm:"column:last_update;type:timestamp;not null;default:CURRENT_TIMESTAMP" json:"last_update"`
	ExpiredDate *time.Time `gorm:"column:expired_date;type:timestamp" json:"expired_date"` // Allow Null
	LastSync    time.Time  `gorm:"column:last_sync;type:timestamp;not null;default:'1901-01-01 00:00:00'" json:"last_sync"`
}

func (SertifikasiISO) TableName() string {
	return "ref.sertifikasi_iso"
}
