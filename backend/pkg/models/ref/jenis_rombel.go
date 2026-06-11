package models

import "time"

type JenisRombel struct {
	JenisRombel   int16      `gorm:"primaryKey;column:jenis_rombel;type:numeric(2,0);not null" json:"jenis_rombel"`
	NmJenisRombel string     `gorm:"column:nm_jenis_rombel;type:varchar(80);not null" json:"nm_jenis_rombel"`
	CreateDate    time.Time  `gorm:"column:create_date;type:timestamp;not null;default:CURRENT_TIMESTAMP" json:"create_date"`
	LastUpdate    time.Time  `gorm:"column:last_update;type:timestamp;not null;default:CURRENT_TIMESTAMP" json:"last_update"`
	ExpiredDate   *time.Time `gorm:"column:expired_date;type:timestamp" json:"expired_date"`
	LastSync      time.Time  `gorm:"column:last_sync;type:timestamp;not null;default:'1901-01-01 00:00:00'" json:"last_sync"`
}

func (JenisRombel) TableName() string {
	return "ref.jenis_rombel"
}
