package models

import (
	"time"

	"github.com/google/uuid"
)

// Variabel mewakili struktur tabel ref.variabel
type Variabel struct {
	// Primary Key
	VariabelID uuid.UUID `gorm:"primaryKey;column:variabel_id;type:uuid;not null" json:"variabel_id"`

	// Informasi Metadata
	Nama          string  `gorm:"column:nama;type:varchar(500);not null" json:"nama"`
	Header        *string `gorm:"column:header;type:varchar(500)" json:"header"`
	Urut          *int16  `gorm:"column:urut;type:int2" json:"urut"`
	StringPattern *string `gorm:"column:string_pattern;type:varchar(500)" json:"string_pattern"`
	Keterangan    *string `gorm:"column:keterangan;type:varchar(500)" json:"keterangan"`
	JenisVariabel string  `gorm:"column:jenis_variabel;type:char(1);not null" json:"jenis_variabel"`

	// Konfigurasi Jenjang (Flag)
	UPaud    int16 `gorm:"column:u_paud;type:numeric(1,0);not null" json:"u_paud"`
	USd      int16 `gorm:"column:u_sd;type:numeric(1,0);not null" json:"u_sd"`
	USmp     int16 `gorm:"column:u_smp;type:numeric(1,0);not null" json:"u_smp"`
	USma     int16 `gorm:"column:u_sma;type:numeric(1,0);not null" json:"u_sma"`
	USmk     int16 `gorm:"column:u_smk;type:numeric(1,0);not null" json:"u_smk"`
	IsTampil int16 `gorm:"column:is_tampil;type:numeric(1,0);not null" json:"is_tampil"`

	// Timestamp System
	CreateDate  time.Time  `gorm:"column:create_date;type:timestamp;not null;default:CURRENT_TIMESTAMP" json:"create_date"`
	LastUpdate  time.Time  `gorm:"column:last_update;type:timestamp;not null;default:CURRENT_TIMESTAMP" json:"last_update"`
	ExpiredDate *time.Time `gorm:"column:expired_date;type:timestamp" json:"expired_date"`
	LastSync    time.Time  `gorm:"column:last_sync;type:timestamp;not null;default:'1901-01-01 00:00:00'" json:"last_sync"`
}

func (Variabel) TableName() string {
	return "ref.variabel"
}
