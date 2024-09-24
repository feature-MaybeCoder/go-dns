package middlewares

import "github.com/gofiber/fiber/v2"

func RegisterMiddlewares(fiber_app *fiber.App) {
	fiber_app.Use(
		GetLoggerMiddleware(),
	)
}
