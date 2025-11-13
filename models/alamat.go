package models

import "time"

type Alamat struct {
	ID          uint      `gorm:"primaryKey" json:"id"`
	ID_User     uint      `gorm:"not null" json:"id_user"`
	JudulAlamat  string    `gorm:"size:255" json:"judul_alamat"`
	NamaPenerima string    `gorm:"size:255" json:"nama_penerima"`
	NoTelp  	string    `gorm:"size:255" json:"no_telp"`
	DetailAlamat  string    `gorm:"type:text" json:"detail_alamat"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`

	User *User `gorm:"foreignKey:ID_User;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;" json:"user,omitempty"`
}

func (Alamat) TableName() string {
	return "alamat"
}