package main

import (
	"github.com/LastDarkNes/go-dns/internal/delivery/rest"
	"github.com/gofiber/fiber/v2"
)

func main() {
	app := fiber.New()
	rest.AddV1RoutesToApp(app)
	app.Listen(":3000")
}
