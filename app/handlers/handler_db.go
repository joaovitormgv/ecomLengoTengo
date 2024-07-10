package handlers

import "github.com/joaovitormgv/ecomLengoTengo/app/models"

// SaveProductToDB salva um produto no banco de dados
func (h *Handlers) SaveProductToDB(product models.Product) error {
	query := `INSERT INTO products (id, title, price, description, category, image) 
              VALUES ($1, $2, $3, $4, $5, $6)`
	_, err := h.DB.Exec(query, product.ID, product.Title, product.Price, product.Description, product.Category, product.Image)
	return err
}
