package authController

import (
	"github.com/gofiber/fiber/v2"
	"libro/apiSchema/authSchema"
	"libro/services/authService"
)

type Controller struct{ service *authService.Service }

func New(service *authService.Service) *Controller { return &Controller{service: service} }

func (ctl *Controller) Register(c *fiber.Ctx) error {
	var req authSchema.RegisterRequest
	if err := c.BodyParser(&req); err != nil {
		return c.Status(400).JSON(fiber.Map{"error": "bad payload"})
	}
	resp, err := ctl.service.Register(req)
	if err != nil {
		return c.Status(400).JSON(fiber.Map{"error": err.Error()})
	}
	return c.JSON(resp)
}
func (ctl *Controller) Login(c *fiber.Ctx) error {
	var req authSchema.LoginRequest
	if err := c.BodyParser(&req); err != nil {
		return c.Status(400).JSON(fiber.Map{"error": "bad payload"})
	}
	resp, err := ctl.service.Login(req)
	if err != nil {
		return c.Status(401).JSON(fiber.Map{"error": err.Error()})
	}
	return c.JSON(resp)
}
func (ctl *Controller) RefreshToken(c *fiber.Ctx) error {
	var req authSchema.RefreshTokenRequest
	if err := c.BodyParser(&req); err != nil {
		return c.Status(400).JSON(fiber.Map{"error": "bad payload"})
	}
	resp, err := ctl.service.RefreshToken(req.RefreshToken)
	if err != nil {
		return c.Status(401).JSON(fiber.Map{"error": err.Error()})
	}
	return c.JSON(resp)
}
func (ctl *Controller) Logout(c *fiber.Ctx) error {
	uid := c.Locals("userId").(uint)
	if err := ctl.service.Logout(uid); err != nil {
		return c.Status(400).JSON(fiber.Map{"error": err.Error()})
	}
	return c.JSON(fiber.Map{"message": "logged out"})
}
func (ctl *Controller) Me(c *fiber.Ctx) error {
	uid := c.Locals("userId").(uint)
	resp, err := ctl.service.Me(uid)
	if err != nil {
		return c.Status(404).JSON(fiber.Map{"error": err.Error()})
	}
	return c.JSON(resp)
}
