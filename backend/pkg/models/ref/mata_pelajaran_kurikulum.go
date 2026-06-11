package models

import (
	"time"

	"github.com/google/uuid"
)

// MataPelajaranKurikulum mewakili struktur tabel ref.mata_pelajaran_kurikulum
type MataPelajaranKurikulum struct {
	// Composite Primary Key
	KurikulumID         int16 `gorm:"primaryKey;column:kurikulum_id;type:int2;not null" json:"kurikulum_id"`
	MataPelajaranID     int32 `gorm:"primaryKey;column:mata_pelajaran_id;type:int4;not null" json:"mata_pelajaran_id"`
	TingkatPendidikanID int16 `gorm:"primaryKey;column:tingkat_pendidikan_id;type:numeric(2,0);not null" json:"tingkat_pendidikan_id"`

	// Informasi Beban & Kurikulum
	JumlahJam         int16      `gorm:"column:jumlah_jam;type:numeric(2,0);not null" json:"jumlah_jam"`
	JumlahJamMaksimum int16      `gorm:"column:jumlah_jam_maksimum;type:numeric(2,0);not null" json:"jumlah_jam_maksimum"`
	StatusDiKurikulum int16      `gorm:"column:status_di_kurikulum;type:numeric(2,0);not null" json:"status_di_kurikulum"`
	Wajib             int16      `gorm:"column:wajib;type:numeric(1,0);not null" json:"wajib"`
	SKS               int16      `gorm:"column:sks;type:numeric(2,0);not null;default:0" json:"sks"`
	APeminatan        int16      `gorm:"column:a_peminatan;type:numeric(1,0);not null" json:"a_peminatan"`
	AreaKompetensi    string     `gorm:"column:area_kompetensi;type:char(1);not null;default:'*'" json:"area_kompetensi"`
	GmpID             *uuid.UUID `gorm:"column:gmp_id;type:uuid" json:"gmp_id"`

	// Foreign Keys untuk relasi/preload
	StatusKurikulum   *StatusDiKurikulum `gorm:"foreignKey:StatusDiKurikulum" json:"status_kurikulum,omitempty"`
	GroupMatpel       *GroupMatpel       `gorm:"foreignKey:GmpID" json:"group_matpel,omitempty"`
	Kurikulum         *Kurikulum         `gorm:"foreignKey:KurikulumID" json:"kurikulum,omitempty"`
	MataPelajaran     *MataPelajaran     `gorm:"foreignKey:MataPelajaranID" json:"mata_pelajaran,omitempty"`
	TingkatPendidikan *TingkatPendidikan `gorm:"foreignKey:TingkatPendidikanID" json:"tingkat_pendidikan,omitempty"`

	// Timestamp System
	CreateDate  time.Time  `gorm:"column:create_date;type:timestamp;not null;default:CURRENT_TIMESTAMP" json:"create_date"`
	LastUpdate  time.Time  `gorm:"column:last_update;type:timestamp;not null;default:CURRENT_TIMESTAMP" json:"last_update"`
	ExpiredDate *time.Time `gorm:"column:expired_date;type:timestamp" json:"expired_date"`
	LastSync    time.Time  `gorm:"column:last_sync;type:timestamp;not null;default:'1901-01-01 00:00:00'" json:"last_sync"`
}

func (MataPelajaranKurikulum) TableName() string {
	return "ref.mata_pelajaran_kurikulum"
}
