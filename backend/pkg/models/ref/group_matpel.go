package models

import "time"

type GroupMatpel struct {
	GmpID               string     `gorm:"primaryKey;column:gmp_id;type:uuid;not null;default:uuid_generate_v4()" json:"gmp_id"`
	NamaGroup           string     `gorm:"column:nama_group;type:varchar(80);not null" json:"nama_group"`
	JumlahJamMaksimum   int16      `gorm:"column:jumlah_jam_maksimum;type:numeric(2,0);not null" json:"jumlah_jam_maksimum"`
	JumlahSksMaksimum   int16      `gorm:"column:jumlah_sks_maksimum;type:numeric(2,0);not null;default:0" json:"jumlah_sks_maksimum"`
	KurikulumID         int16      `gorm:"column:kurikulum_id;type:int2;not null" json:"kurikulum_id"`
	TingkatPendidikanID int16      `gorm:"column:tingkat_pendidikan_id;type:numeric(2,0);not null" json:"tingkat_pendidikan_id"`
	CreateDate          time.Time  `gorm:"column:create_date;type:timestamp;not null;default:CURRENT_TIMESTAMP" json:"create_date"`
	LastUpdate          time.Time  `gorm:"column:last_update;type:timestamp;not null;default:CURRENT_TIMESTAMP" json:"last_update"`
	ExpiredDate         *time.Time `gorm:"column:expired_date;type:timestamp" json:"expired_date"`
	LastSync            time.Time  `gorm:"column:last_sync;type:timestamp;not null;default:'1901-01-01 00:00:00'" json:"last_sync"`

	Kurikulum         *Kurikulum         `gorm:"foreignKey:KurikulumID;references:KurikulumID" json:"kurikulum,omitempty"`
	TingkatPendidikan *TingkatPendidikan `gorm:"foreignKey:TingkatPendidikanID;references:TingkatPendidikanID" json:"tingkat_pendidikan,omitempty"`
}

func (GroupMatpel) TableName() string {
	return "ref.group_matpel"
}
