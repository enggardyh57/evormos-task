package main

import (
	"log"

	"github.com/gofiber/fiber/v2"
	"evormos-task/database"
	"evormos-task/models"
	"evormos-task/routes"
)

func main() {
	
	app := fiber.New()

	
	database.ConnectDB()

	
	if err := database.DB.AutoMigrate(&models.User{}, &models.Toko{},&models.Alamat{},&models.Kategori{},&models.Produk{},&models.Transaksi{},&models.LogProduk{},
		&models.DetailTrx{},); err != nil {
	log.Fatal("Gagal melakukan migrasi:", err)
	}

	
	routes.SetupRoutes(app)

	
	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("API Fiber + GORM sudah aktif di port 3000!")
	})

	
	if err := app.Listen(":3000"); err != nil {
		log.Fatal("Gagal menjalankan server:", err)
	}
}