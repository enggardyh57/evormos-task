package handlers

import (
	"time"

	"evormos-task/database"
	"evormos-task/models"
	"github.com/gofiber/fiber/v2"
)


func GetAlamat(c *fiber.Ctx) error {
	userID, ok := c.Locals("user_id").(uint)
	if !ok {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"error": "User tidak valid"})
	}

	var alamat []models.Alamat
	if err := database.DB.Where("id_user = ?", userID).Find(&alamat).Error; err != nil {
		return c.Status(500).JSON(fiber.Map{"error": "Gagal ambil data alamat"})
	}

	return c.Status(200).JSON(alamat)
}


func CreateAlamat(c *fiber.Ctx) error {
	userID, ok := c.Locals("user_id").(uint)
	if !ok {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"error": "User tidak valid"})
	}

	var data models.Alamat
	if err := c.BodyParser(&data); err != nil {
		return c.Status(400).JSON(fiber.Map{"error": "Request tidak valid"})
	}

	data.ID_User = userID
	data.CreatedAt = time.Now()
	data.UpdatedAt = time.Now()

	if err := database.DB.Create(&data).Error; err != nil {
		return c.Status(500).JSON(fiber.Map{"error": "Gagal menyimpan alamat"})
	}

	return c.Status(201).JSON(fiber.Map{
		"message": "Alamat berhasil ditambahkan",
		"alamat":  data,
	})
}

func GetAlamatByID(c *fiber.Ctx) error {
	userID := c.Locals("user_id").(uint)
	id := c.Params("id")

	var alamat models.Alamat
	if err := database.DB.First(&alamat, id).Error; err != nil {
		return c.Status(404).JSON(fiber.Map{"error": "Alamat tidak ditemukan"})
	}

	if alamat.ID_User != userID {
		return c.Status(403).JSON(fiber.Map{"error": "Tidak boleh akses alamat milik user lain"})
	}

	return c.JSON(alamat)
}

func UpdateAlamat(c *fiber.Ctx) error {
	userID, ok := c.Locals("user_id").(uint)
	if !ok {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"error": "User tidak valid"})
	}

	id := c.Params("id")

	var alamat models.Alamat
	if err := database.DB.First(&alamat, id).Error; err != nil {
		return c.Status(404).JSON(fiber.Map{"error": "Alamat tidak ditemukan"})
	}

	if alamat.ID_User != userID {
		return c.Status(403).JSON(fiber.Map{"error": "Kamu tidak bisa ubah alamat milik user lain"})
	}

	var data map[string]interface{}
	if err := c.BodyParser(&data); err != nil {
		return c.Status(400).JSON(fiber.Map{"error": "Request tidak valid"})
	}

	data["updated_at"] = time.Now()
	if err := database.DB.Model(&alamat).Updates(data).Error; err != nil {
		return c.Status(500).JSON(fiber.Map{"error": "Gagal update alamat"})
	}

	return c.Status(200).JSON(fiber.Map{
		"message": "Alamat berhasil diperbarui",
		"alamat":  alamat,
	})
}

// DELETE hapus alamat
func DeleteAlamat(c *fiber.Ctx) error {
	userID, ok := c.Locals("user_id").(uint)
	if !ok {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"error": "User tidak valid"})
	}

	id := c.Params("id")

	var alamat models.Alamat
	if err := database.DB.First(&alamat, id).Error; err != nil {
		return c.Status(404).JSON(fiber.Map{"error": "Alamat tidak ditemukan"})
	}

	if alamat.ID_User != userID {
		return c.Status(403).JSON(fiber.Map{"error": "Kamu tidak bisa hapus alamat milik user lain"})
	}

	if err := database.DB.Delete(&alamat).Error; err != nil {
		return c.Status(500).JSON(fiber.Map{"error": "Gagal hapus alamat"})
	}

	return c.Status(200).JSON(fiber.Map{"message": "Alamat berhasil dihapus"})
}