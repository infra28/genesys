package models

import "time"

type BidangUsaha struct {
	BidangUsahaID    string     `gorm:"primaryKey;column:bidang_usaha_id;type:char(10);not null" json:"bidang_usaha_id"`
	NamaBidangUsaha  string     `gorm:"column:nama_bidang_usaha;type:varchar(40);not null" json:"nama_bidang_usaha"`
	LevelBidangUsaha *string    `gorm:"column:level_bidang_usaha;type:varchar(20)" json:"level_bidang_usaha"`
	CreateDate       time.Time  `gorm:"column:create_date;type:timestamp;not null;default:CURRENT_TIMESTAMP" json:"create_date"`
	LastUpdate       time.Time  `gorm:"column:last_update;type:timestamp;not null;default:CURRENT_TIMESTAMP" json:"last_update"`
	ExpiredDate      *time.Time `gorm:"column:expired_date;type:timestamp" json:"expired_date"`
	LastSync         time.Time  `gorm:"column:last_sync;type:timestamp;not null;default:'1901-01-01 00:00:00'" json:"last_sync"`
}

func (BidangUsaha) TableName() string {
	return "ref.bidang_usaha"
}
