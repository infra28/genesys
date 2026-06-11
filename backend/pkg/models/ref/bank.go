package models

import "time"

// Bank mewakili struktur tabel ref.bank berdasarkan spesifikasi database
type Bank struct {
	BankID      string     `gorm:"primaryKey;column:bank_id;type:char(3);not null" json:"bank_id"`
	Nama        string     `gorm:"column:nama;type:varchar(20);not null" json:"nama"`
	CreateDate  time.Time  `gorm:"column:create_date;type:timestamp;not null;default:CURRENT_TIMESTAMP" json:"create_date"`
	LastUpdate  time.Time  `gorm:"column:last_update;type:timestamp;not null;default:CURRENT_TIMESTAMP" json:"last_update"`
	ExpiredDate *time.Time `gorm:"column:expired_date;type:timestamp" json:"expired_date"`
	LastSync    time.Time  `gorm:"column:last_sync;type:timestamp;not null;default:'1901-01-01 00:00:00'" json:"last_sync"`
}

func (Bank) TableName() string {
	return "ref.bank"
}
