package handlers

import (
	"log"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	"github.com/joaovitormgv/ecomLengoTengo/app/models"
)

// RegisterSession registra uma nova sessão
func (h *Handlers) RegisterSession(c *fiber.Ctx) error {
	// Gera um ID de sessão único
	sessionID := generateSessionID()

	// Cria um novo objeto de sessão
	session := models.Session{}
	session.SessionID = sessionID.String()
	// Analisa o corpo da requisição em um objeto de sessão
	if err := c.BodyParser(&session); err != nil {
		log.Printf("Falha ao analisar o corpo da requisição: %v", err)
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Falha ao analisar o corpo da requisição",
		})
	}

	session.CreatedAt = time.Now()
	session.UpdatedAt = time.Now()

	// Insere a sessão no banco de dados
	_, err := h.DB.Exec("INSERT INTO sessions (session_id, data, created_at, updated_at) VALUES ($1, $2, $3, $4)", sessionID, session.Data, session.CreatedAt, session.UpdatedAt)
	if err != nil {
		log.Printf("Falha ao registrar a sessão: %v", err)
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "Falha ao registrar a sessão",
		})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"message":    "Sessão registrada com sucesso",
		"session_id": session.SessionID,
		"created_at": session.CreatedAt.Format("2006-01-02 15:04:05"),
		"updated_at": session.UpdatedAt.Format("2006-01-02 15:04:05"),
	})
}

func generateSessionID() uuid.UUID {
	sessionID := uuid.New()
	return sessionID
}

// RegisterSessionNavigationHistory registra um novo histórico de navegação da sessão
func (h *Handlers) RegisterSessionNavigationHistory(c *fiber.Ctx) error {
	// Analisa o corpo da requisição em um objeto de histórico de navegação da sessão
	sessionNavigationHistory := models.SessionNavigationHistory{}
	if err := c.BodyParser(&sessionNavigationHistory); err != nil {
		log.Printf("Falha ao analisar o corpo da requisição: %v", err)
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Falha ao analisar o corpo da requisição",
		})
	}

	sessionNavigationHistory.TimeVisited = time.Now()

	// Verifica se já existe uma linha na tabela com mesmo sessionid e mesmo product id da requisição
	var count int
	err := h.DB.QueryRow("SELECT COUNT(*) FROM session_navigation_history WHERE session_id = $1 AND product_id = $2", sessionNavigationHistory.SessionID, sessionNavigationHistory.ProductID).Scan(&count)
	if err != nil {
		log.Printf("Falha ao verificar a existência do histórico de navegação da sessão: %v", err)
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "Falha ao verificar a existência do histórico de navegação da sessão",
		})
	}

	if count > 0 {
		// Atualiza apenas o horário de time_visited da linha existente
		_, err := h.DB.Exec("UPDATE session_navigation_history SET time_visited = $1 WHERE session_id = $2 AND product_id = $3", sessionNavigationHistory.TimeVisited, sessionNavigationHistory.SessionID, sessionNavigationHistory.ProductID)
		if err != nil {
			log.Printf("Falha ao atualizar o histórico de navegação da sessão: %v", err)
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"message": "Falha ao atualizar o histórico de navegação da sessão",
			})
		}

		return c.Status(fiber.StatusOK).JSON(fiber.Map{
			"message":      "Histórico de navegação da sessão atualizado com sucesso",
			"time_visited": sessionNavigationHistory.TimeVisited.Format("2006-01-02 15:04:05"),
		})
	}
	// Insere o histórico de navegação da sessão no banco de dados
	_, err = h.DB.Exec("INSERT INTO session_navigation_history (session_id, product_id, time_visited, action_taken) VALUES ($1, $2, $3, $4)", sessionNavigationHistory.SessionID, sessionNavigationHistory.ProductID, sessionNavigationHistory.TimeVisited, sessionNavigationHistory.ActionTaken)
	if err != nil {
		log.Printf("Falha ao registrar o histórico de navegação da sessão: %v", err)
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "Falha ao registrar o histórico de navegação da sessão",
		})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"message":      "Histórico de navegação da sessão registrado com sucesso",
		"action_taken": sessionNavigationHistory.ActionTaken,
		"time_visited": sessionNavigationHistory.TimeVisited.Format("2006-01-02 15:04:05"),
		"product_id":   sessionNavigationHistory.ProductID,
		"session_id":   sessionNavigationHistory.SessionID,
	})
}
