package v1

import (
	"encoding/json"
	"fmt"
	"strconv"

	"github.com/LastDarkNes/go-dns/internal/model"
	"github.com/LastDarkNes/go-dns/internal/service"
	"github.com/gofiber/fiber/v2"
)

type DomainService interface {
	Get(string) (*model.Domain, error)
	Create(model.Domain) (*model.Domain, error)
	GetPage(int) []model.Domain
}

func AddDomainRoutesToGroup(group fiber.Router) {
	domain_service := service.NewDomainService()

	group.Get("/domains", func(c *fiber.Ctx) error {
		pageString := c.Query("page")

		page, err := strconv.Atoi(pageString)
		if err != nil {
			fmt.Println("error = ", err)
			return c.SendStatus(400)
		}

		domains := domain_service.GetPage(page)

		result := make([]string, len(domains))

		for i, domain := range domains {
			serialized_domain, err := json.Marshal(domain)
			if err != nil {
				fmt.Println("error = ", err)
				return c.SendStatus(400)
			}
			result[i] = string(serialized_domain)
		}

		return c.JSON(result)
	})

	group.Post("/domains", func(c *fiber.Ctx) error {
		c.Accepts("text/plain", "application/json")

		domain := new(model.Domain)
		err := c.BodyParser(domain)
		if err != nil {
			fmt.Println("error = ", err)
			return c.SendStatus(400)
		}

		new_domain, err := domain_service.Create(*domain)
		if err != nil {
			fmt.Println("error = ", err)
			return c.SendStatus(400)
		}

		result, err := json.Marshal(new_domain)
		if err != nil {
			fmt.Println("error = ", err)
			return c.SendStatus(400)
		}

		return c.JSON(string(result))
	})
}
