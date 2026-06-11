package models

import (
	"time"
)

// StatusKepemilikan mewakili struktur tabel ref.status_kepemilikan
type StatusKepemilikan struct {
	// Key: Primary Key, Type: numeric(1,0) -> int16
	StatusKepemilikanID int16 `gorm:"primaryKey;column:status_kepemilikan_id;type:numeric(1,0);not null" json:"status_kepemilikan_id"`

	// Informasi Status
	Nama string `gorm:"column:nama;type:varchar(20);not null" json:"nama"`

	// Timestamp System
	CreateDate  time.Time  `gorm:"column:create_date;type:timestamp;not null;default:CURRENT_TIMESTAMP" json:"create_date"`
	LastUpdate  time.Time  `gorm:"column:last_update;type:timestamp;not null;default:CURRENT_TIMESTAMP" json:"last_update"`
	ExpiredDate *time.Time `gorm:"column:expired_date;type:timestamp" json:"expired_date"`
	LastSync    time.Time  `gorm:"column:last_sync;type:timestamp;not null;default:'1901-01-01 00:00:00'" json:"last_sync"`
}

func (StatusKepemilikan) TableName() string {
	return "ref.status_kepemilikan"
}
