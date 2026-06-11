package models

import "time"

type JenisAktPd struct {
	IDJnsAktPd  int16      `gorm:"primaryKey;column:id_jns_akt_pd;type:numeric(3,0);not null" json:"id_jns_akt_pd"`
	NmJnsAktPd  string     `gorm:"column:nm_jns_akt_pd;type:varchar(40);not null" json:"nm_jns_akt_pd"`
	KetJnsAktPd *string    `gorm:"column:ket_jns_akt_pd;type:varchar(100)" json:"ket_jns_akt_pd"`
	CreateDate  time.Time  `gorm:"column:create_date;type:timestamp;not null;default:CURRENT_TIMESTAMP" json:"create_date"`
	LastUpdate  time.Time  `gorm:"column:last_update;type:timestamp;not null;default:CURRENT_TIMESTAMP" json:"last_update"`
	ExpiredDate *time.Time `gorm:"column:expired_date;type:timestamp" json:"expired_date"`
	LastSync    time.Time  `gorm:"column:last_sync;type:timestamp;not null;default:'1901-01-01 00:00:00'" json:"last_sync"`
}

func (JenisAktPd) TableName() string {
	return "ref.jenis_akt_pd"
}
