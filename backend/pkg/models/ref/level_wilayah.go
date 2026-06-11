package models

import (
	"time"
)

// LevelWilayah mewakili struktur tabel ref.level_wilayah
type LevelWilayah struct {
	// Key: Primary Key, Type: int2 -> int16
	IDLevelWilayah int16 `gorm:"primaryKey;column:id_level_wilayah;type:int2;not null" json:"id_level_wilayah"`

	// Informasi Level
	// Menggunakan *string karena kolom tidak didefinisikan NOT NULL
	LevelWilayah *string `gorm:"column:level_wilayah;type:varchar(15)" json:"level_wilayah"`

	// Timestamp System
	CreateDate  time.Time  `gorm:"column:create_date;type:timestamp;not null;default:CURRENT_TIMESTAMP" json:"create_date"`
	LastUpdate  time.Time  `gorm:"column:last_update;type:timestamp;not null;default:CURRENT_TIMESTAMP" json:"last_update"`
	ExpiredDate *time.Time `gorm:"column:expired_date;type:timestamp" json:"expired_date"` // Allow Null
	LastSync    time.Time  `gorm:"column:last_sync;type:timestamp;not null;default:'1901-01-01 00:00:00'" json:"last_sync"`
}

func (LevelWilayah) TableName() string {
	return "ref.level_wilayah"
}
