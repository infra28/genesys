package models

import (
	"time"

	"github.com/google/uuid"
)

// StandarSarana mewakili struktur tabel ref.standar_sarana
type StandarSarana struct {
	// Key: Primary Key, Type: uuid
	IDStdSarana uuid.UUID `gorm:"primaryKey;column:id_std_sarana;type:uuid;not null" json:"id_std_sarana"`

	// Informasi Relasional
	JenisPrasaranaID   int32   `gorm:"column:jenis_prasarana_id;type:int4;not null" json:"jenis_prasarana_id"`
	JenisSaranaID      int32   `gorm:"column:jenis_sarana_id;type:int4;not null" json:"jenis_sarana_id"`
	JurusanID          *string `gorm:"column:jurusan_id;type:varchar(25)" json:"jurusan_id"`
	BentukPendidikanID int16   `gorm:"column:bentuk_pendidikan_id;type:int2;not null" json:"bentuk_pendidikan_id"`

	// Spesifikasi
	AHarusAda int16 `gorm:"column:a_harus_ada;type:numeric(1,0);not null" json:"a_harus_ada"`

	// Foreign Keys (untuk mempermudah Preload)
	JenisPrasarana   *JenisPrasarana   `gorm:"foreignKey:JenisPrasaranaID" json:"jenis_prasarana,omitempty"`
	JenisSarana      *JenisSarana      `gorm:"foreignKey:JenisSaranaID" json:"jenis_sarana,omitempty"`
	Jurusan          *Jurusan          `gorm:"foreignKey:JurusanID" json:"jurusan,omitempty"`
	BentukPendidikan *BentukPendidikan `gorm:"foreignKey:BentukPendidikanID" json:"bentuk_pendidikan,omitempty"`

	// Timestamp System
	CreateDate  time.Time  `gorm:"column:create_date;type:timestamp;not null;default:CURRENT_TIMESTAMP" json:"create_date"`
	LastUpdate  time.Time  `gorm:"column:last_update;type:timestamp;not null;default:CURRENT_TIMESTAMP" json:"last_update"`
	ExpiredDate *time.Time `gorm:"column:expired_date;type:timestamp" json:"expired_date"`
	LastSync    time.Time  `gorm:"column:last_sync;type:timestamp;not null;default:'1901-01-01 00:00:00'" json:"last_sync"`
}

func (StandarSarana) TableName() string {
	return "ref.standar_sarana"
}
