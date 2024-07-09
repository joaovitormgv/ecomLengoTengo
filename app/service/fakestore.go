// service/fakestore.go

package service

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type Product struct {
	ID          int     `json:"id"`
	Title       string  `json:"title"`
	Price       float64 `json:"price"`
	Description string  `json:"description"`
	Category    string  `json:"category"`
	Image       string  `json:"image"`
}

// GetProducts faz uma requisição GET para a API Fake Store e retorna os produtos obtidos
func GetProducts(limit int) ([]Product, error) {
	url := fmt.Sprintf("https://fakestoreapi.com/products?limit=%d", limit)

	// Faz a requisição HTTP GET
	resp, err := http.Get(url)
	if err != nil {
		return nil, fmt.Errorf("erro ao fazer requisição HTTP: %v", err)
	}
	defer resp.Body.Close()

	// Verifica se a resposta da API foi bem-sucedida (código 200)
	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("falha na requisição: %s", resp.Status)
	}

	// Decodifica a resposta JSON
	var products []Product
	if err := json.NewDecoder(resp.Body).Decode(&products); err != nil {
		return nil, fmt.Errorf("erro ao decodificar resposta JSON: %v", err)
	}

	return products, nil
}

// Make a request to the Fake Store API and return the product and an error
func GetProduct(id int) (Product, error) {
	url := fmt.Sprintf("https://fakestoreapi.com/products/%d", id)

	// Make the HTTP GET request
	resp, err := http.Get(url)
	if err != nil {
		return Product{}, fmt.Errorf("error making HTTP request: %v", err)
	}
	defer resp.Body.Close()

	// Check if the API response was successful (status code 200)
	if resp.StatusCode != http.StatusOK {
		return Product{}, fmt.Errorf("request failed: %s", resp.Status)
	}

	// Decode the JSON response
	var product Product
	if err := json.NewDecoder(resp.Body).Decode(&product); err != nil {
		return Product{}, fmt.Errorf("error decoding JSON response: %v", err)
	}

	return product, nil
}

// Get products in a specific category
func GetProductsByCategory(category string) ([]Product, error) {
	url := fmt.Sprintf("https://fakestoreapi.com/products/category/%s", category)

	// Make the HTTP GET request
	resp, err := http.Get(url)
	if err != nil {
		return nil, fmt.Errorf("error making HTTP request: %v", err)
	}
	defer resp.Body.Close()

	// Check if the API response was successful (status code 200)
	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("request failed: %s", resp.Status)
	}

	// Decode the JSON response
	var products []Product
	if err := json.NewDecoder(resp.Body).Decode(&products); err != nil {
		return nil, fmt.Errorf("error decoding JSON response: %v", err)
	}

	return products, nil
}

// Get all categories
func GetCategories() ([]string, error) {
	url := "https://fakestoreapi.com/products/categories"

	// Make the HTTP GET request
	resp, err := http.Get(url)
	if err != nil {
		return nil, fmt.Errorf("error making HTTP request: %v", err)
	}
	defer resp.Body.Close()

	// Check if the API response was successful (status code 200)
	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("request failed: %s", resp.Status)
	}

	// Decode the JSON response
	var categories []string
	if err := json.NewDecoder(resp.Body).Decode(&categories); err != nil {
		return nil, fmt.Errorf("error decoding JSON response: %v", err)
	}

	return categories, nil
}

// Exemplos de uso

// func main() {
// 	products, err := service.GetProductsByCategory("women's clothing")
// 	// products, err := service.GetProducts(20)
// 	// product, err := service.GetProduct(15)
// 	if err != nil {
// 		log.Fatalf("Erro ao obter produtos da Fake Store API: %v", err)
// 	}

// 	// products := []service.Product{product}
// 	for _, p := range products {
// 		fmt.Printf("ID: %d\n", p.ID)
// 		fmt.Printf("Title: %s\n", p.Title)
// 		fmt.Printf("Price: %.2f\n", p.Price)
// 		fmt.Printf("Description: %s\n", p.Description)
// 		fmt.Printf("Category: %s\n", p.Category)
// 		fmt.Printf("Image: %s\n", p.Image)
// 		fmt.Println("-----------------------------------")
// 	}
// }

// func main() {
// 	categories, err := service.GetCategories()
// 	if err != nil {
// 		log.Fatalf("Erro ao obter categorias da Fake Store API: %v", err)
// 	}

// 	for _, c := range categories {
// 		fmt.Println(c)
// 	}
// }
