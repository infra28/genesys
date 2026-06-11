package models

import "time"

type JenjangKepengawasan struct {
	JenjangKepengawasanID int16  `gorm:"primaryKey;column:jenjang_kepengawasan_id;type:numeric(2,0);not null" json:"jenjang_kepengawasan_id"`
	Nama                  string `gorm:"column:nama;type:varchar(50);not null" json:"nama"`
	Tk                    int16  `gorm:"column:jenjang_kepengawasan_tk;type:numeric(1,0);not null" json:"jenjang_kepengawasan_tk"`
	Sd                    int16  `gorm:"column:jenjang_kepengawasan_sd;type:numeric(1,0);not null" json:"jenjang_kepengawasan_sd"`
	Smp                   int16  `gorm:"column:jenjang_kepengawasan_smp;type:numeric(1,0);not null" json:"jenjang_kepengawasan_smp"`
	Sma                   int16  `gorm:"column:jenjang_kepengawasan_sma;type:numeric(1,0);not null" json:"jenjang_kepengawasan_sma"`
	Smk                   int16  `gorm:"column:jenjang_kepengawasan_smk;type:numeric(1,0);not null" json:"jenjang_kepengawasan_smk"`
	Slb                   int16  `gorm:"column:jenjang_kepengawasan_slb;type:numeric(1,0);not null" json:"jenjang_kepengawasan_slb"`

	CreateDate  time.Time  `gorm:"column:create_date;type:timestamp;not null;default:CURRENT_TIMESTAMP" json:"create_date"`
	LastUpdate  time.Time  `gorm:"column:last_update;type:timestamp;not null;default:CURRENT_TIMESTAMP" json:"last_update"`
	ExpiredDate *time.Time `gorm:"column:expired_date;type:timestamp" json:"expired_date"`
	LastSync    time.Time  `gorm:"column:last_sync;type:timestamp;not null;default:'1901-01-01 00:00:00'" json:"last_sync"`
}

func (JenjangKepengawasan) TableName() string {
	return "ref.jenjang_kepengawasan"
}
