package models

import "time"

type AlasanLayakPip struct {
	IDLayakPip     int16      `gorm:"primaryKey;column:id_layak_pip;type:numeric(2,0);not null" json:"id_layak_pip"`
	AlasanLayakPip string     `gorm:"column:alasan_layak_pip;type:varchar(100);not null" json:"alasan_layak_pip"`
	CreateDate     time.Time  `gorm:"column:create_date;type:timestamp;not null;default:CURRENT_TIMESTAMP" json:"create_date"`
	LastUpdate     time.Time  `gorm:"column:last_update;type:timestamp;not null;default:CURRENT_TIMESTAMP" json:"last_update"`
	ExpiredDate    *time.Time `gorm:"column:expired_date;type:timestamp" json:"expired_date"`
	LastSync       time.Time  `gorm:"column:last_sync;type:timestamp;not null;default:'1901-01-01 00:00:00'" json:"last_sync"`
}

func (AlasanLayakPip) TableName() string {
	return "ref.alasan_layak_pip"
}
