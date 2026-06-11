package models

import "time"

type BentukPendidikan struct {
	BentukPendidikanID   int16      `gorm:"primaryKey;column:bentuk_pendidikan_id;type:int2;not null" json:"bentuk_pendidikan_id"`
	Nama                 string     `gorm:"column:nama;type:varchar(50);not null" json:"nama"`
	JenjangPaud          int16      `gorm:"column:jenjang_paud;type:numeric(1,0);not null" json:"jenjang_paud"`
	JenjangTk            int16      `gorm:"column:jenjang_tk;type:numeric(1,0);not null" json:"jenjang_tk"`
	JenjangSd            int16      `gorm:"column:jenjang_sd;type:numeric(1,0);not null" json:"jenjang_sd"`
	JenjangSmp           int16      `gorm:"column:jenjang_smp;type:numeric(1,0);not null" json:"jenjang_smp"`
	JenjangSma           int16      `gorm:"column:jenjang_sma;type:numeric(1,0);not null" json:"jenjang_sma"`
	JenjangTinggi        int16      `gorm:"column:jenjang_tinggi;type:numeric(1,0);not null" json:"jenjang_tinggi"`
	DirektoratPembinaan  *string    `gorm:"column:direktorat_pembinaan;type:varchar(40)" json:"direktorat_pembinaan"`
	Aktif                int16      `gorm:"column:aktif;type:numeric(1,0);not null" json:"aktif"`
	FormalitasPendidikan string     `gorm:"column:formalitas_pendidikan;type:char(1);not null" json:"formalitas_pendidikan"`
	CreateDate           time.Time  `gorm:"column:create_date;type:timestamp;not null;default:CURRENT_TIMESTAMP" json:"create_date"`
	LastUpdate           time.Time  `gorm:"column:last_update;type:timestamp;not null;default:CURRENT_TIMESTAMP" json:"last_update"`
	ExpiredDate          *time.Time `gorm:"column:expired_date;type:timestamp" json:"expired_date"`
	LastSync             time.Time  `gorm:"column:last_sync;type:timestamp;not null;default:'1901-01-01 00:00:00'" json:"last_sync"`
}

func (BentukPendidikan) TableName() string {
	return "ref.bentuk_pendidikan"
}
