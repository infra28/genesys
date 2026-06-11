package models

import "time"

type JenisLK struct {
	IDJenisLK   string     `gorm:"primaryKey;column:id_jenis_lk;type:char(6);not null" json:"id_jenis_lk"`
	NmJenisLK   string     `gorm:"column:nm_jenis_lk;type:varchar(160);not null" json:"nm_jenis_lk"`
	KetJenisLK  *string    `gorm:"column:ket_jenis_lk;type:varchar(200)" json:"ket_jenis_lk"`
	CreateDate  time.Time  `gorm:"column:create_date;type:timestamp;not null;default:CURRENT_TIMESTAMP" json:"create_date"`
	LastUpdate  time.Time  `gorm:"column:last_update;type:timestamp;not null;default:CURRENT_TIMESTAMP" json:"last_update"`
	ExpiredDate *time.Time `gorm:"column:expired_date;type:timestamp" json:"expired_date"`
	LastSync    time.Time  `gorm:"column:last_sync;type:timestamp;not null;default:'1901-01-01 00:00:00'" json:"last_sync"`
}

func (JenisLK) TableName() string {
	return "ref.jenis_lk"
}
