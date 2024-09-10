package product

import (
	"database/sql"
	"fmt"
	"strings"

	"github.com/Tutuacs/learn_golang/api.git/types"
)

type Store struct {
	db *sql.DB
}

func NewStore(db *sql.DB) *Store {
	return &Store{db: db}
}

func (s *Store) GetProducts() ([]types.Product, error) {
	sql := "SELECT * FROM products"
	rows, err := s.db.Query(sql)
	if err != nil {
		return nil, err
	}

	products := make([]types.Product, 0)
	for rows.Next() {
		prod, err := scanRowsIntoProduct(rows)
		if err != nil {
			continue
		}

		products = append(products, *prod)
	}

	return products, nil
}

func (s *Store) GetProductsByIDs(productIDs []int) ([]types.Product, error) {
	placeHolders := strings.Repeat(",?", len(productIDs)-1)
	sql := "SELECT * FROM products WHERE id IN (?%s)"
	query := fmt.Sprintf(sql, placeHolders)

	args := make([]interface{}, len(productIDs))
	for i, v := range productIDs {
		args[i] = v
	}

	rows, err := s.db.Query(query, args...)
	if err != nil {
		return nil, err
	}

	products := []types.Product{}
	for rows.Next() {
		prod, err := scanRowsIntoProduct(rows)
		if err != nil {
			continue
		}

		products = append(products, *prod)
	}

	return products, nil
}

func (s *Store) CreateProduct(product types.Product) error {

	sql := "INSERT INTO products (name, description, image, price, quantity) VALUES (?,?,?,?,?)"
	_, err := s.db.Exec(sql, product.Name, product.Description, product.Image, product.Price, product.Quantity)
	if err != nil {
		return err
	}

	return nil
}

func scanRowsIntoProduct(rows *sql.Rows) (*types.Product, error) {
	product := new(types.Product)

	err := rows.Scan(
		&product.ID,
		&product.Name,
		&product.Description,
		&product.Image,
		&product.Price,
		&product.Quantity,
		&product.CreatedAt,
	)
	if err != nil {
		return nil, err
	}

	return product, nil
}

func (s *Store) UpdateProduct(product types.Product) error {

	sql := "UPDATE products SET name = ?, price = ?, image = ?, description = ? quantity = ? WHERE id = ?"

	_, err := s.db.Exec(sql, product.Name, product.Price, product.Image, product.Description, product.Quantity, product.ID)
	if err != nil {
		return err
	}

	return nil

}
