package models

import "time"

type BatasWaktuRapor struct {
	SemesterID      string     `gorm:"primaryKey;column:semester_id;type:char(5);not null" json:"semester_id"`
	TglRaporMulai   time.Time  `gorm:"column:tgl_rapor_mulai;type:date;not null" json:"tgl_rapor_mulai"`
	TglRaporSelesai time.Time  `gorm:"column:tgl_rapor_selesai;type:date;not null" json:"tgl_rapor_selesai"`
	TglUsmMulai     *time.Time `gorm:"column:tgl_usm_mulai;type:date" json:"tgl_usm_mulai"`
	TglUsmSelesai   *time.Time `gorm:"column:tgl_usm_selesai;type:date" json:"tgl_usm_selesai"`
	CreateDate      time.Time  `gorm:"column:create_date;type:timestamp;not null" json:"create_date"`
	LastUpdate      time.Time  `gorm:"column:last_update;type:timestamp;not null" json:"last_update"`
	ExpiredDate     *time.Time `gorm:"column:expired_date;type:timestamp" json:"expired_date"`
	LastSync        time.Time  `gorm:"column:last_sync;type:timestamp;not null" json:"last_sync"`

	Semester *Semester `gorm:"foreignKey:SemesterID;references:SemesterID" json:"semester,omitempty"`
}

func (BatasWaktuRapor) TableName() string {
	return "ref.batas_waktu_rapor"
}
