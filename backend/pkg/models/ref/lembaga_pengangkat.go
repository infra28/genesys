package models

import (
	"time"
)

// LembagaPengangkat mewakili struktur tabel ref.lembaga_pengangkat
type LembagaPengangkat struct {
	// Key: Primary Key, Type: numeric(2,0) -> int16
	LembagaPengangkatID int16 `gorm:"primaryKey;column:lembaga_pengangkat_id;type:numeric(2,0);not null" json:"lembaga_pengangkat_id"`

	// Informasi Lembaga
	Nama string `gorm:"column:nama;type:varchar(100);not null" json:"nama"`

	// Timestamp System
	CreateDate  time.Time  `gorm:"column:create_date;type:timestamp;not null;default:CURRENT_TIMESTAMP" json:"create_date"`
	LastUpdate  time.Time  `gorm:"column:last_update;type:timestamp;not null;default:CURRENT_TIMESTAMP" json:"last_update"`
	ExpiredDate *time.Time `gorm:"column:expired_date;type:timestamp" json:"expired_date"` // Allow Null
	LastSync    time.Time  `gorm:"column:last_sync;type:timestamp;not null;default:'1901-01-01 00:00:00'" json:"last_sync"`
}

func (LembagaPengangkat) TableName() string {
	return "ref.lembaga_pengangkat"
}
