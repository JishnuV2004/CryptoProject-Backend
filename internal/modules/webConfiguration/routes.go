package webconfiguration

import (
	"github.com/gofiber/fiber/v2"
)

func WebConfigRoutes(r fiber.Router, webService FeatureService) {
	
	webcontroller := NewController(webService)

	admin := r.Group("/admin")

	admin.Get("/features", webcontroller.GetFeatures)
	admin.Put("/features/:id", webcontroller.UpdateFeature)
}