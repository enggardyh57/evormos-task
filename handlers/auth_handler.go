package handlers

import (
	"time"
	"github.com/gofiber/fiber/v2"
	"golang.org/x/crypto/bcrypt"
	"github.com/golang-jwt/jwt/v4"
	"evermos-task/database"
	"evermos-task/models"
)

func Register(c *fiber.Ctx) error {
	user := new(models.User)
	if err := c.BodyParser(user); err != nil {
		return c.Status(400).JSON(fiber.Map{
			"error":   "Invalid input",
			"details": err.Error(),
		})
	}

	
	hash, _ := bcrypt.GenerateFromPassword([]byte(user.Kata_sandi), 14)
	user.Kata_sandi = string(hash)

	
	if user.IsAdmin == false {
		user.IsAdmin = false
	}

	
	if err := database.DB.Create(&user).Error; err != nil {
		return c.Status(500).JSON(fiber.Map{"error": err.Error()})
	}

	
	toko := models.Toko{
		Nama_Toko: user.Nama + " Store",
		ID_User:   user.ID,
	}
	database.DB.Create(&toko)

	return c.Status(201).JSON(fiber.Map{
		"message": "Registrasi berhasil dan toko dibuat otomatis",
		"user":    user,
		"toko":    toko,
	})
}

func Login(c *fiber.Ctx) error {
	var input struct {
		Email    string `json:"email"`
		Kata_sandi string `json:"kata_sandi"`
	}

	if err := c.BodyParser(&input); err != nil {
		return c.Status(400).JSON(fiber.Map{"error": "Gagal parsing data"})
	}

	var user models.User
	if err := database.DB.Where("email = ?", input.Email).First(&user).Error; err != nil {
		return c.Status(401).JSON(fiber.Map{"error": "Email tidak ditemukan"})
	}

	if err := bcrypt.CompareHashAndPassword([]byte(user.Kata_sandi), []byte(input.Kata_sandi)); err != nil {
		return c.Status(401).JSON(fiber.Map{"error": "Password salah"})
	}

	
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"user_id":  user.ID,
		"is_admin": user.IsAdmin, 
		"exp":      time.Now().Add(time.Hour * 24).Unix(),
	})

	tokenString, _ := token.SignedString([]byte("supersecretkey"))

	return c.JSON(fiber.Map{
		"message": "Login berhasil",
		"token":   tokenString,
		"is_admin": user.IsAdmin,
	})
}