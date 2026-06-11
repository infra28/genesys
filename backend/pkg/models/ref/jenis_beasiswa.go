package models

import "time"

type JenisBeasiswa struct {
	JenisBeasiswaID int32  `gorm:"primaryKey;column:jenis_beasiswa_id;type:int4;not null" json:"jenis_beasiswa_id"`
	SumberDanaID    int16  `gorm:"column:sumber_dana_id;type:numeric(3,0);not null" json:"sumber_dana_id"`
	Nama            string `gorm:"column:nama;type:varchar(50);not null" json:"nama"`
	UntukPd         int16  `gorm:"column:untuk_pd;type:numeric(1,0);not null" json:"untuk_pd"`
	UntukPtk        int16  `gorm:"column:untuk_ptk;type:numeric(1,0);not null" json:"untuk_ptk"`

	CreateDate  time.Time  `gorm:"column:create_date;type:timestamp;not null;default:CURRENT_TIMESTAMP" json:"create_date"`
	LastUpdate  time.Time  `gorm:"column:last_update;type:timestamp;not null;default:CURRENT_TIMESTAMP" json:"last_update"`
	ExpiredDate *time.Time `gorm:"column:expired_date;type:timestamp" json:"expired_date"`
	LastSync    time.Time  `gorm:"column:last_sync;type:timestamp;not null;default:'1901-01-01 00:00:00'" json:"last_sync"`

	SumberDana *SumberDana `gorm:"foreignKey:SumberDanaID;references:SumberDanaID" json:"sumber_dana,omitempty"`
}

func (JenisBeasiswa) TableName() string {
	return "ref.jenis_beasiswa"
}
