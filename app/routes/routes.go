package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/joaovitormgv/ecomLengoTengo/app/handlers"
	"github.com/joaovitormgv/ecomLengoTengo/app/middleware"
)

func Setup(app *fiber.App, h *handlers.Handlers) {
	// Métodos para manipular pedidos
	// app.Get("api/orders" )

	// Métodos para manipular usuários
	app.Post("/api/login", h.LoginUser)
	app.Post("/api/logout", middleware.AuthRequired(h.Store), h.LogoutUser)
}
