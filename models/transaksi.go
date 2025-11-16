package models

import "time"

type Transaksi struct {
	ID               uint      `gorm:"primaryKey" json:"id"`
	ID_User          uint      `gorm:"not null" json:"id_user"`
	ID_Alamat        uint      `gorm:"not null" json:"id_alamat"` 
	HargaTotal       float64   `json:"harga_total"`
	KodeInvoice      string    `gorm:"size:255" json:"kode_invoice"`
	MethodBayar      string    `gorm:"size:255" json:"method_bayar"`
	CreatedAt        time.Time `json:"created_at"`
	UpdatedAt        time.Time `json:"updated_at"`

	User   *User   `gorm:"foreignKey:ID_User;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;" json:"user,omitempty"`
	Alamat *Alamat `gorm:"foreignKey:ID_Alamat;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;" json:"alamat,omitempty"`
	DetailTrx []DetailTrx `gorm:"foreignKey:IDTrx" json:"detail_trx,omitempty"`

}

func (Transaksi) TableName() string {
	return "trx"
}
