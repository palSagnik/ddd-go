package memory

import (
	"domain/aggregate"
	"domain/domain/product"
	"sync"

	"github.com/google/uuid"
)

type MemoryProductRepository struct {
	products map[uuid.UUID]aggregate.Product
	sync.Mutex
}

func New() *MemoryProductRepository {
	return &MemoryProductRepository{
		products: make(map[uuid.UUID]aggregate.Product),
	}
}

func (mempr *MemoryProductRepository) GetAll() ([]aggregate.Product, error) {
	var products []aggregate.Product
	for _, product := range mempr.products {
		products = append(products, product)
	}

	return products, nil
}

func (mempr *MemoryProductRepository) GetByID(productID uuid.UUID) (aggregate.Product, error) {
	if product, ok := mempr.products[productID]; ok {
		return product, nil
	}

	return aggregate.Product{}, product.ErrProductNotFound
}