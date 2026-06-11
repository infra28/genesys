package models

import "time"

type SumberDana struct {
	SumberDanaID int16  `gorm:"primaryKey;column:sumber_dana_id;type:numeric(3,0);not null" json:"sumber_dana_id"`
	Nama         string `gorm:"column:nama;type:varchar(50);not null" json:"nama"`

	CreateDate  time.Time  `gorm:"column:create_date;type:timestamp;not null;default:CURRENT_TIMESTAMP" json:"create_date"`
	LastUpdate  time.Time  `gorm:"column:last_update;type:timestamp;not null;default:CURRENT_TIMESTAMP" json:"last_update"`
	ExpiredDate *time.Time `gorm:"column:expired_date;type:timestamp" json:"expired_date"`
	LastSync    time.Time  `gorm:"column:last_sync;type:timestamp;not null;default:'1901-01-01 00:00:00'" json:"last_sync"`
	// JenisBeasiswas []JenisBeasiswa `gorm:"foreignKey:SumberDanaID;references:SumberDanaID" json:"jenis_beasiswas,omitempty"`
}

func (SumberDana) TableName() string {
	return "ref.sumber_dana"
}
