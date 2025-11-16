package handlers

import (
	"fmt"
	"time"

	"github.com/gofiber/fiber/v2"
	"evermos-task/database"
	"evermos-task/models"
)

type ItemInput struct {
	IDProduk  uint `json:"id_produk"`
	Kuantitas uint `json:"kuantitas"`
}

type CreateTrxInput struct {
	ID_User     uint        `json:"id_user"`
	ID_Alamat   uint        `json:"id_alamat"`
	MethodBayar string      `json:"method_bayar"`
	Items       []ItemInput `json:"items"`
}

func CreateTransaksi(c *fiber.Ctx) error {
	var input CreateTrxInput
	if err := c.BodyParser(&input); err != nil {
		return c.Status(400).JSON(fiber.Map{"error": "Input tidak valid"})
	}

	
	trx := models.Transaksi{
		ID_User:     input.ID_User,
		ID_Alamat:   input.ID_Alamat,
		MethodBayar: input.MethodBayar,
		KodeInvoice: fmt.Sprintf("INV-%d", time.Now().Unix()),
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
	}

	if err := database.DB.Create(&trx).Error; err != nil {
		return c.Status(500).JSON(fiber.Map{"error": "Gagal membuat transaksi"})
	}

	var totalHarga float64 = 0

		for _, item := range input.Items {
		var p models.Produk
		if err := database.DB.First(&p, item.IDProduk).Error; err != nil {
			continue
		}

		
		log := models.LogProduk{
			IDProduk:      p.ID,
			NamaProduk:    p.NamaProduk,
			Slug:          p.Slug,
			HargaReseller: p.HargaReseller,
			HargaKonsumen: p.HargaKonsumen,
			Deskripsi:     p.Deskripsi,
			IDToko:        p.IDToko,
			IDKategori:    p.IDKategori,
			CreatedAt:     time.Now(),
			UpdatedAt:     time.Now(),
		}

		database.DB.Create(&log)

		
		hargaFloat := parseHarga(p.HargaKonsumen)
		totalHarga += float64(item.Kuantitas) * hargaFloat

		detail := models.DetailTrx{
			IDTrx:     trx.ID,
			IDLog:     log.ID,
			IDToko:    p.IDToko,
			Kuantitas: item.Kuantitas,
			Harga:     p.HargaKonsumen,
			CreatedAt: time.Now(),
			UpdatedAt: time.Now(),
		}

		database.DB.Create(&detail)
	}

	
	database.DB.Model(&trx).Update("harga_total", totalHarga)

	return c.JSON(fiber.Map{
		"message": "Transaksi berhasil dibuat",
		"data":    trx,
	})
}

func parseHarga(s string) float64 {
	var h float64
	fmt.Sscanf(s, "%f", &h)
	return h
}

func GetAllTransaksi(c *fiber.Ctx) error {
	var trx []models.Transaksi

	if err := database.DB.
		Preload("User").
		Preload("Alamat").
		Preload("DetailTrx").
		Preload("DetailTrx.LogProduk").
		Find(&trx).Error; err != nil {
		return c.Status(500).JSON(fiber.Map{"error": err.Error()})
	}

	return c.JSON(trx)
}

func GetTransaksiByID(c *fiber.Ctx) error {
	id := c.Params("id")
	var trx models.Transaksi

	if err := database.DB.
		Preload("User").
		Preload("Alamat").
		Preload("DetailTrx").
		Preload("DetailTrx.LogProduk").
		First(&trx, id).Error; err != nil {
		return c.Status(404).JSON(fiber.Map{"error": "Transaksi tidak ditemukan"})
	}

	return c.JSON(trx)
}

