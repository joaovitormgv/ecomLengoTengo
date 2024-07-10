package handlers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/joaovitormgv/ecomLengoTengo/app/models"
)

func (h *Handlers) GetProductById(c *fiber.Ctx) error {
	id, err := c.ParamsInt("id")
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	product := &models.Product{}

	row := h.DB.QueryRow("SELECT * FROM products WHERE id = $1", id)
	err = row.Scan(&product.ID, &product.Title, &product.Description, &product.Price, &product.Category, &product.Image)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	return c.JSON(product)
}
