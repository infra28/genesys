package models

import "time"

type EkstraKurikuler struct {
	IDEkskul    int32      `gorm:"primaryKey;column:id_ekskul;type:int4;not null" json:"id_ekskul"`
	NmEkskul    string     `gorm:"column:nm_ekskul;type:varchar(80);not null" json:"nm_ekskul"`
	USd         int16      `gorm:"column:u_sd;type:numeric(1,0);not null;default:0" json:"u_sd"`
	USmp        int16      `gorm:"column:u_smp;type:numeric(1,0);not null;default:0" json:"u_smp"`
	USma        int16      `gorm:"column:u_sma;type:numeric(1,0);not null;default:0" json:"u_sma"`
	USmk        int16      `gorm:"column:u_smk;type:numeric(1,0);not null;default:0" json:"u_smk"`
	CreateDate  time.Time  `gorm:"column:create_date;type:timestamp;not null;default:CURRENT_TIMESTAMP" json:"create_date"`
	LastUpdate  time.Time  `gorm:"column:last_update;type:timestamp;not null;default:CURRENT_TIMESTAMP" json:"last_update"`
	ExpiredDate *time.Time `gorm:"column:expired_date;type:timestamp" json:"expired_date"`
	LastSync    time.Time  `gorm:"column:last_sync;type:timestamp;not null;default:'1901-01-01 00:00:00'" json:"last_sync"`
}

func (EkstraKurikuler) TableName() string {
	return "ref.ekstra_kurikuler"
}
