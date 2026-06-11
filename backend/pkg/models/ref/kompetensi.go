package models

import (
	"time"

	"github.com/google/uuid"
)

// Kompetensi mewakili struktur tabel ref.kompetensi
type Kompetensi struct {
	// Primary Key
	IDKomp uuid.UUID `gorm:"primaryKey;column:id_komp;type:uuid;not null" json:"id_komp"`

	// Informasi Kompetensi
	Desk     string `gorm:"column:desk;type:text;not null" json:"desk"`
	Nmr      string `gorm:"column:nmr;type:varchar(5)" json:"nmr"`
	Kelompok string `gorm:"column:kelompok;type:char(1);not null" json:"kelompok"`
	Versi    int32  `gorm:"column:versi;type:int4;not null" json:"versi"`

	// Relasi Hierarki & Konfigurasi
	IDIntiDasar         *uuid.UUID `gorm:"column:id_inti_dasar;type:uuid" json:"id_inti_dasar"`
	LevelKomp           *int16     `gorm:"column:level_komp;type:numeric(3,0)" json:"level_komp"`
	TingkatPendidikanID int16      `gorm:"column:tingkat_pendidikan_id;type:numeric(2,0);not null" json:"tingkat_pendidikan_id"`
	KurikulumID         int16      `gorm:"column:kurikulum_id;type:int2;not null" json:"kurikulum_id"`
	MataPelajaranID     int32      `gorm:"column:mata_pelajaran_id;type:int4;not null" json:"mata_pelajaran_id"`

	// Self-Reference & Foreign Keys
	ParentKompetensi  *Kompetensi        `gorm:"foreignKey:IDIntiDasar" json:"parent_kompetensi,omitempty"`
	TingkatPendidikan *TingkatPendidikan `gorm:"foreignKey:TingkatPendidikanID" json:"tingkat_pendidikan,omitempty"`
	Kurikulum         *Kurikulum         `gorm:"foreignKey:KurikulumID" json:"kurikulum,omitempty"`
	MataPelajaran     *MataPelajaran     `gorm:"foreignKey:MataPelajaranID" json:"mata_pelajaran,omitempty"`

	// Timestamp System
	CreateDate  time.Time  `gorm:"column:create_date;type:timestamp;not null;default:CURRENT_TIMESTAMP" json:"create_date"`
	LastUpdate  time.Time  `gorm:"column:last_update;type:timestamp;not null;default:CURRENT_TIMESTAMP" json:"last_update"`
	ExpiredDate *time.Time `gorm:"column:expired_date;type:timestamp" json:"expired_date"`
	LastSync    time.Time  `gorm:"column:last_sync;type:timestamp;not null;default:'1901-01-01 00:00:00'" json:"last_sync"`
}

func (Kompetensi) TableName() string {
	return "ref.kompetensi"
}
