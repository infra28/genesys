package models

import "time"

type KelompokBidang struct {
	LevelBidangID        string          `gorm:"primaryKey;column:level_bidang_id;type:varchar(5);not null" json:"level_bidang_id"`
	NamaLevelBidang      string          `gorm:"column:nama_level_bidang;type:varchar(100);not null" json:"nama_level_bidang"`
	UntukSma             int16           `gorm:"column:untuk_sma;type:numeric(1,0);not null" json:"untuk_sma"`
	UntukSmk             int16           `gorm:"column:untuk_smk;type:numeric(1,0);not null" json:"untuk_smk"`
	UntukPt              int16           `gorm:"column:untuk_pt;type:numeric(1,0);not null" json:"untuk_pt"`
	UntukSlb             int16           `gorm:"column:untuk_slb;type:numeric(1,0);not null;default:0" json:"untuk_slb"`
	UntukSmklb           int16           `gorm:"column:untuk_smklb;type:numeric(1,0);not null;default:0" json:"untuk_smklb"`
	LevelBidangInduk     *string         `gorm:"column:level_bidang_induk;type:varchar(5)" json:"level_bidang_induk"`
	CreateDate           time.Time       `gorm:"column:create_date;type:timestamp;not null;default:CURRENT_TIMESTAMP" json:"create_date"`
	LastUpdate           time.Time       `gorm:"column:last_update;type:timestamp;not null;default:CURRENT_TIMESTAMP" json:"last_update"`
	ExpiredDate          *time.Time      `gorm:"column:expired_date;type:timestamp" json:"expired_date"`
	LastSync             time.Time       `gorm:"column:last_sync;type:timestamp;not null;default:'1901-01-01 00:00:00'" json:"last_sync"`
	ParentKelompokBidang *KelompokBidang `gorm:"foreignKey:LevelBidangInduk;references:LevelBidangID" json:"parent_kelompok_bidang,omitempty"`
}

func (KelompokBidang) TableName() string {
	return "ref.kelompok_bidang"
}
