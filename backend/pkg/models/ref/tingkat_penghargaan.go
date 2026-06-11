package models

import (
	"time"
)

// TingkatPenghargaan mewakili struktur tabel ref.tingkat_penghargaan
type TingkatPenghargaan struct {
	// Key: Primary Key, Type: int4 -> int32
	TingkatPenghargaanID int32 `gorm:"primaryKey;column:tingkat_penghargaan_id;type:int4;not null" json:"tingkat_penghargaan_id"`

	// Informasi Tingkat Penghargaan
	Nama string `gorm:"column:nama;type:varchar(50);not null" json:"nama"`

	// Timestamp System
	CreateDate  time.Time  `gorm:"column:create_date;type:timestamp;not null;default:CURRENT_TIMESTAMP" json:"create_date"`
	LastUpdate  time.Time  `gorm:"column:last_update;type:timestamp;not null;default:CURRENT_TIMESTAMP" json:"last_update"`
	ExpiredDate *time.Time `gorm:"column:expired_date;type:timestamp" json:"expired_date"` // Allow Null
	LastSync    time.Time  `gorm:"column:last_sync;type:timestamp;not null;default:'1901-01-01 00:00:00'" json:"last_sync"`
}

func (TingkatPenghargaan) TableName() string {
	return "ref.tingkat_penghargaan"
}
