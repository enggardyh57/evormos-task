package main

import (
    "log"

    "github.com/gofiber/fiber/v2"
    "evormos-task/database"
)

func main() {
    app := fiber.New()

    database.ConnectDB()

    app.Get("/", func(c *fiber.Ctx) error {
        return c.SendString("Fiber sudah jalan")
    })
    app.Listen(":3000")
}
