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

// Exemplo de uso:
// func main() {
//     products, err := service.GetProducts(5)
//     if err != nil {
//         log.Fatalf("Erro ao obter produtos da Fake Store API: %v", err)
//     }
//
//     for _, p := range products {
//         fmt.Printf("ID: %d\n", p.ID)
//         fmt.Printf("Title: %s\n", p.Title)
//         fmt.Printf("Price: %.2f\n", p.Price)
//         fmt.Printf("Description: %s\n", p.Description)
//         fmt.Printf("Category: %s\n", p.Category)
//         fmt.Printf("Image: %s\n", p.Image)
//         fmt.Println("-----------------------------------")
//     }
// }
