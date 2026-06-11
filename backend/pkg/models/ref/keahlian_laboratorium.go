package models

import "time"

type KeahlianLaboratorium struct {
	KeahlianLaboratoriumID int16      `gorm:"primaryKey;column:keahlian_laboratorium_id;type:int2;not null" json:"keahlian_laboratorium_id"`
	Nama                   string     `gorm:"column:nama;type:varchar(50);not null" json:"nama"`
	CreateDate             time.Time  `gorm:"column:create_date;type:timestamp;not null;default:CURRENT_TIMESTAMP" json:"create_date"`
	LastUpdate             time.Time  `gorm:"column:last_update;type:timestamp;not null;default:CURRENT_TIMESTAMP" json:"last_update"`
	ExpiredDate            *time.Time `gorm:"column:expired_date;type:timestamp" json:"expired_date"`
	LastSync               time.Time  `gorm:"column:last_sync;type:timestamp;not null;default:'1901-01-01 00:00:00'" json:"last_sync"`
}

func (KeahlianLaboratorium) TableName() string {
	return "ref.keahlian_laboratorium"
}
