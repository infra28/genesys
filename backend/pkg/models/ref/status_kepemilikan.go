package models

import (
	"time"
)

// KepemilikanSarpras mewakili struktur tabel ref.kepemilikan_sarpras
type KepemilikanSarpras struct {
	// Key: Primary Key, Type: numeric(1,0) -> int16
	KepemilikanSarprasID int16 `gorm:"primaryKey;column:kepemilikan_sarpras_id;type:numeric(1,0);not null" json:"kepemilikan_sarpras_id"`

	// Informasi Kepemilikan
	Nama string `gorm:"column:nama;type:varchar(20);not null" json:"nama"`

	// Timestamp System
	CreateDate  time.Time  `gorm:"column:create_date;type:timestamp;not null;default:CURRENT_TIMESTAMP" json:"create_date"`
	LastUpdate  time.Time  `gorm:"column:last_update;type:timestamp;not null;default:CURRENT_TIMESTAMP" json:"last_update"`
	ExpiredDate *time.Time `gorm:"column:expired_date;type:timestamp" json:"expired_date"`
	LastSync    time.Time  `gorm:"column:last_sync;type:timestamp;not null;default:'1901-01-01 00:00:00'" json:"last_sync"`
}

func (KepemilikanSarpras) TableName() string {
	return "ref.kepemilikan_sarpras"
}
