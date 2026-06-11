package models

import "time"

type KelompokUsaha struct {
	KelompokUsahaID   string     `gorm:"primaryKey;column:kelompok_usaha_id;type:char(8);not null" json:"kelompok_usaha_id"`
	NamaKelompokUsaha string     `gorm:"column:nama_kelompok_usaha;type:varchar(60);not null" json:"nama_kelompok_usaha"`
	CreateDate        time.Time  `gorm:"column:create_date;type:timestamp;not null;default:CURRENT_TIMESTAMP" json:"create_date"`
	LastUpdate        time.Time  `gorm:"column:last_update;type:timestamp;not null;default:CURRENT_TIMESTAMP" json:"last_update"`
	ExpiredDate       *time.Time `gorm:"column:expired_date;type:timestamp" json:"expired_date"`
	LastSync          time.Time  `gorm:"column:last_sync;type:timestamp;not null;default:'1901-01-01 00:00:00'" json:"last_sync"`
}

func (KelompokUsaha) TableName() string {
	return "ref.kelompok_usaha"
}
