package userController

import (
	"github.com/gofiber/fiber/v2"
	"libro/apiSchema/userSchema"
	"libro/services/authService"
)

type Controller struct{ service *authService.Service }

func New(service *authService.Service) *Controller { return &Controller{service: service} }

func (ctl *Controller) GetProfile(c *fiber.Ctx) error {
	resp, err := ctl.service.GetProfile(c.Locals("userId").(uint))
	if err != nil {
		return c.Status(404).JSON(fiber.Map{"error": err.Error()})
	}
	return c.JSON(resp)
}
func (ctl *Controller) UpdateProfile(c *fiber.Ctx) error {
	var req userSchema.UpdateProfileRequest
	if err := c.BodyParser(&req); err != nil {
		return c.Status(400).JSON(fiber.Map{"error": "bad payload"})
	}
	resp, err := ctl.service.UpdateProfile(c.Locals("userId").(uint), req.Name)
	if err != nil {
		return c.Status(400).JSON(fiber.Map{"error": err.Error()})
	}
	return c.JSON(resp)
}
func (ctl *Controller) ChangePassword(c *fiber.Ctx) error {
	var req userSchema.ChangePasswordRequest
	if err := c.BodyParser(&req); err != nil {
		return c.Status(400).JSON(fiber.Map{"error": "bad payload"})
	}
	if err := ctl.service.ChangePassword(c.Locals("userId").(uint), req.CurrentPassword, req.NewPassword); err != nil {
		return c.Status(400).JSON(fiber.Map{"error": err.Error()})
	}
	return c.JSON(fiber.Map{"message": "password updated"})
}
