package handler

import (
	"database/sql"
	"log"

	"github.com/gofiber/fiber"
	"github.com/pascaloseko/go-rest-fibre-postgres/database"
	"github.com/pascaloseko/go-rest-fibre-postgres/model"
)

// GetAllProducts from db
func GetAllProducts(c *fiber.Ctx) {
	rows, err := database.DB.Query("SELECT id, name, description, category, amount FROM products order by name")
	if err != nil {
		c.Status(500).JSON(&fiber.Map{
			"success": false,
			"error":   err,
		})
		return
	}

	defer rows.Close()
	result := model.Products{}

	for rows.Next() {
		product := model.Product{}
		err := rows.Scan(&product.ID, &product.Name, &product.Description, &product.Category, &product.Amount)
		if err != nil {
			c.Status(500).JSON(&fiber.Map{
				"success": false,
				"error":   err,
			})
			return
		}
		result.Products = append(result.Products, product)
	}

	if err := c.JSON(&fiber.Map{
		"success":  true,
		"products": result,
		"message":  "All product returned successfully",
	}); err != nil {
		c.Status(500).JSON(&fiber.Map{
			"success": false,
			"message": err,
		})
		return
	}
}

// GetSingleProduct from db
func GetSingleProduct(c *fiber.Ctx) {
	id := c.Params("id")
	product := model.Product{}

	row, err := database.DB.Query("SELECT * FROM products WHERE id = $1", id)
	if err != nil {
		c.Status(500).JSON(&fiber.Map{
			"success": false,
			"message": err,
		})
		return
	}

	defer row.Close()
	for row.Next() {
		switch err := row.Scan(&product.ID, &product.Amount, &product.Name, &product.Description, &product.Category); err {
		case sql.ErrNoRows:
			log.Println("No rows were returned!")
			c.Status(500).JSON(&fiber.Map{
				"success": false,
				"message": err,
			})
		case nil:
			log.Println(product.Name, product.Description, product.Category, product.Amount)
		default:
			//   panic(err)
			c.Status(500).JSON(&fiber.Map{
				"success": false,
				"message": err,
			})
		}
	}

	if err := c.JSON(&fiber.Map{
		"success": false,
		"message": "Successfully fetched product",
		"product": product,
	}); err != nil {
		c.Status(500).JSON(&fiber.Map{
			"success": false,
			"message": err,
		})
		return
	}
}

// CreateProduct handler
func CreateProduct(c *fiber.Ctx) {
	// Instantiate new Product struct
	p := new(model.Product)
	//  Parse body into product struct
	if err := c.BodyParser(p); err != nil {
		log.Println(err)
		c.Status(400).JSON(&fiber.Map{
			"success": false,
			"message": err,
		})
		return
	}
	// Insert Product into database
	_, err := database.DB.Query("INSERT INTO products (name, description, category, amount) VALUES ($1, $2, $3, $4)", p.Name, p.Description, p.Category, p.Amount)
	if err != nil {
		c.Status(500).JSON(&fiber.Map{
			"success": false,
			"message": err,
		})
		return
	}

	// Return Product in JSON format
	if err := c.Status(201).JSON(&fiber.Map{
		"success": true,
		"message": "Product successfully created",
		"product": p,
	}); err != nil {
		c.Status(500).JSON(&fiber.Map{
			"success": false,
			"message": "Error creating product",
		})
		return
	}
}

// UpdateProduct handler
func UpdateProduct(c *fiber.Ctx) {
	id := c.Params("id")
	// Instantiate new Product struct
	p := new(model.Product)
	//  Parse body into product struct
	if err := c.BodyParser(p); err != nil {
		log.Println(err)
		c.Status(400).JSON(&fiber.Map{
			"success": false,
			"message": err,
		})
		return
	}
	// Update Product database
	_, err := database.DB.Query("UPDATE products SET name=$1, description=$2, category=$3, amount=$4 WHERE id = $5", p.Name, p.Description, p.Category, p.Amount, id)
	if err != nil {
		c.Status(500).JSON(&fiber.Map{
			"success": false,
			"message": err,
		})
		return
	}

	// Return Product in JSON format
	if err := c.Status(200).JSON(&fiber.Map{
		"success": true,
		"message": "Product successfully updated",
		"product": p,
	}); err != nil {
		c.Status(500).JSON(&fiber.Map{
			"success": false,
			"message": "Error creating product",
		})
		return
	}
}

// DeleteProduct from db
func DeleteProduct(c *fiber.Ctx) {
	id := c.Params("id")
	// query product table in database
	_, err := database.DB.Query("DELETE FROM products WHERE id = $1", id)
	if err != nil {
		c.Status(500).JSON(&fiber.Map{
			"success": false,
			"error":   err,
		})
		return
	}
	// return product in JSON format
	if err := c.JSON(&fiber.Map{
		"success": true,
		"message": "product deleted successfully",
	}); err != nil {
		c.Status(500).JSON(&fiber.Map{
			"success": false,
			"error":   err,
		})
		return
	}
}
