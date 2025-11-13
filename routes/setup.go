package routes

import (
	"github.com/gofiber/fiber/v2"
	"evormos-task/handlers"
	"evormos-task/middlewares"
	
)

func SetupRoutes(app *fiber.App) {
	api := app.Group("/api")

	api.Post("/register", handlers.Register)
	api.Post("/login", handlers.Login)

	protected := api.Group("/toko", middlewares.Protected())
	protected.Get("/", func(c *fiber.Ctx) error {
		return c.JSON(fiber.Map{"message": "Kamu berhasil mengakses endpoint toko yang dilindungi JWT"})
	})

	user := api.Group("/user", middlewares.Protected())
	user.Get("/", handlers.GetUserProfile)
	user.Get("/:id", handlers.GetUserByID)
	user.Put("/update", handlers.UpdateUser)
	user.Delete("/:id", handlers.DeleteUser)

	alamat := api.Group("/alamat", middlewares.Protected())
	alamat.Get("/", handlers.GetAlamat)
	alamat.Post("/", handlers.CreateAlamat)
	alamat.Get("/:id", handlers.GetAlamatByID)
	alamat.Put("/:id", handlers.UpdateAlamat)
	alamat.Delete("/:id", handlers.DeleteAlamat)
}