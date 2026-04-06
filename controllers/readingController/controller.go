package readingController

import (
	"strconv"

	"github.com/gofiber/fiber/v2"
	"libro/apiSchema/readingSchema"
	"libro/services/readingService"
)

type Controller struct{ service *readingService.Service }

func New(service *readingService.Service) *Controller { return &Controller{service: service} }

func (ctl *Controller) GetCurrentReadingBooks(c *fiber.Ctx) error {
	resp, err := ctl.service.GetCurrentReadingBooks(c.Locals("userId").(uint))
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"error": err.Error()})
	}
	return c.JSON(fiber.Map{"items": resp})
}
func (ctl *Controller) UpdateReadingProgress(c *fiber.Ctx) error {
	id, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return c.Status(400).JSON(fiber.Map{"error": "invalid id"})
	}
	var req readingSchema.UpdateReadingProgressRequest
	if err := c.BodyParser(&req); err != nil {
		return c.Status(400).JSON(fiber.Map{"error": "bad payload"})
	}
	resp, err := ctl.service.UpdateReadingProgress(c.Locals("userId").(uint), uint(id), req.CurrentPage)
	if err != nil {
		return c.Status(400).JSON(fiber.Map{"error": err.Error()})
	}
	return c.JSON(resp)
}
