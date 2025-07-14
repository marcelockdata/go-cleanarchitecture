package repository

import (
	"cleanarchitecture/internal/domain"
	"database/sql"
	"regexp"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/assert"
)

func TestCreateProduct(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	repo := NewProductRepository(db)

	product := domain.Product{Name: "Test Product", Price: 10.0}

	mock.ExpectExec("INSERT INTO products").
		WithArgs(product.Name, product.Price).
		WillReturnResult(sqlmock.NewResult(1, 1))

	err = repo.Create(product)
	assert.NoError(t, err)

	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expectations: %s", err)
	}
}

func TestGetAllProducts(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	repo := NewProductRepository(db)

	rows := sqlmock.NewRows([]string{"id", "name", "price"}).
		AddRow(1, "Product 1", 10.0).
		AddRow(2, "Product 2", 20.0)

	mock.ExpectQuery("SELECT id, name, price FROM products").WillReturnRows(rows)

	products, err := repo.GetAll()
	assert.NoError(t, err)
	assert.Len(t, products, 2)
	assert.Equal(t, "Product 1", products[0].Name)
	assert.Equal(t, 10.0, products[0].Price)
	assert.Equal(t, "Product 2", products[1].Name)
	assert.Equal(t, 20.0, products[1].Price)

	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expectations: %s", err)
	}
}

func TestGetProductById(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	repo := NewProductRepository(db)

	row := sqlmock.NewRows([]string{"id", "name", "price"}).
		AddRow(1, "Product 1", 10.0)

	mock.ExpectQuery("SELECT id, name, price FROM products WHERE id = ?").
		WithArgs(1).
		WillReturnRows(row)

	product, err := repo.GetById(1)
	assert.NoError(t, err)
	assert.Equal(t, "Product 1", product.Name)
	assert.Equal(t, 10.0, product.Price)

	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expectations: %s", err)
	}
}

func TestUpdateProduct(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	repo := NewProductRepository(db)

	product := domain.Product{Name: "Updated Product", Price: 15.0, ID: 1}

	query := regexp.QuoteMeta("UPDATE products SET name = ?, price = ? WHERE id = ?")
	mock.ExpectExec(query).
		WithArgs(product.Name, product.Price, product.ID).
		WillReturnResult(sqlmock.NewResult(1, 1))

	err = repo.Update(product.ID, product)
	assert.NoError(t, err)

	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expectations: %s", err)
	}
}

func TestDeleteProduct(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	repo := NewProductRepository(db)

	mock.ExpectExec("DELETE FROM products WHERE id = ?").
		WithArgs(1).
		WillReturnResult(sqlmock.NewResult(1, 1))

	err = repo.Delete(1)
	assert.NoError(t, err)

	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expectations: %s", err)
	}
}

func TestGetProductByIdNotFound(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	repo := NewProductRepository(db)

	mock.ExpectQuery("SELECT id, name, price FROM products WHERE id = ?").
		WithArgs(1).
		WillReturnError(sql.ErrNoRows)

	_, err = repo.GetById(1)
	assert.Error(t, err)
	assert.Equal(t, "product not found", err.Error())

	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expectations: %s", err)
	}
}
