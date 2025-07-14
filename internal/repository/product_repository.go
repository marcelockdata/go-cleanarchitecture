package repository

import (
	"cleanarchitecture/internal/domain"
	"database/sql"
	"errors"

	_ "github.com/mattn/go-sqlite3"
)

type ProductRepository interface {
	Create(product domain.Product) error
	GetAll() ([]domain.Product, error)
	GetById(id int) (domain.Product, error)
	Update(id int, product domain.Product) error
	Delete(id int) error
}

type productRepository struct {
	db *sql.DB
}

func NewProductRepository(db *sql.DB) ProductRepository {
	return &productRepository{db: db}
}

func (r *productRepository) Create(product domain.Product) error {
	_, err := r.db.Exec("INSERT INTO products (name, price) VALUES (?, ?)", product.Name, product.Price)
	return err
}
func (r *productRepository) GetAll() ([]domain.Product, error) {
	rows, err := r.db.Query("SELECT id, name, price FROM products")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var products []domain.Product
	for rows.Next() {
		var product domain.Product
		if err := rows.Scan(&product.ID, &product.Name, &product.Price); err != nil {
			return nil, err
		}
		products = append(products, product)
	}

	return products, nil
}

func (r *productRepository) GetById(id int) (domain.Product, error) {
	query := "SELECT id, name, price FROM products WHERE id = ?"
	row := r.db.QueryRow(query, id)

	var product domain.Product
	if err := row.Scan(&product.ID, &product.Name, &product.Price); err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return domain.Product{}, errors.New("product not found")
		}
		return domain.Product{}, err
	}

	return product, nil
}

func (r *productRepository) Update(id int, product domain.Product) error {
	query := "UPDATE products SET name = ?, price = ? WHERE id = ?"
	result, err := r.db.Exec(query, product.Name, product.Price, id)
	if err != nil {
		return err
	}

	affected, err := result.RowsAffected()
	if err != nil {
		return err
	}

	if affected == 0 {
		return errors.New("no rows affected")
	}

	return nil
}

func (r *productRepository) Delete(id int) error {
	query := "DELETE FROM products WHERE id = ?"
	result, err := r.db.Exec(query, id)
	if err != nil {
		return err
	}

	affected, err := result.RowsAffected()
	if err != nil {
		return err
	}

	if affected == 0 {
		return errors.New("no rows affected")
	}

	return nil
}

func InitDB(filepath string) (*sql.DB, error) {
	db, err := sql.Open("sqlite3", filepath)
	if err != nil {
		return nil, err
	}
	if _, err := db.Exec(`CREATE TABLE IF NOT EXISTS products (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		name TEXT NOT NULL,
		price REAL NOT NULL
	);`); err != nil {
		return nil, err
	}

	if _, err := db.Exec(`CREATE TABLE IF NOT EXISTS users (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		username TEXT NOT NULL,
		password TEXT NOT NULL
	);`); err != nil {
		return nil, err
	}
	return db, nil
}
