package models

import "time"

type JenisKeluar struct {
	JenisKeluarID string `gorm:"primaryKey;column:jenis_keluar_id;type:char(1);not null" json:"jenis_keluar_id"`

	KetKeluar string `gorm:"column:ket_keluar;type:varchar(40);not null" json:"ket_keluar"`

	KeluarPd int16 `gorm:"column:keluar_pd;type:numeric(1,0);not null" json:"keluar_pd"`

	KeluarPtk int16 `gorm:"column:keluar_ptk;type:numeric(1,0);not null" json:"keluar_ptk"`

	CreateDate  time.Time  `gorm:"column:create_date;type:timestamp;not null;default:CURRENT_TIMESTAMP" json:"create_date"`
	LastUpdate  time.Time  `gorm:"column:last_update;type:timestamp;not null;default:CURRENT_TIMESTAMP" json:"last_update"`
	ExpiredDate *time.Time `gorm:"column:expired_date;type:timestamp" json:"expired_date"`
	LastSync    time.Time  `gorm:"column:last_sync;type:timestamp;not null;default:'1901-01-01 00:00:00'" json:"last_sync"`
}

func (JenisKeluar) TableName() string {
	return "ref.jenis_keluar"
}
