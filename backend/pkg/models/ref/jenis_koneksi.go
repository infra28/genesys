package models

import "time"

type JenisKoneksi struct {
	JenisKoneksiID int16      `gorm:"primaryKey;column:jenis_koneksi_id;type:numeric(2,0);not null" json:"jenis_koneksi_id"`
	JenisKoneksi   string     `gorm:"column:jenis_koneksi;type:varchar(30);not null" json:"jenis_koneksi"`
	CreateDate     time.Time  `gorm:"column:create_date;type:timestamp;not null;default:CURRENT_TIMESTAMP" json:"create_date"`
	LastUpdate     time.Time  `gorm:"column:last_update;type:timestamp;not null;default:CURRENT_TIMESTAMP" json:"last_update"`
	ExpiredDate    *time.Time `gorm:"column:expired_date;type:timestamp" json:"expired_date"`
	LastSync       time.Time  `gorm:"column:last_sync;type:timestamp;not null;default:'1901-01-01 00:00:00'" json:"last_sync"`
}

func (JenisKoneksi) TableName() string {
	return "ref.jenis_koneksi"
}
