package main

import (
	// "fmt"
	"time"

	"github.com/gofiber/fiber/v2"
)

func main() {
    app := fiber.New()

    app.Get("/", func(c *fiber.Ctx) error {
		
    })

    app.Listen(":3000")
}