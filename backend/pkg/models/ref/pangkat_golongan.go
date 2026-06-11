package models

import (
	"time"
)

// PangkatGolongan mewakili struktur tabel ref.pangkat_golongan
type PangkatGolongan struct {
	// Key: Primary Key, Type: numeric(2,0) -> int16
	PangkatGolonganID int16 `gorm:"primaryKey;column:pangkat_golongan_id;type:numeric(2,0);not null" json:"pangkat_golongan_id"`

	// Informasi Pangkat
	Kode string `gorm:"column:kode;type:varchar(5);not null" json:"kode"`
	Nama string `gorm:"column:nama;type:varchar(20);not null" json:"nama"`

	// Timestamp System
	CreateDate  time.Time  `gorm:"column:create_date;type:timestamp;not null;default:CURRENT_TIMESTAMP" json:"create_date"`
	LastUpdate  time.Time  `gorm:"column:last_update;type:timestamp;not null;default:CURRENT_TIMESTAMP" json:"last_update"`
	ExpiredDate *time.Time `gorm:"column:expired_date;type:timestamp" json:"expired_date"` // Allow Null
	LastSync    time.Time  `gorm:"column:last_sync;type:timestamp;not null;default:'1901-01-01 00:00:00'" json:"last_sync"`
}

func (PangkatGolongan) TableName() string {
	return "ref.pangkat_golongan"
}
