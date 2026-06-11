package models

import "time"

type JenisLembaga struct {
	JenisLembagaID  int32      `gorm:"primaryKey;column:jenis_lembaga_id;type:numeric(5,0);not null" json:"jenis_lembaga_id"`
	Nama            string     `gorm:"column:nama;type:varchar(80);not null" json:"nama"`
	TempatPengawas  int16      `gorm:"column:tempat_pengawas;type:numeric(1,0);not null" json:"tempat_pengawas"`
	SimpulPendataan int16      `gorm:"column:simpul_pendataan;type:numeric(1,0);not null" json:"simpul_pendataan"`
	CreateDate      time.Time  `gorm:"column:create_date;type:timestamp;not null;default:CURRENT_TIMESTAMP" json:"create_date"`
	LastUpdate      time.Time  `gorm:"column:last_update;type:timestamp;not null;default:CURRENT_TIMESTAMP" json:"last_update"`
	ExpiredDate     *time.Time `gorm:"column:expired_date;type:timestamp" json:"expired_date"`
	LastSync        time.Time  `gorm:"column:last_sync;type:timestamp;not null;default:'1901-01-01 00:00:00'" json:"last_sync"`
}

func (JenisLembaga) TableName() string {
	return "ref.jenis_lembaga"
}
