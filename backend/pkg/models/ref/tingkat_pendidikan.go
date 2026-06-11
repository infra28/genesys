package models

import "time"

type TingkatPendidikan struct {
	TingkatPendidikanID int16              `gorm:"primaryKey;column:tingkat_pendidikan_id;type:numeric(2,0);not null" json:"tingkat_pendidikan_id"`
	Kode                string             `gorm:"column:kode;type:varchar(5);not null" json:"kode"`
	Nama                string             `gorm:"column:nama;type:varchar(20);not null" json:"nama"`
	JenjangPendidikanID int16              `gorm:"column:jenjang_pendidikan_id;type:numeric(2,0);not null" json:"jenjang_pendidikan_id"`
	CreateDate          time.Time          `gorm:"column:create_date;type:timestamp;not null;default:CURRENT_TIMESTAMP" json:"create_date"`
	LastUpdate          time.Time          `gorm:"column:last_update;type:timestamp;not null;default:CURRENT_TIMESTAMP" json:"last_update"`
	ExpiredDate         *time.Time         `gorm:"column:expired_date;type:timestamp" json:"expired_date"`
	LastSync            time.Time          `gorm:"column:last_sync;type:timestamp;not null;default:'1901-01-01 00:00:00'" json:"last_sync"`
	JenjangPendidikan   *JenjangPendidikan `gorm:"foreignKey:JenjangPendidikanID;references:JenjangPendidikanID" json:"jenjang_pendidikan,omitempty"`
}

func (TingkatPendidikan) TableName() string {
	return "ref.tingkat_pendidikan"
}
