package models

import (
	"time"
)

// MstWilayah mewakili struktur tabel ref.mst_wilayah
type MstWilayah struct {
	// Primary Key
	KodeWilayah string `gorm:"primaryKey;column:kode_wilayah;type:char(8);not null" json:"kode_wilayah"`

	// Informasi Geografis
	Nama           string  `gorm:"column:nama;type:varchar(60);not null" json:"nama"`
	IDLevelWilayah int16   `gorm:"column:id_level_wilayah;type:int2;not null" json:"id_level_wilayah"`
	MstKodeWilayah *string `gorm:"column:mst_kode_wilayah;type:char(8)" json:"mst_kode_wilayah"`
	NegaraID       string  `gorm:"column:negara_id;type:char(2);not null" json:"negara_id"`
	AsalWilayah    *string `gorm:"column:asal_wilayah;type:char(8)" json:"asal_wilayah"`

	// Kode Referensi Eksternal
	KodeBps   *string `gorm:"column:kode_bps;type:char(7)" json:"kode_bps"`
	KodeDagri *string `gorm:"column:kode_dagri;type:char(10)" json:"kode_dagri"`
	KodeKeu   *string `gorm:"column:kode_keu;type:char(10)" json:"kode_keu"`

	// Hierarki Pemetaan
	IDProv    *string `gorm:"column:id_prov;type:char(8)" json:"id_prov"`
	IDKabkota *string `gorm:"column:id_kabkota;type:char(8)" json:"id_kabkota"`
	IDKec     *string `gorm:"column:id_kec;type:char(8)" json:"id_kec"`

	// Flags
	ADesa      int16 `gorm:"column:a_desa;type:numeric(1,0);not null;default:0" json:"a_desa"`
	AKelurahan int16 `gorm:"column:a_kelurahan;type:numeric(1,0);not null;default:0" json:"a_kelurahan"`
	A35        int16 `gorm:"column:a_35;type:numeric(1,0);not null;default:0" json:"a_35"`
	AUrban     int16 `gorm:"column:a_urban;type:numeric(1,0);not null;default:0" json:"a_urban"`

	KategoriDesaID *int16 `gorm:"column:kategori_desa_id;type:numeric(2,0)" json:"kategori_desa_id"`

	// Relasi untuk Preload
	ParentWilayah *MstWilayah   `gorm:"foreignKey:MstKodeWilayah;references:KodeWilayah" json:"parent_wilayah,omitempty"`
	LevelWilayah  *LevelWilayah `gorm:"foreignKey:IDLevelWilayah" json:"level_wilayah,omitempty"`
	KategoriDesa  *KategoriDesa `gorm:"foreignKey:KategoriDesaID" json:"kategori_desa,omitempty"`
	Negara        *Negara       `gorm:"foreignKey:NegaraID" json:"negara,omitempty"`

	// Timestamp System
	CreateDate  time.Time  `gorm:"column:create_date;type:timestamp;not null;default:CURRENT_TIMESTAMP" json:"create_date"`
	LastUpdate  time.Time  `gorm:"column:last_update;type:timestamp;not null;default:CURRENT_TIMESTAMP" json:"last_update"`
	ExpiredDate *time.Time `gorm:"column:expired_date;type:timestamp" json:"expired_date"`
	LastSync    time.Time  `gorm:"column:last_sync;type:timestamp;not null;default:'1901-01-01 00:00:00'" json:"last_sync"`
}

func (MstWilayah) TableName() string {
	return "ref.mst_wilayah"
}
