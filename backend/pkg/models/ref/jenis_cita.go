package models

import "time"

type JenisCita struct {
	IDCita int32  `gorm:"primaryKey;column:id_cita;type:numeric(5,0);not null" json:"id_cita"`
	NmCita string `gorm:"column:nm_cita;type:varchar(150);not null" json:"nm_cita"`

	CreateDate  time.Time  `gorm:"column:create_date;type:timestamp;not null;default:CURRENT_TIMESTAMP" json:"create_date"`
	LastUpdate  time.Time  `gorm:"column:last_update;type:timestamp;not null;default:CURRENT_TIMESTAMP" json:"last_update"`
	ExpiredDate *time.Time `gorm:"column:expired_date;type:timestamp" json:"expired_date"`
	LastSync    time.Time  `gorm:"column:last_sync;type:timestamp;not null;default:'1901-01-01 00:00:00'" json:"last_sync"`
}

func (JenisCita) TableName() string {
	return "ref.jenis_cita"
}
