package models

import (
	"time"
)

// Pekerjaan mewakili struktur tabel ref.pekerjaan
type Pekerjaan struct {
	// Key: Primary Key, Type: int4 -> int32
	PekerjaanID int32 `gorm:"primaryKey;column:pekerjaan_id;type:int4;not null" json:"pekerjaan_id"`

	// Informasi Pekerjaan
	Nama           *string `gorm:"column:nama;type:varchar(25)" json:"nama"`
	AWirausaha     int16   `gorm:"column:a_wirausaha;type:numeric(1,0);not null;default:0" json:"a_wirausaha"`
	APejabatPublik int16   `gorm:"column:a_pejabat_publik;type:numeric(1,0);not null;default:0" json:"a_pejabat_publik"`

	// Timestamp System
	CreateDate  time.Time  `gorm:"column:create_date;type:timestamp;not null;default:CURRENT_TIMESTAMP" json:"create_date"`
	LastUpdate  time.Time  `gorm:"column:last_update;type:timestamp;not null;default:CURRENT_TIMESTAMP" json:"last_update"`
	ExpiredDate *time.Time `gorm:"column:expired_date;type:timestamp" json:"expired_date"` // Allow Null
	LastSync    time.Time  `gorm:"column:last_sync;type:timestamp;not null;default:'1901-01-01 00:00:00'" json:"last_sync"`
}

func (Pekerjaan) TableName() string {
	return "ref.pekerjaan"
}
