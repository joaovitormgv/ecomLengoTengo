package handlers

import (
	"database/sql"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	"github.com/joaovitormgv/ecomLengoTengo/app/models"
)

func (h *Handlers) GetOwnOrders(c *fiber.Ctx) error {
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

func (h *Handlers) CreateOrder(c *fiber.Ctx) error {
	sess, err := h.Store.Get(c)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	userID := sess.Get("user_id")
	order := &models.Order{}
	if err := c.BodyParser(order); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	// Implementar um get product by id
	// De modo a verificar se o produto existe e pegar as suas informações

	// fazer função de validação de dados

	orderID := uuid.New()

	_, err = h.DB.Exec("INSERT INTO orders (order_id, user_id, product_id, product_name, category, quantity, price, order_date) VALUES ($1, $2, $3, $4, $5, $6, $7, $8)", orderID, userID, order.ProductID, order.ProductName, order.Category, order.Quantity, order.Price, time.Now().Format("2006-01-02"))
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	return c.JSON(fiber.Map{
		"order_id": orderID,
	})
}
