package models

import "time"

type JenisPrasarana struct {
	JenisPrasaranaID int32      `gorm:"primaryKey;column:jenis_prasarana_id;type:int4;not null" json:"jenis_prasarana_id"`
	Nama             string     `gorm:"column:nama;type:varchar(60);not null" json:"nama"`
	AUnitOrganisasi  int16      `gorm:"column:a_unit_organisasi;type:numeric(1,0);not null;default:0" json:"a_unit_organisasi"`
	ATanah           int16      `gorm:"column:a_tanah;type:numeric(1,0);not null;default:0" json:"a_tanah"`
	ABangunan        int16      `gorm:"column:a_bangunan;type:numeric(1,0);not null;default:0" json:"a_bangunan"`
	ARuang           int16      `gorm:"column:a_ruang;type:numeric(1,0);not null;default:0" json:"a_ruang"`
	ASub             int16      `gorm:"column:a_sub;type:numeric(1,0);not null" json:"a_sub"`
	CreateDate       time.Time  `gorm:"column:create_date;type:timestamp;not null;default:CURRENT_TIMESTAMP" json:"create_date"`
	LastUpdate       time.Time  `gorm:"column:last_update;type:timestamp;not null;default:CURRENT_TIMESTAMP" json:"last_update"`
	ExpiredDate      *time.Time `gorm:"column:expired_date;type:timestamp" json:"expired_date"`
	LastSync         time.Time  `gorm:"column:last_sync;type:timestamp;not null;default:'1901-01-01 00:00:00'" json:"last_sync"`
}

func (JenisPrasarana) TableName() string {
	return "ref.jenis_prasarana"
}
