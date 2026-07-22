package webconfiguration

import (
	"cryptox/internal/modules/market"
	"cryptox/packages/utils"
	"strconv"

	"github.com/gofiber/contrib/websocket"
	"github.com/gofiber/fiber/v2"
)

type Controller struct {
	service FeatureService
}

func NewController(service FeatureService) *Controller {
	return &Controller{service: service}
}

func (h *Controller) GetFeatures(c *fiber.Ctx) error {

	features, err := h.service.GetFeatures()

	if err != nil {
		return utils.Error(c, 500, "error", err.Error())
	}
	return utils.Success(c, 200, "success", features)
}

func (h *Controller) UpdateFeature(c *fiber.Ctx) error {

	idStr := c.Params("id")

	id64, err := strconv.ParseUint(idStr, 10, 32)
	if err != nil {
		return utils.Error(c, 400, "invalid id", err.Error())
	}

	id := uint(id64)

	var req struct {
		Enabled bool `json:"enabled"`
	}

	if err := c.BodyParser(&req); err != nil {
		return utils.Error(c, 400, "error", "invalid request")
	}

	err = h.service.Update(id, req.Enabled)

	if err != nil {
		return utils.Error(c, 500, "error", err.Error())
	}
	return utils.Success(c, 200, "success", "updated")
}

func WebSocketHandler(hub *market.Hub) fiber.Handler {

	return websocket.New(
		func(c *websocket.Conn){

			hub.AddGlobal(c)

			defer hub.RemoveGlobal(c)

			for {
				_,_,err:=
				c.ReadMessage()

				if err!=nil{
					break
				}
			}
		},
	)
}