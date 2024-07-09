package main

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/lib/pq"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/session"
	"github.com/joaovitormgv/ecomLengoTengo/app/handlers"
	"github.com/joaovitormgv/ecomLengoTengo/app/middleware"
	"github.com/joaovitormgv/ecomLengoTengo/app/routes"
)

const (
	user     = "postgres"
	password = "123456"
	dbname   = "lengotengo"
)

var connectionString = fmt.Sprintf("postgres://%s:%s@localhost:5433/%s?sslmode=disable", user, password, dbname)

func main() {
	db, err := sql.Open("postgres", connectionString)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	store := session.New()

	h := &handlers.Handlers{
		Store: store,
		DB:    db,
	}

	app := fiber.New()
	app.Use(middleware.CorsMiddleware())
	routes.Setup(app, h)
	app.Listen(":3000")
}
