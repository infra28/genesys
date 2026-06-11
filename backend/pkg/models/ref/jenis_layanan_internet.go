package models

import "time"

type JenisLayananInternet struct {
	JenisLayananInternetID int16      `gorm:"primaryKey;column:jenis_layanan_internet_id;type:numeric(2,0);not null" json:"jenis_layanan_internet_id"`
	JenisLayanan           string     `gorm:"column:jenis_layanan;type:varchar(30);not null" json:"jenis_layanan"`
	CreateDate             time.Time  `gorm:"column:create_date;type:timestamp;not null;default:CURRENT_TIMESTAMP" json:"create_date"`
	LastUpdate             time.Time  `gorm:"column:last_update;type:timestamp;not null;default:CURRENT_TIMESTAMP" json:"last_update"`
	ExpiredDate            *time.Time `gorm:"column:expired_date;type:timestamp" json:"expired_date"`
	LastSync               time.Time  `gorm:"column:last_sync;type:timestamp;not null;default:'1901-01-01 00:00:00'" json:"last_sync"`
}

func (JenisLayananInternet) TableName() string {
	return "ref.jenis_layanan_internet"
}
