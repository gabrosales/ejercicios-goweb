package product

import (
	"ejercicios-goweb/internal/domain"
	"errors"
)

var (
	ErrAlreadyExist = errors.New("Already exist")
	ErrNotExist     = errors.New("Does not exist")
)

// controller
type Service interface {
	GetAll() ([]domain.Product, error)
	GetByID(id int) (domain.Product, error)
	SearchProductsByPrice(priceGt int) ([]domain.Product, error)

	Create(name string, quantity int, code_value string, is_published bool, expiration string, price float64) (domain.Product, error)
	Update(id int, name string, quantity int, code_value string, is_published bool, expiration string, price float64) (domain.Product, error)
}

func NewService(rp Repository) Service {
	return &service{rp: rp}
}

type service struct {
	// repo
	rp Repository

	// external api's
	// ...
}

// read
func (sv *service) GetAll() ([]domain.Product, error) {
	return sv.rp.GetAll()
}
func (sv *service) GetByID(id int) (domain.Product, error) {
	return sv.rp.GetByID(id)
}
func (sv *service) SearchProductsByPrice(priceGt int) ([]domain.Product, error) {
	return sv.rp.SearchProductsByPrice(priceGt)
}

// write
func (sv *service) Create(name string, quantity int, code_value string, is_published bool, expiration string, price float64) (domain.Product, error) {
	if sv.rp.ExistCode(code_value) {
		return domain.Product{}, ErrAlreadyExist
	}

	product := domain.Product{
		Name:         name,
		Quantity:     quantity,
		Code_value:   code_value,
		Is_published: is_published,
		Expiration:   expiration,
		Price:        price,
	}

	lastID, err := sv.rp.Create(product)
	if err != nil {
		return domain.Product{}, err
	}

	product.ID = lastID

	return product, nil
}

func (sv *service) Update(id int, name string, quantity int, code_value string, is_published bool, expiration string, price float64) (domain.Product, error) {
	if !sv.rp.ExistId(id) {
		return domain.Product{}, ErrNotExist
	}

	product := domain.Product{
		ID:           id,
		Name:         name,
		Quantity:     quantity,
		Code_value:   code_value,
		Is_published: is_published,
		Expiration:   expiration,
		Price:        price,
	}

	err := sv.rp.Update(product)
	if err != nil {
		return domain.Product{}, err
	}

	result, err := sv.GetByID(id)

	if err != nil {
		return domain.Product{}, err
	}

	return result, nil
}
