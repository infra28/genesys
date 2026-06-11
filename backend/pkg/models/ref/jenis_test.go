package models

import "time"

type JenisTest struct {
	JenisTestID int16   `gorm:"primaryKey;column:jenis_test_id;type:numeric(3,0);not null" json:"jenis_test_id"`
	JenisTest   string  `gorm:"column:jenis_test;type:varchar(30);not null" json:"jenis_test"`
	Keterangan  *string `gorm:"column:keterangan;type:varchar(80)" json:"keterangan"`
	NilaiMaks   float32 `gorm:"column:nilai_maks;type:numeric(6,2);not null" json:"nilai_maks"`
	KetSkor1    *string `gorm:"column:ket_skor1;type:varchar(80)" json:"ket_skor1"`
	KetSkor2    *string `gorm:"column:ket_skor2;type:varchar(80)" json:"ket_skor2"`
	KetSkor3    *string `gorm:"column:ket_skor3;type:varchar(80)" json:"ket_skor3"`
	KetSkor4    *string `gorm:"column:ket_skor4;type:varchar(80)" json:"ket_skor4"`
	KetSkor5    *string `gorm:"column:ket_skor5;type:varchar(80)" json:"ket_skor5"`

	CreateDate  time.Time  `gorm:"column:create_date;type:timestamp;not null;default:CURRENT_TIMESTAMP" json:"create_date"`
	LastUpdate  time.Time  `gorm:"column:last_update;type:timestamp;not null;default:CURRENT_TIMESTAMP" json:"last_update"`
	ExpiredDate *time.Time `gorm:"column:expired_date;type:timestamp" json:"expired_date"`
	LastSync    time.Time  `gorm:"column:last_sync;type:timestamp;not null;default:'1901-01-01 00:00:00'" json:"last_sync"`
}

func (JenisTest) TableName() string {
	return "ref.jenis_test"
}
