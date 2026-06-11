package models

import (
	"time"
)

// Negara mewakili struktur tabel ref.negara
type Negara struct {
	// Key: Primary Key, Type: char(2)
	NegaraID string `gorm:"primaryKey;column:negara_id;type:char(2);not null" json:"negara_id"`

	// Informasi Negara
	Nama       string `gorm:"column:nama;type:varchar(45);not null" json:"nama"`
	LuarNegeri int16  `gorm:"column:luar_negeri;type:numeric(1,0);not null" json:"luar_negeri"`

	// Timestamp System
	CreateDate  time.Time  `gorm:"column:create_date;type:timestamp;not null;default:CURRENT_TIMESTAMP" json:"create_date"`
	LastUpdate  time.Time  `gorm:"column:last_update;type:timestamp;not null;default:CURRENT_TIMESTAMP" json:"last_update"`
	ExpiredDate *time.Time `gorm:"column:expired_date;type:timestamp" json:"expired_date"` // Allow Null
	LastSync    time.Time  `gorm:"column:last_sync;type:timestamp;not null;default:'1901-01-01 00:00:00'" json:"last_sync"`
}

func (Negara) TableName() string {
	return "ref.negara"
}
