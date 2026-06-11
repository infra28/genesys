package models

import (
	"time"
)

// SumberListrik mewakili struktur tabel ref.sumber_listrik
type SumberListrik struct {
	// Key: Primary Key, Type: numeric(2,0) -> int16
	SumberListrikID int16 `gorm:"primaryKey;column:sumber_listrik_id;type:numeric(2,0);not null" json:"sumber_listrik_id"`

	// Informasi Sumber Listrik
	Nama string `gorm:"column:nama;type:varchar(50);not null" json:"nama"`

	// Timestamp System
	CreateDate  time.Time  `gorm:"column:create_date;type:timestamp;not null;default:CURRENT_TIMESTAMP" json:"create_date"`
	LastUpdate  time.Time  `gorm:"column:last_update;type:timestamp;not null;default:CURRENT_TIMESTAMP" json:"last_update"`
	ExpiredDate *time.Time `gorm:"column:expired_date;type:timestamp" json:"expired_date"`
	LastSync    time.Time  `gorm:"column:last_sync;type:timestamp;not null;default:'1901-01-01 00:00:00'" json:"last_sync"`
}

func (SumberListrik) TableName() string {
	return "ref.sumber_listrik"
}
