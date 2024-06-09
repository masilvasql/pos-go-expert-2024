package entity

import (
	"errors"
	"time"

	"github.com/masilvasql/pos-go-expert-2024/APIs/pkg/entity"
)

var (
	ErrIdIsRequired  = errors.New("id is required")
	ErrInvalidId     = errors.New("invalid Id")
	ErrNameRequired  = errors.New("name is required")
	ErrPriceRequired = errors.New("price is required")
	ErrInvalidPrice  = errors.New("invalid Price")
)

type Product struct {
	ID         entity.ID `json:"id"`
	Name       string    `json:"name"`
	Price      int       `json:"price"`
	CreatedAtd time.Time `json:"created_at"`
}

func NewProduct(name string, price int) (*Product, error) {
	product := &Product{
		ID:         entity.NewId(),
		Name:       name,
		Price:      price,
		CreatedAtd: time.Now(),
	}

	err := product.Validate()
	if err != nil {
		return nil, err
	}
	return product, nil
}

func (p *Product) Validate() error {
	if p.ID.String() == "" {
		return ErrIdIsRequired
	}

	if _, err := entity.ParseID(p.ID.String()); err != nil {
		return ErrInvalidId
	}

	if p.Name == "" {
		return ErrNameRequired
	}

	if p.Price == 0 {
		return ErrPriceRequired
	}

	if p.Price < 0 {
		return ErrInvalidPrice
	}

	return nil
}
