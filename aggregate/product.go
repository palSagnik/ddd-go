package aggregate

import (
	"domain/entity"
	"errors"

	"github.com/google/uuid"
)

var (
	ErrMissingValues = errors.New("missing values")
)

type Product struct {
	item     *entity.Item
	price    float64
	quantity int
}

func NewProduct(name string, desc string, price float64) (Product, error) {
	if len(name) == 0 || len(desc) == 0 || price == float64(0) {
		return Product{}, ErrMissingValues
	}

	return Product{
		item: &entity.Item{
			Name: name,
			Description: desc,
		},
		price: price,
		quantity: 0,
	}, nil
}

func (p Product) GetID() uuid.UUID {
	return p.item.ID
}

func (p Product) GetItem() *entity.Item {
	return p.item
}

func (p Product) GetName() string {
	return p.item.Name
}