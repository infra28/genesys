package models

import (
	"time"
)

// StatusKeaktifanPegawai mewakili struktur tabel ref.status_keaktifan_pegawai
type StatusKeaktifanPegawai struct {
	// Key: Primary Key, Type: numeric(2,0) -> int16
	StatusKeaktifanID int16 `gorm:"primaryKey;column:status_keaktifan_id;type:numeric(2,0);not null" json:"status_keaktifan_id"`

	// Informasi Status
	Nama string `gorm:"column:nama;type:varchar(30);not null" json:"nama"`

	// Timestamp System
	CreateDate  time.Time  `gorm:"column:create_date;type:timestamp;not null;default:CURRENT_TIMESTAMP" json:"create_date"`
	LastUpdate  time.Time  `gorm:"column:last_update;type:timestamp;not null;default:CURRENT_TIMESTAMP" json:"last_update"`
	ExpiredDate *time.Time `gorm:"column:expired_date;type:timestamp" json:"expired_date"` // Allow Null
	LastSync    time.Time  `gorm:"column:last_sync;type:timestamp;not null;default:'1901-01-01 00:00:00'" json:"last_sync"`
}

func (StatusKeaktifanPegawai) TableName() string {
	return "ref.status_keaktifan_pegawai"
}
