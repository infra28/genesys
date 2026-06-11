package models

import "time"

type JenisPtk struct {
	JenisPtkID  int16      `gorm:"primaryKey;column:jenis_ptk_id;type:numeric(2,0);not null" json:"jenis_ptk_id"`
	JenisPtk    string     `gorm:"column:jenis_ptk;type:varchar(30);not null" json:"jenis_ptk"`
	CreateDate  time.Time  `gorm:"column:create_date;type:timestamp;not null;default:CURRENT_TIMESTAMP" json:"create_date"`
	LastUpdate  time.Time  `gorm:"column:last_update;type:timestamp;not null;default:CURRENT_TIMESTAMP" json:"last_update"`
	ExpiredDate *time.Time `gorm:"column:expired_date;type:timestamp" json:"expired_date"`
	LastSync    time.Time  `gorm:"column:last_sync;type:timestamp;not null;default:'1901-01-01 00:00:00'" json:"last_sync"`
	// JabatanPtks []JabatanPtk `gorm:"foreignKey:JenisPtkID;references:JenisPtkID" json:"jabatan_ptks,omitempty"`
}

func (JenisPtk) TableName() string {
	return "ref.jenis_ptk"
}
