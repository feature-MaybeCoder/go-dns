package middlewares

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
)

func GetLoggerMiddleware() func(*fiber.Ctx) error {
	new_logger := logger.New(logger.Config{
		Format: "${time} | [${status}] ${method} - ${path}\n",
	})
	return new_logger
}
