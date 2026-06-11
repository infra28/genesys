package models

import "time"

type KebutuhanKhusus struct {
	KebutuhanKhususID int32      `gorm:"primaryKey;column:kebutuhan_khusus_id;type:int4;not null" json:"kebutuhan_khusus_id"`
	KebutuhanKhusus   string     `gorm:"column:kebutuhan_khusus;type:varchar(40);not null" json:"kebutuhan_khusus"`
	KkA               int16      `gorm:"type:numeric(1,0);not null" json:"-"`
	KkB               int16      `gorm:"type:numeric(1,0);not null" json:"-"`
	KkC               int16      `gorm:"type:numeric(1,0);not null" json:"-"`
	KkC1              int16      `gorm:"type:numeric(1,0);not null" json:"-"`
	KkD               int16      `gorm:"type:numeric(1,0);not null" json:"-"`
	KkD1              int16      `gorm:"type:numeric(1,0);not null" json:"-"`
	KkE               int16      `gorm:"type:numeric(1,0);not null" json:"-"`
	KkF               int16      `gorm:"type:numeric(1,0);not null" json:"-"`
	KkH               int16      `gorm:"type:numeric(1,0);not null" json:"-"`
	KkI               int16      `gorm:"type:numeric(1,0);not null" json:"-"`
	KkJ               int16      `gorm:"type:numeric(1,0);not null" json:"-"`
	KkK               int16      `gorm:"type:numeric(1,0);not null" json:"-"`
	KkN               int16      `gorm:"type:numeric(1,0);not null" json:"-"`
	KkO               int16      `gorm:"type:numeric(1,0);not null" json:"-"`
	KkP               int16      `gorm:"type:numeric(1,0);not null" json:"-"`
	KkQ               int16      `gorm:"type:numeric(1,0);not null" json:"-"`
	UntukLembaga      int16      `gorm:"column:untuk_lembaga;type:numeric(1,0);not null;default:1" json:"untuk_lembaga"`
	UntukPtk          int16      `gorm:"column:untuk_ptk;type:numeric(1,0);not null;default:1" json:"untuk_ptk"`
	UntukPd           int16      `gorm:"column:untuk_pd;type:numeric(1,0);not null;default:1" json:"untuk_pd"`
	CreateDate        time.Time  `gorm:"column:create_date;type:timestamp;not null;default:CURRENT_TIMESTAMP" json:"create_date"`
	LastUpdate        time.Time  `gorm:"column:last_update;type:timestamp;not null;default:CURRENT_TIMESTAMP" json:"last_update"`
	ExpiredDate       *time.Time `gorm:"column:expired_date;type:timestamp" json:"expired_date"`
	LastSync          time.Time  `gorm:"column:last_sync;type:timestamp;not null;default:'1901-01-01 00:00:00'" json:"last_sync"`
}

func (KebutuhanKhusus) TableName() string {
	return "ref.kebutuhan_khusus"
}
