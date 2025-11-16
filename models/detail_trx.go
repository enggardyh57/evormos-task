package models

import "time"

type DetailTrx struct {
	ID        uint      `json:"id" gorm:"primaryKey"`
	IDTrx     uint      `json:"id_trx"`
	IDLog     uint      `json:"id_log"`
	IDToko    uint      `json:"id_toko"`
	Kuantitas uint      `json:"kuantitas"`
	Harga     string    `json:"harga" gorm:"size:255"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`

	LogProduk *LogProduk `gorm:"foreignKey:IDLog;references:ID" json:"log_produk,omitempty"`
	Transaksi *Transaksi `gorm:"foreignKey:IDTrx;references:ID" json:"transaksi,omitempty"`
	Toko      *Toko      `gorm:"foreignKey:IDToko;references:ID" json:"toko,omitempty"`
}

func (DetailTrx) TableName() string {
	return "detail_trx"
}
