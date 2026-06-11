package models

import "time"

type JabatanTugasPtk struct {
	JabatanPtkID      int32      `gorm:"primaryKey;column:jabatan_ptk_id;type:numeric(5,0);not null" json:"jabatan_ptk_id"`
	Nama              string     `gorm:"column:nama;type:varchar(50);not null" json:"nama"`
	JabatanUtama      int16      `gorm:"column:jabatan_utama;type:numeric(1,0);not null" json:"jabatan_utama"`
	TugasTambahanGuru int16      `gorm:"column:tugas_tambahan_guru;type:numeric(1,0);not null" json:"tugas_tambahan_guru"`
	JumlahJamDiakui   *int16     `gorm:"column:jumlah_jam_diakui;type:numeric(2,0)" json:"jumlah_jam_diakui"`
	HarusReferUnitOrg int16      `gorm:"column:harus_refer_unit_org;type:numeric(1,0);not null;default:0" json:"harus_refer_unit_org"`
	CreateDate        time.Time  `gorm:"column:create_date;type:timestamp;not null;default:CURRENT_TIMESTAMP" json:"create_date"`
	LastUpdate        time.Time  `gorm:"column:last_update;type:timestamp;not null;default:CURRENT_TIMESTAMP" json:"last_update"`
	ExpiredDate       *time.Time `gorm:"column:expired_date;type:timestamp" json:"expired_date"`
	LastSync          time.Time  `gorm:"column:last_sync;type:timestamp;not null;default:'1901-01-01 00:00:00'" json:"last_sync"`
}

func (JabatanTugasPtk) TableName() string {
	return "ref.jabatan_tugas_ptk"
}
