package handlers

import (
	"database/sql"
	"strconv"

	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	"github.com/joaovitormgv/ecomLengoTengo/app/models"
)

type CategoryView struct {
	Category string
	Views    int
}

// Modificação na função PopularProductsHandler para aceitar um parâmetro de consulta 'limit'
func (h *Handlers) PopularProductsHandler(c *fiber.Ctx) error {
	// Ler o parâmetro 'limit' da URL, com um valor padrão de 5 se não for especificado
	limitParam := c.Query("limit", "5")
	limit, err := strconv.Atoi(limitParam)
	if err != nil {
		return err
	}

	// Passar o limite para a função RecommendPopularProducts
	ids, err := h.RecommendPopularProducts(limit)
	if err != nil {
		return err
	}

	products := []models.Product{}
	for _, id := range ids {
		product := models.Product{}
		row := h.DB.QueryRow("SELECT * FROM products WHERE id = $1", id)
		err = row.Scan(&product.ID, &product.Title, &product.Description, &product.Price, &product.Category, &product.Image)
		if err != nil {
			return err
		}
		products = append(products, product)
	}

	return c.JSON(products)
}

// Modificação na função RecommendPopularProducts para aceitar um limite
func (h *Handlers) RecommendPopularProducts(limit int) ([]int, error) {
	// Usar o argumento 'limit' na consulta SQL
	query := `
	SELECT                 
		p.id, 
		COUNT(o.product_id) AS total_sales
	FROM 
		products p
	JOIN 
		orders o ON p.id = o.product_id
	GROUP BY 
		p.id
	ORDER BY 
		total_sales DESC
	LIMIT $1;
`

	rows, err := h.DB.Query(query, limit)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var products []int
	for rows.Next() {
		var product int
		var totalSales int
		if err := rows.Scan(&product, &totalSales); err != nil {
			return nil, err
		}
		products = append(products, product)
	}

	return products, nil
}

// Função para selecionar os produtos mais vendidos por categoria
func (h *Handlers) PopularProductsByCategoryHandler(c *fiber.Ctx) error {
	// Ler o parâmetro 'category' da URL
	category := c.Params("category")

	// Passar a categoria para a função RecommendPopularProductsByCategory
	ids, err := h.RecommendPopularProductsByCategory(category, 5)
	if err != nil {
		return err
	}

	products := []models.Product{}
	for _, id := range ids {
		product := models.Product{}
		row := h.DB.QueryRow("SELECT * FROM products WHERE id = $1", id)
		err = row.Scan(&product.ID, &product.Title, &product.Description, &product.Price, &product.Category, &product.Image)
		if err != nil {
			return err
		}
		products = append(products, product)
	}

	return c.JSON(products)
}

// Função para recomendar produtos mais vendidos por categoria
func (h *Handlers) RecommendPopularProductsByCategory(category string, limit int) ([]int, error) {
	// Consulta SQL para selecionar os produtos mais vendidos em uma categoria específica
	query := `
	SELECT                 
		p.id, 
		COUNT(o.product_id) AS total_sales
	FROM 
		products p
	JOIN 
		orders o ON p.id = o.product_id
	WHERE 
		p.category = $1
	GROUP BY 
		p.id
	ORDER BY 
		total_sales DESC
	LIMIT $2;
`

	rows, err := h.DB.Query(query, category, limit)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var products []int
	for rows.Next() {
		var product int
		var totalSales int
		if err := rows.Scan(&product, &totalSales); err != nil {
			return nil, err
		}
		products = append(products, product)
	}

	return products, nil
}

func getTopCategoriesForUser(db *sql.DB, userID int) ([]CategoryView, error) {
	query := `
        SELECT p.category, COUNT(*) AS views
        FROM user_navigation_history unh
        JOIN products p ON unh.product_id = p.id
        WHERE unh.user_id = $1
        GROUP BY p.category
        ORDER BY views DESC
        LIMIT 4;
    `
	rows, err := db.Query(query, userID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var categories []CategoryView
	for rows.Next() {
		var cv CategoryView
		if err := rows.Scan(&cv.Category, &cv.Views); err != nil {
			return nil, err
		}
		categories = append(categories, cv)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return categories, nil
}

func (h *Handlers) RecommendProductsBasedOnCategoryViews(userID int) ([]models.Product, error) {
	// Obter as categorias mais vistas pelo usuário
	categories, err := getTopCategoriesForUser(h.DB, userID)
	if err != nil {
		return nil, err
	}

	var recommendedProducts []models.Product

	// Para cada categoria, recomendar produtos com base no número de visualizações
	for _, categoryView := range categories {
		// Calcular o limite de produtos a serem recomendados com base nas visualizações
		// Por exemplo, se a categoria foi vista 2 vezes, recomendar 6 produtos (2 * 3)
		limit := categoryView.Views * 3

		// Obter os IDs dos produtos recomendados para a categoria
		productIDs, err := h.RecommendPopularProductsByCategory(categoryView.Category, limit)
		if err != nil {
			return nil, err
		}

		// Para cada ID de produto, buscar os detalhes do produto e adicionar à lista de recomendados
		for _, id := range productIDs {
			product := models.Product{}
			row := h.DB.QueryRow("SELECT * FROM products WHERE id = $1", id)
			err = row.Scan(&product.ID, &product.Title, &product.Description, &product.Price, &product.Category, &product.Image)
			if err != nil {
				return nil, err
			}
			recommendedProducts = append(recommendedProducts, product)
		}
	}

	return recommendedProducts, nil
}

func (h *Handlers) RecommendProductsBasedOnCategoryViewsHandler(c *fiber.Ctx) error {
	sess, err := h.Store.Get(c)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	userID := sess.Get("user_id").(int)

	recommendedProducts, err := h.RecommendProductsBasedOnCategoryViews(userID)
	if err != nil {
		return err
	}

	return c.JSON(recommendedProducts)
}

func getTopCategoriesForSession(db *sql.DB, sessionID uuid.UUID) ([]CategoryView, error) {
	query := `
        SELECT p.category, COUNT(*) AS views
        FROM session_navigation_history snh
        JOIN products p ON snh.product_id = p.id
        WHERE snh.session_id = $1
        GROUP BY p.category
        ORDER BY views DESC
        LIMIT 4;
    `
	rows, err := db.Query(query, sessionID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var categories []CategoryView
	for rows.Next() {
		var cv CategoryView
		if err := rows.Scan(&cv.Category, &cv.Views); err != nil {
			return nil, err
		}
		categories = append(categories, cv)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return categories, nil
}

func (h *Handlers) RecommendProductsBasedOnCategoryViewsForSession(sessionID uuid.UUID) ([]models.Product, error) {
	// Obter as categorias mais vistas na sessão
	categories, err := getTopCategoriesForSession(h.DB, sessionID)
	if err != nil {
		return nil, err
	}

	var recommendedProducts []models.Product

	// Para cada categoria, recomendar produtos com base no número de visualizações
	for _, categoryView := range categories {
		// Calcular o limite de produtos a serem recomendados com base nas visualizações
		// Por exemplo, se a categoria foi vista 2 vezes, recomendar 6 produtos (2 * 3)
		limit := categoryView.Views * 3

		// Obter os IDs dos produtos recomendados para a categoria
		productIDs, err := h.RecommendPopularProductsByCategory(categoryView.Category, limit)
		if err != nil {
			return nil, err
		}

		// Para cada ID de produto, buscar os detalhes do produto e adicionar à lista de recomendados
		for _, id := range productIDs {
			product := models.Product{}
			row := h.DB.QueryRow("SELECT * FROM products WHERE id = $1", id)
			err = row.Scan(&product.ID, &product.Title, &product.Description, &product.Price, &product.Category, &product.Image)
			if err != nil {
				return nil, err
			}
			recommendedProducts = append(recommendedProducts, product)
		}
	}

	return recommendedProducts, nil
}

func (h *Handlers) RecommendProductsBasedOnSessionViewsHandler(c *fiber.Ctx) error {
	// Extrair sessionID da query
	sessionIDStr := c.Query("sessionID")
	if sessionIDStr == "" {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "sessionID is required",
		})
	}

	sessionID, err := uuid.Parse(sessionIDStr)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "invalid sessionID format",
		})
	}

	// Obter produtos recomendados com base nas visualizações da sessão
	recommendedProducts, err := h.RecommendProductsBasedOnCategoryViewsForSession(sessionID)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	return c.JSON(recommendedProducts)
}
