package handlers

import (
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/joaovitormgv/ecomLengoTengo/app/models"
)

// registrar página visitada
func (h *Handlers) RegisterPageVisited(c *fiber.Ctx) error {
	var requestBody models.UserNavigationHistory
	if err := c.BodyParser(&requestBody); err != nil {
		return err
	}

	var id int64
	var count int
	row := h.DB.QueryRow("SELECT COUNT(*) FROM user_navigation_history WHERE user_id = $1 AND product_id = $2",
		requestBody.UserID, requestBody.ProductID)
	err := row.Scan(&count)
	if err != nil {
		return err
	}

	if count > 0 {
		_, err = h.DB.Exec("UPDATE user_navigation_history SET time_visited = $1, action_taken = $2 WHERE user_id = $3 AND product_id = $4",
			time.Now(), requestBody.ActionTaken, requestBody.UserID, requestBody.ProductID)
		if err != nil {
			return err
		}
		return c.JSON(fiber.Map{
			"message":      "Histórico de navegação do usuário atualizado com sucesso",
			"time_visited": time.Now().Format("2006-01-02 15:04:05"),
			"action_taken": requestBody.ActionTaken,
		})

	} else {
		row = h.DB.QueryRow("INSERT INTO user_navigation_history (user_id, product_id, time_visited, action_taken) VALUES ($1, $2, $3, $4) RETURNING id",
			requestBody.UserID, requestBody.ProductID, time.Now(), requestBody.ActionTaken)
		err = row.Scan(&id)
		if err != nil {
			return err
		}
	}

	return c.JSON(fiber.Map{
		"id": id,
	})
}
