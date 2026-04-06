package mainController

import "github.com/gofiber/fiber/v2"

type Controller struct{}

func New() *Controller                            { return &Controller{} }
func (ctl *Controller) Health(c *fiber.Ctx) error { return c.JSON(fiber.Map{"status": "ok"}) }
func (ctl *Controller) DashboardSummary(c *fiber.Ctx) error {
	return c.JSON(fiber.Map{"summary": "Libro API running"})
}
