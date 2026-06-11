package models

import (
	"time"
)

// Mulok mewakili struktur tabel ref.mulok
type Mulok struct {
	// Composite Primary Key
	KodeWilayah     string `gorm:"primaryKey;column:kode_wilayah;type:char(8);not null" json:"kode_wilayah"`
	MataPelajaranID int32  `gorm:"primaryKey;column:mata_pelajaran_id;type:int4;not null" json:"mata_pelajaran_id"`

	// Informasi Muatan Lokal
	SkMulok    string    `gorm:"column:sk_mulok;type:varchar(80);not null" json:"sk_mulok"`
	TglSkMulok time.Time `gorm:"column:tgl_sk_mulok;type:date;not null" json:"tgl_sk_mulok"`

	// Foreign Keys
	// Menambahkan relasi ke tabel master untuk mempermudah Preload
	MstWilayah    *MstWilayah    `gorm:"foreignKey:KodeWilayah;references:KodeWilayah" json:"mst_wilayah,omitempty"`
	MataPelajaran *MataPelajaran `gorm:"foreignKey:MataPelajaranID;references:MataPelajaranID" json:"mata_pelajaran,omitempty"`

	// Timestamp System
	CreateDate  time.Time  `gorm:"column:create_date;type:timestamp;not null;default:CURRENT_TIMESTAMP" json:"create_date"`
	LastUpdate  time.Time  `gorm:"column:last_update;type:timestamp;not null;default:CURRENT_TIMESTAMP" json:"last_update"`
	ExpiredDate *time.Time `gorm:"column:expired_date;type:timestamp" json:"expired_date"`
	LastSync    time.Time  `gorm:"column:last_sync;type:timestamp;not null;default:'1901-01-01 00:00:00'" json:"last_sync"`
}

func (Mulok) TableName() string {
	return "ref.mulok"
}
