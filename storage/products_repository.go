package storage

import (
	"encoding/json"
	"errors"

	"viniti.us/hashout/config/db"
	"viniti.us/hashout/models/checkout"
	customErr "viniti.us/hashout/models/errors"
)

type ProductRepository struct {
	db db.Database
}

func NewProductRepository(d db.Database) ProductRepository {
	return ProductRepository{db: d}
}

func (r ProductRepository) FindAll(items []checkout.Item) (mergedItems []checkout.Item, err error) {
	var notFound []int32
	for _, i := range items {
		if p, found := r.db.FindProductByID(i.Product.ID); found {
			mergedItems = append(mergedItems, i.Merge(p))
		} else {
			notFound = append(notFound, i.Product.ID)
		}
	}

	if len(notFound) > 0 {
		p, _ := json.Marshal(notFound)
		err = &customErr.NotFound{Entity: "Product", Err: errors.New("the following products were not found: " + string(p))}
	}

	return
}

func (r ProductRepository) FindLastByIsGift(isGift bool) (product checkout.Product, err error) {
	product, _ = r.db.FindLastProductByIsGift(isGift)
	return
}
