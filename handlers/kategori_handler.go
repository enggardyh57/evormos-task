package handlers

import (
	"time"
	"evermos-task/database"
	"evermos-task/models"
	"github.com/gofiber/fiber/v2"
)


func GetAllKategori(c *fiber.Ctx) error {
	var kategori []models.Kategori
	if err := database.DB.Find(&kategori).Error; err != nil {
		return c.Status(500).JSON(fiber.Map{"error": "Gagal mengambil data kategori"})
	}
	return c.JSON(kategori)
}


func CreateKategori(c *fiber.Ctx) error {
	isAdmin := c.Locals("is_admin").(bool)
	if !isAdmin {
		return c.Status(403).JSON(fiber.Map{"error": "Hanya admin yang dapat menambah kategori"})
	}

	var data models.Kategori
	if err := c.BodyParser(&data); err != nil {
		return c.Status(400).JSON(fiber.Map{"error": "Request tidak valid"})
	}

	data.CreatedAt = time.Now()
	data.UpdatedAt = time.Now()

	if err := database.DB.Create(&data).Error; err != nil {
		return c.Status(500).JSON(fiber.Map{"error": "Gagal menyimpan kategori"})
	}

	return c.JSON(fiber.Map{
		"message": "Kategori berhasil ditambahkan",
		"data":    data,
	})
}

func GetKategoriByID(c *fiber.Ctx) error {
	id := c.Params("id")

	var kategori models.Kategori
	if err := database.DB.First(&kategori, id).Error; err != nil {
		return c.Status(404).JSON(fiber.Map{"error": "Kategori tidak ditemukan"})
	}

	return c.JSON(kategori)
}


func UpdateKategori(c *fiber.Ctx) error {
	isAdmin := c.Locals("is_admin").(bool)
	if !isAdmin {
		return c.Status(403).JSON(fiber.Map{"error": "Hanya admin yang dapat mengubah kategori"})
	}

	id := c.Params("id")
	var kategori models.Kategori

	if err := database.DB.First(&kategori, id).Error; err != nil {
		return c.Status(404).JSON(fiber.Map{"error": "Kategori tidak ditemukan"})
	}

	var data map[string]interface{}
	if err := c.BodyParser(&data); err != nil {
		return c.Status(400).JSON(fiber.Map{"error": "Request tidak valid"})
	}

	data["updated_at"] = time.Now()
	if err := database.DB.Model(&kategori).Updates(data).Error; err != nil {
		return c.Status(500).JSON(fiber.Map{"error": "Gagal mengupdate kategori"})
	}

	return c.JSON(fiber.Map{
		"message": "Kategori berhasil diperbarui",
		"data":    kategori,
	})
}


func DeleteKategori(c *fiber.Ctx) error {
	isAdmin := c.Locals("is_admin").(bool)
	if !isAdmin {
		return c.Status(403).JSON(fiber.Map{"error": "Hanya admin yang dapat menghapus kategori"})
	}

	id := c.Params("id")
	var kategori models.Kategori

	if err := database.DB.First(&kategori, id).Error; err != nil {
		return c.Status(404).JSON(fiber.Map{"error": "Kategori tidak ditemukan"})
	}

	if err := database.DB.Delete(&kategori).Error; err != nil {
		return c.Status(500).JSON(fiber.Map{"error": "Gagal menghapus kategori"})
	}

	return c.JSON(fiber.Map{"message": "Kategori berhasil dihapus"})
}