package models

import "time"

type JenisKerusakan struct {
	KerusakanID int16  `gorm:"primaryKey;column:kerusakan_id;type:numeric(2,0);not null" json:"kerusakan_id"`
	Klasifikasi string `gorm:"column:klasifikasi;type:varchar(30);not null" json:"klasifikasi"`
	UBangunan   int16  `gorm:"column:u_bangunan;type:numeric(1,0);not null" json:"u_bangunan"`
	URuang      int16  `gorm:"column:u_ruang;type:numeric(1,0);not null" json:"u_ruang"`

	CreateDate  time.Time  `gorm:"column:create_date;type:timestamp;not null;default:CURRENT_TIMESTAMP" json:"create_date"`
	LastUpdate  time.Time  `gorm:"column:last_update;type:timestamp;not null;default:CURRENT_TIMESTAMP" json:"last_update"`
	ExpiredDate *time.Time `gorm:"column:expired_date;type:timestamp" json:"expired_date"`
	LastSync    time.Time  `gorm:"column:last_sync;type:timestamp;not null;default:'1901-01-01 00:00:00'" json:"last_sync"`
}

func (JenisKerusakan) TableName() string {
	return "ref.jenis_kerusakan"
}
