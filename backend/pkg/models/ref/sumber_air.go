package models

import (
	"time"
)

// SumberAir mewakili struktur tabel ref.sumber_air
type SumberAir struct {
	// Key: Primary Key, Type: numeric(2,0) -> int16
	SumberAirID int16 `gorm:"primaryKey;column:sumber_air_id;type:numeric(2,0);not null" json:"sumber_air_id"`

	// Informasi Sumber Air
	Nama        string `gorm:"column:nama;type:varchar(25);not null" json:"nama"`
	SumberAir   *int16 `gorm:"column:sumber_air;type:numeric(1,0)" json:"sumber_air"`
	SumberMinum *int16 `gorm:"column:sumber_minum;type:numeric(1,0)" json:"sumber_minum"`

	// Timestamp System
	CreateDate  *time.Time `gorm:"column:create_date;type:timestamp" json:"create_date"`
	LastUpdate  *time.Time `gorm:"column:last_update;type:timestamp" json:"last_update"`
	ExpiredDate *time.Time `gorm:"column:expired_date;type:timestamp" json:"expired_date"`
	LastSync    *time.Time `gorm:"column:last_sync;type:timestamp" json:"last_sync"`
}

func (SumberAir) TableName() string {
	return "ref.sumber_air"
}
