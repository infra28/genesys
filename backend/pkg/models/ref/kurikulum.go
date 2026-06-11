package models

import "time"

type Kurikulum struct {
	KurikulumID         int16              `gorm:"primaryKey;column:kurikulum_id;type:int2;not null" json:"kurikulum_id"`
	NamaKurikulum       string             `gorm:"column:nama_kurikulum;type:varchar(120);not null" json:"nama_kurikulum"`
	MulaiBerlaku        time.Time          `gorm:"column:mulai_berlaku;type:date;not null" json:"mulai_berlaku"`
	SistemSks           int16              `gorm:"column:sistem_sks;type:numeric(1,0);not null;default:0" json:"sistem_sks"`
	TotalSks            int16              `gorm:"column:total_sks;type:numeric(3,0);not null;default:0" json:"total_sks"`
	JenjangPendidikanID int16              `gorm:"column:jenjang_pendidikan_id;type:numeric(2,0);not null" json:"jenjang_pendidikan_id"`
	JurusanID           *string            `gorm:"column:jurusan_id;type:varchar(25)" json:"jurusan_id"`
	CreateDate          time.Time          `gorm:"column:create_date;type:timestamp;not null;default:CURRENT_TIMESTAMP" json:"create_date"`
	LastUpdate          time.Time          `gorm:"column:last_update;type:timestamp;not null;default:CURRENT_TIMESTAMP" json:"last_update"`
	ExpiredDate         *time.Time         `gorm:"column:expired_date;type:timestamp" json:"expired_date"`
	LastSync            time.Time          `gorm:"column:last_sync;type:timestamp;not null;default:'1901-01-01 00:00:00'" json:"last_sync"`
	JenjangPendidikan   *JenjangPendidikan `gorm:"foreignKey:JenjangPendidikanID;references:JenjangPendidikanID" json:"jenjang_pendidikan,omitempty"`
	Jurusan             *Jurusan           `gorm:"foreignKey:JurusanID;references:JurusanID" json:"jurusan,omitempty"`
}

// TableName menentukan nama skema dan tabel spesifik di PostgreSQL
func (Kurikulum) TableName() string {
	return "ref.kurikulum"
}
