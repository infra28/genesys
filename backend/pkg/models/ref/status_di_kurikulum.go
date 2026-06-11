package models

import (
	"time"
)

// StatusDiKurikulum mewakili struktur tabel ref.status_di_kurikulum
type StatusDiKurikulum struct {
	// Key: Primary Key, Type: numeric(2,0) -> int16
	StatusDiKurikulumID int16 `gorm:"primaryKey;column:status_di_kurikulum;type:numeric(2,0);not null" json:"status_di_kurikulum"`

	// Informasi Status
	KetStatDiKurikulum string `gorm:"column:ket_stat_di_kurikulum;type:varchar(40);not null" json:"ket_stat_di_kurikulum"`

	// Timestamp System
	CreateDate  time.Time  `gorm:"column:create_date;type:timestamp;not null;default:CURRENT_TIMESTAMP" json:"create_date"`
	LastUpdate  time.Time  `gorm:"column:last_update;type:timestamp;not null;default:CURRENT_TIMESTAMP" json:"last_update"`
	ExpiredDate *time.Time `gorm:"column:expired_date;type:timestamp" json:"expired_date"`
	LastSync    time.Time  `gorm:"column:last_sync;type:timestamp;not null;default:'1901-01-01 00:00:00'" json:"last_sync"`
}

func (StatusDiKurikulum) TableName() string {
	return "ref.status_di_kurikulum"
}
