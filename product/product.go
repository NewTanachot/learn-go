package product

import (
	"database/sql"
	db "github.com/NewTanachot/learn-go/database"
)

type Product struct {
	Id    int
	Name  string
	Price int
}

var dbContext *sql.DB

// this init func is a special function in Go that is executed automatically upon package initialization
func init() {
	tempDb, connectionError := db.Connect()

	if connectionError != nil {
		panic(connectionError)
	}

	dbContext = tempDb
}

func CreateProduct(product *Product) error {
	_, insertError := dbContext.Exec(
		"INSERT INTO public.products(name, price) VALUES ($1, $2);",
		product.Name,
		product.Price)

	return insertError
}

func GetProduct(id int) (*Product, error) {
	product := new(Product)
	row := dbContext.QueryRow(
		"SELECT id, name, price FROM products WHERE id = $1;", id)

	err := row.Scan(&product.Id, &product.Name, &product.Price)

	if err != nil {
		return product, err
	}

	return product, nil
}

func GetProducts() (*[]Product, error) {
	rows, err := dbContext.Query("SELECT id, name, price FROM products")

	if err != nil {
		return nil, err
	}

	var products []Product

	for rows.Next() {
		var product Product
		err := rows.Scan(&product.Id, &product.Name, &product.Price)

		if err != nil {
			return nil, err
		}

		products = append(products, product)
	}

	if rows.Err() != nil {
		return nil, err
	}

	return &products, nil
}

func UpdateProduct(id int, name string, price int) (*Product, error) {

	product := new(Product)
	row := dbContext.QueryRow(
		"UPDATE products SET name = $1, price = $2 WHERE id = $3 RETURNING id, name, price;",
		name, price, id)

	err := row.Scan(&product.Id, &product.Name, &product.Price)

	if err != nil {
		return product, err
	}

	return product, nil

	// _, err := db.Exec(
	// 	`UPDATE products SET name = $1, price = $2 WHERE id = $5;`,
	// 	name, price, id)

	// return err
}

func DeleteProduct(id int) error {
	_, err := dbContext.Exec(`DELETE FROM products WHERE id = $1;`, id)
	return err
}
