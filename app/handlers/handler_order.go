package handlers

import (
	"database/sql"

	"github.com/gofiber/fiber/v2"
	"github.com/joaovitormgv/ecomLengoTengo/app/models"
)

func (h *Handlers) GetOrders(c *fiber.Ctx) error {
	sess, err := h.Store.Get(c)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	userID := sess.Get("user_id")

	var rows *sql.Rows
	rows, err = h.DB.Query("SELECT * FROM orders WHERE user_id = $1", userID)

	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": err.Error(),
		})
	}
	defer rows.Close()

	orders := []models.Order{}
	for rows.Next() {
		order := models.Order{}
		err = rows.Scan(&order.OrderID, &order.UserID, &order.ProductID, &order.ProductName, &order.Category, &order.Quantity, &order.Price, &order.OrderDate)
		if err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"error": err.Error(),
			})
		}
		orders = append(orders, order)
	}

	return c.JSON(orders)
}
