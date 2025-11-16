package handlers

import (
	"strings"
	"time"

	"github.com/gofiber/fiber/v2"
	"evormos-task/database"
	"evormos-task/models"
)

func GetMyToko(c *fiber.Ctx) error {
	userID := c.Locals("user_id")
	if userID == nil {
		return c.Status(401).JSON(fiber.Map{"error": "Unauthorized"})
	}

	var toko models.Toko
	if err := database.DB.Where("id_user = ?", userID).Preload("User").First(&toko).Error; err != nil {
		return c.Status(404).JSON(fiber.Map{"error": "Toko tidak ditemukan"})
	}

	return c.JSON(toko)
}

func UpdateMyToko(c *fiber.Ctx) error {
	idToko := c.Params("id_toko")

	var toko models.Toko
	if err := database.DB.First(&toko, idToko).Error; err != nil {
		return c.Status(404).JSON(fiber.Map{"error": "Toko tidak ditemukan"})
	}

	userID := c.Locals("user_id")
	if userID == nil || toko.ID_User != userID.(uint) {
		return c.Status(403).JSON(fiber.Map{"error": "Akses ditolak"})
	}

	var data map[string]interface{}
	if err := c.BodyParser(&data); err != nil {
		return c.Status(400).JSON(fiber.Map{"error": "Data tidak valid"})
	}

	data["updated_at"] = time.Now()

	if err := database.DB.Model(&toko).Updates(data).Error; err != nil {
		return c.Status(500).JSON(fiber.Map{"error": "Gagal update toko"})
	}

	return c.JSON(fiber.Map{
		"message": "Toko berhasil diperbarui",
		"toko":    toko,
	})
}

func GetTokoByID(c *fiber.Ctx) error {
	idToko := c.Params("id_toko")

	var toko models.Toko
	if err := database.DB.Preload("User").First(&toko, idToko).Error; err != nil {
		return c.Status(404).JSON(fiber.Map{"error": "Toko tidak ditemukan"})
	}

	return c.JSON(toko)
}

func GetAllToko(c *fiber.Ctx) error {
	query := database.DB.Model(&models.Toko{}).Preload("User")

	nama := c.Query("nama")
	if nama == "" {
		nama = c.Query("nama_toko")
	}

	if nama != "" {
		clean := strings.TrimSpace(strings.ToLower(nama))
		query = query.Where("LOWER(nama_toko) LIKE ?", "%"+clean+"%")
	}

	var tokos []models.Toko
	if err := query.Find(&tokos).Error; err != nil {
		return c.Status(500).JSON(fiber.Map{"error": "Gagal mengambil data toko"})
	}

	return c.JSON(tokos)
}

