package models

import (
	"time"
)

// JenjangPendidikan mewakili struktur tabel ref.jenjang_pendidikan
type JenjangPendidikan struct {
	// Key: Primary Key, Type: numeric(2,0) -> int16
	JenjangPendidikanID int16 `gorm:"primaryKey;column:jenjang_pendidikan_id;type:numeric(2,0);not null" json:"jenjang_pendidikan_id"`

	// Informasi Jenjang
	Nama           string `gorm:"column:nama;type:varchar(25);not null" json:"nama"`
	JenjangLembaga int16  `gorm:"column:jenjang_lembaga;type:numeric(1,0);not null" json:"jenjang_lembaga"`
	JenjangOrang   int16  `gorm:"column:jenjang_orang;type:numeric(1,0);not null" json:"jenjang_orang"`

	// Timestamp System
	CreateDate  time.Time  `gorm:"column:create_date;type:timestamp;not null;default:CURRENT_TIMESTAMP" json:"create_date"`
	LastUpdate  time.Time  `gorm:"column:last_update;type:timestamp;not null;default:CURRENT_TIMESTAMP" json:"last_update"`
	ExpiredDate *time.Time `gorm:"column:expired_date;type:timestamp" json:"expired_date"`
	LastSync    time.Time  `gorm:"column:last_sync;type:timestamp;not null;default:'1901-01-01 00:00:00'" json:"last_sync"`
}

func (JenjangPendidikan) TableName() string {
	return "ref.jenjang_pendidikan"
}
