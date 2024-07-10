// service/fakestore.go

package service

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/joaovitormgv/ecomLengoTengo/app/models"
)

// GetProducts faz uma requisição GET para a API Fake Store e retorna os produtos obtidos
func GetProducts(limit int) ([]models.Product, error) {
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
	var products []models.Product
	if err := json.NewDecoder(resp.Body).Decode(&products); err != nil {
		return nil, fmt.Errorf("erro ao decodificar resposta JSON: %v", err)
	}

	return products, nil
}

// Make a request to the Fake Store API and return the product and an error
func GetProduct(id int) (models.Product, error) {
	url := fmt.Sprintf("https://fakestoreapi.com/products/%d", id)

	// Make the HTTP GET request
	resp, err := http.Get(url)
	if err != nil {
		return models.Product{}, fmt.Errorf("error making HTTP request: %v", err)
	}
	defer resp.Body.Close()

	// Check if the API response was successful (status code 200)
	if resp.StatusCode != http.StatusOK {
		return models.Product{}, fmt.Errorf("request failed: %s", resp.Status)
	}

	// Decode the JSON response
	var product models.Product
	if err := json.NewDecoder(resp.Body).Decode(&product); err != nil {
		return models.Product{}, fmt.Errorf("error decoding JSON response: %v", err)
	}

	return product, nil
}

// Get products in a specific category
func GetProductsByCategory(category string) ([]models.Product, error) {
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
	var products []models.Product
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
