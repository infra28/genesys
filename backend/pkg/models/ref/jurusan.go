package models

import "time"

type Jurusan struct {
	JurusanID           string             `gorm:"primaryKey;column:jurusan_id;type:varchar(25);not null" json:"jurusan_id"`
	NamaJurusan         string             `gorm:"column:nama_jurusan;type:varchar(100);not null" json:"nama_jurusan"`
	UntukSma            int16              `gorm:"column:untuk_sma;type:numeric(1,0);not null" json:"untuk_sma"`
	UntukSmk            int16              `gorm:"column:untuk_smk;type:numeric(1,0);not null" json:"untuk_smk"`
	UntukPt             int16              `gorm:"column:untuk_pt;type:numeric(1,0);not null" json:"untuk_pt"`
	UntukSlb            int16              `gorm:"column:untuk_slb;type:numeric(1,0);not null;default:0" json:"untuk_slb"`
	UntukSmklb          int16              `gorm:"column:untuk_smklb;type:numeric(1,0);not null;default:0" json:"untuk_smklb"`
	JenjangPendidikanID *int16             `gorm:"column:jenjang_pendidikan_id;type:numeric(2,0)" json:"jenjang_pendidikan_id"`
	JurusanInduk        *string            `gorm:"column:jurusan_induk;type:varchar(25)" json:"jurusan_induk"`
	LevelBidangID       string             `gorm:"column:level_bidang_id;type:varchar(5);not null" json:"level_bidang_id"`
	CreateDate          time.Time          `gorm:"column:create_date;type:timestamp;not null;default:CURRENT_TIMESTAMP" json:"create_date"`
	LastUpdate          time.Time          `gorm:"column:last_update;type:timestamp;not null;default:CURRENT_TIMESTAMP" json:"last_update"`
	ExpiredDate         *time.Time         `gorm:"column:expired_date;type:timestamp" json:"expired_date"`
	LastSync            time.Time          `gorm:"column:last_sync;type:timestamp;not null;default:'1901-01-01 00:00:00'" json:"last_sync"`
	ParentJurusan       *Jurusan           `gorm:"foreignKey:JurusanInduk;references:JurusanID" json:"parent_jurusan,omitempty"`
	JenjangPendidikan   *JenjangPendidikan `gorm:"foreignKey:JenjangPendidikanID;references:JenjangPendidikanID" json:"jenjang_pendidikan,omitempty"`
	KelompokBidang      *KelompokBidang    `gorm:"foreignKey:LevelBidangID;references:LevelBidangID" json:"kelompok_bidang,omitempty"`
}

func (Jurusan) TableName() string {
	return "ref.jurusan"
}
