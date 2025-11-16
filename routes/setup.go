package routes

import (
	"github.com/gofiber/fiber/v2"
	"evermos-task/handlers"
	"evermos-task/middlewares"
)

func SetupRoutes(app *fiber.App) {

	auth := app.Group("/auth")
	auth.Post("/login", handlers.Login)
	auth.Post("/register", handlers.Register)

	user := app.Group("/user", middlewares.Protected())
	user.Get("/", handlers.GetUserProfile)
	user.Put("/", handlers.UpdateUser)

	alamat := user.Group("/alamat")
	alamat.Get("/", handlers.GetAlamat)
	alamat.Post("/", handlers.CreateAlamat)
	alamat.Get("/:id", handlers.GetAlamatByID)
	alamat.Put("/:id", handlers.UpdateAlamat)
	alamat.Delete("/:id", handlers.DeleteAlamat)

	category := app.Group("/category", middlewares.Protected())
	category.Get("/", handlers.GetAllKategori)
	category.Post("/", handlers.CreateKategori)
	category.Get("/:id", handlers.GetKategoriByID)
	category.Put("/:id", handlers.UpdateKategori)
	category.Delete("/:id", handlers.DeleteKategori)

	toko := app.Group("/toko", middlewares.Protected())
	toko.Get("/my", handlers.GetMyToko)
	toko.Put("/:id_toko", handlers.UpdateMyToko)
	toko.Get("/:id_toko", handlers.GetTokoByID)
	toko.Get("/", handlers.GetAllToko)

	product := app.Group("/product", middlewares.Protected())
	product.Get("/", handlers.GetAllProduk)
	product.Get("/:id", handlers.GetProdukByID)
	product.Post("/", handlers.CreateProduk)
	product.Put("/:id", handlers.UpdateProduk)
	product.Delete("/:id", handlers.DeleteProduk)

	tr := app.Group("/trx")
	tr.Get("/", handlers.GetAllTransaksi)
	tr.Get("/:id", handlers.GetTransaksiByID)
	tr.Post("/", handlers.CreateTransaksi)
	
}
