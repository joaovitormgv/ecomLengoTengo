package handlers

import (
	"database/sql"

	"github.com/gofiber/fiber/v2"
	"github.com/joaovitormgv/ecomLengoTengo/app/models"
)

func (h *Handlers) LoginUser(c *fiber.Ctx) error {
	// Verificar se há uma sessão ativa
	sess, err := h.Store.Get(c)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	if sess.Get("user_id") != nil {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"error": "Ainda há uma sessão ativa",
		})
	}

	user := &models.User{}
	err = c.BodyParser(user)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	row := h.DB.QueryRow("SELECT password, id FROM users WHERE username = $1", user.Username)
	var receivedPassword string
	err = row.Scan(&receivedPassword, &user.ID)
	if err != nil {
		if err == sql.ErrNoRows {
			return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
				"error": "Usuário não encontrado",
			})
		} else {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"error": err.Error(),
			})
		}
	}

	if receivedPassword != user.Password {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"error": "Senha incorreta",
		})
	}

	sess, err = h.Store.Get(c)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	sess.Set("user_id", user.ID)
	err = sess.Save()
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	return c.JSON(fiber.Map{
		"message": "Login efetuado com sucesso",
	})
}

func (h *Handlers) LogoutUser(c *fiber.Ctx) error {
	sess, err := h.Store.Get(c)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	sess.Destroy()
	err = sess.Save()
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	return c.JSON(fiber.Map{
		"message": "Logout efetuado com sucesso",
	})
}
