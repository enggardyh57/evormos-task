package handlers

import (
	"time"

	"evormos-task/database"
	"evormos-task/models"
	"github.com/gofiber/fiber/v2"
)

func GetUserProfile(c *fiber.Ctx) error {
	userID, ok := c.Locals("user_id").(uint)
	if !ok {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"error": "User ID tidak ditemukan di token",
		})
	}

	var user models.User
	if err := database.DB.First(&user, userID).Error; err != nil {
		return c.Status(404).JSON(fiber.Map{
			"error": "User tidak ditemukan",
		})
	}

	return c.Status(200).JSON(user)
}

func GetUserByID(c *fiber.Ctx) error {
	id := c.Params("id")

	var user models.User
	if err := database.DB.First(&user, id).Error; err != nil {
		return c.Status(404).JSON(fiber.Map{"error": "User tidak ditemukan"})
	}

	return c.JSON(user)
}

func UpdateUser(c *fiber.Ctx) error {
	userID, ok := c.Locals("user_id").(uint)
	if !ok {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"error": "User ID tidak ditemukan di token",
		})
	}

	var user models.User
	if err := database.DB.First(&user, userID).Error; err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"error": "User tidak ditemukan",
		})
	}

	var data map[string]interface{}
	if err := c.BodyParser(&data); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Request tidak valid",
		})
	}

	
	if nama, ok := data["nama"].(string); ok {
		user.Nama = nama
	}
	if notelp, ok := data["notelp"].(string); ok {
		user.Notelp = notelp
	}
	if tanggal, ok := data["tanggal_lahir"].(string); ok {
		user.Tanggal_Lahir = tanggal
	}
	if jk, ok := data["jenis_kelamin"].(string); ok {
		user.Jenis_Kelamin = jk
	}
	if tentang, ok := data["tentang"].(string); ok {
		user.Tentang = tentang
	}
	if pekerjaan, ok := data["pekerjaan"].(string); ok {
		user.Pekerjaan = pekerjaan
	}
	if email, ok := data["email"].(string); ok {
		user.Email = email
	}
	if prov, ok := data["id_provinsi"].(string); ok {
		user.ID_Provinsi = prov
	}
	if kota, ok := data["id_kota"].(string); ok {
		user.ID_Kota = kota
	}

	user.UpdatedAt = time.Now()

	if err := database.DB.Save(&user).Error; err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Gagal memperbarui data user",
		})
	}

	return c.Status(200).JSON(fiber.Map{
		"message": "Data user berhasil diperbarui",
		"user":    user,
	})
}

func DeleteUser(c *fiber.Ctx) error {
	id := c.Params("id")

	var user models.User
	if err := database.DB.First(&user, id).Error; err != nil {
		return c.Status(404).JSON(fiber.Map{
			"error": "User tidak ditemukan",
		})
	}

	if err := database.DB.Delete(&user).Error; err != nil {
		return c.Status(500).JSON(fiber.Map{
			"error": "Gagal menghapus user",
		})
	}

	return c.Status(200).JSON(fiber.Map{
		"message": "User berhasil dihapus",
	})
}
