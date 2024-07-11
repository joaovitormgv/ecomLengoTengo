package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/joaovitormgv/ecomLengoTengo/app/handlers"
	"github.com/joaovitormgv/ecomLengoTengo/app/middleware"
)

func Setup(app *fiber.App, h *handlers.Handlers) {
	// Métodos para manipular pedidos
	app.Get("api/orders", middleware.AuthRequired(h.Store), h.GetOwnOrders)
	app.Post("api/orders", middleware.AuthRequired(h.Store), h.CreateOrder)

	// Métodos para manipular usuários
	app.Post("/api/register", h.RegisterUser)
	app.Post("/api/login", h.LoginUser)
	app.Post("/api/logout", middleware.AuthRequired(h.Store), h.LogoutUser)

	// Métodos para manipular produtos
	app.Get("/api/product/:id", h.GetProductById)
	app.Get("/api/products", h.GetProducts)

	// Métodos para manipular sessões
	app.Post("/api/sessions", h.RegisterSession)

	// Métodos para manipular histórico de navegação de sessões
	app.Post("/api/sessions/navigation", h.RegisterSessionNavigationHistory)

	// Métodos para manipular histórico de navegação de usuários
	app.Post("/api/user/navigation", middleware.AuthRequired(h.Store), h.RegisterPageVisited)
}
