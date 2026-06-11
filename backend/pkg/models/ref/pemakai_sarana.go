package models

import (
	"time"
)

// PemakaiSarana mewakili struktur tabel ref.pemakai_sarana
type PemakaiSarana struct {
	// Composite Primary Key
	JenisSaranaID int32  `gorm:"primaryKey;column:jenis_sarana_id;type:int4;not null" json:"jenis_sarana_id"`
	JurusanID     string `gorm:"primaryKey;column:jurusan_id;type:varchar(25);not null" json:"jurusan_id"`

	// Foreign Keys
	// Relasi ke tabel master
	Jurusan     *Jurusan     `gorm:"foreignKey:JurusanID;references:JurusanID" json:"jurusan,omitempty"`
	JenisSarana *JenisSarana `gorm:"foreignKey:JenisSaranaID;references:JenisSaranaID" json:"jenis_sarana,omitempty"`

	// Timestamp System
	CreateDate  time.Time  `gorm:"column:create_date;type:timestamp;not null;default:CURRENT_TIMESTAMP" json:"create_date"`
	LastUpdate  time.Time  `gorm:"column:last_update;type:timestamp;not null;default:CURRENT_TIMESTAMP" json:"last_update"`
	ExpiredDate *time.Time `gorm:"column:expired_date;type:timestamp" json:"expired_date"`
	LastSync    time.Time  `gorm:"column:last_sync;type:timestamp;not null;default:'1901-01-01 00:00:00'" json:"last_sync"`
}

func (PemakaiSarana) TableName() string {
	return "ref.pemakai_sarana"
}
