package models

import "time"

type ErrorType struct {
	IDType        int32      `gorm:"primaryKey;column:idtype;type:int4;not null" json:"idtype"`
	KategoriError *int32     `gorm:"column:kategori_error;type:int4" json:"kategori_error"`
	Keterangan    *string    `gorm:"column:keterangan;type:varchar(255)" json:"keterangan"`
	CreateDate    time.Time  `gorm:"column:create_date;type:timestamp;not null;default:CURRENT_TIMESTAMP" json:"create_date"`
	LastUpdate    time.Time  `gorm:"column:last_update;type:timestamp;not null;default:CURRENT_TIMESTAMP" json:"last_update"`
	ExpiredDate   *time.Time `gorm:"column:expired_date;type:timestamp" json:"expired_date"`
	LastSync      time.Time  `gorm:"column:last_sync;type:timestamp;not null;default:'1901-01-01 00:00:00'" json:"last_sync"`
}

func (ErrorType) TableName() string {
	return "ref.errortype"
}
