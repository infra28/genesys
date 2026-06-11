package models

import "time"

type Semester struct {
	SemesterID     string       `gorm:"primaryKey;column:semester_id;type:char(5);not null" json:"semester_id"`
	TahunAjaranID  int16        `gorm:"column:tahun_ajaran_id;type:numeric(4,0);not null" json:"tahun_ajaran_id"`
	Nama           string       `gorm:"column:nama;type:varchar(20);not null" json:"nama"`
	Semester       int16        `gorm:"column:semester;type:numeric(1,0);not null" json:"semester"`
	PeriodeAktif   int16        `gorm:"column:periode_aktif;type:numeric(1,0);not null" json:"periode_aktif"`
	TanggalMulai   time.Time    `gorm:"column:tanggal_mulai;type:date;not null" json:"tanggal_mulai"`
	TanggalSelesai time.Time    `gorm:"column:tanggal_selesai;type:date;not null" json:"tanggal_selesai"`
	CreateDate     time.Time    `gorm:"column:create_date;type:timestamp;not null;default:CURRENT_TIMESTAMP" json:"create_date"`
	LastUpdate     time.Time    `gorm:"column:last_update;type:timestamp;not null;default:CURRENT_TIMESTAMP" json:"last_update"`
	ExpiredDate    *time.Time   `gorm:"column:expired_date;type:timestamp" json:"expired_date"`
	LastSync       time.Time    `gorm:"column:last_sync;type:timestamp;not null;default:'1901-01-01 00:00:00'" json:"last_sync"`
	TahunAjaran    *TahunAjaran `gorm:"foreignKey:TahunAjaranID;references:TahunAjaranID" json:"tahun_ajaran,omitempty"`
}

func (Semester) TableName() string {
	return "ref.semester"
}
