package models

import "time"

type JenisSarana struct {
	JenisSaranaID       int32      `gorm:"primaryKey;column:jenis_sarana_id;type:int4;not null" json:"jenis_sarana_id"`
	Nama                string     `gorm:"column:nama;type:varchar(60);not null" json:"nama"`
	Kelompok            *string    `gorm:"column:kelompok;type:varchar(50)" json:"kelompok"`
	Keterangan          *string    `gorm:"column:keterangan;type:varchar(128)" json:"keterangan"`
	PerluPenempatan     int16      `gorm:"column:perlu_penempatan;type:numeric(1,0);not null" json:"perlu_penempatan"`
	AAlat               int16      `gorm:"column:a_alat;type:numeric(1,0);not null;default:0" json:"a_alat"`
	AAngkutan           int16      `gorm:"column:a_angkutan;type:numeric(1,0);not null;default:0" json:"a_angkutan"`
	SpmQtyMinPerSiswa   float32    `gorm:"column:spm_qty_min_per_siswa;type:numeric(3,1);not null;default:-1" json:"spm_qty_min_per_siswa"`
	SpmQtyMinPerSekolah int16      `gorm:"column:spm_qty_min_per_sekolah;type:numeric(4,0);not null;default:-1" json:"spm_qty_min_per_sekolah"`
	CreateDate          time.Time  `gorm:"column:create_date;type:timestamp;not null;default:CURRENT_TIMESTAMP" json:"create_date"`
	LastUpdate          time.Time  `gorm:"column:last_update;type:timestamp;not null;default:CURRENT_TIMESTAMP" json:"last_update"`
	ExpiredDate         *time.Time `gorm:"column:expired_date;type:timestamp" json:"expired_date"`
	LastSync            time.Time  `gorm:"column:last_sync;type:timestamp;not null;default:'1901-01-01 00:00:00'" json:"last_sync"`
}

func (JenisSarana) TableName() string {
	return "ref.jenis_sarana"
}
