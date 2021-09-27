package storage

import (
	"errors"

	"viniti.us/hashout/config/db"
	"viniti.us/hashout/models/checkout"
	customErr "viniti.us/hashout/models/errors"
)

type ProductRepository struct {
	db db.ProductsDataset
}

func NewProductRepository(d db.ProductsDataset) ProductRepository {
	return ProductRepository{db: d}
}

func (r ProductRepository) Find(items []checkout.Item) (mergedItems []checkout.Item, err error) {
	var notFound []int32
	for _, i := range items {
		if p, found := r.db.ByID[i.Product.ID]; found {
			mergedItems = append(mergedItems, i.Merge(p))
		} else {
			notFound = append(notFound, p.ID)
		}
	}

	if len(notFound) > 0 {
		err = &customErr.NotFound{Entity: "Product", Err: errors.New("the following products were not found: " + string(notFound))}
	}

	return
}
