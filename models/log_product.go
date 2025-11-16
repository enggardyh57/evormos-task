package models

import "time"

type LogProduk struct {
	ID             uint      `json:"id" gorm:"primaryKey"`
	IDProduk       uint      `json:"id_produk"`
	NamaProduk     string    `json:"nama_produk" gorm:"size:255"`
	Slug           string    `json:"slug" gorm:"size:255"`
	HargaReseller  string    `json:"harga_reseller" gorm:"size:255"`
	HargaKonsumen  string    `json:"harga_konsumen" gorm:"size:255"`
	Deskripsi      string    `json:"deskripsi" gorm:"type:text"`
	IDToko         uint      `json:"id_toko"`
	IDKategori     uint      `json:"id_kategori"`
	CreatedAt      time.Time `json:"created_at"`
	UpdatedAt      time.Time `json:"updated_at"`

	Produk   *Produk   `gorm:"foreignKey:IDProduk;references:ID" json:"produk,omitempty"`
	Toko     *Toko     `gorm:"foreignKey:IDToko;references:ID" json:"toko,omitempty"`
	Kategori *Kategori `gorm:"foreignKey:IDKategori;references:ID" json:"kategori,omitempty"`
}

func (LogProduk) TableName() string {
	return "log_produk"
}
