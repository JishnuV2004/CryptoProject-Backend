package middleware

import (
	webconfiguration "cryptox/internal/modules/webConfiguration"
	"cryptox/packages/utils"

	"github.com/gofiber/fiber/v2"
)

func FeatureMiddleware(service webconfiguration.FeatureService, feature string) fiber.Handler {
		return func(c *fiber.Ctx) error {

		var req struct {
			Email string `json:"email"`
			Password string `json:"password"`
		}
		if err := c.BodyParser(&req); err != nil {
			return c.Status(400).JSON(fiber.Map{
				"error": utils.Error(c, 400, "invalid input", nil),
			})
		}

		if role := service.Find(req.Email); role == "admin" {
			return c.Next()
		}

		if !service.IsEnabled(feature) {
			return utils.Error(c, 403, "Sorry " + feature + " disabled for App maintanance", nil)
		}
		return c.Next()
	}
}