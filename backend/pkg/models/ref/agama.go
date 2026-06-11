package models

import "time"

type Agama struct {
	AgamaID     int16      `gorm:"primaryKey;column:agama_id;type:smallint;not null" json:"agama_id"`
	Nama        string     `gorm:"column:nama;type:varchar(25);not null" json:"nama"`
	CreateDate  time.Time  `gorm:"column:create_date;type:timestamp;not null;default:CURRENT_TIMESTAMP" json:"create_date"`
	LastUpdate  time.Time  `gorm:"column:last_update;type:timestamp;not null;default:CURRENT_TIMESTAMP" json:"last_update"`
	ExpiredDate *time.Time `gorm:"column:expired_date;type:timestamp" json:"expired_date"`
	LastSync    time.Time  `gorm:"column:last_sync;type:timestamp;not null;default:'1901-01-01 00:00:00'" json:"last_sync"`
}

// TableName menentukan nama skema dan tabel spesifik di PostgreSQL
func (Agama) TableName() string {
	return "ref.agama"
}
