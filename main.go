package main

import (
	"ecomVoraz/app/service"
	"fmt"
	"log"
)

// import (
// 	"github.com/gofiber/fiber/v2"
// )

// func main() {
// 	app := fiber.New()

// 	app.Get("/", func(c *fiber.Ctx) error {
// 		return c.SendString("Hello, World!")
// 	})

// 	app.Listen(":3000")
// }

func main() {
	products, err := service.GetProducts(5)
	if err != nil {
		log.Fatalf("Erro ao obter produtos da Fake Store API: %v", err)
	}

	for _, p := range products {
		fmt.Printf("ID: %d\n", p.ID)
		fmt.Printf("Title: %s\n", p.Title)
		fmt.Printf("Price: %.2f\n", p.Price)
		fmt.Printf("Description: %s\n", p.Description)
		fmt.Printf("Category: %s\n", p.Category)
		fmt.Printf("Image: %s\n", p.Image)
		fmt.Println("-----------------------------------")
	}
}
