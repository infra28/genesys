package models

import "time"

type JabatanPtk struct {
	JabatanPtkID int32      `gorm:"primaryKey;column:jabatan_ptk_id;type:numeric(5,0);not null" json:"jabatan_ptk_id"`
	JenisPtkID   int16      `gorm:"column:jenis_ptk_id;type:numeric(2,0);not null" json:"jenis_ptk_id"`
	JabatanPtk   string     `gorm:"column:jabatan_ptk;type:varchar(100);not null" json:"jabatan_ptk"`
	JabatanKode  *string    `gorm:"column:jabatan_kode;type:varchar(20)" json:"jabatan_kode"`
	CreateDate   time.Time  `gorm:"column:create_date;type:timestamp;not null;default:CURRENT_TIMESTAMP" json:"create_date"`
	LastUpdate   time.Time  `gorm:"column:last_update;type:timestamp;not null;default:CURRENT_TIMESTAMP" json:"last_update"`
	ExpiredDate  *time.Time `gorm:"column:expired_date;type:timestamp" json:"expired_date"`
	LastSync     time.Time  `gorm:"column:last_sync;type:timestamp;not null;default:'1901-01-01 00:00:00'" json:"last_sync"`
	JenisPtk     *JenisPtk  `gorm:"foreignKey:JenisPtkID;references:JenisPtkID" json:"jenis_ptk,omitempty"`
}

func (JabatanPtk) TableName() string {
	return "ref.jabatan_ptk"
}
