package main

import (
	// "time"

	"fmt"

	"github.com/gofiber/fiber/v2"
)

func main() {
	app := fiber.New()

	app.Get("/:new/:abc", func(c *fiber.Ctx) error {
		fmt.Println(c.AllParams())
		fmt.Println("********************")
		fmt.Println(c.BaseURL())
		fmt.Println("********************")
		fmt.Println(c.Body())
		fmt.Println("********************")
		fmt.Println(c.BodyRaw())
		fmt.Println("********************")
		fmt.Println(c.Fresh())
		fmt.Println("********************")
		fmt.Println(c.GetReqHeaders())
		fmt.Println("********************")
		fmt.Println(c.GetRespHeaders())
		fmt.Println("********************")
		fmt.Println(c.Hostname())
		fmt.Println("********************")
		fmt.Println(c.IP())
		fmt.Println("********************")
		fmt.Println(c.IPs())
		fmt.Println("********************")
		fmt.Println(c.IsFromLocal())
		fmt.Println("********************")
		fmt.Println(c.OriginalURL())
		fmt.Println("********************")
		fmt.Println(c.Port())
		fmt.Println("********************")
		fmt.Println(c.Protocol())
		fmt.Println("********************")
		fmt.Println(c.Queries())
		fmt.Println("********************")
		fmt.Println(c.String())
		fmt.Println("********************")
		fmt.Println(c.Accepts("text/plain", "application/json"),"*****")
		fmt.Println("********************")
		fmt.Println("********************")
		fmt.Println("********************")

		return c.SendString("Hello, World!")
	})

	app.Listen("127.0.0.1:3000")

}
