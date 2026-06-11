package models

import (
	"time"
)

// Penghasilan mewakili struktur tabel ref.penghasilan
type Penghasilan struct {
	// Key: Primary Key, Type: int4 -> int32
	PenghasilanID int32 `gorm:"primaryKey;column:penghasilan_id;type:int4;not null" json:"penghasilan_id"`

	// Informasi Kategori Penghasilan
	Nama       string `gorm:"column:nama;type:varchar(40);not null" json:"nama"`
	BatasBawah int32  `gorm:"column:batas_bawah;type:int4;not null" json:"batas_bawah"`
	BatasAtas  int32  `gorm:"column:batas_atas;type:int4;not null" json:"batas_atas"`

	// Timestamp System
	CreateDate  time.Time  `gorm:"column:create_date;type:timestamp;not null;default:CURRENT_TIMESTAMP" json:"create_date"`
	LastUpdate  time.Time  `gorm:"column:last_update;type:timestamp;not null;default:CURRENT_TIMESTAMP" json:"last_update"`
	ExpiredDate *time.Time `gorm:"column:expired_date;type:timestamp" json:"expired_date"`
	LastSync    time.Time  `gorm:"column:last_sync;type:timestamp;not null;default:'1901-01-01 00:00:00'" json:"last_sync"`
}

func (Penghasilan) TableName() string {
	return "ref.penghasilan"
}
