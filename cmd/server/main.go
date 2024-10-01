package main

import (
	"github.com/LastDarkNes/go-dns/internal/delivery/rest"
	"github.com/LastDarkNes/go-dns/pkg/middlewares"
	"github.com/gofiber/fiber/v2"
)

func main() {
	app := fiber.New()
	middlewares.RegisterMiddlewares(app)
	rest.AddV1RoutesToApp(app)

	app.Listen(":3000")
}
