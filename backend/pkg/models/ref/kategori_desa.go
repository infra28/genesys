package models

import "time"

type KategoriDesa struct {
	KategoriDesaID int16      `gorm:"primaryKey;column:kategori_desa_id;type:numeric(2,0);not null" json:"kategori_desa_id"`
	Nama           string     `gorm:"column:nama;type:varchar(30);not null" json:"nama"`
	CreateDate     time.Time  `gorm:"column:create_date;type:timestamp;not null" json:"create_date"`
	LastUpdate     time.Time  `gorm:"column:last_update;type:timestamp;not null" json:"last_update"`
	ExpiredDate    *time.Time `gorm:"column:expired_date;type:timestamp" json:"expired_date"`
	LastSync       time.Time  `gorm:"column:last_sync;type:timestamp;not null" json:"last_sync"`
}

func (KategoriDesa) TableName() string {
	return "ref.kategori_desa"
}
