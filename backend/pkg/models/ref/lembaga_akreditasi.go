package models

import (
	"time"
)

type LembagaAkreditasi struct {
	LaID string `gorm:"primaryKey;column:la_id;type:char(5);not null" json:"la_id"`

	Nama       string    `gorm:"column:nama;type:varchar(100);not null" json:"nama"`
	LaTglMulai time.Time `gorm:"column:la_tgl_mulai;type:date;not null" json:"la_tgl_mulai"`
	LaKet      *string   `gorm:"column:la_ket;type:varchar(250)" json:"la_ket"`

	AlamatJalan   string  `gorm:"column:alamat_jalan;type:varchar(80);not null" json:"alamat_jalan"`
	Rt            *int16  `gorm:"column:rt;type:numeric(2,0)" json:"rt"`
	Rw            *int16  `gorm:"column:rw;type:numeric(2,0)" json:"rw"`
	NamaDusun     *string `gorm:"column:nama_dusun;type:varchar(60)" json:"nama_dusun"`
	DesaKelurahan string  `gorm:"column:desa_kelurahan;type:varchar(60);not null" json:"desa_kelurahan"`

	KodeWilayah string      `gorm:"column:kode_wilayah;type:char(8);not null" json:"kode_wilayah"`
	MstWilayah  *MstWilayah `gorm:"foreignKey:KodeWilayah;references:KodeWilayah" json:"mst_wilayah,omitempty"`

	KodePos *string  `gorm:"column:kode_pos;type:char(5)" json:"kode_pos"`
	Lintang *float64 `gorm:"column:lintang;type:numeric(18,12)" json:"lintang"`
	Bujur   *float64 `gorm:"column:bujur;type:numeric(18,12)" json:"bujur"`

	NomorTelepon *string `gorm:"column:nomor_telepon;type:varchar(20)" json:"nomor_telepon"`
	NomorFax     *string `gorm:"column:nomor_fax;type:varchar(20)" json:"nomor_fax"`
	Email        *string `gorm:"column:email;type:varchar(60)" json:"email"`
	Website      *string `gorm:"column:website;type:varchar(100)" json:"website"`

	CreateDate  time.Time  `gorm:"column:create_date;type:timestamp;not null;default:CURRENT_TIMESTAMP" json:"create_date"`
	LastUpdate  time.Time  `gorm:"column:last_update;type:timestamp;not null;default:CURRENT_TIMESTAMP" json:"last_update"`
	ExpiredDate *time.Time `gorm:"column:expired_date;type:timestamp" json:"expired_date"`
	LastSync    time.Time  `gorm:"column:last_sync;type:timestamp;not null;default:'1901-01-01 00:00:00'" json:"last_sync"`
}

func (LembagaAkreditasi) TableName() string {
	return "ref.lembaga_akreditasi"
}
