package models

import "time"

type JenisSertifikasi struct {
	IDJenisSertifikasi int16  `gorm:"primaryKey;column:id_jenis_sertifikasi;type:numeric(3,0);not null" json:"id_jenis_sertifikasi"`
	JenisSertifikasi   string `gorm:"column:jenis_sertifikasi;type:varchar(30);not null" json:"jenis_sertifikasi"`
	ProfGuru           int16  `gorm:"column:prof_guru;type:numeric(1,0);not null" json:"prof_guru"`
	KepalaSekolah      int16  `gorm:"column:kepala_sekolah;type:numeric(1,0);not null" json:"kepala_sekolah"`
	Laboran            int16  `gorm:"column:laboran;type:numeric(1,0);not null" json:"laboran"`
	APd                int16  `gorm:"column:a_pd;type:numeric(1,0);not null;default:0" json:"a_pd"`

	KebutuhanKhususID int32 `gorm:"column:kebutuhan_khusus_id;type:int4;not null" json:"kebutuhan_khusus_id"`

	CreateDate  time.Time  `gorm:"column:create_date;type:timestamp;not null;default:CURRENT_TIMESTAMP" json:"create_date"`
	LastUpdate  time.Time  `gorm:"column:last_update;type:timestamp;not null;default:CURRENT_TIMESTAMP" json:"last_update"`
	ExpiredDate *time.Time `gorm:"column:expired_date;type:timestamp" json:"expired_date"`
	LastSync    time.Time  `gorm:"column:last_sync;type:timestamp;not null;default:'1901-01-01 00:00:00'" json:"last_sync"`

	KebutuhanKhusus *KebutuhanKhusus `gorm:"foreignKey:KebutuhanKhususID;references:KebutuhanKhususID" json:"kebutuhan_khusus,omitempty"`
}

func (JenisSertifikasi) TableName() string {
	return "ref.jenis_sertifikasi"
}
