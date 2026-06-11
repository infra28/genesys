package models

import (
	"time"
)

// VariabelValue mewakili struktur tabel ref.variabel_value
type VariabelValue struct {
	// ID Variabel (UUID/String)
	VariabelID string `gorm:"column:variabel_id;type:char(36);not null" json:"variabel_id"`

	// ID Nilai (Primary Key atau bagian dari Composite Key)
	ValueID int32 `gorm:"column:value_id;type:int4;not null" json:"value_id"`

	// Nama/Label Nilai
	ValueName string `gorm:"column:value_name;type:varchar(200);not null" json:"value_name"`

	// Timestamp System
	CreateDate  time.Time  `gorm:"column:create_date;type:timestamp;not null" json:"create_date"`
	LastUpdate  time.Time  `gorm:"column:last_update;type:timestamp;not null" json:"last_update"`
	ExpiredDate *time.Time `gorm:"column:expired_date;type:timestamp" json:"expired_date"`
	LastSync    time.Time  `gorm:"column:last_sync;type:timestamp;not null" json:"last_sync"`
}

func (VariabelValue) TableName() string {
	return "ref.variabel_value"
}
