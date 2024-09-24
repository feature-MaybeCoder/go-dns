package rest

import (
	v1 "github.com/LastDarkNes/go-dns/internal/delivery/rest/v1"
	"github.com/gofiber/fiber/v2"
)

func AddV1RoutesToApp(fiber_app *fiber.App) {
	v1_group := fiber_app.Group("/v1")
	v1.AddDomainRoutesToGroup(v1_group)
}
