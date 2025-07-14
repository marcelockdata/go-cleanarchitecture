package usecase

import (
	"cleanarchitecture/internal/domain"
	"cleanarchitecture/internal/repository"
	"errors"
)

type ProductUsecase interface {
	Create(product domain.Product) error
	GetAll() ([]domain.Product, error)
	GetById(id int) (domain.Product, error)
	Update(id int, product domain.Product) error
	Delete(id int) error
}

type productUsecase struct {
	repo repository.ProductRepository
}

func NewProductUsecase(repo repository.ProductRepository) ProductUsecase {
	return &productUsecase{repo: repo}
}

func (u *productUsecase) Create(product domain.Product) error {
	if product.Name == "" || product.Price <= 0 {
		return errors.New("invalid product data")
	}
	return u.repo.Create(product)
}

func (u *productUsecase) GetAll() ([]domain.Product, error) {
	return u.repo.GetAll()
}

func (u *productUsecase) GetById(id int) (domain.Product, error) {
	if id <= 0 {
		return domain.Product{}, errors.New("invalid product ID")
	}
	return u.repo.GetById(id)
}

func (u *productUsecase) Update(id int, product domain.Product) error {
	if id <= 0 || product.Name == "" || product.Price <= 0 {
		return errors.New("invalid data for update")
	}
	return u.repo.Update(id, product)
}

func (u *productUsecase) Delete(id int) error {
	if id <= 0 {
		return errors.New("invalid product ID")
	}
	return u.repo.Delete(id)
}
