package wishlistController

import (
	"strconv"

	"github.com/gofiber/fiber/v2"
	"libro/apiSchema/purchaseLinkSchema"
	"libro/apiSchema/wishlistSchema"
	"libro/pkg/pagination"
	"libro/services/wishlistService"
)

type Controller struct{ service *wishlistService.Service }

func New(service *wishlistService.Service) *Controller { return &Controller{service: service} }

func getID(c *fiber.Ctx, name string) (uint, error) {
	id, err := strconv.Atoi(c.Params(name))
	return uint(id), err
}
func (ctl *Controller) CreateWishlistItem(c *fiber.Ctx) error {
	var req wishlistSchema.CreateWishlistRequest
	if err := c.BodyParser(&req); err != nil {
		return c.Status(400).JSON(fiber.Map{"error": "bad payload"})
	}
	resp, err := ctl.service.CreateWishlistItem(c.Locals("userId").(uint), req)
	if err != nil {
		return c.Status(400).JSON(fiber.Map{"error": err.Error()})
	}
	return c.JSON(resp)
}
func (ctl *Controller) GetWishlist(c *fiber.Ctx) error {
	req := pagination.Parse(c)
	resp, err := ctl.service.GetWishlist(c.Locals("userId").(uint), req)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"error": err.Error()})
	}
	return c.JSON(resp)
}
func (ctl *Controller) GetWishlistItemByID(c *fiber.Ctx) error {
	id, err := getID(c, "id")
	if err != nil {
		return c.Status(400).JSON(fiber.Map{"error": "invalid id"})
	}
	resp, err := ctl.service.GetWishlistItemByID(c.Locals("userId").(uint), id)
	if err != nil {
		return c.Status(404).JSON(fiber.Map{"error": err.Error()})
	}
	return c.JSON(resp)
}
func (ctl *Controller) UpdateWishlistItem(c *fiber.Ctx) error {
	id, err := getID(c, "id")
	if err != nil {
		return c.Status(400).JSON(fiber.Map{"error": "invalid id"})
	}
	var req wishlistSchema.UpdateWishlistRequest
	if err := c.BodyParser(&req); err != nil {
		return c.Status(400).JSON(fiber.Map{"error": "bad payload"})
	}
	resp, err := ctl.service.UpdateWishlistItem(c.Locals("userId").(uint), id, req)
	if err != nil {
		return c.Status(400).JSON(fiber.Map{"error": err.Error()})
	}
	return c.JSON(resp)
}
func (ctl *Controller) DeleteWishlistItem(c *fiber.Ctx) error {
	id, err := getID(c, "id")
	if err != nil {
		return c.Status(400).JSON(fiber.Map{"error": "invalid id"})
	}
	if err := ctl.service.DeleteWishlistItem(c.Locals("userId").(uint), id); err != nil {
		return c.Status(404).JSON(fiber.Map{"error": err.Error()})
	}
	return c.SendStatus(204)
}
func (ctl *Controller) AddPurchaseLink(c *fiber.Ctx) error {
	id, err := getID(c, "id")
	if err != nil {
		return c.Status(400).JSON(fiber.Map{"error": "invalid id"})
	}
	var req purchaseLinkSchema.CreatePurchaseLinkRequest
	if err := c.BodyParser(&req); err != nil {
		return c.Status(400).JSON(fiber.Map{"error": "bad payload"})
	}
	resp, err := ctl.service.AddPurchaseLink(c.Locals("userId").(uint), id, req)
	if err != nil {
		return c.Status(400).JSON(fiber.Map{"error": err.Error()})
	}
	return c.JSON(resp)
}
func (ctl *Controller) UpdatePurchaseLink(c *fiber.Ctx) error {
	id, err := getID(c, "id")
	if err != nil {
		return c.Status(400).JSON(fiber.Map{"error": "invalid id"})
	}
	linkID, err := getID(c, "linkId")
	if err != nil {
		return c.Status(400).JSON(fiber.Map{"error": "invalid linkId"})
	}
	var req purchaseLinkSchema.UpdatePurchaseLinkRequest
	if err := c.BodyParser(&req); err != nil {
		return c.Status(400).JSON(fiber.Map{"error": "bad payload"})
	}
	resp, err := ctl.service.UpdatePurchaseLink(c.Locals("userId").(uint), id, linkID, req)
	if err != nil {
		return c.Status(400).JSON(fiber.Map{"error": err.Error()})
	}
	return c.JSON(resp)
}
func (ctl *Controller) DeletePurchaseLink(c *fiber.Ctx) error {
	id, err := getID(c, "id")
	if err != nil {
		return c.Status(400).JSON(fiber.Map{"error": "invalid id"})
	}
	linkID, err := getID(c, "linkId")
	if err != nil {
		return c.Status(400).JSON(fiber.Map{"error": "invalid linkId"})
	}
	if err := ctl.service.DeletePurchaseLink(c.Locals("userId").(uint), id, linkID); err != nil {
		return c.Status(404).JSON(fiber.Map{"error": err.Error()})
	}
	return c.SendStatus(204)
}
