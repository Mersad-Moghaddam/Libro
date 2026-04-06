package bookController

import (
	"strconv"

	"github.com/gofiber/fiber/v2"
	"libro/apiSchema/bookSchema"
	"libro/pkg/pagination"
	"libro/services/bookService"
)

type Controller struct{ service *bookService.Service }

func New(service *bookService.Service) *Controller { return &Controller{service: service} }

func getID(c *fiber.Ctx, name string) (uint, error) {
	id, err := strconv.Atoi(c.Params(name))
	return uint(id), err
}

func (ctl *Controller) CreateBook(c *fiber.Ctx) error {
	var req bookSchema.CreateBookRequest
	if err := c.BodyParser(&req); err != nil {
		return c.Status(400).JSON(fiber.Map{"error": "bad payload"})
	}
	resp, err := ctl.service.CreateBook(c.Locals("userId").(uint), req)
	if err != nil {
		return c.Status(400).JSON(fiber.Map{"error": err.Error()})
	}
	return c.JSON(resp)
}
func (ctl *Controller) GetBooks(c *fiber.Ctx) error {
	req := pagination.Parse(c)
	resp, err := ctl.service.GetBooks(c.Locals("userId").(uint), req)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"error": err.Error()})
	}
	return c.JSON(resp)
}
func (ctl *Controller) GetBookByID(c *fiber.Ctx) error {
	id, err := getID(c, "id")
	if err != nil {
		return c.Status(400).JSON(fiber.Map{"error": "invalid id"})
	}
	resp, err := ctl.service.GetBookByID(c.Locals("userId").(uint), id)
	if err != nil {
		return c.Status(404).JSON(fiber.Map{"error": err.Error()})
	}
	return c.JSON(resp)
}
func (ctl *Controller) UpdateBook(c *fiber.Ctx) error {
	id, err := getID(c, "id")
	if err != nil {
		return c.Status(400).JSON(fiber.Map{"error": "invalid id"})
	}
	var req bookSchema.UpdateBookRequest
	if err := c.BodyParser(&req); err != nil {
		return c.Status(400).JSON(fiber.Map{"error": "bad payload"})
	}
	resp, err := ctl.service.UpdateBook(c.Locals("userId").(uint), id, req)
	if err != nil {
		return c.Status(400).JSON(fiber.Map{"error": err.Error()})
	}
	return c.JSON(resp)
}
func (ctl *Controller) DeleteBook(c *fiber.Ctx) error {
	id, err := getID(c, "id")
	if err != nil {
		return c.Status(400).JSON(fiber.Map{"error": "invalid id"})
	}
	if err := ctl.service.DeleteBook(c.Locals("userId").(uint), id); err != nil {
		return c.Status(404).JSON(fiber.Map{"error": err.Error()})
	}
	return c.SendStatus(204)
}
func (ctl *Controller) UpdateBookStatus(c *fiber.Ctx) error {
	id, err := getID(c, "id")
	if err != nil {
		return c.Status(400).JSON(fiber.Map{"error": "invalid id"})
	}
	var req bookSchema.UpdateBookStatusRequest
	if err := c.BodyParser(&req); err != nil {
		return c.Status(400).JSON(fiber.Map{"error": "bad payload"})
	}
	resp, err := ctl.service.UpdateBookStatus(c.Locals("userId").(uint), id, req.Status)
	if err != nil {
		return c.Status(400).JSON(fiber.Map{"error": err.Error()})
	}
	return c.JSON(resp)
}
func (ctl *Controller) UpdateBookBookmark(c *fiber.Ctx) error {
	id, err := getID(c, "id")
	if err != nil {
		return c.Status(400).JSON(fiber.Map{"error": "invalid id"})
	}
	var req bookSchema.UpdateBookBookmarkRequest
	if err := c.BodyParser(&req); err != nil {
		return c.Status(400).JSON(fiber.Map{"error": "bad payload"})
	}
	resp, err := ctl.service.UpdateBookBookmark(c.Locals("userId").(uint), id, req.CurrentPage)
	if err != nil {
		return c.Status(400).JSON(fiber.Map{"error": err.Error()})
	}
	return c.JSON(resp)
}
