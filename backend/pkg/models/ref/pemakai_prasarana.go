package models

import (
	"time"
)

// PemakaiPrasarana mewakili struktur tabel ref.pemakai_prasarana
type PemakaiPrasarana struct {
	// Composite Primary Key
	JenisPrasaranaID int32  `gorm:"primaryKey;column:jenis_prasarana_id;type:int4;not null" json:"jenis_prasarana_id"`
	JurusanID        string `gorm:"primaryKey;column:jurusan_id;type:varchar(25);not null" json:"jurusan_id"`

	// Informasi Spesifikasi
	JmlStdMin int32 `gorm:"column:jml_std_min;type:numeric(5,0);not null;default:0" json:"jml_std_min"`

	// Foreign Keys
	// Relasi ke tabel master untuk mempermudah operasional Preload
	Jurusan        *Jurusan        `gorm:"foreignKey:JurusanID;references:JurusanID" json:"jurusan,omitempty"`
	JenisPrasarana *JenisPrasarana `gorm:"foreignKey:JenisPrasaranaID;references:JenisPrasaranaID" json:"jenis_prasarana,omitempty"`

	// Timestamp System
	CreateDate  time.Time  `gorm:"column:create_date;type:timestamp;not null;default:CURRENT_TIMESTAMP" json:"create_date"`
	LastUpdate  time.Time  `gorm:"column:last_update;type:timestamp;not null;default:CURRENT_TIMESTAMP" json:"last_update"`
	ExpiredDate *time.Time `gorm:"column:expired_date;type:timestamp" json:"expired_date"`
	LastSync    time.Time  `gorm:"column:last_sync;type:timestamp;not null;default:'1901-01-01 00:00:00'" json:"last_sync"`
}

func (PemakaiPrasarana) TableName() string {
	return "ref.pemakai_prasarana"
}
