package models

import "time"

type JenisKesejahteraan struct {
	JenisKesejahteraanID int32      `gorm:"primaryKey;column:jenis_kesejahteraan_id;type:int4;not null" json:"jenis_kesejahteraan_id"`
	Nama                 string     `gorm:"column:nama;type:varchar(50);not null" json:"nama"`
	Penyelenggara        string     `gorm:"column:penyelenggara;type:varchar(100);not null" json:"penyelenggara"`
	UPtk                 int16      `gorm:"column:u_ptk;type:numeric(1,0);not null;default:1" json:"u_ptk"`
	UPd                  int16      `gorm:"column:u_pd;type:numeric(1,0);not null;default:0" json:"u_pd"`
	CreateDate           time.Time  `gorm:"column:create_date;type:timestamp;not null;default:CURRENT_TIMESTAMP" json:"create_date"`
	LastUpdate           time.Time  `gorm:"column:last_update;type:timestamp;not null;default:CURRENT_TIMESTAMP" json:"last_update"`
	ExpiredDate          *time.Time `gorm:"column:expired_date;type:timestamp" json:"expired_date"`
	LastSync             time.Time  `gorm:"column:last_sync;type:timestamp;not null;default:'1901-01-01 00:00:00'" json:"last_sync"`
}

func (JenisKesejahteraan) TableName() string {
	return "ref.jenis_kesejahteraan"
}
