package models

import "time"

type GelarAkademik struct {
	GelarAkademikID int32      `gorm:"primaryKey;column:gelar_akademik_id;type:int4;not null" json:"gelar_akademik_id"`
	Kode            string     `gorm:"column:kode;type:varchar(10);not null" json:"kode"`
	Nama            string     `gorm:"column:nama;type:varchar(40);not null" json:"nama"`
	PosisiGelar     int16      `gorm:"column:posisi_gelar;type:numeric(1,0);not null" json:"posisi_gelar"`
	CreateDate      time.Time  `gorm:"column:create_date;type:timestamp;not null;default:CURRENT_TIMESTAMP" json:"create_date"`
	LastUpdate      time.Time  `gorm:"column:last_update;type:timestamp;not null;default:CURRENT_TIMESTAMP" json:"last_update"`
	ExpiredDate     *time.Time `gorm:"column:expired_date;type:timestamp" json:"expired_date"`
	LastSync        time.Time  `gorm:"column:last_sync;type:timestamp;not null;default:'1901-01-01 00:00:00'" json:"last_sync"`
}

func (GelarAkademik) TableName() string {
	return "ref.gelar_akademik"
}
