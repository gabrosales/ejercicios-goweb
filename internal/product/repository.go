package product

import (
	"ejercicios-goweb/internal/domain"
	"errors"
	"fmt"
)

var (
	ErrNotFound = errors.New("Product not found")
)

type Repository interface {
	// read
	GetAll() ([]domain.Product, error)
	GetByID(id int) (domain.Product, error)
	ExistCode(url string) bool
	ExistId(id int) bool
	SearchProductsByPrice(priceGt int) ([]domain.Product, error)
	// write
	Create(domain.Product) (int, error)
	Update(domain.Product) (domain.Product, error)
	Delete(id int) error
}

type repository struct {
	db *[]domain.Product
	// config
	lastID int
}

func NewRepository(db *[]domain.Product, lastID int) Repository {
	return &repository{db: db, lastID: lastID}
}

// read
func (r *repository) GetAll() ([]domain.Product, error) {
	return *r.db, nil
}
func (r *repository) GetByID(id int) (domain.Product, error) {
	for _, pro := range *r.db {
		if pro.ID == id {
			return pro, nil
		}
	}

	return domain.Product{}, fmt.Errorf("%w. %s", ErrNotFound, "Product does not exist")
}

func (r *repository) SearchProductsByPrice(priceGt int) ([]domain.Product, error) {
	var productsFiltered []domain.Product

	for _, pro := range *r.db {
		if pro.Price > float64(priceGt) {
			productsFiltered = append(productsFiltered, pro)
		}
	}

	if productsFiltered != nil {
		return productsFiltered, nil
	}

	return []domain.Product{}, fmt.Errorf("%w. %s", ErrNotFound, "Products does not exist")
}

func (r *repository) ExistCode(code_value string) bool {
	for _, pro := range *r.db {
		if pro.Code_value == code_value {
			return true
		}
	}

	return false
}

func (r *repository) ExistId(id int) bool {
	for _, pro := range *r.db {
		if pro.ID == id {
			return true
		}
	}

	return false
}

// write
func (r *repository) Create(product domain.Product) (int, error) {
	r.lastID++
	product.ID = r.lastID
	*r.db = append(*r.db, product)

	return r.lastID, nil
}

func (r *repository) Update(product domain.Product) (domain.Product, error) {

	for i, pro := range *r.db {
		if pro.ID == product.ID {
			(*r.db)[i] = product
			return product, nil
		}
	}

	return domain.Product{}, fmt.Errorf("%w. %s", ErrNotFound, "Product does not exist")
}

func (r *repository) Delete(id int) error {
	for i, pro := range *r.db {
		if pro.ID == id {
			(*r.db) = append((*r.db)[:i], (*r.db)[i+1:]...)
			return nil
		}
	}
	return errors.New("product not found")
}
