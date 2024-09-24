package main

import (
	"encoding/json"
	"fmt"
	"strconv"

	"github.com/LastDarkNes/go-dns/controller"
	"github.com/LastDarkNes/go-dns/model"
	"github.com/LastDarkNes/go-dns/service"
	"github.com/gofiber/fiber/v2"
)

func main() {
	app := fiber.New()

	var domain_controller = controller.DomainController{
		Service: service.DomainService{},
	}

	app.Get("/domains", func(c *fiber.Ctx) error {
		pageString := c.Query("page")

		page, err := strconv.Atoi(pageString)
		if err != nil {
			fmt.Println("error = ", err)
			return c.SendStatus(400)
		}

		domains := domain_controller.GetPaginated(page)

		var result []string

		for _, domain := range domains {
			serialized_domain, err := json.Marshal(domain)

			if err != nil {
				fmt.Println("error = ", err)
				return c.SendStatus(400)
			}

			result = append(result, string(serialized_domain))
		}

		if err != nil {
			fmt.Println("error = ", err)
			return c.SendStatus(400)
		}

		return c.JSON(result)
	})

	app.Post("/domains", func(c *fiber.Ctx) error {
		c.Accepts("text/plain", "application/json")

		domain := new(model.Domain)
		err := c.BodyParser(domain)
		if err != nil {
			fmt.Println("error = ", err)
			return c.SendStatus(400)
		}

		new_domain, err := domain_controller.Create(*domain)
		if err != nil {
			fmt.Println("error = ", err)
			return c.SendStatus(400)
		}

		result, err := json.Marshal(new_domain)
		if err != nil {
			fmt.Println("error = ", err)
			return c.SendStatus(400)
		}
		println(string(result))
		return c.JSON(string(result))
	})

	app.Listen(":3000")
}
