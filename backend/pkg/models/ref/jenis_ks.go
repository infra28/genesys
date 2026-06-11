package models

import "time"

type JenisKS struct {
	IDJnsKs  int32   `gorm:"primaryKey;column:id_jns_ks;type:numeric(6,0);not null" json:"id_jns_ks"`
	NmJnsKs  string  `gorm:"column:nm_jns_ks;type:varchar(100);not null" json:"nm_jns_ks"`
	KetJnsKs *string `gorm:"column:ket_jns_ks;type:varchar(250)" json:"ket_jns_ks"`

	CreateDate  time.Time  `gorm:"column:create_date;type:timestamp;not null;default:CURRENT_TIMESTAMP" json:"create_date"`
	LastUpdate  time.Time  `gorm:"column:last_update;type:timestamp;not null;default:CURRENT_TIMESTAMP" json:"last_update"`
	ExpiredDate *time.Time `gorm:"column:expired_date;type:timestamp" json:"expired_date"`
	LastSync    time.Time  `gorm:"column:last_sync;type:timestamp;not null;default:'1901-01-01 00:00:00'" json:"last_sync"`
}

func (JenisKS) TableName() string {
	return "ref.jenis_ks"
}
