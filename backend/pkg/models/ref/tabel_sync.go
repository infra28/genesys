package models

// TableSync mewakili struktur tabel ref.table_sync untuk kontrol sinkronisasi
type TableSync struct {
	// Primary Key
	Table_Name string `gorm:"primaryKey;column:table_name;type:varchar(30);not null" json:"table_name"`

	// Informasi Sinkronisasi
	TableAlias   *string `gorm:"column:table_alias;type:varchar(50)" json:"table_alias"`
	SyncType     string  `gorm:"column:sync_type;type:char(1);not null" json:"sync_type"`
	SyncSeq      int16   `gorm:"column:sync_seq;type:numeric(4,0);not null" json:"sync_seq"`
	KolomKecuali *string `gorm:"column:kolom_kecuali;type:varchar(200)" json:"kolom_kecuali"`

	// Status & Pengaturan
	TableStatus    *int16  `gorm:"column:table_status;type:int2" json:"table_status"`
	TableKet       *string `gorm:"column:table_ket;type:varchar(100)" json:"table_ket"`
	JmlThread      *int16  `gorm:"column:jml_thread;type:int2;default:5" json:"jml_thread"`
	BarisPerThread *int32  `gorm:"column:baris_per_thread;type:int4;default:500" json:"baris_per_thread"`
	OrderEkstra    *string `gorm:"column:order_ekstra;type:varchar(100)" json:"order_ekstra"`
	ATableAktif    int16   `gorm:"column:a_table_aktif;type:numeric(1,0);not null;default:1" json:"a_table_aktif"`
}

func (TableSync) TableName() string {
	return "ref.table_sync"
}
