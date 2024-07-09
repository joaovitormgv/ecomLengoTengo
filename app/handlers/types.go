package handlers

import (
	"database/sql"

	"github.com/gofiber/fiber/v2/middleware/session"
)

type Handlers struct {
	Store *session.Store
	DB    *sql.DB
}
