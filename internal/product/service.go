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
	Delete(id int) error
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

	p, err := sv.GetByID(id)

	if err != nil {
		return domain.Product{}, ErrNotExist
	}

	if name != "" {
		p.Name = name
	}
	if code_value != "" {
		p.Code_value = code_value
	}
	if expiration != "" {
		p.Expiration = expiration
	}
	if quantity > 0 {
		p.Quantity = quantity
	}
	if price > 0 {
		p.Price = price
	}

	result, err := sv.rp.Update(p)
	if err != nil {
		return domain.Product{}, err
	}

	return result, nil
}

func (sv *service) Delete(id int) error {
	err := sv.rp.Delete(id)
	if err != nil {
		return err
	}
	return nil
}
