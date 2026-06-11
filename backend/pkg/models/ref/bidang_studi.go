package models

import "time"

type BidangStudi struct {
	BidangStudiID         int32        `gorm:"primaryKey;column:bidang_studi_id;type:int4;not null" json:"bidang_studi_id"`
	KelompokBidangStudiID *int32       `gorm:"column:kelompok_bidang_studi_id;type:int4" json:"kelompok_bidang_studi_id"`
	Kode                  *string      `gorm:"column:kode;type:varchar(30)" json:"kode"`
	BidangStudi           string       `gorm:"column:bidang_studi;type:varchar(40);not null" json:"bidang_studi"`
	Kelompok              int16        `gorm:"column:kelompok;type:numeric(1,0);not null" json:"kelompok"`
	JenjangPaud           int16        `gorm:"column:jenjang_paud;type:numeric(1,0);not null" json:"jenjang_paud"`
	JenjangTk             int16        `gorm:"column:jenjang_tk;type:numeric(1,0);not null" json:"jenjang_tk"`
	JenjangSd             int16        `gorm:"column:jenjang_sd;type:numeric(1,0);not null" json:"jenjang_sd"`
	JenjangSmp            int16        `gorm:"column:jenjang_smp;type:numeric(1,0);not null" json:"jenjang_smp"`
	JenjangSma            int16        `gorm:"column:jenjang_sma;type:numeric(1,0);not null" json:"jenjang_sma"`
	JenjangTinggi         int16        `gorm:"column:jenjang_tinggi;type:numeric(1,0);not null" json:"jenjang_tinggi"`
	ASertKomp             int16        `gorm:"column:a_sert_komp;type:numeric(1,0);not null;default:0" json:"a_sert_komp"`
	ASertProfesi          int16        `gorm:"column:a_sert_profesi;type:numeric(1,0);not null;default:0" json:"a_sert_profesi"`
	CreateDate            time.Time    `gorm:"column:create_date;type:timestamp;not null;default:CURRENT_TIMESTAMP" json:"create_date"`
	LastUpdate            time.Time    `gorm:"column:last_update;type:timestamp;not null;default:CURRENT_TIMESTAMP" json:"last_update"`
	ExpiredDate           *time.Time   `gorm:"column:expired_date;type:timestamp" json:"expired_date"`
	LastSync              time.Time    `gorm:"column:last_sync;type:timestamp;not null;default:'1901-01-01 00:00:00'" json:"last_sync"`
	ParentBidangStudi     *BidangStudi `gorm:"foreignKey:KelompokBidangStudiID;references:BidangStudiID" json:"parent_bidang_studi,omitempty"`
}

func (BidangStudi) TableName() string {
	return "ref.bidang_studi"
}
